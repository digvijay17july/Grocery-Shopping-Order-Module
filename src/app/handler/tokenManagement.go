package handler

import (
	"github.com/dgrijalva/jwt-go"
	"net/http"
)

// Create the JWT key used to create the signature
var jwtKey = []byte("my_secret_key")

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

type Credentials struct {
	username string `json:"username"`
	password string `json:"password"`
}
func validate(tknStr string) bool {


	// Initialize a new instance of `Claims`
	claims := &Claims{}

	isValid := false
	if tknStr ==""{
 	return  isValid
	}

	// Parse the JWT string and store the result in `claims`.
	// Note that we are passing the key in this method as well. This method will return an error
	// if the token is invalid (if it has expired according to the expiry time we set on sign in),
	// or if the signature does not match
	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if tkn==nil || !tkn.Valid  {
		//w.WriteHeader(http.StatusUnauthorized)
		return isValid
	}
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			//w.WriteHeader(http.StatusUnauthorized)
			return isValid
		}
		//w.WriteHeader(http.StatusBadRequest)
		return isValid
	}
	isValid=true

	// Finally, return the welcome message to the user, along with their
	// username given in the token
	//w.Write([]byte(fmt.Sprintf("Welcome %s!", claims.Username)))
	return isValid
}

func  authenticate( w http.ResponseWriter, r *http.Request) bool{

	result:=validate(r.Header.Get("jwt-token"))
	return result

}