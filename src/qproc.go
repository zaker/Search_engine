package main

import (
	"math"
	"sort"
// 	"strconv"
)

type terms map[string]termW

type termW struct{
	
	
	doc int
	num int
	weight float64
	
}

func unique(in []int)(out []int){
	sa := sort.IntArray(in)
	

	
	for i := range sa {
		chk := true
		for j := range out {
			if sa[i] == out[j] {
// 				t++
				chk = false
				break
				
			}
		}
		if chk{
			out = append(out,sa[i])
			
// 			tot++
		}
		
	}
	return
}
func totalUnique(im InvertMap,term string)(tot int){
	
// 	tot := 0
// 	b := len(im[term]) 
	tot = len(im[term])
	return 
	
}

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

// func weigh_querry(dm DocMap,qm QuerryMap,im InvertMap,id,iq int )(wght float64){
// 	
// 	
// 	for i := range qm[iq].S{
// 		tot := totalUnique(im,qm[iq].S[i])
// // 		if tot > 0 {
// // 			println(tot,qm[iq].I,qm[iq].S[i])
// // 		}
// 		wght += weigh_term(dm,qm[iq].S[i],tot,id)
// 	}
// 	if wght > 0.0 {
// 		println( wght)
// 	}
// 	
// 	return
// 	
// }



func QuerryProc(dm DocMap,qm QuerryMap,im InvertMap){
	qs := qm[1].S
	
	for i:= range qs {
		term_docs := im[qs[i]]
		println(qs[i])
		tot := len(term_docs)
		var w float64
		for j := range term_docs {
			d,n := term_docs.Get(j)
// 			println(i,j,qs[i],term_docs[j])
// 			println(d,n,tot)
// 			weigh_querry(dm,qm,im,term_docs[j],1)
			w += weigh_term(dm,qs[i],d,n, tot)
// 			break
		}
		print(w)
	}
}