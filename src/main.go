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


func (dm DocMap) AddTo(s string) (err os.Error) {

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
func DocReader() (err os.Error) {

	Docs, err := contents("../data/cran.all.1400")

	if err != nil {
		return
	}

	DocStrings := strings.Split(Docs, ".I", -1)

	// 	println(DocStrings[3])
	dm := make(DocMap)
	for i := range DocStrings {
		// 	for i = 0; i < 2; i++{
		dm.AddTo(DocStrings[i])

	}

	// 	println("|",dm[52].I,"|")
	// 	println("|",dm[52].T,"|")
	// 	println("|",dm[52].A,"|")
	// 	println("|",dm[52].B,"|")
	println("|", dm[52].W, "|")

	for i := range dm {
		// 	for i := 0 ; i < len(dm);i++{
		if len(dm[i].A) < 4 {
			println(dm[i].A, dm[i].I)
		}
	}

	return nil
}


func main() {

	DocReader()
	// 	println(i)
}
