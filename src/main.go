package main

import (
	"./docmap"
	"./invertmap"
	"./qproc"
	"./querrymap"
	"flag"
	"github.com/davecheney/profile"
)

var startQuerry = flag.Bool("q", false, "Going through querries")

// var startWebserver = flag.Bool("w", true, "Starting webserver")

func main() {
	defer profile.Start(profile.CPUProfile).Stop()
	flag.Parse()
	dm := docmap.NewDocMap()
	err := dm.DocReader()

	if err != nil {
		println(err)
		return
	}
	qm := querrymap.NewQuerryMap()

	qm.QuerryReader()

	im := invertmap.NewInvertMap()

	err = im.DocMToInM(dm)

	if err != nil {
		println(err)
	}

	if *startQuerry {
		qproc.QuerriesProc(dm, qm, im)
	} else {
		println("No webservice")
	}

}
