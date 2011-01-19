package main

import (
	"math"
	"sort"
	// 	"strconv"
)

type weightTable map[int][]float64

type tWeigths []tWeigth


type tWeigth struct {
	doc    int
	weight float64
}


func (tw tWeigths) Len() int {

	return len(tw)
}
func (tw tWeigths) Less(i, j int) bool {

	return tw[i].weight > tw[j].weight
}
func (tw tWeigths) Swap(i, j int) {

	tw[i], tw[j] = tw[j], tw[i]
}
func (tw tWeigths) Sort() {

	sort.Sort(tw)
}


// type tWeigth interface{
// 	doc int
// 	weight float64
// 	
// }
type docWeight map[int]float64


func weigh_term(dm DocMap, term string, id, num, tot int) (wght float64) {

	ldoc := len(dm[id].S)
	if ldoc <= 0 {
		// 		println(id,dm[id].W)
		return 0.0
	}
	// 	terms := 0
	// 	
	// 	for i := range dm[id].S{
	// 		if term == dm[id].S[i] {
	// 			terms++
	// 		}
	// 		
	// 	}

	tf := float64(num) / float64(ldoc)
	// 	println(num,ldoc,tf)
	tot_doc := len(dm)
	// 	println("tot docs:",tot_doc,term)
	l := math.Log(float64(tot_doc) / float64(tot))

	wght = l * float64(tf)
	// 	println(wght)
	return

}


func QuerryProc(dm DocMap, qm QuerryMap, im InvertMap) {
	qs := qm[1].S
	// 	dw := make(docWeight, 1)
	wT := make(weightTable)
	for i := range qs {
		term_docs := im[qs[i]]
		// 		println(qs[i])
		tot := len(term_docs)

		for j := range term_docs {

			term_docs := im[qs[i]]
			d, n := term_docs.Get(j)
			if len(wT[d]) < len(qs)-1 {
				wT[d] = make([]float64, len(qs))
			}
			wT[d][i] = weigh_term(dm, qs[i], d, n, tot)
			// 			break
		}
	}
	tW := make(tWeigths, 1)
	for i := range wT {
		var sum float64
		for j := range wT[i] {
			sum += wT[i][j]
			// 			print(wT[i][j]," ")
		}
		tW = append(tW, tWeigth{doc: i, weight: sum})
		// 		println(i,sum)
	}

	tW.Sort()
	tW = tW[:20]

	for i := range tW {
		println(i, tW[i].doc, tW[i].weight)
	}

}
