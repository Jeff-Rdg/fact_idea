package main

import (
	"encoding/json"
	"fmt"
	"github.com/jeff-rdg/idea_pattern/httpRequest"
	"github.com/jeff-rdg/idea_pattern/integrationPost"
	"github.com/jeff-rdg/idea_pattern/integrationPost/integrationRequest"
	"github.com/jeff-rdg/idea_pattern/integrationPost/integrationResponse"
)

func main() {

	res, err := httpRequest.New("https://jsonplaceholder.typicode.com/posts", httpRequest.GET).
		SetParams(map[string]interface{}{
			"userId": 1,
		}).
		Send()

	fmt.Println(string(res))

	post := &integrationRequest.PostRequest{
		Title:  "Testando integração",
		Body:   "numero 01",
		UserId: 1,
	}
	result, err := integrationPost.Execute(post)
	if err != nil {
		fmt.Println(err)
	}

	response, err := integrationResponse.FactoryResponse{}.New(post.RequestType())
	err = json.Unmarshal(result, response)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(response.ResponseToString())

}
