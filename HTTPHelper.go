package crema

import (
	"net/http"
)

func PopulateSingleValueQueries(r *http.Request, conditions map[string]string) map[string]string {
	r.ParseForm()
	request := r.Form

	for key, val := range request {
		conditions[key] = val[0]
	}

	return conditions
}

func PopulateParams(params map[string]string, conditions map[string]string) map[string]string {
	for key, val := range params {
		conditions[key] = val
	}

	return conditions
}
