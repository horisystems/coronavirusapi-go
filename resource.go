package client

import (
	"strconv"
	"strings"

	simplejson "github.com/bitly/go-simplejson"
)

type Resource struct {
	Name string
	Path string
}

func NewResource(name, path string) *Resource {

	r := &Resource{
		Name: name,
		Path: path,
	}
	return r
}

func (r *Resource) GetAll() (*simplejson.Json, error) {
	return getClient().http.request("GET", r.Path, nil)
}

func (r *Resource) GetById(id int) (*simplejson.Json, error) {
	return getClient().http.request("GET", formatPath(r.Path, strconv.Itoa(id)), nil)
}

func (r *Resource) Create(data map[string]interface{}) {
	getClient().http.request("POST", r.Path, data)
}

func (r *Resource) DeleteById(id int) {
	getClient().http.request("DELETE", formatPath(r.Path, strconv.Itoa(id)), nil)
}

func (r *Resource) UpdateById(id int, data map[string]interface{}) {
	getClient().http.request("PUT", formatPath(r.Path, strconv.Itoa(id)), data)
}

func formatPath(paths ...string) string {
	var str string
	for _, path := range paths {
		str += strings.Trim(path, "/") + "/"
	}
	return strings.TrimRight(str, "/")
}
