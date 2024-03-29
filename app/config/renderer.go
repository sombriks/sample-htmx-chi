package config

import (
	"github.com/osteele/liquid"
	"log"
	"net/http"
	"os"
)

var engine = liquid.NewEngine()

func Render(writer http.ResponseWriter, tplPath string, data map[string]interface{}) {
	content, err := os.ReadFile(tplPath)
	if err != nil {
		log.Printf("ERROR: %v\n", err)
		writer.WriteHeader(500)
		writer.Write([]byte(err.Error()))
		return
	}
	out, err := engine.ParseAndRender(content, data)
	if err != nil {
		log.Printf("ERROR: %v\n", err)
		writer.WriteHeader(500)
		writer.Write([]byte(err.Error()))
	} else {
		writer.Write(out)
	}
}
