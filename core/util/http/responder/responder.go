package responder

import (
	net_http "net/http"
)

func Respond(w net_http.ResponseWriter, data ...interface{}) error {
	if data[0] == nil {
		return respondWithStatus(w, net_http.StatusNoContent)
	}

	switch d0 := data[0].(type) {
	case string:
		return respondWithStatusAndRawData(w, net_http.StatusOK, []byte(d0))
	case *string:
		return respondWithStatusAndRawData(w, net_http.StatusOK, []byte(*d0))
	case []byte:
		return respondWithStatusAndRawData(w, net_http.StatusOK, d0)
	case *[]byte:
		return respondWithStatusAndRawData(w, net_http.StatusOK, *d0)
	}
	return nil
}

func respondWithStatus(w net_http.ResponseWriter, status int) error {
	w.WriteHeader(status)
	return nil
}

func respondWithStatusAndRawData(w net_http.ResponseWriter, status int, body []byte) error {
	if err := respondWithStatus(w, status); err != nil {
		return err
	}
	_, err := w.Write(body)
	return err
}
