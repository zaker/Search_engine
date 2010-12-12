package main


import (
	// 	"fmt"
	"os"
	"strings"
	"strconv"
)

type Querry struct {
	I int
	W string
}

type QuerryMap map[int]Querry




func sToQuerry(s string) (*Querry, os.Error) {

	tmp := strings.Split(s, ".W", -1)

	tmp[0] = strings.TrimSpace(tmp[0])

	I, err := strconv.Atoi(tmp[0])
	W := strings.TrimSpace(tmp[1])

	if I == 0 {
		return nil, os.NewError("Empty doc")
	}
	if err != nil {
		return nil, err
	}
	

	return &Querry{I: I, W: W}, nil
}

func NewQuerryMap()(QuerryMap){
	
	return make(QuerryMap)
}

func (qm QuerryMap) addTo(s string) (err os.Error) {

	if s == "" {
		return os.NewError("Empty string")
	}
	d, err := sToQuerry(s)

	if err != nil {
		return err
	}

	qm[d.I] = *d

	return nil
}
func (qm QuerryMap)QuerryReader() ( err os.Error) {

	Querries, err := contents("../data/cran.qry")

	if err != nil {
		return 
	}
	
	if qm == nil{
		
		return os.NewError("QuerryMap not initialized")
	}

	QuerryStrings := strings.Split(Querries, ".I", -1)

	
	for i := range QuerryStrings {
// 	for i := 0; i < 2; i++{
// 		println("w",i,QuerryStrings[i])
		qm.addTo(QuerryStrings[i])

	}

	for i := range qm {
// 		println("r",i)
// 			for i := 0 ; i < len(dm);i++{
		println(qm[i].I, qm[i].W)
	}
	return nil
}
