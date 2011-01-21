package main

import (
	// 	"sort"
	"os"
	"flag"
)


var startQuerry = flag.Bool("q", false, "Going through querries")
// var startWebserver = flag.Bool("w", true, "Starting webserver")

func main() {
	flag.Parse()
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

	if *startQuerry {
		QuerriesProc(dm, qm, im)
	} else {
		wServer(dm, im)
	}

}
