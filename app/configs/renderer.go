package configs

import (
	"github.com/osteele/liquid"
	"log"
	"net/http"
	"os"
	"path"
)

type RenderConfig struct {
	engine       *liquid.Engine
	TemplatePath string
	AssetsPath   string
}

type Config struct {
	RenderConfig
	DbConfig
}

func NewConfig() *Config {
	var config Config
	config.engine = liquid.NewEngine()

	var set bool
	config.TemplatePath, set = os.LookupEnv("TEMPLATE_PATH")
	if !set {
		config.TemplatePath = "app/templates"
	}
	config.AssetsPath, set = os.LookupEnv("ASSETS_PATH")
	if !set {
		config.AssetsPath = "app/assets"
	}

	config.Open()
	config.InitSchema()

	return &config
}

// Render - find, process and write liquid template TODO find a Chi plugin for that
func (config *Config) Render(writer http.ResponseWriter, tplPath string, data map[string]interface{}) {
	content, err := os.ReadFile(path.Join(config.TemplatePath, tplPath))
	if err != nil {
		log.Printf("ERROR: %v\n", err)
		writer.WriteHeader(500)
		writer.Write([]byte(err.Error()))
		return
	}
	out, err := config.engine.ParseAndRender(content, data)
	if err != nil {
		log.Printf("ERROR: %v\n", err)
		writer.WriteHeader(500)
		writer.Write([]byte(err.Error()))
	} else {
		writer.Write(out)
	}
}
