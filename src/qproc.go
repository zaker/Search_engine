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

func (tw tWeigths) Sum(tWs []tWeigths) {

	for i := range tWs {
		if i == 0 {
			copy(tw, tWs[i])
			continue
		}
		for j := range tWs[i] {
			println(tWs[i][j].weight)
			for k := range tw {
				if tWs[i][j].doc == tw[k].doc {
					tw[k].doc += tWs[i][j].doc
				}
			}

		}

	}
	return
}
// type tWeigth interface{
// 	doc int
// 	weight float64
// 	
// }
type docWeight map[int]float64


func weigh_term(dm DocMap, id, num, tot int) (wght float64) {

	ldoc := len(dm[id].S)
	if ldoc <= 0 {
		return 0.0
	}

	tf := float64(num) / float64(ldoc)
	// 	println("in doc",ldoc,"num",num,"tot",tot)
	tot_doc := len(dm)
	l := math.Log(float64(tot_doc) / float64(tot))
	// 	println("tot_doc",tot_doc,"l",l)
	wght = l * float64(tf)
	return

}

type Terms map[int]int

type Term struct {
	doc int
	num int
}

func exists(dm DocMap, rf []Reference) (terms Terms) {
	terms = make(Terms)
	for i := range rf {
		_, ok := dm[rf[i].DocNo]
		if ok {
			terms[rf[i].DocNo]++
		}
	}
	return
}

func qProc(dm DocMap, im InvertMap, qs []string) (wT weightTable) {

	wT = make(weightTable)
	for i := range qs {
		term_docs := exists(dm, im[qs[i]])
		tot := len(term_docs)

		for j := range term_docs {
			d, n := j, term_docs[j]

			if len(wT[d]) < len(qs) {
				wT[d] = make([]float64, len(qs))
			}
			// 			println(qs[i])
			wT[d][i] = weigh_term(dm, d, n, tot)
		}
	}
	return
}
func (wT weightTable) Sum2Slice() (tW tWeigths) {
	for i := range wT {
		var sum float64
		for j := range wT[i] {
			sum += wT[i][j]
		}
		tW = append(tW, tWeigth{doc: i, weight: sum})
	}

	tW.Sort()
	return
}
func (wT weightTable) Add(wT2 weightTable) {
	for i := range wT2 {

		// 		var sum float64
		// 		_,ok := wT[i]
		// 		if ok{
		for j := range wT2[i] {
			if len(wT[i]) < 1 {
				wT[i] = make([]float64, 1)
			}
			wT[i][0] += wT2[i][j]
		}
		// 		}
		// 		tW = append(tW, tWeigth{doc: i, weight: sum})
	}

	// 	tW.Sort()
	return
}

func QuerryProc(dm DocMap, im InvertMap, qs []string) (outW tWeigths) {

	wT := qProc(dm, im, qs)
	tW := wT.Sum2Slice()
	dm2 := NewDocMap()
	if len(tW) > 5 {
		tW = tW[:5]
	}
	for i := range tW {
		// 		println("first feed",tW[i].doc)
		dm2[tW[i].doc] = dm[tW[i].doc]
	}

	wT2 := make(weightTable)
	for i := range tW {
		tempW := qProc(dm2, im, dm[tW[i].doc].S)

		wT2.Add(tempW)

	}
	outW = wT2.Sum2Slice()
	// 	for i := range tW2 {
	// // 		println("rel feed",tW2[i].doc)
	// // 	}
	return
}

func QuerriesProc(dm DocMap, qm QuerryMap, im InvertMap) {

	for i := range qm {
		// 		println("q",i)
		QuerryProc(dm, im, qm[i].S)
	}

}
