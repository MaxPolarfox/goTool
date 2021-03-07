package errors

import (
	"encoding/json"
	"net/http"
)

func RespondWithError(rw http.ResponseWriter, statusCode int, message string) {

	response := Error{
		// we will omit StatusCode in the body, because one can grab it from the response itself
		Message: message,
	}
	js, err := json.Marshal(&response)
	if err != nil {
		failedToMarshalError := Error{
			Message: err.Error(),
		}
		failedJS, _ := json.Marshal(&failedToMarshalError)
		rw.Header().Set("Content-Type", "application/json")
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write(failedJS)
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(statusCode)
	rw.Write(js)
}