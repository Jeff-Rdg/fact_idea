package httpRequest

import "net/http"

type HttpMethod string

const (
	GET    HttpMethod = http.MethodGet
	POST   HttpMethod = http.MethodPost
	PATCH  HttpMethod = http.MethodPatch
	PUT    HttpMethod = http.MethodPut
	DELETE HttpMethod = http.MethodDelete
)

func isValidMethod(method string) bool {
	isValid := map[HttpMethod]bool{
		GET:    true,
		POST:   true,
		PATCH:  true,
		PUT:    true,
		DELETE: true,
	}

	_, ok := isValid[HttpMethod(method)]
	return ok
}

func (h HttpMethod) String() string {
	return string(h)
}
