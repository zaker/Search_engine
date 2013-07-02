package webserv

import (
	"../docmap"
	"../invertmap"
	"../qproc"
	"../utils"
	"flag"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"strconv"
)

var addr = flag.String("addr", ":1400", "http service address") // as in cran 1400
var fmap = template.FormatterMap{
	"html":     template.HTMLFormatter,
	"url+html": UrlHtmlFormatter,
}
var templ = template.MustParse(templateStr, fmap)

var q_in = make(chan string)
var a_out = make(chan string)

func handleQuerry(dm docmap.DocMap, im invertmap.InvertMap) {
	for {
		s := <-q_in
		qs := utils.CleanS(s)
		res := qproc.QuerryProcFeedback(dm, im, qs)
		// 	println(s)

		out := ""
		for i := range res {

			tmp := strconv.Itoa(i+1) + ". doc[" + strconv.Itoa(res[i].doc) + "]  " + dm[res[i].doc].W + "<br><br>"
			out += tmp
			// 		       println(tmp)
		}
		a_out <- out
	}

}

func wServer(dm docmap.DocMap, im InvertMap) {
	println("webserver")
	flag.Parse()
	// 	fmt.Println("%d\n",7)
	go handleQuerry(dm, im)
	http.Handle("/", http.HandlerFunc(Search))
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Exit("ListenAndServe:", err)
	}
}

func QR(w http.ResponseWriter, req *http.Request) {
	templ.Execute(req.FormValue("s"), w)
}

func Search(w http.ResponseWriter, req *http.Request) {
	// 	println(req.FormValue("s"))
	s := req.FormValue("s")
	q_in <- s
	templ.Execute(s, w)
	res := <-a_out
	fmt.Fprintln(w, "Results here <br>"+res)
}

func UrlHtmlFormatter(w io.Writer, fmt string, v ...interface{}) {
	template.HTMLEscape(w, []byte(http.URLEscape(v[0].(string))))
	// 	fmt.Fprintln(w, "dsdsd\nasdfasdf\tasdfasdf\"tile\"")
}

const templateStr = `
<html>
<head>
<title>Cran set Searcher</title>
</head>
<body>
{.section @}
<img src="http://chart.apis.google.com/chart?chs=300x300&cht=qr&choe=UTF-8&chl={@|url+html}"
/>
<br>
{@|html}
<br>
<br>
{.end}
<form action="/" name=f method="GET"><input maxLength=1024 size=70
name=s value="" title="Text to Search"><input type=submit
value="search" name=search>
</form>
</body>
</html>
`
