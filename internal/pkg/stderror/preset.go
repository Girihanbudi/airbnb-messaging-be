package stderror

import "net/http"

var (
	DEF_SERVER_500 = New(http.StatusUnauthorized, "DEF_SERVER_500", "Failed to get message")
	DEF_SERVER_503 = New(http.StatusServiceUnavailable, "DEF_SERVER_503", "Service unavailable to response")
	DEF_AUTH_401   = New(http.StatusUnauthorized, "DEF_AUTH_401", "Failed to authorize user")
	DEF_DATA_400   = New(http.StatusBadRequest, "DEF_DATA_400", "Failed to bind request")
)
