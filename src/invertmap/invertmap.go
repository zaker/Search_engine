package invertmap

import (
	"../docmap"
	"../utils"
	"bytes"
	"encoding/gob"
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

func (im InvertMap) AddStemTo(doc []string, DocNo int) (err error) {

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
func (im InvertMap) DeleteStem(s string) (err error) {

	im[s] = nil
	return
}

// Returns size of map
func (im InvertMap) LenDocs(key string) (l int) {

	return len(im[key])

}

//  Loads index from file
func (im InvertMap) Load(im_filename string) (err error) {

	str, err := utils.Contents(im_filename)

	b := bytes.NewBufferString(str)
	dec := gob.NewDecoder(b)

	dec.Decode(&im)

	return nil
}

// Stores the inverted index to file
func (im InvertMap) Save(im_filename string) (err error) {

	b := new(bytes.Buffer)

	enc := gob.NewEncoder(b)
	err = enc.Encode(&im)

	if err != nil {
		fmt.Printf("encode %s\n", err)
		return
	}

	err = utils.Write_to(im_filename, b.Bytes())

	if err != nil {
		fmt.Printf("write %s\n", err)
		return
	}

	return nil
}

// Takes a Doc Map and adds to the inverted map
func (im *InvertMap) DocMToInM(dm docmap.DocMap) (err error) {

	im_filename := "../tmp/im_file"

	if b, err := utils.ExistsQ(im_filename); b {
		println("exist")
		return im.Load(im_filename)
	} else {
		if err != nil {
			return err
		}
	}
	i := 0
	for _, j := range dm {

		if i%30 == 0 {
			fmt.Printf("\n")
		}
		fmt.Printf(".")
		err = im.AddStemTo(j.S, j.I)
		i++
	}
	fmt.Printf("\n")
	if err != nil {
		fmt.Println(err)
		return
	}

	return im.Save(im_filename)

}
