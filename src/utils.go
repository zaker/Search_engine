package main

import (
	"os"
	"path"
	"fmt"
	"regexp"
	"strings"
)

var stopWords = []string{"a", "able", "about", "across", "after", "all", "almost",
	"also", "am", "among", "an", "and", "any", "are", "as", "at", "be", "because", "been", "but",
	"by", "can", "cannot", "could", "dear", "did", "do", "does", "either", "else", "ever", "every",
	"for", "from", "get", "got", "had", "has", "have", "he", "her", "hers", "him", "his", "how", "however",
	"i", "if", "in", "into", "is", "it", "its", "just", "least", "let", "like", "likely", "may", "me",
	"might", "most", "must", "my", "neither", "no", "nor", "not", "of", "off", "often", "on", "only",
	"or", "other", "our", "own", "rather", "said", "say", "says", "she", "should", "since", "so", "some",
	"than", "that", "the", "their", "them", "then", "there", "these", "they", "this", "tis", "to", "too",
	"twas", "us", "wants", "was", "we", "were", "what", "when", "where", "which", "while", "who", "whom",
	"why", "will", "with", "would", "yet", "you", "your"}

func isStopWord(s string) bool {

	for i := range stopWords {
		if stopWords[i] == s {
			return true
		}
	}
	return false
}
func cleanS(s string) (out []string) {

	// 	remove symbols and numbers
	r, _ := regexp.Compile("[^a-z]")

	s = strings.ToLower(s)
	s = r.ReplaceAllString(s, " ")

	tmp := strings.Fields(s)
	st := NewStemmer()
	for i := range tmp {

		if len(tmp[i]) > 3 && !isStopWord(tmp[i]) {
			stem, _ := st.Stem(tmp[i])
			out = append(out, stem)
		}
	}

	return
}

func contents(filename string) (string, os.Error) {
	f, err := os.Open(filename, os.O_RDONLY, 0)
	if err != nil {
		return "", err
	}
	defer f.Close() // f.Close will run when we're finished.

	var result []byte
	buf := make([]byte, 100)
	for {
		n, err := f.Read(buf[0:])
		result = append(result, buf[0:n]...) // append is discussed later.
		if err != nil {
			if err == os.EOF {
				break
			}
			return "", err // f will be closed if we return here.
		}
	}
	return string(result), nil // f will be closed if we return here.
}
func contentsB(filename string) ([]byte, os.Error) {
	f, err := os.Open(filename, os.O_RDONLY, 0)
	if err != nil {
		return nil, err
	}
	defer f.Close() // f.Close will run when we're finished.

	var result []byte
	buf := make([]byte, 100)
	for {
		n, err := f.Read(buf[0:])
		result = append(result, buf[0:n]...) // append is discussed later.
		if err != nil {
			if err == os.EOF {
				break
			}
			return nil, err // f will be closed if we return here.
		}
	}
	return result, nil // f will be closed if we return here.
}

func write_to(filename string, buffer []byte) (err os.Error) {

	f, err := os.Open(filename, os.O_CREAT|os.O_RDWR, 0777)
	defer f.Close()
	if err != nil {
		fmt.Printf("Open %s\n", err.String())
		return
	}
	_, err = f.Write(buffer)

	if err != nil {
		fmt.Printf("Write %s\n", err.String())
		return
	}
	return nil
}

func existsQ(filename string) bool {

	if path.Glob(filename) == nil {
		return false
	}

	return true

}
