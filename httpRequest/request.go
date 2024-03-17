package httpRequest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

type HttpRequest struct {
	url         string
	method      HttpMethod
	contentType *ContentType
	body        []byte
	token       *string
	params      map[string]interface{}
	basicAuth   *basicAuth
	client      *http.Client
}

type basicAuth struct {
	username string
	password string
}

func New(url string, method HttpMethod) Requestier {
	return &HttpRequest{
		url:    url,
		method: method,
		client: &http.Client{},
	}
}

type Requestier interface {
	SetBody(body []byte) *HttpRequest
	SetToken(token string) *HttpRequest
	SetParams(params map[string]interface{}) *HttpRequest
	SetBasicAuth(username, password string) *HttpRequest
	SetContentType(content ContentType) *HttpRequest
	Send() ([]byte, error)
}

func (r *HttpRequest) SetBody(body []byte) *HttpRequest {
	r.body = body
	return r
}

func (r *HttpRequest) SetParams(params map[string]interface{}) *HttpRequest {
	r.params = params
	return r
}

func (r *HttpRequest) SetBasicAuth(username, password string) *HttpRequest {
	r.basicAuth.username = username
	r.basicAuth.password = password
	return r
}

func (r *HttpRequest) SetToken(token string) *HttpRequest {
	r.token = &token
	return r
}

func (r *HttpRequest) SetContentType(content ContentType) *HttpRequest {
	r.contentType = &content
	return r
}

func (r *HttpRequest) Send() ([]byte, error) {
	payload := bytes.NewBuffer(r.body)

	parsedUrl, err := url.Parse(r.url)
	if err != nil {
		return nil, err
	}

	queries := parsedUrl.Query()

	if r.params != nil {
		for key, value := range r.params {
			res, _ := json.Marshal(&value)
			queries.Add(key, string(res))
		}
	}

	parsedUrl.RawQuery = queries.Encode()

	req, err := http.NewRequest(r.method.String(), parsedUrl.String(), payload)
	if err != nil {
		return nil, err
	}

	if r.contentType != nil {
		req.Header.Add("Content-Type", r.contentType.String())
	}

	if r.token != nil {
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", *r.token))
	}

	if r.basicAuth != nil {
		req.SetBasicAuth(r.basicAuth.username, r.basicAuth.password)
	}

	resp, err := r.client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	defer resp.Body.Close()

	return body, nil
}
