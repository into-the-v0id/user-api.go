package service

import (
	"encoding/json"
	"fmt"
	"github.com/samber/lo"
	"net/http"
	"strings"
)

func RespondData(writer http.ResponseWriter, data interface{}, request *http.Request) error {
	acceptHeader := request.Header.Get("Accept")

	if strings.Contains(acceptHeader, "application/json") || strings.Contains(acceptHeader, "application/*") || strings.Contains(acceptHeader, "*/*") {
		jsonBytes, err := json.Marshal(data)
		if err != nil {
			return err
		}

		jsonString := string(jsonBytes)

		writer.Header().Set("Content-Type", "application/json; charset=utf-8")
		lo.Must(fmt.Fprint(writer, jsonString))
		return nil
	}

	http.Error(writer, "406 Not Acceptable", http.StatusNotAcceptable)
	return nil
}
