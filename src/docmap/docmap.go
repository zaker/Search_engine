package docmap

import (
	// 	"fmt"
	"../utils"
	"errors"
	"strconv"
	"strings"
)

type Doc struct {
	I int
	T string
	A string
	B string
	W string
	S []string
}

type DocMap map[int]Doc

// Converts the cran1400 file to a map of documents
func sToDoc(s string) (*Doc, error) {

	s = strings.Replace(s, "\n", " ", -1)
	tmp := strings.Split(s, ".T")

	tmp[0] = strings.TrimSpace(tmp[0])
	I, err := strconv.Atoi(tmp[0])

	if I == 0 {
		return nil, errors.New("Empty doc")
	}
	if err != nil {
		return nil, err
	}

	tmp = strings.Split(tmp[1], ".A")

	T := strings.TrimSpace(tmp[0])

	tmp = strings.Split(tmp[1], ".B")

	A := strings.TrimSpace(tmp[0])

	tmp = strings.Split(tmp[1], ".W")

	B := strings.TrimSpace(tmp[0])

	W := strings.TrimSpace(tmp[1])

	S := utils.CleanS(W)

	return &Doc{I: I, T: T, A: A, B: B, W: W, S: S}, nil
}

// Creates a new document map
func NewDocMap() DocMap {

	return make(DocMap)
}

// Adds instances to the document map
func (dm DocMap) addTo(s string) (err error) {

	if s == "" {
		return errors.New("Empty string")
	}
	d, err := sToDoc(s)

	if err != nil {
		return err
	}

	dm[d.I] = *d

	return nil
}

// Reads from the cran1400 file and maps it
func (dm DocMap) DocReader() (err error) {

	Docs, err := utils.Contents("../data/cran.all.1400")

	if err != nil {
		return
	}

	if dm == nil {

		return errors.New("DocMap not initialized")
	}

	DocStrings := strings.Split(Docs, ".I")

	for i := range DocStrings {
		dm.addTo(DocStrings[i])

	}

	return nil
}
