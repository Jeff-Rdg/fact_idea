package integrationRequest

import (
	"encoding/json"
	"reflect"
)

type PostRequest struct {
	Title  string `json:"title"`
	Body   string `json:"body"`
	UserId int    `json:"userId"`
}

func (i *PostRequest) RequestToString() string {
	req, _ := json.Marshal(i)

	return string(req)
}

func (i PostRequest) RequestType() reflect.Type {
	return reflect.TypeOf(i)
}
