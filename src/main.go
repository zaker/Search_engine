package main

import (
	"fmt"
	"os"
)

type Doc struct{
	
	I int
	T string
	A string
	B string
	W string
	
}

type DocMap map[int] Doc


func DocReader() (i int64, err os.Error) {

	document_file, err := os.Open("../data/cran.all.1400", os.O_RDONLY, 0666)
	defer document_file.Close()

	dm := make(DocMap,1)
	
	d := new(Doc)
	
	d.I = 1
	d.T = "title"
	d.A = "author"
	d.B = "book"
	d.W = "words"
	
	
	dm[1]= *d
	
	file_info, err := document_file.Stat()

	fmt.Printf("The File %s is of length %d\n", file_info.Name, file_info.Size)

	i = file_info.Size
	println(dm[1].W)
	
	return i, err
}


func main() {

	i, _ := DocReader()
	println(i)
}
