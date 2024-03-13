package main

import (
	"encoding/json"
	"fmt"
	"github.com/jeff-rdg/idea_pattern/integrationPost"
	"github.com/jeff-rdg/idea_pattern/integrationPost/integrationRequest"
	"github.com/jeff-rdg/idea_pattern/integrationPost/integrationResponse"
)

func main() {
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
