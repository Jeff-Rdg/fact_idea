package integrationResponse

import (
	"encoding/json"
	"errors"
	"github.com/jeff-rdg/idea_pattern/integrationPost"
	"github.com/jeff-rdg/idea_pattern/integrationPost/integrationRequest"
	"reflect"
)

type PostResponse struct {
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
	UserId int    `json:"user_id"`
}

func (i *PostResponse) ResponseToString() string {
	req, _ := json.Marshal(i)

	return string(req)
}

type FactoryInterface interface {
	New(reqType reflect.Type) (integrationPost.Responsier, error)
}

type FactoryResponse struct {
}

func (f FactoryResponse) New(reqType reflect.Type) (integrationPost.Responsier, error) {
	switch reqType {
	case integrationRequest.PostRequest{}.RequestType():
		return &PostResponse{}, nil

	default:
		return nil, errors.New("invalid type")
	}
}
