package main

import (
// 	"sort"
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


// 	QuerriesProc(dm, qm, im)
	wServer(dm,im)

}
