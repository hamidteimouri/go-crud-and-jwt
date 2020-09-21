package response

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func JSON(ResponseWriter http.ResponseWriter, statusCode int, data interface{}) {
	ResponseWriter.WriteHeader(statusCode)
	err := json.NewEncoder(ResponseWriter).Encode(data)

	if err != nil {
		fmt.Fprint(ResponseWriter, "%s", err.Error())
	}
}
func ERROR(ResponseWriter http.ResponseWriter, statusCode int, err error) {
	if err != nil {
		JSON(ResponseWriter, statusCode, struct {
			Error string `json:"message"`
		}{
			Error: err.Error(),
		})
	}
	JSON(ResponseWriter, http.StatusBadRequest, nil)
}
