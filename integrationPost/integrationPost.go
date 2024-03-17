package integrationPost

import (
	"encoding/json"
	"github.com/jeff-rdg/idea_pattern/httpRequest"
	"reflect"
)

type Requestier interface {
	RequestToString() string
	RequestType() reflect.Type
}

type Responsier interface {
	ResponseToString() string
}

func Execute(request Requestier) ([]byte, error) {
	url := "https://jsonplaceholder.typicode.com/posts"

	payload, _ := json.Marshal(request)

	return httpRequest.New(url, httpRequest.POST).SetContentType(httpRequest.JSON).SetBody(payload).Send()
}
