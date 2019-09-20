package main

import (
	"github.com/erikdubbelboer/fasthttp"
	"google.golang.org/appengine"
)

func main() {
	fasthttp.ListenAndServe(":8080", Router().InitRouter().HandleRequest)
	appengine.Main()
}
