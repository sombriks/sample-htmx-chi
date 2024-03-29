package controllers

import (
	"log"
	"net/http"
	"os"
)

// content - cache js content so we rad it just once
var content []byte = nil

func HtmxHandler(writer http.ResponseWriter, _ *http.Request) {
	if content == nil {
		var err error
		content, err = os.ReadFile("app/assets/htmx-1.9.11.min.js")
		if err != nil {
			log.Fatal(err)
		}
	}
	writer.Write(content)
}
