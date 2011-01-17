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
//	this is a go port of said algorithm
// 	*/
// 
package main

import (
	// 	"strings"
	"os"
)
// 	/* stemmer is a structure for a few local bits of data,
// 	*/
// 
type Stemmer struct {
	b string
	k int
	j int
}
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
func NewStemmer() *Stemmer {

	// 	st = make(Stemmer)
	return new(Stemmer)
}

func (st *Stemmer) printS() {
	println(st.b)
	println(st.k, st.j)
}
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

// 	/* r(z, s) is used further down. */
// 
func (st *Stemmer) r(s string) {
	if st.m() > 0 {
		st.setto(s)
	}
}

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

// 	/* step1c(z) turns terminal y to i when there is another vowel in the stem. */
// 
func (st *Stemmer) step1c() {

	if st.ends("y") && st.vowelinstem() {
		st.b = st.b[:len(st.b)-1]
		st.b += "i"
	}
}

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

// 	/* step4(z) takes off -ant, -ence etc., in context <c>vcvc<v>. */
// 
func (st *Stemmer) step4() {
	if st.k < 3 {
		return
	}

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

	if len(in) < 3 {
		out = in
		return 
	}
	st.step1ab()
	st.step1c()
	st.step2()
	st.step3()
	st.step4()
	st.step5()

	out = st.b[:st.k+1]
	return out, err
}
