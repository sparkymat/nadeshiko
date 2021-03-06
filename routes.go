/*
Package nadeshiko routes
Built in Router and Server

Nadeshiko comes with own very simple router, and own server.

Heading

More data
*/
package nadeshiko

import (
	"io"
	"net/http"

	"github.com/kirillrdy/nadeshiko/html"

	"code.google.com/p/go.net/websocket"
)

const get = "GET"
const post = "POST"

type routes []route

func (routes *routes) get(path string, handler func(http.ResponseWriter, *http.Request)) {
	route := route{path, get, handler}
	*routes = append(*routes, route)
}

func (routes *routes) post(path string, handler func(http.ResponseWriter, *http.Request)) {
	route := route{path, post, handler}
	*routes = append(*routes, route)
}

func (routes *routes) nadeshiko(path string, handler func(*Document)) {
	httpHandler := func(response http.ResponseWriter, request *http.Request) {
		page := html.Html().Children(
			html.Head().Children(
				Scripts()...,
			),
			html.Body(),
		)
		io.WriteString(response, page.String())
	}

	customeWebsocketServer := websocketServer(handler)

	//TODO fix .websocket path
	routes.get(path+".websocket", websocket.Handler(customeWebsocketServer).ServeHTTP)
	routes.get(path, httpHandler)
}

// Handler returns returns http.Handler for given nadeshiko handler so that it can be used with router of your choice
func Handler(nadeshikoHandler func(*Document)) http.Handler {
	return websocket.Handler(websocketServer(nadeshikoHandler))
}

func (routes *routes) webSocket(path string, handler func(*Document)) {
	customeWebsocketServer := websocketServer(handler)
	routes.get(path+".websocket", websocket.Handler(customeWebsocketServer).ServeHTTP)
}
