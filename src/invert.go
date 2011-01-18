package main


import (
// 	"strings"
	"os"
	"bytes"
	"gob"
	"fmt"
// 	"regexp"
)


type InvertMap map[string][]int


func NewInvertMap() InvertMap {
	return make(InvertMap)
}


func (im InvertMap) AddTo(doc string, index int) (err os.Error) {

	words := cleanS(doc)

	for i := range words {
		im[words[i]] = append(im[words[i]], index)
		
	}
// 		// 		println(words[i])
// 		if im[words[i]] == nil {
// 			im[words[i]] = append(im[words[i]], index)
// 		} else {
// 			chk := true
// // 			for j := range im[words[i]] {
// 
// // 				if im[words[i]][j] == index {
// // 					chk = false
// // 				}
// // 			}
// 
// 			if chk {
// 				im[words[i]] = append(im[words[i]], index)
// 			}
// 
// 		}
// 
// 	}

	return nil
}

func (im InvertMap) LenDocs(key string) (l int) {

	ar := im[key]

	return len(ar)

}


func (im *InvertMap) Load(im_filename string) (err os.Error) {

	str, err := contents(im_filename)

	b := bytes.NewBufferString(str)
	dec := gob.NewDecoder(b)

	dec.Decode(im)

	return nil
}

func (im *InvertMap) Save(im_filename string) (err os.Error) {

	b := new(bytes.Buffer)

	enc := gob.NewEncoder(b)

	err = enc.Encode(im)

	if err != nil {
		fmt.Printf("encode %s\n", err.String())
		return
	}

	err = write_to(im_filename, b.Bytes())

	if err != nil {
		fmt.Printf("write %s\n", err.String())
		return
	}

	return nil
}

// Takes a Doc Map and adds to the inverted map
func (im *InvertMap) DocMToInM(dm DocMap) (err os.Error) {

	im_filename := "../tmp/im_file"

	if existsQ(im_filename) {
		println("exist")
		return im.Load(im_filename)
	}

	for i := range dm {

		err = im.AddTo(dm[i].W, dm[i].I)
	}

	if err != nil {
		fmt.Printf("%s\n", err.String())
		return
	}

	return im.Save(im_filename)

}
