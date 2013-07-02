package querrymap

import (
	// 	"fmt"
	"../utils"
	"errors"
	"strconv"
	"strings"
)

type Querry struct {
	I int
	W string
	S []string
}

type QuerryMap map[int]Querry

func (qm QuerryMap) Print(i int) {

	println(qm[i].I)
	println(qm[i].W)
	for j := range qm[i].S {

		println(qm[i].S[j])
	}

}

func sToQuerry(s string) (*Querry, error) {

	tmp := strings.Split(s, ".W")

	tmp[0] = strings.TrimSpace(tmp[0])

	I, err := strconv.Atoi(tmp[0])
	W := strings.TrimSpace(tmp[1])

	if I == 0 {
		return nil, errors.New("Empty doc")
	}
	if err != nil {
		return nil, err
	}

	S := utils.CleanS(W)
	// 	println(W)
	// 	qm.Print(4)
	// 	S := make([]string,2)

	return &Querry{I: I, W: W, S: S}, nil
}

func NewQuerryMap() QuerryMap {

	return make(QuerryMap)
}

func (qm QuerryMap) addTo(s string) (err error) {

	if s == "" {
		return errors.New("Empty string")
	}
	d, err := sToQuerry(s)

	if err != nil {
		return err
	}

	qm[d.I] = *d

	return nil
}
func (qm QuerryMap) QuerryReader() (err error) {

	Querries, err := utils.Contents("../data/cran.qry")

	if err != nil {
		return
	}

	if qm == nil {

		return errors.New("QuerryMap not initialized")
	}

	QuerryStrings := strings.Split(Querries, ".I")

	for i := range QuerryStrings {
		// 	for i := 0; i < 2; i++{
		// 		println("w",i,QuerryStrings[i])
		qm.addTo(QuerryStrings[i])

	}

	// 	for i := range qm {
	// 		// 		println("r",i)
	// 		// 			for i := 0 ; i < len(dm);i++{
	// 		println(qm[i].I, qm[i].W)
	// 	}
	return nil
}
