// 	/* This is the Porter stemming algorithm, coded up as thread-safe ANSI C
// 	by the author.
// 
// 	It may be be regarded as cononical, in that it follows the algorithm
// 	presented in
// 
// 	Porter, 1980, An algorithm for suffix stripping, Program, Vol. 14,
// 	no. 3, pp 130-137,
// 
// 	only differing from it at the points maked --DEPARTURE-- below.
// 
// 	See also http://www.tartarus.org/~martin/PorterStemmer
// 
// 	The algorithm as described in the paper could be exactly replicated
// 	by adjusting the points of DEPARTURE, but this is barely necessary,
// 	because (a) the points of DEPARTURE are definitely improvements, and
// 	(b) no encoding of the Porter stemmer I have seen is anything like
// 	as exact as this version, even with the points of DEPARTURE!
// 
// 	You can compile it on Unix with 'gcc -O3 -o stem stem.c' after which
// 	'stem' takes a list of inputs and sends the stemmed equivalent to
// 	stdout.
// 
// 	The algorithm as encoded here is particularly fast.
// 
// 	Release 2 (the more old-fashioned, non-thread-safe version may be
// 	regarded as release 1.)
// 	*/
// 
package main

import (
	// 	"strings"
	"os"
)
// 	#include <stdlib.h>  /* for malloc, free */
// 	#include <string.h>  /* for memcmp, memmove */
// 
// 	/* You will probably want to move the following declarations to a central
// 	header file.
// 	*/
// 
// 	struct stemmer;
// 
// 	extern struct stemmer * create_stemmer(void);
// 	extern void free_stemmer(struct stemmer * z);
// 
// 	extern int stem(struct stemmer * z, char * b, int k);
// 
// 
// 
// 	/* The main part of the stemming algorithm starts here.
// 	*/
// 
// 	#define TRUE 1
// 	#define FALSE 0
// 
// 	/* stemmer is a structure for a few local bits of data,
// 	*/
// 
type Stemmer struct {
	b string
	k int
	j int
}
// 	struct stemmer {
// 	char * b;       /* buffer for word to be stemmed */
// 	int k;          /* offset to the end of the string */
// 	int j;          /* a general offset into the string */
// 	};
// 
// 
// 	/* Member b is a buffer holding a word to be stemmed. The letters are in
// 	b[0], b[1] ... ending at b[z->k]. Member k is readjusted downwards as
// 	the stemming progresses. Zero termination is not in fact used in the
// 	algorithm.
// 
// 	Note that only lower case sequences are stemmed. Forcing to lower case
// 	should be done before stem(...) is called.
// 
// 
// 	Typical usage is:
// 
// 	struct stemmer * z = create_stemmer();
// 	char b[] = "pencils";
// 	int res = stem(z, b, 6);
// 		/- stem the 7 characters of b[0] to b[6]. The result, res,
// 		will be 5 (the 's' is removed). -/
// 	free_stemmer(z);
// 	*/
// 
// 
func NewStemmer() *Stemmer {

	// 	st = make(Stemmer)
	return new(Stemmer)
}

func (st *Stemmer) printS() {
	println(st.b)
	println(st.k, st.j)
}
// 	extern struct stemmer * create_stemmer(void)
// 	{
// 	return (struct stemmer *) malloc(sizeof(struct stemmer));
// 	/* assume malloc succeeds */
// 	}
// 
// 	extern void free_stemmer(struct stemmer * z)
// 	{
// 	free(z);
// 	}
// 
// 
// 	/* cons(z, i) is TRUE <=> b[i] is a consonant. ('b' means 'z->b', but here
// 	and below we drop 'z->' in comments.
// 	*/
// 
func (st *Stemmer) cons(i int) bool {
	// 	ch := st.b[i]
	// 	println(ch)
	switch st.b[i] {
	case 'a', 'e', 'i', 'o', 'u', 'y':
		return false
	}
	return true
}
// 	static int cons(struct stemmer * z, int i)
// 	{  switch (z->b[i])
// 	{  case 'a': case 'e': case 'i': case 'o': case 'u': return FALSE;
// 	case 'y': return (i == 0) ? TRUE : !cons(z, i - 1);
// 	default: return TRUE;
// 	}
// 	}
// 
// 	/* m(z) measures the number of consonant sequences between 0 and j. if c is
// 	a consonant sequence and v a vowel sequence, and <..> indicates arbitrary
// 	presence,
// 
// 	<c><v>       gives 0
// 	<c>vc<v>     gives 1
// 	<c>vcvc<v>   gives 2
// 	<c>vcvcvc<v> gives 3
// 	....
// 	*/
// 
func (st *Stemmer) m() int {
	n := 0
	i := 0
	j := st.j

	for ; ; i++ {

		if i > j {
			return n
		}
		if !st.cons(i) {
			break
		}

	}

	i++
	for {
		for ; ; i++ {

			if i > j {
				return n
			}

			if st.cons(i) {
				break
			}

		}

		i++
		n++
		for ; ; i++ {
			if i > j {
				return n
			}
			if !st.cons(i) {
				break
			}
		}
		i++
	}
	return n
}
// 	static int m(struct stemmer * z)
// 	{  int n = 0;
// 	int i = 0;
// 	int j = z->j;
// 	while(TRUE)
// 	{  if (i > j) return n;
// 	if (! cons(z, i)) break; i++;
// 	}
// 	i++;
// 	while(TRUE)
// 	{  while(TRUE)
// 	{  if (i > j) return n;
// 		if (cons(z, i)) break;
// 		i++;
// 	}
// 	i++;
// 	n++;
// 	while(TRUE)
// 	{  if (i > j) return n;
// 		if (! cons(z, i)) break;
// 		i++;
// 	}
// 	i++;
// 	}
// 	}
// 
// 	/* vowelinstem(z) is TRUE <=> 0,...j contains a vowel */
// 
func (st *Stemmer) vowelinstem() bool {
	j := st.j

	for i := 0; i <= j; i++ {
		if !st.cons(i) {
			return true
		}
	}
	return false

}
// 	static int vowelinstem(struct stemmer * z)
// 	{
// 	int j = z->j;
// 	int i; for (i = 0; i <= j; i++) if (! cons(z, i)) return TRUE;
// 	return FALSE;
// 	}
// 
// 	/* doublec(z, j) is TRUE <=> j,(j-1) contain a double consonant. */
// 
func (st *Stemmer) doublec(j int) bool {
	if j < 1 {
		return false
	}

	if st.b[j] != st.b[j-1] {
		return false
	}
	return st.cons(j)
}
// 	static int doublec(struct stemmer * z, int j)
// 	{
// 	char * b = z->b;
// 	if (j < 1) return FALSE;
// 	if (b[j] != b[j - 1]) return FALSE;
// 	return cons(z, j);
// 	}
// 
// 	/* cvc(z, i) is TRUE <=> i-2,i-1,i has the form consonant - vowel - consonant
// 	and also if the second c is not w,x or y. this is used when trying to
// 	restore an e at the end of a short word. e.g.
// 
// 	cav(e), lov(e), hop(e), crim(e), but
// 	snow, box, tray.
// 
// 	*/
// 
func (st *Stemmer) cvc(i int) bool {
	if i < 2 || !st.cons(i) || st.cons(i-1) || !st.cons(i-2) {
		return false
	}
	// 	ch := st.b[i]
	switch st.b[i] {
	case 'w', 'x', 'y':
		return false
	}
	return true
}
// 	static int cvc(struct stemmer * z, int i)
// 	{  if (i < 2 || !cons(z, i) || cons(z, i - 1) || !cons(z, i - 2)) return FALSE;
// 		{  int ch = z->b[i];
// 			if (ch  == 'w' || ch == 'x' || ch == 'y') return FALSE;
// 		}
// 		return TRUE;
// 	}
// 
// 	/* ends(z, s) is TRUE <=> 0,...k ends with the string s. */
// 
func (st *Stemmer) ends(s string) bool {
	// 	ends(s) is TRUE <=> k0,...k ends with the string s."""
	length := len(s)
	// 	println("before tiny",st.b,s,length,st.k)
	if st.k < length {
		return false
	}
	if s[length-1] != st.b[st.k] { // tiny speed-up
		return false
	}
	// 	println(s,length,st.k)

	if length > (st.k + 1) {
		return false
	}
	// 	println(s,length,st.k)
	if st.b[st.k-length+1:st.k+1] != s {
		return false
	}
	// 	println(s,length,st.k)
	st.j = st.k - length
	return true
	// 	return strings.HasSuffix(st.b,suf)
}
// 	static int ends(struct stemmer * z, char * s)
// 	{  int length = s[0];
// 	char * b = z->b;
// 	int k = z->k;
// 	if (s[length] != b[k]) return FALSE; /* tiny speed-up */
// 	if (length > k + 1) return FALSE;
// 	if (memcmp(b + k - length + 1, s + 1, length) != 0) return FALSE;
// 	z->j = k-length;
// 	return TRUE;
// 	}
// 
// 	/* setto(z, s) sets (j+1),...k to the characters in the string s, readjusting
// 	k. */
// 
func (st *Stemmer) setto(s string) {
	// 	st.b  = st.b[:len(rep)]
	length := len(s)
	if st.j >= st.k {
		return
	}
	// 	append(st.b, rep)

	st.b = st.b[:st.j+1] + s + st.b[st.j+length+1:]
	st.k = st.j + length
}

// 	static void setto(struct stemmer * z, char * s)
// 	{  int length = s[0];
// 	int j = z->j;
// 	memmove(z->b + j + 1, s + 1, length);
// 	z->k = j+length;
// 	}
// 
// 	/* r(z, s) is used further down. */
// 
func (st *Stemmer) r(s string) {
	if st.m() > 0 {
		st.setto(s)
	}
}
// 	static void r(struct stemmer * z, char * s) { if (m(z) > 0) setto(z, s); }
// 
// 	/* step1ab(z) gets rid of plurals and -ed or -ing. e.g.
// 
// 	caresses  ->  caress
// 	ponies    ->  poni
// 	ties      ->  ti
// 	caress    ->  caress
// 	cats      ->  cat
// 
// 	feed      ->  feed
// 	agreed    ->  agree
// 	disabled  ->  disable
// 
// 	matting   ->  mat
// 	mating    ->  mate
// 	meeting   ->  meet
// 	milling   ->  mill
// 	messing   ->  mess
// 
// 	meetings  ->  meet
// 
// 	*/
// 
func (st *Stemmer) step1ab() {
	// 	println(st.b,st.k)
	if st.b[st.k] == 's' {
		if st.ends("sses") {
			st.k -= 2
		}
		if st.ends("ies") {
			st.setto("i")
		}
		if st.b[st.k-1] != 's' {
			st.k--
		}

	}
	// 	println(st.b,st.k)
	if st.ends("eed") {
	} else if st.ends("ed") || (st.ends("ing") && st.vowelinstem()) {
		st.k = st.j
		if st.ends("at") {
			st.setto("ate")
		} else if st.ends("bl") {
			st.setto("ble")
		} else if st.ends("iz") {
			st.setto("ize")
		} else if st.doublec(st.k) {
			st.k--

			ch := st.b[st.k]
			switch ch {
			case 'l', 's', 'z':
				st.k++
			}
		}
	} else if st.m() == 1 && st.cvc(st.k) {
		st.setto("e")
	}
}
// 	static void step1ab(struct stemmer * z)
// 	{
// 	char * b = z->b;
// 	if (b[z->k] == 's')
// 	{  if (ends(z, "\04" "sses")) z->k -= 2; else
// 	if (ends(z, "\03" "ies")) setto(z, "\01" "i"); else
// 	if (b[z->k - 1] != 's') z->k--;
// 	}
// 	if (ends(z, "\03" "eed")) { if (m(z) > 0) z->k--; } else
// 	if ((ends(z, "\02" "ed") || ends(z, "\03" "ing")) && vowelinstem(z))
// 	{  z->k = z->j;
// 	if (ends(z, "\02" "at")) setto(z, "\03" "ate"); else
// 	if (ends(z, "\02" "bl")) setto(z, "\03" "ble"); else
// 	if (ends(z, "\02" "iz")) setto(z, "\03" "ize"); else
// 	if (doublec(z, z->k))
// 	{  z->k--;
// 		{  int ch = b[z->k];
// 		if (ch == 'l' || ch == 's' || ch == 'z') z->k++;
// 		}
// 	}
// 	else if (m(z) == 1 && cvc(z, z->k)) setto(z, "\01" "e");
// 	}
// 	}
// 
// 	/* step1c(z) turns terminal y to i when there is another vowel in the stem. */
// 
func (st *Stemmer) step1c() {

	if st.ends("y") && st.vowelinstem() {
		st.b = st.b[:len(st.b)-1]
		st.b += "i"
	}
}
// 	static void step1c(struct stemmer * z)
// 	{
// 	if (ends(z, "\01" "y") && vowelinstem(z)) z->b[z->k] = 'i';
// 	}
// 
// 
// 	/* step2(z) maps double suffices to single ones. so -ization ( = -ize plus
// 	-ation) maps to -ize etc. note that the string before the suffix must give
// 	m(z) > 0. */
// 
func (st *Stemmer) step2() {
	if st.k < 3 {
		return
	} //think it is needed to avoid to short words

	switch st.b[st.k-1] {
	case 'a':
		if st.ends("ational") {
			st.r("ate")
			break
		}
		if st.ends("tional") {
			st.r("tion")
			break
		}

	case 'c':
		if st.ends("enci") {
			st.r("ence")
			break
		}
		if st.ends("anci") {
			st.r("ance")
			break
		}

	case 'e':
		if st.ends("izer") {
			st.r("ize")
			break
		}

	case 'l':
		if st.ends("bli") {
			st.r("ble")
			break
		}
		if st.ends("alli") {
			st.r("al")
			break
		}
		if st.ends("entli") {
			st.r("ent")
			break
		}
		if st.ends("eli") {
			st.r("e")
			break
		}
		if st.ends("ousli") {
			st.r("ous")
			break
		}
	case 'o':
		if st.ends("ization") {
			st.r("ize")
			break
		}
		if st.ends("ation") {
			st.r("ate")
			break
		}
		if st.ends("ator") {
			st.r("ate")
			break
		}
	case 's':
		if st.ends("alism") {
			st.r("al")
			break
		}
		if st.ends("iveness") {
			st.r("ive")
			break
		}
		if st.ends("fulness") {
			st.r("ful")
			break
		}
		if st.ends("ousness") {
			st.r("ous")
			break
		}
	case 't':
		// 		println("t2:",st.b,st.k,st.j)
		if st.ends("aliti") {
			st.r("al")
			break
		}
		if st.ends("iviti") {
			st.r("ive")
			break
		}
		if st.ends("biliti") {
			st.r("ble")
			break
		}
	case 'g':
		if st.ends("logi") {
			st.r("log")
			break
		}
	}
}
// 	static void step2(struct stemmer * z) { switch (z->b[z->k-1])
// 	{
// 	case 'a': if (ends(z, "\07" "ational")) { r(z, "\03" "ate"); break; }
// 		if (ends(z, "\06" "tional")) { r(z, "\04" "tion"); break; }
// 		break;
// 	case 'c': if (ends(z, "\04" "enci")) { r(z, "\04" "ence"); break; }
// 		if (ends(z, "\04" "anci")) { r(z, "\04" "ance"); break; }
// 		break;
// 	case 'e': if (ends(z, "\04" "izer")) { r(z, "\03" "ize"); break; }
// 		break;
// 	case 'l': if (ends(z, "\03" "bli")) { r(z, "\03" "ble"); break; } /*-DEPARTURE-*/
// 
// 	/* To match the published algorithm, replace this line with
// 	case 'l': if (ends(z, "\04" "abli")) { r(z, "\04" "able"); break; } */
// 
// 		if (ends(z, "\04" "alli")) { r(z, "\02" "al"); break; }
// 		if (ends(z, "\05" "entli")) { r(z, "\03" "ent"); break; }
// 		if (ends(z, "\03" "eli")) { r(z, "\01" "e"); break; }
// 		if (ends(z, "\05" "ousli")) { r(z, "\03" "ous"); break; }
// 		break;
// 	case 'o': if (ends(z, "\07" "ization")) { r(z, "\03" "ize"); break; }
// 		if (ends(z, "\05" "ation")) { r(z, "\03" "ate"); break; }
// 		if (ends(z, "\04" "ator")) { r(z, "\03" "ate"); break; }
// 		break;
// 	case 's': if (ends(z, "\05" "alism")) { r(z, "\02" "al"); break; }
// 		if (ends(z, "\07" "iveness")) { r(z, "\03" "ive"); break; }
// 		if (ends(z, "\07" "fulness")) { r(z, "\03" "ful"); break; }
// 		if (ends(z, "\07" "ousness")) { r(z, "\03" "ous"); break; }
// 		break;
// 	case 't': if (ends(z, "\05" "aliti")) { r(z, "\02" "al"); break; }
// 		if (ends(z, "\05" "iviti")) { r(z, "\03" "ive"); break; }
// 		if (ends(z, "\06" "biliti")) { r(z, "\03" "ble"); break; }
// 		break;
// 	case 'g': if (ends(z, "\04" "logi")) { r(z, "\03" "log"); break; } /*-DEPARTURE-*/
// 
// 	/* To match the published algorithm, delete this line */
// 
// 	} }
// 
// 	/* step3(z) deals with -ic-, -full, -ness etc. similar strategy to step2. */
// 
func (st *Stemmer) step3() {
	if st.k < 3 {
		return
	}
	// 	
	switch st.b[st.k] {
	case 'e':
		if st.ends("icate") {
			st.r("ic")
		}
		if st.ends("ative") {
			st.r("")
		}
		if st.ends("alize") {
			st.r("al")
		}

	case 'i':
		if st.ends("iciti") {
			st.r("ic")
		}
	case 'l':
		if st.ends("ical") {
			st.r("ic")
		}
		if st.ends("ful") {
			st.r("")
		}
	case 's':
		if st.ends("ness") {
			st.r("")
		}
	}
}
// 	static void step3(struct stemmer * z) { switch (z->b[z->k])
// 	{
// 	case 'e': if (ends(z, "\05" "icate")) { r(z, "\02" "ic"); break; }
// 		if (ends(z, "\05" "ative")) { r(z, "\00" ""); break; }
// 		if (ends(z, "\05" "alize")) { r(z, "\02" "al"); break; }
// 		break;
// 	case 'i': if (ends(z, "\05" "iciti")) { r(z, "\02" "ic"); break; }
// 		break;
// 	case 'l': if (ends(z, "\04" "ical")) { r(z, "\02" "ic"); break; }
// 		if (ends(z, "\03" "ful")) { r(z, "\00" ""); break; }
// 		break;
// 	case 's': if (ends(z, "\04" "ness")) { r(z, "\00" ""); break; }
// 		break;
// 	} }
// 
// 	/* step4(z) takes off -ant, -ence etc., in context <c>vcvc<v>. */
// 
func (st *Stemmer) step4() {
	if st.k < 3 {
		return
	}
	// 	st.printS()
	// 	println("Step4")
	// 	st.j = st.k
	switch st.b[st.k-1] {
	case 'a':
		if st.ends("al") {
			break
		}
		return
	case 'c':
		if st.ends("ance") {
			break
		}
		if st.ends("ence") {
			break
		}
		return
	case 'e':
		if st.ends("er") {
			break
		}
	case 'i':
		if st.ends("ic") {
			break
		}
		return
	case 'l':
		if st.ends("able") {
			break
		}
		if st.ends("ible") {
			break
		}
		return
	case 'n':
		if st.ends("ant") {
			break
		}
		if st.ends("ement") {
			break
		}
		if st.ends("ment") {
			break
		}
		if st.ends("ent") {
			break
		}
		return
	case 'o':
		// 		st.printS()
		// 		println("here",st.b,st.k,st.j)
		if st.k <= 3 {
			return
		}
		if st.ends("ion") && (st.b[st.j] == 's' || st.b[st.j] == 't') {
			break
		}
		if st.ends("ou") {
			break
		}
		return

	case 's':
		if st.ends("ism") {
			break
		}
		return

	case 't':
		if st.ends("ate") {
			break
		}
		if st.ends("iti") {
			break
		}
		return

	case 'u':
		if st.ends("ous") {
			break
		}
		return

	case 'v':
		if st.ends("ive") {
			break
		}
		return

	case 'z':
		if st.ends("ize") {
			break
		}
		return
	default:
		return
	}
	if st.m() > 1 {
		st.k = st.j
	}
}
// 	static void step4(struct stemmer * z)
// 	{  switch (z->b[z->k-1])
// 	{  case 'a': if (ends(z, "\02" "al")) break; return;
// 	case 'c': if (ends(z, "\04" "ance")) break;
// 		if (ends(z, "\04" "ence")) break; return;
// 	case 'e': if (ends(z, "\02" "er")) break; return;
// 	case 'i': if (ends(z, "\02" "ic")) break; return;
// 	case 'l': if (ends(z, "\04" "able")) break;
// 		if (ends(z, "\04" "ible")) break; return;
// 	case 'n': if (ends(z, "\03" "ant")) break;
// 		if (ends(z, "\05" "ement")) break;
// 		if (ends(z, "\04" "ment")) break;
// 		if (ends(z, "\03" "ent")) break; return;
// 	case 'o': if (ends(z, "\03" "ion") && (z->b[z->j] == 's' || z->b[z->j] == 't')) break;
// 		if (ends(z, "\02" "ou")) break; return;
// 		/* takes care of -ous */
// 	case 's': if (ends(z, "\03" "ism")) break; return;
// 	case 't': if (ends(z, "\03" "ate")) break;
// 		if (ends(z, "\03" "iti")) break; return;
// 	case 'u': if (ends(z, "\03" "ous")) break; return;
// 	case 'v': if (ends(z, "\03" "ive")) break; return;
// 	case 'z': if (ends(z, "\03" "ize")) break; return;
// 	default: return;
// 	}
// 	if (m(z) > 1) z->k = z->j;
// 	}
// 
// 	/* step5(z) removes a final -e if m(z) > 1, and changes -ll to -l if
// 	m(z) > 1. */
// 
func (st *Stemmer) step5() {
	if st.k < 3 {
		return
	}
	st.j = st.k

	if st.b[st.k] == 'e' {
		if st.m() > 1 || (st.m() == 1 && !st.cvc(st.k-1)) {
			st.k--
		}
	}
	if st.b[st.k] == 'l' && st.doublec(st.k) && st.m() > 1 {
		st.k--
	}

}
// 	static void step5(struct stemmer * z)
// 	{
// 	char * b = z->b;
// 	z->j = z->k;
// 	if (b[z->k] == 'e')
// 	{  int a = m(z);
// 	if (a > 1 || a == 1 && !cvc(z, z->k - 1)) z->k--;
// 	}
// 	if (b[z->k] == 'l' && doublec(z, z->k) && m(z) > 1) z->k--;
// 	}
// 
// 	/* In stem(z, b, k), b is a char pointer, and the string to be stemmed is
// 	from b[0] to b[k] inclusive.  Possibly b[k+1] == '\0', but it is not
// 	important. The stemmer adjusts the characters b[0] ... b[k] and returns
// 	the new end-point of the string, k'. Stemming never increases word
// 	length, so 0 <= k' <= k.
// 	*/
// 
func (st *Stemmer) Stem(in string) (out string, err os.Error) {
	st.b = in
	st.k = len(in) - 1
	st.j = st.k
	// 	println("stem this :",in)
	st.step1ab()
	st.step1c()
	st.step2()
	st.step3()
	st.step4()
	st.step5()

	out = st.b[:st.k+1]
	// 	println("stemmed :",out)
	return out, err
}
// 	extern int stem(struct stemmer * z, char * b, int k)
// 	{
// 	if (k <= 1) return k; /*-DEPARTURE-*/
// 	z->b = b; z->k = k; /* copy the parameters into z */
// 
// 	/* With this line, strings of length 1 or 2 don't go through the
// 	stemming process, although no mention is made of this in the
// 	published algorithm. Remove the line to match the published
// 	algorithm. */
// 
// 	step1ab(z); step1c(z); step2(z); step3(z); step4(z); step5(z);
// 	return z->k;
// 	}
// 
// 	/*--------------------stemmer definition ends here------------------------*/
// 
// 	#include <stdio.h>
// 	#include <stdlib.h>      /* for malloc, free */
// 	#include <ctype.h>       /* for isupper, islower, tolower */
// 
// 	static char * s;         /* buffer for words tobe stemmed */
// 
// 	#define INC 50           /* size units in which s is increased */
// 	static int i_max = INC;  /* maximum offset in s */
// 
// 	#define LETTER(ch) (isupper(ch) || islower(ch))
// 
// 	void stemfile(struct stemmer * z, FILE * f)
// 	{  while(TRUE)
// 	{  int ch = getc(f);
// 	if (ch == EOF) return;
// 	if (LETTER(ch))
// 	{  int i = 0;
// 		while(TRUE)
// 		{  if (i == i_max)
// 		{  i_max += INC;
// 		s = realloc(s, i_max + 1);
// 		}
// 		ch = tolower(ch); /* forces lower case */
// 
// 		s[i] = ch; i++;
// 		ch = getc(f);
// 		if (!LETTER(ch)) { ungetc(ch,f); break; }
// 		}
// 		s[stem(z, s, i - 1) + 1] = 0;
// 		/* the previous line calls the stemmer and uses its result to
// 		zero-terminate the string in s */
// 		printf("%s",s);
// 	}
// 	else putchar(ch);
// 	}
// 	}
// 
// 	int main(int argc, char * argv[])
// 	{  int i;
// 
// 	struct stemmer * z = create_stemmer();
// 
// 	s = (char *) malloc(i_max + 1);
// 	for (i = 1; i < argc; i++)
// 	{  FILE * f = fopen(argv[i],"r");
// 	if (f == 0) { fprintf(stderr,"File %s not found\n",argv[i]); exit(1); }
// 	stemfile(z, f);
// 	}
// 	free(s);
// 
// 	free_stemmer(z);
// 
// 	return 0;
// 	}
