package main

import (
	// 	"fmt"
	"os"
	"strings"
	"strconv"
)

type Doc struct {
	I int
	T string
	A string
	B string
	W string
}

type DocMap map[int]Doc


func sToDoc(s string) (*Doc, os.Error) {

	s = strings.Replace(s, "\n", " ", -1)
	tmp := strings.Split(s, ".T", -1)

	tmp[0] = strings.TrimSpace(tmp[0])
	I, err := strconv.Atoi(tmp[0])

	if I == 0 {
		return nil, os.NewError("Empty doc")
	}
	if err != nil {
		return nil, err
	}

	tmp = strings.Split(tmp[1], ".A", -1)

	T := strings.TrimSpace(tmp[0])

	tmp = strings.Split(tmp[1], ".B", -1)

	A := strings.TrimSpace(tmp[0])

	tmp = strings.Split(tmp[1], ".W", -1)

	B := strings.TrimSpace(tmp[0])

	W := strings.TrimSpace(tmp[1])

	return &Doc{I: I, T: T, A: A, B: B, W: W}, nil
}


func NewDocMap() DocMap {

	return make(DocMap)
}

func (dm DocMap) addTo(s string) (err os.Error) {

	if s == "" {
		return os.NewError("Empty string")
	}
	d, err := sToDoc(s)

	if err != nil {
		return err
	}

	dm[d.I] = *d

	return nil
}
func (dm DocMap) DocReader() (err os.Error) {

	Docs, err := contents("../data/cran.all.1400")

	if err != nil {
		return
	}

	if dm == nil {

		return os.NewError("DocMap not initialized")
	}

	DocStrings := strings.Split(Docs, ".I", -1)

	// 	println(DocStrings[3])

	for i := range DocStrings {
		// 	for i = 0; i < 2; i++{
		dm.addTo(DocStrings[i])

	}

	return nil
}
