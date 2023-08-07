package contract

import "net/http"

type TemplateServiceRestInterface interface {
	Add(w http.ResponseWriter, r *http.Request)
}
