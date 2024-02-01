package transport

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/Farhan-slurrp/go-car/database"
	"github.com/Farhan-slurrp/go-car/internal"
	"github.com/Farhan-slurrp/go-car/internal/config"
	"github.com/Farhan-slurrp/go-car/pkg/user/endpoints"
)

func NewHTTPHandler(ep endpoints.Set) http.Handler {
	database.ConnectDB()
	database.DB.AutoMigrate(&internal.User{})
	config.GoogleConfig()

	m := http.NewServeMux()

	m.HandleFunc("/login", handleLogin)
	m.HandleFunc("/callback", handleCallback)

	return m
}

func generateStateOauthCookie(w http.ResponseWriter) string {
	var expiration = time.Now().Add(365 * 24 * time.Hour)

	b := make([]byte, 16)
	rand.Read(b)
	state := base64.URLEncoding.EncodeToString(b)
	cookie := http.Cookie{Name: "oauthstate", Value: state, Expires: expiration}
	http.SetCookie(w, &cookie)

	return state
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	// Create oauthState cookie
	oauthState := generateStateOauthCookie(w)
	u := config.AppConfig.GoogleLoginConfig.AuthCodeURL(oauthState)
	http.Redirect(w, r, u, http.StatusTemporaryRedirect)
}

func handleCallback(w http.ResponseWriter, r *http.Request) {
	// Read oauthState from Cookie
	oauthState, _ := r.Cookie("oauthstate")

	if r.FormValue("state") != oauthState.Value {
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	data, err := getUserDataFromGoogle(r.FormValue("code"))
	if err != nil {
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	// GetOrCreate User in your db.
	// Redirect or response with a token.
	user := internal.User{}
	userResp := internal.UserResponse{}
	err = json.Unmarshal(data, &userResp)
	if err != nil {
		fmt.Printf("%v", err)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	user.Email = userResp.Email
	user.Name = userResp.Name
	user.Role = "user"
	database.DB.Create(&user)
	fmt.Fprintf(w, "UserInfo: %s\n", user)
}

func getUserDataFromGoogle(code string) ([]byte, error) {
	// Use code to get token and get user info from Google.

	token, err := config.AppConfig.GoogleLoginConfig.Exchange(context.Background(), code)
	if err != nil {
		return nil, fmt.Errorf("code exchange wrong: %s", err.Error())
	}
	response, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		return nil, fmt.Errorf("failed getting user info: %s", err.Error())
	}
	defer response.Body.Close()
	userData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed read response: %s", err.Error())
	}
	return userData, nil
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if e, ok := response.(error); ok && e != nil {
		encodeError(ctx, e, w)
		return nil
	}
	return json.NewEncoder(w).Encode(response)
}

func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	switch err {
	// case util.ErrUnknown:
	// 	w.WriteHeader(http.StatusNotFound)
	// case util.ErrInvalidArgument:
	// 	w.WriteHeader(http.StatusBadRequest)
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}
