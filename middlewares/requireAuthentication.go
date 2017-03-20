package middlewares

import (
	"net/http"
	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"SimpleChat/secrets"
)

func RequireAuthentication(rw http.ResponseWriter, req *http.Request, next http.HandlerFunc) {
	token, err := request.ParseFromRequest(
		req,
		request.AuthorizationHeaderExtractor,
		func(token *jwt.Token) (interface{}, error) {
			return []byte(secrets.JwtSecret), nil
		},
	)

	if err == nil && token.Valid {
		next(rw, req)
	} else {
		rw.WriteHeader(http.StatusUnauthorized)
	}
}