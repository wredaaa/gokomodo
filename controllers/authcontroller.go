package controllers

import (
	"API-GOKOMODO/database"
	"API-GOKOMODO/entities"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type Authentication struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Token struct {
	Email       string `json:"email"`
	TokenString string `json:"token"`
}

/*
	Do this function to generate new hash password.

	Args:
		password: plain password

	Return:
		hashing password

*
*/
func GeneratehashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

/*
	Do this function to compare plain password with hash password.

	Args:
		password: plain password
		hash: hash password

	Return:
		true/false

*
*/
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

/*
	Do this function to generate JWT token.

	Args:
		email: email user
		role: hash password

	Return:
		token JWT

*
*/
func GenerateJWT(email string) (string, error) {
	var mySigningKey = []byte(os.Getenv("API_SECRET"))
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["email"] = email
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		fmt.Errorf("Something went Wrong: %s", err.Error())
		return "", err
	}

	return tokenString, nil
}

/*
	Do this function to check user is authorized or not.

*
*/
func IsAuthorized(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var valToken = ""
		var mySigningKey = []byte(os.Getenv("API_SECRET"))
		bearerToken := r.Header.Get("Authorization")

		if len(strings.Split(bearerToken, " ")) == 2 {
			valToken = strings.Split(bearerToken, " ")[1]
		} else {
			json.NewEncoder(w).Encode("No Token Found")
			return
		}

		token, err := jwt.Parse(valToken, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("There was an error in parsing token.")
			}
			return mySigningKey, nil
		})

		if err != nil {
			json.NewEncoder(w).Encode("Your Token has been expired.")
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			Pretty(claims)
			handler.ServeHTTP(w, r)
			return
		}
		json.NewEncoder(w).Encode("Not Authorized.")
	}
}

/*
	Do this function to login.
	if email or password wrong then error

	Args:
		email: email user
		password: password user

	Returns:
		user data & token JWT

*
*/
func LoginBuyer(w http.ResponseWriter, r *http.Request) {
	connection := database.GetDatabase()
	defer database.CloseDatabase(connection)

	var authDetails Authentication

	err := json.NewDecoder(r.Body).Decode(&authDetails)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode("Error in reading payload.")
		return
	}

	var authBuyer entities.Buyer
	connection.Where("email = 	?", authDetails.Email).First(&authBuyer)

	if authBuyer.Email == "" {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode("Username or Password is incorrect")
		return
	}

	check := CheckPasswordHash(authDetails.Password, authBuyer.Password)

	if !check {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode("Username or Password is incorrect")
		return
	}

	validToken, err := GenerateJWT(authBuyer.Email)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode("Failed to generate token")
		return
	}

	var token Token
	token.Email = authBuyer.Email
	token.TokenString = validToken
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(token)
}

/*
	Do this function to login.
	if email or password wrong then error

	Args:
		email: email user
		password: password user

	Returns:
		user data & token JWT

*
*/
func LoginSeller(w http.ResponseWriter, r *http.Request) {
	connection := database.GetDatabase()
	defer database.CloseDatabase(connection)

	var authDetails Authentication

	err := json.NewDecoder(r.Body).Decode(&authDetails)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode("Error in reading payload.")
		return
	}

	var authSeller entities.Seller
	connection.Where("email = 	?", authDetails.Email).First(&authSeller)

	if authSeller.Email == "" {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode("Username or Password is incorrect")
		return
	}

	check := CheckPasswordHash(authDetails.Password, authSeller.Password)

	if !check {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode("Username or Password is incorrect")
		return
	}

	validToken, err := GenerateJWT(authSeller.Email)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode("Failed to generate token")
		return
	}

	var token Token
	token.Email = authSeller.Email
	token.TokenString = validToken
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(token)
}

// Pretty display the claims licely in the terminal
func Pretty(data interface{}) {
	b, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println(string(b))
}
