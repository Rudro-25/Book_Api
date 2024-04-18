package authHandler

//
//import (
//	"encoding/json"
//	"fmt"
//	jwt "github.com/golang-jwt/jwt/v5"
//	"net/http"
//	"time"
//)
//
//var jwtKey = []byte("secret_key")
//
//var users = map[string]string{
//	"user1": "password1",
//	"user2": "password2",
//}
//
//type Credentials struct {
//	Username string `json:"username"`
//	Password string `json:"password"`
//}
//
//type Claims struct {
//	Username string `json:"username"`
//	jwt.RegisteredClaims
//}
//
//func Login(w http.ResponseWriter, r *http.Request) {
//	var credentials Credentials
//	err := json.NewDecoder(r.Body).Decode(&credentials)
//	if err != nil {
//		w.WriteHeader(http.StatusBadRequest)
//		return
//	}
//
//	expectedPassword, ok := users[credentials.Username]
//
//	if !ok || expectedPassword != credentials.Password {
//		w.WriteHeader(http.StatusUnauthorized)
//		return
//	}
//
//	expirationTime := time.Now().Add(time.Minute * 5)
//
//	claims := &Claims{
//		Username: credentials.Username,
//		RegisteredClaims: jwt.RegisteredClaims{
//			ExpiresAt: jwt.NewNumericDate(expirationTime),
//		},
//	}
//
//	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
//	tokenString, err := token.SignedString(jwtKey)
//
//	if err != nil {
//		w.WriteHeader(http.StatusInternalServerError)
//		return
//	}
//
//	http.SetCookie(w,
//		&http.Cookie{
//			Name:    "token",
//			Value:   tokenString,
//			Expires: expirationTime,
//		})
//
//}
//
//func Home(w http.ResponseWriter, r *http.Request) {
//	cookie, err := r.Cookie("token")
//	if err != nil {
//		if err == http.ErrNoCookie {
//			w.WriteHeader(http.StatusUnauthorized)
//			return
//		}
//		w.WriteHeader(http.StatusBadRequest)
//		return
//	}
//
//	tokenStr := cookie.Value
//
//	claims := &Claims{}
//
//	tkn, err := jwt.ParseWithClaims(tokenStr, claims,
//		func(t *jwt.Token) (interface{}, error) {
//			return jwtKey, nil
//		})
//
//	if err != nil {
//		if err == jwt.ErrSignatureInvalid {
//			w.WriteHeader(http.StatusUnauthorized)
//			return
//		}
//		w.WriteHeader(http.StatusBadRequest)
//		return
//	}
//
//	if !tkn.Valid {
//		w.WriteHeader(http.StatusUnauthorized)
//		return
//	}
//
//	w.Write([]byte(fmt.Sprintf("Hello, %s", claims.Username)))
//
//}
