package main

import "os"


func main() {

	
	var err os.Error
	
	dm := NewDocMap()
	dm.DocReader()

	if err != nil{
		return
	}
	qm := NewQuerryMap()

	qm.QuerryReader()

	im := NewInvertMap()

// 	for i := range dm {
// 
// 		im.AddTo(dm[i].W, dm[i].I)
// 	}
	
	err = im.DocMToInM(dm)
	
	
	println("done indexing! ", err.String())

// 	var words []string
// 	for i := range im {
// 		// 		key, _ := im[i]
// 		// 		print(i, " : ")
// 
// 		words = append(words, i)
// 		// 		for j := range key {
// 		// 			print(key[j], " ")
// 		// 		}
// 		// 		println()
// 	}
// 
// 	println(len(im), len(words))
	
	
	
// 	sa := sort.StringArray(words)
	// 	sort_words := sort.SortStrings(sa)
	// 	println(i)
}
