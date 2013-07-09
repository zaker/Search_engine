package qproc

import (
	"../docmap"
	"../invertmap"
	"../querrymap"
	"../utils"
	"math"
	"sort"
	// 	"fmt"
	"strconv"
	"sync"
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

func weigh_term(dm docmap.DocMap, id, num, tot int) (wght float64) {

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
	//num int
}

func exists(dm docmap.DocMap, rf []invertmap.Reference) (terms Terms) {
	terms = make(Terms)
	termArray := make([]int, 0)
	for _, ref := range rf {
		_, ok := dm[ref.DocNo]
		if ok {
			termArray = append(termArray, ref.DocNo)
		}
	}
	sort.Ints(termArray)
	if len(termArray) > 0 {
		t1 := termArray[0]
		sum := 0
		for _, t2 := range termArray {
			if t1 == t2 {
				sum++
			} else {
				t1 = t2
				terms[t1] = sum
				sum = 0
			}
		}
	}
	return
}

func qProc(dm docmap.DocMap, im invertmap.InvertMap, qs []string) (wT weightTable) {

	wT = make(weightTable)
	var wg sync.WaitGroup
	var lck sync.Mutex
	for i, qsT := range qs {
		wg.Add(1)
		go func(totQ, iT int, qStr string) {
			defer wg.Done()
			term_docs := exists(dm, im[qStr])
			tot := len(term_docs)

			for d, td := range term_docs {
				lck.Lock()
				if len(wT[d]) < totQ {
					wT[d] = make([]float64, totQ)
				}
				lck.Unlock()
				// 			println(qs[i])
				wghtT := weigh_term(dm, d, td, tot)
				lck.Lock()
				wT[d][iT] = wghtT
				lck.Unlock()
			}
		}(len(qs), i, qsT)
	}
	wg.Wait()
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
func (wT weightTable) Mul(wT2 weightTable) {
	for i := range wT2 {

		// 		var sum float64
		// 		_,ok := wT[i]
		// 		if ok{
		for j := range wT2[i] {
			if len(wT[i]) < 1 {
				wT[i] = make([]float64, 1)
			}
			wT[i][0] += 0.8 * wT2[i][j]
		}
		// 		}
		// 		tW = append(tW, tWeigth{doc: i, weight: sum})
	}

	// 	tW.Sort()
	return
}

func QuerryProc(dm docmap.DocMap, im invertmap.InvertMap, qs []string) (outW tWeigths) {

	wT := qProc(dm, im, qs)
	outW = wT.Sum2Slice()

	if len(outW) > 20 {
		outW = outW[:20]
	}
	return
}

func QuerryProcFeedback(dm docmap.DocMap, im invertmap.InvertMap, qs []string) (outW tWeigths) {

	tW := QuerryProc(dm, im, qs)
	dm2 := docmap.NewDocMap()
	for i := range tW {
		// 		println("first feed",tW[i].doc)
		dm2[tW[i].doc] = dm[tW[i].doc]
	}

	wT2 := make(weightTable)
	for i := range tW {
		tempW := qProc(dm2, im, dm[tW[i].doc].S)

		wT2.Mul(tempW)

	}
	tW2 := wT2.Sum2Slice()
	// 	if len(outW) > 5 {
	// // 		outW = outW[:5]
	// 	}
	for i := range tW2 {
		// 		println(tW[i].weight,tW2[i].weight)
		outW = append(outW, tWeigth{doc: tW2[i].doc, weight: tW[i].weight * tW2[i].weight})
	}
	return
}

func QuerriesProc(dm docmap.DocMap, qm querrymap.QuerryMap, im invertmap.InvertMap) {

	println("querry parser")
	i := 0
	outS := ""

	for j := 0; j < len(qm); {
		_, ok := qm[i]
		if ok {
			// 			println(i, j)
			j++
			// 			outW := QuerryProc(dm, im, qm[i].S)

			outW := QuerryProcFeedback(dm, im, qm[i].S)
			for k := range outW {
				tmp := strconv.Itoa(j) +
					" 0 " + strconv.Itoa(outW[k].doc) +
					" 0 " + strconv.FormatFloat(outW[k].weight, 'f', -1, 64) +
					" testRun\n"
				outS += tmp
				// 				println(tmp)
			}

		}
		i++
	}
	utils.Write_to("trec_eval2", []byte(outS))

}
