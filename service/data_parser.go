package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func ParseData(request *http.Request, data any, writer http.ResponseWriter) error {
	contentTypeHeader := request.Header.Get("Content-Type")

	if strings.Contains(contentTypeHeader, "application/json") || strings.Contains(contentTypeHeader, "application/*") || strings.Contains(contentTypeHeader, "*/*") {
		decoder := json.NewDecoder(request.Body)
		decoder.DisallowUnknownFields()

		err := decoder.Decode(&data)
		if err != nil {
			http.Error(writer, "400 Bad Request", http.StatusBadRequest)
			return err
		}

		return nil
	}

	http.Error(writer, "415 Unsupported Media Type", http.StatusUnsupportedMediaType)
	return fmt.Errorf("content negotiation failed")
}
