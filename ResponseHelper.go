package crema

import (
	"net/http"
	"strconv"
)

func GenericHTTPResponse(code int) Response {
	var response Response

	if code == http.StatusOK {
		response.Status = 1
		response.Message = "Success " + strconv.Itoa(code)
	} else {
		response.Status = 0
		response.Message = "Error " + strconv.Itoa(code)
	}

	return response
}
