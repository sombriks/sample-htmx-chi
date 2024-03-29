package controllers

import (
	"net/http"
	"sample-htmx-chi/app/config"
)

func IndexHandler(writer http.ResponseWriter, request *http.Request) {
	config.Render(writer, "app/templates/index.liquid", nil)
}

func ListHandler(writer http.ResponseWriter, request *http.Request) {}

func InsertHandler(writer http.ResponseWriter, request *http.Request) {}

func FindHandler(writer http.ResponseWriter, request *http.Request) {}

func UpdateHandler(writer http.ResponseWriter, request *http.Request) {}

func DeleteHandler(writer http.ResponseWriter, request *http.Request) {}
