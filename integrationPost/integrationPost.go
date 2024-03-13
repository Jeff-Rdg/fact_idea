package integrationPost

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
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

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}
	defer resp.Body.Close()

	return body, nil

}
