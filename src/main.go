package main

import (
	"sort"
	"os"
)
 

func main() {

	var err os.Error

	dm := NewDocMap()
	dm.DocReader()

	if err != nil {
		return
	}
	qm := NewQuerryMap()

	qm.QuerryReader()

	im := NewInvertMap()

	err = im.DocMToInM(dm)

	if err != nil {
		println(err.String())
	}

	var words []string
	for i := range im {
		// 		key, _ := im[i]
		// 		print(i, " : ")
		// 
		words = append(words, i)
		// 		for j := range key {
		// 		print(im.LenDocs(i))
		// 		}
		// 		println()
	}
	// 

	// 	sa := sort.StringArray(words[0:])
	sort.SortStrings(words)

// 	for i := range words {
// 		println(words[i])
// 	}
	println(len(im), len(words))
// 	println(i)
	println(qm[4].W)
// 	st map[string][]int
	for i := range qm[4].S{
		q := qm[4].S[i]
// 		println()

		a := im[q]
		
		for j:= range a{
			print(a[j],",")
		}
		println()
		
	}
	
	qm.Print(3)
		
}
