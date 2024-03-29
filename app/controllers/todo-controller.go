package controllers

import (
	"github.com/go-chi/chi/v5"
	"net/http"
	"sample-htmx-chi/app/configs"
	"sample-htmx-chi/app/models"
	"sample-htmx-chi/app/services"
)

type TodoController struct {
	config  *configs.Config
	service *services.TodoService
}

func NewTodoController(config *configs.Config, service *services.TodoService) *TodoController {
	var controller TodoController
	controller.config = config
	controller.service = service
	return &controller
}

func (controller *TodoController) IndexHandler(writer http.ResponseWriter, request *http.Request) {
	q := request.URL.Query().Get("q")
	controller.response(writer, "index.liquid", q)
}

func (controller *TodoController) ListHandler(writer http.ResponseWriter, request *http.Request) {
	q := request.URL.Query().Get("q")
	controller.response(writer, "todos/list.liquid", q)

}

func (controller *TodoController) InsertHandler(writer http.ResponseWriter, request *http.Request) {
	todo, err := models.TodoFromForm(request)
	if err != nil {
		writer.WriteHeader(422)
		writer.Write([]byte(err.Error()))
		return
	}
	controller.service.Insert(todo)
	q := request.URL.Query().Get("q")
	controller.response(writer, "todos/list.liquid", q)
}

func (controller *TodoController) FindHandler(writer http.ResponseWriter, request *http.Request) {
	id := chi.URLParam(request, "id")
	todo, err := controller.service.Find(id)
	if err != nil {
		writer.WriteHeader(500)
		writer.Write([]byte(err.Error()))
		return
	}
	data := map[string]interface{}{"todo": todo}
	controller.config.Render(writer, "todos/detail.liquid", data)
}

func (controller *TodoController) UpdateHandler(writer http.ResponseWriter, request *http.Request) {
	id := chi.URLParam(request, "id")
	todo, err := models.TodoFromForm(request)
	if err != nil {
		writer.WriteHeader(422)
		writer.Write([]byte(err.Error()))
		return
	}
	todo.SetId(id)
	controller.service.Update(todo)
	q := request.URL.Query().Get("q")
	controller.response(writer, "todos/list.liquid", q)
}

func (controller *TodoController) DeleteHandler(writer http.ResponseWriter, request *http.Request) {
	id := chi.URLParam(request, "id")
	controller.service.Delete(id)
	q := request.URL.Query().Get("q")
	controller.response(writer, "todos/list.liquid", q)
}

func (controller *TodoController) response(writer http.ResponseWriter, view string, q string) {
	list, err := controller.service.List(q)
	if err != nil {
		writer.WriteHeader(500)
		writer.Write([]byte(err.Error()))
		return
	}
	data := map[string]interface{}{"todos": list}
	controller.config.Render(writer, view, data)
}
