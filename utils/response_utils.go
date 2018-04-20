package utils

import (
	"encoding/json"
	"net/http"
)

func ProcessJson(res http.ResponseWriter, httpCode int, data interface{}) {

	response, _ := json.Marshal(data)
	res.WriteHeader(httpCode)
	res.Write(response)
}

func ProcessError(res http.ResponseWriter, err error) {

	errString := err.Error()

	switch errString {

	case "Internal Server Error: sql: no rows in result set":
		ProcessJson(res, http.StatusNotFound, map[string]string{"result": "Not Found"})

	// case INCORRECT_USERNAME_PASSWORD:
	// 	processJson(res, http.StatusUnauthorized, map[string]string{"result": INCORRECT_USERNAME_PASSWORD})

	// case INVALID_TOKEN:
	// 	processJson(res, http.StatusUnauthorized, map[string]string{"result": INVALID_TOKEN})

	// case ID_ALREADY_PRESENT:
	// 	processJson(res, http.StatusBadRequest, map[string]string{"result": ID_ALREADY_PRESENT})

	// case ID_DOES_NOT_EXIST:
	// 	processJson(res, http.StatusNotFound, map[string]string{"result": ID_DOES_NOT_EXIST})

	// case BAD_REQUEST:
	// 	processJson(res, http.StatusNotFound, map[string]string{"result": BAD_REQUEST})

	default:
		ProcessJson(res, http.StatusInternalServerError, map[string]string{"result": errString})
	}

}
