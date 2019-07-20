package controllers

import (
	"crypto/hmac"
	"crypto/sha512"
	"dtos"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"models"
	"net/http"
	"strconv"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

//JwtClaims struct
type JwtClaims struct {
	Name string `json:"name"`
	jwt.StandardClaims
}

type RealToekn struct {
	Token string `json:"Token"`
}

//Login user
func Login(w http.ResponseWriter, r *http.Request) {
	var user dtos.UserDto
	dbUser := models.User{}
	err := json.NewDecoder(r.Body).Decode(&user)
	check(err)

	rows, err := DB.Query("select * from users")
	for rows.Next() {
		err = rows.Scan(&dbUser.Id,
			&dbUser.Username,
			&dbUser.PasswordHash,
			&dbUser.PasswordSalt,
		)
		if user.Username == dbUser.Username {
			if !ValidPassword([]byte(user.Password), []byte(dbUser.PasswordHash), []byte(dbUser.PasswordSalt)) {
				http.Error(w, "wrong password", 400)
				return
			}
			var realToken RealToekn
			realToken.Token, err = CreateJwtToken(dbUser.Username, strconv.Itoa(dbUser.Id), user.Password)
			check(err)
			token, _ := json.Marshal(realToken)
			fmt.Fprint(w, string(token))
			return
		}
		http.Error(w, "username doesn't exist", 400)
		return
	}
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	if !CheckToken(r) {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	return
}

//Register new user
func Register(w http.ResponseWriter, r *http.Request) {

}

//CreateJwtToken is creating a token to return the client
func CreateJwtToken(userName string, id string, password string) (string, error) {
	claims := JwtClaims{
		userName,
		jwt.StandardClaims{
			Id:        id,
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
		},
	}
	rawToken := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	token, err := rawToken.SignedString(GetSecretKey())
	if err != nil {
		return "", err
	}
	return token, nil

}

//ValidPassword checks if the password is valid
func ValidPassword(password, passwordHash, passwordSalt []byte) bool {
	mac := hmac.New(sha512.New, passwordSalt)
	mac.Write(password)
	expectedPassword := mac.Sum(nil)
	return hmac.Equal(passwordHash, expectedPassword)
}

//GetSecretKey gets the special key for the encryption
func GetSecretKey() []byte {
	key, err := ioutil.ReadFile("settings.txt")
	check(err)
	return key
}

//CheckToken validates the token
func CheckToken(r *http.Request) bool {
	reqToken := r.Header.Get("Authorization")
	splitToken := strings.Split(reqToken, "Bearer ")
	if len(splitToken) < 2 {
		return false
	}
	reqToken = splitToken[1]
	token, err := jwt.Parse(reqToken, func(token *jwt.Token) (interface{}, error) {
		return GetSecretKey(), nil
	})
	if err == nil && token.Valid {
		return true
	} else {
		return false
	}
}
