package jobapps

import (
	"encoding/json"
	"fmt"
	"github.com/nikhilmalhotra123/apps/jobapps/model"
  "github.com/nikhilmalhotra123/apps/jobapps/service"
	"net/http"
  "strings"
	"io/ioutil"
	"github.com/gorilla/mux"
	jwt "github.com/dgrijalva/jwt-go"
)

// InsertApplicationHandler function
func InsertApplicationHandler(w http.ResponseWriter, r *http.Request)  {
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

		var app model.Application
		body, _ := ioutil.ReadAll(r.Body)
		if err := json.Unmarshal(body, &app); err != nil {
			res.Error = err.Error()
			json.NewEncoder(w).Encode(res)
			return
		}

		if user.Username != app.Username {
			res.Error = "User cannt insert app for other user"
			json.NewEncoder(w).Encode(res)
      return
		}

    err = jobapps.InsertApplication(&app);
    if err != nil {
      res.Error = "Failed to insert app"
			json.NewEncoder(w).Encode(res)
      return
    }

		res.Result = "Successfully inserted application"
		json.NewEncoder(w).Encode(res)
		return
	}

	res.Error = err.Error()
	json.NewEncoder(w).Encode(res)
	return
}

// GetApplicationHandler function
func GetApplicationHandler(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")

	tokenString := r.Header.Get("Authorization")
  tokenSplit := strings.Split(tokenString, " ")[1]
	token, err := jwt.Parse(tokenSplit, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method")
		}
		return []byte("secret"), nil
	})

	vars := mux.Vars(r)
	var user model.User
	var res model.Response

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		user.Username = claims["username"].(string)
		user.Firstname = claims["firstname"].(string)
		user.Lastname = claims["lastname"].(string)

    result, err := jobapps.GetApplicationByID(vars["id"])
    if err != nil {
      res.Error = "Cannot find application"
      return
    }
		res.Result = result
		json.NewEncoder(w).Encode(res)
		return
	}

	res.Error = err.Error()
	json.NewEncoder(w).Encode(res)
	return
}

// GetAllApplicationsHandler function
func GetAllApplicationsHandler(w http.ResponseWriter, r *http.Request) {
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

    result, err := jobapps.GetAllApplicationsByUsername(user.Username)
    if err != nil {
      res.Error = "Cannot find applications"
      return
    }
		res.Result = result
		json.NewEncoder(w).Encode(res)
		return
	}

	res.Error = err.Error()
	json.NewEncoder(w).Encode(res)
	return
}

// UpdateApplicationHandler function
func UpdateApplicationHandler(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")

	tokenString := r.Header.Get("Authorization")
  tokenSplit := strings.Split(tokenString, " ")[1]
	token, err := jwt.Parse(tokenSplit, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method")
		}
		return []byte("secret"), nil
	})

	vars := mux.Vars(r)
	var res model.Response
	var user model.User

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		user.Username = claims["username"].(string)
		user.Firstname = claims["firstname"].(string)
		user.Lastname = claims["lastname"].(string)

		var app model.Application
		body, _ := ioutil.ReadAll(r.Body)

		if err := json.Unmarshal(body, &app); err != nil {
			res.Error = err.Error()
			json.NewEncoder(w).Encode(res)
			return
		}

		appOwner, err := jobapps.GetApplicationByID(app.ID)
		if err != nil {
			res.Error = "Can not find application"
			return
		}

		if user.Username != appOwner.Username {
			res.Error = "User cannt update app for other user"
			json.NewEncoder(w).Encode(res)
			return
		}

		if app.ID != vars["id"] {
			res.Error = "App id does not match api path"
			json.NewEncoder(w).Encode(res)
			return
		}

    if err := jobapps.UpdateApplication(&app); err != nil {
      res.Error = "Failed to update app"
			json.NewEncoder(w).Encode(res)
      return
    }
		res.Result = "Applicatin Updated Successfully"
		json.NewEncoder(w).Encode(res)
		return
	}

	res.Error = err.Error()
	json.NewEncoder(w).Encode(res)
	return
}

// DeleteApplicationHandler function
func DeleteApplicationHandler(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")

	tokenString := r.Header.Get("Authorization")
  tokenSplit := strings.Split(tokenString, " ")[1]
	token, err := jwt.Parse(tokenSplit, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method")
		}
		return []byte("secret"), nil
	})

	vars := mux.Vars(r)
	var user model.User
	var res model.Response

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		user.Username = claims["username"].(string)
		user.Firstname = claims["firstname"].(string)
		user.Lastname = claims["lastname"].(string)

		appOwner, err := jobapps.GetApplicationByID(vars["id"])
		if err != nil {
			res.Error = "Can not find application"
			return
		}

		if user.Username != appOwner.Username {
			res.Error = "User cannt delete app for other user"
			json.NewEncoder(w).Encode(res)
			return
		}

    if err = jobapps.DeleteApplicationByID(vars["id"]); err != nil {
      res.Error = "Failed to delete app"
			json.NewEncoder(w).Encode(res)
      return
    }

		res.Result = "Application Deleted Successfully"
		json.NewEncoder(w).Encode(res)
		return
	}

	res.Error = err.Error()
	json.NewEncoder(w).Encode(res)
	return
}
