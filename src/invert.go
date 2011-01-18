package main


import (
	"strings"
	"strconv"
	"os"
	"bytes"
	"gob"
	"fmt"
	// 	"regexp"
)


type InvertMap map[string]Inv


type Inv []string

func (in Inv) Get(i int) (d, n int) {

	spl := strings.Split(in[i], ":", -1)
	d, _ = strconv.Atoi(spl[0])
	n, _ = strconv.Atoi(spl[1])
	return
}
func (in Inv) Set(i, d, n int) {

	// 	spl := strings.Split(in[i],":",-1)
	// 	d,_ = strconv.Atoi(spl[0])
	// 	n,_ = strconv.Atoi(spl[1])
	in[i] = strconv.Itoa(d) + ":" + strconv.Itoa(n)
	return
}

func NewInvertMap() InvertMap {
	return make(InvertMap)
}


// func (im InvertMap) AddTo(doc string, index int) (err os.Error) {
// 
// 	words := cleanS(doc)
// 
// 	for i := range words {
// 		im[words[i]] = append(im[words[i]], index)
// 		
// 	}
// }
func (im InvertMap) AddStemTo(doc []string, index int) (err os.Error) {

	// 	words := cleanS(doc)

	for i := range doc {
		// 		im[doc[i]] = append(im[doc[i]], index)
		// 		
		// 	}
		// 	return
		// }

		// 		println(words[i])

		// 		if this is the first time the word is added
		if im[doc[i]] == nil {
			// 			in := new(invert)
			// 			in.doc = i
			// 			println("test")
			in := strconv.Itoa(index) + ":" + strconv.Itoa(1)
			im[doc[i]] = append(im[doc[i]], in)
		} else {
			chk := true
			for j := range im[doc[i]] {
				// 				if there already exist an inverted doc, just add to num
				// 				println(im[doc[i]][j])
				// 				spl := strings.Split(im[doc[i]][j],":",-1)
				// 				d,_ := strconv.Atoi(spl[0])
				// 				n,_ := strconv.Atoi(spl[1])
				d, n := im[doc[i]].Get(j)
				if d == index {
					n++
					im[doc[i]][j] = strconv.Itoa(d) + ":" + strconv.Itoa(n)
					chk = false
				}

			}

			if chk {

				in := strconv.Itoa(index) + ":" + strconv.Itoa(1)
				im[doc[i]] = append(im[doc[i]], in)
				// 			println("test")
				// 				im[doc[i]] = append(im[doc[i]], index)
			}

		}

	}
	return nil
}

func (im InvertMap) LenDocs(key string) (l int) {

	// 	ar := im[key]

	return len(im[key])

}


func (im InvertMap) Load(im_filename string) (err os.Error) {

	str, err := contents(im_filename)

	b := bytes.NewBufferString(str)
	dec := gob.NewDecoder(b)

	dec.Decode(&im)

	return nil
}

func (im InvertMap) Save(im_filename string) (err os.Error) {

	b := new(bytes.Buffer)

	enc := gob.NewEncoder(b)
	err = enc.Encode(&im)

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
	j := 0
	for i := range dm {
		j++
		println(j)
		err = im.AddStemTo(dm[i].S, dm[i].I)
	}

	if err != nil {
		fmt.Printf("%s\n", err.String())
		return
	}

	return im.Save(im_filename)

}
