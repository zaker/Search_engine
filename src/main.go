package main

import "sort"

func main() {

	dm := NewDocMap()
	dm.DocReader()

	qm := NewQuerryMap()

	qm.QuerryReader()

	im := NewInvertMap()

	for i := range dm {

		im.AddTo(dm[i].W, dm[i].I)
	}
	println("done indexing!")
	
	var words []string;
	for i := range im {
// 		key, _ := im[i]
// 		print(i, " : ")
		
		words = append(words,i)
// 		for j := range key {
// 			print(key[j], " ")
// 		}
// 		println()
	}
	
	println(len(im), len(words))
	sa := sort.StringArray(words)
// 	sort_words := sort.SortStrings(sa)
	// 	println(i)
}
