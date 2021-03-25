package jobapps

import (
	"encoding/json"
	"fmt"
	"github.com/nikhilmalhotra123/apps/jobapps/model"
  "github.com/nikhilmalhotra123/apps/jobapps/service"
	"io/ioutil"
	"net/http"
  "strings"
	jwt "github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

// SignUpHandler function
func SignUpHandler(w http.ResponseWriter, r *http.Request)  {
  w.Header().Set("Content-Type", "application/json")

  var user model.User
  var res model.Response
  body, _ := ioutil.ReadAll(r.Body)

  if err := json.Unmarshal(body, &user); err != nil {
    res.Error = err.Error()
    json.NewEncoder(w).Encode(res)
    return
  }

  if _, err := jobapps.GetUserByUsername(user.Username); err != nil {
    if err.Error() == "mongo: no documents in result" {
      hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

      if err != nil {
				res.Error = "Error While Hashing Password, Try Again"
				json.NewEncoder(w).Encode(res)
				return
			}

      user.Password = string(hash)

      if err := jobapps.InsertUser(&user); err != nil {
        res.Error = "Error creating user. Try again"
        json.NewEncoder(w).Encode(res)
        return
      }

      res.Result = "Account Created Succesfully"
      json.NewEncoder(w).Encode(res)
			return
    }

    res.Error = err.Error()
    json.NewEncoder(w).Encode(res)
    return
  }

  res.Error = "Username already Exists!"
  json.NewEncoder(w).Encode(res)
  return
}

// LoginHandler function
func LoginHandler(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")

  var user model.User
  var res model.Response

  body, _ := ioutil.ReadAll(r.Body)
  if err := json.Unmarshal(body, &user); err != nil {
    res.Error = err.Error()
    json.NewEncoder(w).Encode(res)
    return
  }

  var result *model.User
  result, err := jobapps.GetUserByUsername(user.Username)

  if err != nil {
    res.Error = "Invalid username"
    json.NewEncoder(w).Encode(res)
    return
  }

  if err := bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(user.Password)); err != nil {
    res.Error = "Invalid password"
    json.NewEncoder(w).Encode(res)
    return
  }

  token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username":  result.Username,
		"firstname": result.Firstname,
		"lastname":  result.Lastname,
	})

  tokenString, err := token.SignedString([]byte("secret"))

  if err != nil {
    res.Error = "Error while generating token,Try again"
    json.NewEncoder(w).Encode(res)
    return
  }

  result.Token = tokenString
  result.Password = ""
	res.Result = result
  json.NewEncoder(w).Encode(res)
}

// ProfileHandler function
func ProfileHandler(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")

	tokenString := r.Header.Get("Authorization")
  tokenSplit := strings.Split(tokenString, " ")[1]
	token, err := jwt.Parse(tokenSplit, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method")
		}
		return []byte("secret"), nil
	})

	var user model.User
	var res model.Response

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		user.Username = claims["username"].(string)
		user.Firstname = claims["firstname"].(string)
		user.Lastname = claims["lastname"].(string)

    var result *model.User
    result, err = jobapps.GetUserByUsername(user.Username)

    if err != nil {
      res.Error = "User does not exist"
      json.NewEncoder(w).Encode(res)
      return
    }

    result.Password = ""
		res.Result = result
		json.NewEncoder(w).Encode(res)
		return
	}

	res.Error = err.Error()
	json.NewEncoder(w).Encode(res)
	return
}

// DeleteHandler function
func DeleteHandler(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")

	tokenString := r.Header.Get("Authorization")
  tokenSplit := strings.Split(tokenString, " ")[1]
	token, err := jwt.Parse(tokenSplit, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method")
		}
		return []byte("secret"), nil
	})

  var username string
	var res model.Response

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		username = claims["username"].(string)
    if err := jobapps.DeleteUserByUsername(username); err != nil {
			res.Error = err.Error()
			json.NewEncoder(w).Encode(res)
			return
    }

    res.Result = "Account Deleted Successfully"
    json.NewEncoder(w).Encode(res)
		return
	}

	res.Error = err.Error()
	json.NewEncoder(w).Encode(res)
	return
}
