package models

import (
	"net/http"
	"strconv"
	"time"
)

type Todo struct {
	Id          uint64 `gorm:"primaryKey"`
	Description string
	Done        bool
	created     time.Time `gorm:"autoCreateTime"`
}

func (t *Todo) SetId(id string) {
	result, err := strconv.Atoi(id)
	if err != nil {
		t.Id = 0
	}
	t.Id = uint64(result)
}

func TodoFromForm(request *http.Request) (Todo, error) {
	err := request.ParseForm()
	return Todo{
		Id:          getUint64(request, "id", 0),
		Description: request.PostForm.Get("description"),
		Done:        getBool(request, "done", false),
		created:     getTime(request, "created", time.Now()),
	}, err
}

func getUint64(request *http.Request, key string, orElse uint64) uint64 {
	val := request.PostForm.Get(key)
	result, err := strconv.Atoi(val)
	if err != nil {
		return orElse
	}
	return uint64(result)
}

func getBool(request *http.Request, key string, orElse bool) bool {
	val := request.PostForm.Get(key)
	result, err := strconv.ParseBool(val)
	if err != nil {
		return orElse
	}
	return result
}

func getTime(request *http.Request, key string, orElse time.Time) time.Time {
	val := request.PostForm.Get(key)
	result, err := time.Parse(time.RFC3339, val)
	if err != nil {
		return orElse
	}
	return result

}
