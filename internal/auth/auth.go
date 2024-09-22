package auth

import (
	"errors"
	"net/http"
	"strings"
)

func GetApiKey(header http.Header)(string,error) {

	val := header.Get("Authorization")
	if val == " " {
		return "", errors.New("no authorization info found")
	}

	vals := strings.Split(val, " ")

	if vals[0] != "ApiKey" || len(vals) != 2 {
		return "", errors.New("malformed auth header")
	}

	return vals[1],nil
}
