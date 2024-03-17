package parser

import (
	"encoding/json"
	"io"
	"net/http"
)

func ParseBody(req *http.Request, out interface{}) error {
	bodyInBytes, err := io.ReadAll(req.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(bodyInBytes, out)
	if err != nil {
		return err
	}

	return nil
}
