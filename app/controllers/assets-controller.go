package controllers

import (
	"log"
	"net/http"
	"os"
	"path"
)

// content - cache js content so we rad it just once
var content = map[string][]byte{}

func (controller *TodoController) AssetsHandler(writer http.ResponseWriter, req *http.Request) {
	asset := req.URL.Path
	if content[asset] == nil {
		var err error
		content[asset], err = os.ReadFile(path.Join(controller.config.AssetsPath, asset))
		if err != nil {
			log.Fatal(err)
		}
	}
	writer.Write(content[asset])
}
