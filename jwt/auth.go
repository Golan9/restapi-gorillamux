package jwt

import (
	"encoding/base64"
	"net/http"
	cfg "restapi-gorillamux/config"
	m "restapi-gorillamux/modules/users/models"
	"strings"
	"time"

	"fmt"

	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	jwt "github.com/dgrijalva/jwt-go"
)

type jwtClaims struct {
	username string
	jwt.StandardClaims
}

// BasicAuthRequest ..
type basicAuthRequest struct {
	username string
	password string
}

// MySigningKey ..
var mySigningKey = []byte("privateKeys")

var config = cfg.Config{}

// JwtMiddleware ..
var JwtMiddleware = jwtmiddleware.New(jwtmiddleware.Options{
	ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	},
	SigningMethod: jwt.SigningMethodHS512,
})

// CreateJwtToken ..
func CreateJwtToken() (string, error) {
	var user m.User
	claims := jwtClaims{
		user.Username,
		jwt.StandardClaims{
			Id:        "main_user_id",
			ExpiresAt: time.Now().Add(10 * time.Minute).Unix(),
		},
	}
	rawToken := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)

	token, err := rawToken.SignedString(mySigningKey)
	if err != nil {
		return "", err
	}
	return token, nil
}

// BasicAuth ..
func BasicAuth(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)

		config.Read()

		var ba basicAuthRequest

		ba.username = config.Username
		ba.password = config.Password

		username, password, ok := r.BasicAuth()
		fmt.Println("debugging", config)
		s := strings.SplitN(r.Header.Get("Authorization"), " ", 2)
		if len(s) != 2 {
			http.Error(w, "Not authorized", 401)
			return
		}

		b, err := base64.StdEncoding.DecodeString(s[1])
		if err != nil {
			http.Error(w, err.Error(), 401)
			return
		}

		pair := strings.SplitN(string(b), ":", 2)
		if len(pair) != 2 {
			http.Error(w, "Not authorized", 401)
			return
		}

		if username != ba.username || password != ba.password || ok != true {
			http.Error(w, "Not authorized", 401)
			return
		}

		if pair[0] != ba.username || pair[1] != ba.password {
			http.Error(w, "Not authorized", 401)
			return
		}

		h.ServeHTTP(w, r)
	}
}
