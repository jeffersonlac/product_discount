package helpers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func JSONResponse(w http.ResponseWriter, data interface{}, httpStatus int) {
	w.WriteHeader(httpStatus)

	bytes, e := json.Marshal(data)
	if e != nil {
		fmt.Println("Fail on parse response to JSON")
	}

	if _, e := w.Write(bytes); e != nil {
		fmt.Println("Fail on write response")
	}
}
