package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/MaxPolarfox/goTools/errors"
)

// parseResponse ensures the response status codes, and ability to decode response body
func parseResponse(res *http.Response, respStruct interface{}) error {
	js, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return &errors.Error{
			Message:    err.Error(),
		}
	}

	switch res.StatusCode {
	case http.StatusOK, http.StatusCreated:
		if err := json.Unmarshal(js, respStruct); err != nil {
			return &errors.Error{
				Message:    string(js),
			}
		}
		return nil
	case http.StatusNoContent:
		return nil
	case http.StatusNotFound, http.StatusBadRequest, http.StatusConflict:
		return &errors.Error{
			Message:    string(js),
		}
	default:
		return &errors.Error{
			Message:    fmt.Sprintf("unexpected response code: %v, message: %v", res.StatusCode, string(js)),
		}
	}
}
