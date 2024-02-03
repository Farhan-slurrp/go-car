package authentication

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/Farhan-slurrp/go-car/internal/config"
	"golang.org/x/oauth2"
)

func GenerateStateOauthCookie(w http.ResponseWriter) string {
	var expiration = time.Now().Add(365 * 24 * time.Hour)

	b := make([]byte, 16)
	rand.Read(b)
	state := base64.URLEncoding.EncodeToString(b)
	cookie := http.Cookie{Name: "oauthstate", Value: state, Expires: expiration}
	http.SetCookie(w, &cookie)

	return state
}

func GetToken(code string) (*oauth2.Token, error) {
	token, err := config.AppConfig.GoogleLoginConfig.Exchange(context.Background(), code)
	if err != nil {
		return nil, fmt.Errorf("code exchange wrong: %s", err.Error())
	}

	return token, nil
}

func GetUserDataFromGoogle(token string) ([]byte, error) {
	response, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token)

	if err != nil {
		return nil, fmt.Errorf("failed getting user info: %s", err.Error())
	}
	defer response.Body.Close()
	userData, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, fmt.Errorf("failed read response: %s", err.Error())
	}
	return userData, nil
}
