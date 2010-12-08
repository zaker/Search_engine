package main

import (
// 	"fmt"
	"os"
	"strings"
	"strconv"
)

type Doc struct{
	
	I int
	T string
	A string
	B string
	W string
	
}

type DocMap map[int] Doc

func contents(filename string) (string, os.Error) {
    f, err := os.Open(filename, os.O_RDONLY, 0)
    if err != nil {
        return "", err
    }
    defer f.Close()  // f.Close will run when we're finished.

    var result []byte
    buf := make([]byte, 100)
    for {
        n, err := f.Read(buf[0:])
        result = append(result, buf[0:n]...) // append is discussed later.
        if err != nil {
            if err == os.EOF {
                break
            }
            return "", err  // f will be closed if we return here.
        }
    }
    return string(result), nil // f will be closed if we return here.
}
func (dm DocMap)AddTo(a int, s string) (err os.Error){
	
	if len(s) == 0{
		return
	}
	d := new(Doc)
	
	tmp := strings.Split(s,".T",-1)
	
// 	for i := range(tmp){
// 		println(i,tmp[i])
// 	}
	tmp[0] = strings.TrimSpace(tmp[0])
	I,_ := strconv.Atoi(tmp[0])
// 	println("I = ",I)
	
	if I == 0 {
		return
	}
	
	tmp = strings.Split(tmp[1],".A",-1)
	
	T := strings.TrimSpace(tmp[0])
	
	tmp = strings.Split(tmp[1],".B",-1)
	
	A := strings.TrimSpace(tmp[0])
	
	tmp = strings.Split(tmp[1],".W",-1)
	
	B := strings.TrimSpace(tmp[0])
	
	W := strings.TrimSpace(tmp[1])
	
// 	println(I,T,A,B,W)
	d.I = I
	d.T = T
	d.A = A
	d.B = B
	d.W = W
	dm[I] = *d
	
	return
}
func DocReader() (i int, err os.Error) {

	Docs , err := contents("../data/cran.all.1400")

	if err != nil {
		return 0, err
	}
	
	DocStrings := strings.Split(Docs,".I",-1)
	
// 	println(DocStrings[3])
	dm := make(DocMap)
	for i = 0; i < len(DocStrings); i++{
// 	for i = 0; i < 2; i++{
		dm.AddTo(i,DocStrings[i])
		
	}
	
	
	println("|",dm[52].I,"|")
	println("|",dm[52].T,"|")
	println("|",dm[52].A,"|")
	println("|",dm[52].B,"|")
	println("|",dm[52].W,"|")
	return i, err
}


func main() {

	DocReader()
// 	println(i)
}
