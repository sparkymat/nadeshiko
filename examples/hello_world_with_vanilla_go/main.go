package main

import (
	"io"
	"net/http"

	"github.com/kirillrdy/nadeshiko"
	"github.com/kirillrdy/nadeshiko/html"
	"github.com/kirillrdy/nadeshiko/jquery"
	"github.com/sparkymat/webdsl/css"
)

func handler(document *nadeshiko.Document) {
	document.JQuery(css.Body).HTML(html.H1().Text("Hello World !!!"))
}

func httpHandler(response http.ResponseWriter, request *http.Request) {

	page := html.Html().Children(
		html.Head().Children(
			nadeshiko.Scripts()...,
		),
		html.Body(),
	)
	io.WriteString(response, page.String())

}

func main() {
	http.HandleFunc("/", httpHandler)

	http.HandleFunc(jquery.WebPath, jquery.FileHandler)
	http.HandleFunc(nadeshiko.JsWebPath, nadeshiko.JsHandler)

	//XXX for now just have path + .websocket pattern
	http.Handle("/.websocket", nadeshiko.Handler(handler))
	http.ListenAndServe(":3000", nil)
}
