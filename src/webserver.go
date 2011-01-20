package main

import (
    "flag"
    "http"
    "io"
    "log"
    "template"
)

var addr = flag.String("addr", ":1400", "http service address") // as in cran 1400
var fmap = template.FormatterMap{
    "html": template.HTMLFormatter,
    "url+html": UrlHtmlFormatter,
    "str": UrlHtmlFormatter2,
}
var templ = template.MustParse(templateStr, fmap)

func wServer(dm DocMap,im InvertMap) {
    flag.Parse()
    http.Handle("/", http.HandlerFunc(Search))
    err := http.ListenAndServe(*addr, nil)
    if err != nil {
        log.Exit("ListenAndServe:", err)
    }
}

func QR(w http.ResponseWriter, req *http.Request) {
    templ.Execute(req.FormValue("s"), w)
}

func Search(w http.ResponseWriter, req *http.Request){
	templ.Execute(req.FormValue("s"), w)
}

func UrlHtmlFormatter(w io.Writer, fmt string, v ...interface{}) {
    template.HTMLEscape(w, []byte(http.URLEscape(v[0].(string))))
}
func UrlHtmlFormatter2(w io.Writer, fmt string, v ...interface{}) {
    template.HTMLEscape(w, []byte("hello <br>"))
    template.StringFormatter(w, "world",v...)
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
{@|str}
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