package main

import (
	"math"
	"sort"
// 	"strconv"
)

type tWeigths []tWeigth

type tWeigth struct{
	doc int
	weight float64
	
}


func (tw tWeigths)Len() int{
	
	return len(tw)
}
func (tw tWeigths)Less(i,j int) bool{
	
	return tw[i].weight < tw[j].weight
}
func (tw tWeigths)Swap(i,j int){
	
	tw[i],tw[j] = tw[j],tw[i]
}
func (tw tWeigths)Sort(){
	
	sort.Sort(tw)
}


// type tWeigth interface{
// 	doc int
// 	weight float64
// 	
// }
type docWeight map[int]float64


func weigh_term(dm DocMap,term string ,id ,num , tot int )(wght float64){
	
	
	ldoc := len(dm[id].S)
	if ldoc <= 0{
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
	
	tf := float64(num)/float64(ldoc)
// 	println(num,ldoc,tf)
	tot_doc := len(dm)
// 	println("tot docs:",tot_doc,term)
	l := math.Log(float64(tot_doc)/float64(tot))
	
	wght = l*float64(tf)
// 	println(wght)
	return
	
}




func QuerryProc(dm DocMap,qm QuerryMap,im InvertMap){
	qs := qm[1].S
	dw := make(docWeight,1)
	for i:= range qs {
		term_docs := im[qs[i]]
// 		println(qs[i])
		tot := len(term_docs)
// 		var w float64
		for j := range term_docs {
			d,n := term_docs.Get(j)
// 			println(i,j,qs[i],term_docs[j])
// 			println(d,n,tot)
// 			weigh_querry(dm,qm,im,term_docs[j],1)
			dw[d] += weigh_term(dm,qs[i],d,n, tot)
// 			break
		}
// 		print(w)
	}
	tw := make(tWeigths,1)
	for i := range dw {

		tw = append(tw,tWeigth{doc:i,weight:dw[i]})
		
	}
	tw.Sort()
	tw = tw[len(tw)-20:]
	for i := range tw {
// 		r  := len(tw) - i -1
		println(tw[i].doc,tw[i].weight, qm[1].W)
		println(dm[tw[i].doc].W)
	}
}