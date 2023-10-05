package core

import "net/http"

type RequestWrapper interface {
	GetPathParameter(param string) string
	GetRequest() *http.Request //helpful, at least for now

}
