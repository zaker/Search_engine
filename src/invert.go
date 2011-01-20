package main


import (
	"os"
	"bytes"
	"gob"
	"fmt"
)


type Reference struct {
	DocNo  int
	WordNo int
}

type InvertMap map[string][]Reference

// Creates a new inverted index
func NewInvertMap() InvertMap {
	return make(InvertMap)
}


func (im InvertMap) AddStemTo(doc []string, DocNo int) (err os.Error) {

	// 	words := cleanS(doc)
	for wordNo, word := range doc {
		ref, ok := im[word]
		if !ok {
			ref = nil
		}
		im[word] = append(ref, Reference{DocNo, wordNo})
	}
	return
}
// Deletes map entry
func (im InvertMap) DeleteStem(s string) (err os.Error) {

	im[s] = nil, false
	return
}

// Returns size of map
func (im InvertMap) LenDocs(key string) (l int) {

	return len(im[key])

}

//  Loads index from file
func (im InvertMap) Load(im_filename string) (err os.Error) {

	str, err := contents(im_filename)

	b := bytes.NewBufferString(str)
	dec := gob.NewDecoder(b)

	dec.Decode(&im)

	return nil
}
// Stores the inverted index to file
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
