package main


import (
	"strings"
	"os"
)


type InvertMap map[string][]int


func NewInvertMap() InvertMap {
	return make(InvertMap)
}

func cleanS(s string)( out []string){
	
// 	remove symbols and numbers
	
	s  = strings.Replace(s,"."," ",-1)
	s  = strings.Replace(s,"'"," ",-1)
	s  = strings.Replace(s,","," ",-1)
	s  = strings.Replace(s,"/"," ",-1)
	s  = strings.Replace(s,"-"," ",-1)
	s  = strings.Replace(s,"("," ",-1)
	s  = strings.Replace(s,")"," ",-1)
	s  = strings.Replace(s,"0"," ",-1)
	s  = strings.Replace(s,"1"," ",-1)
	s  = strings.Replace(s,"2"," ",-1)
	s  = strings.Replace(s,"3"," ",-1)
	s  = strings.Replace(s,"4"," ",-1)
	s  = strings.Replace(s,"5"," ",-1)
	s  = strings.Replace(s,"6"," ",-1)
	s  = strings.Replace(s,"7"," ",-1)
	s  = strings.Replace(s,"8"," ",-1)
	s  = strings.Replace(s,"9"," ",-1)
	
	tmp := strings.Fields(s)
	
	for i := range tmp{
		
		if len(tmp[i]) > 1 {
			out = append(out,tmp[i])
		}
	}
	
	return
}
	
	
	

func (im InvertMap) AddTo(doc string, index int) (err os.Error) {

	// TODO: make better stemmer
	words := cleanS(doc)

	for i := range words {

		// 		println(words[i])
		if im[words[i]] == nil {
			im[words[i]] = append(im[words[i]], index)
		} else {
			chk := true
			for j := range im[words[i]] {

				if im[words[i]][j] == index {
					chk = false
				}
			}

			if chk {
				im[words[i]] = append(im[words[i]], index)
			}

		}

	}

	// 	for i := range im["no"]{
	// 		
	// 		print(im["no"][i])
	// 	}
	// 	println()

	return nil
}
