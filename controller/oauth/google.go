package controller

import (
	"encoding/json"
	"fmt"
	"github.com/dchest/uniuri"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"io/ioutil"
	"net/http"
	"os"
)

//GoogleUser struct is here
type GoogleUser struct {
	ID            string `json:"id"`
	Email         string `json:"email"`
	VerifiedEmail bool   `json:"verified_email"`
	Name          string `json:"name"`
	GivenName     string `json:"given_name"`
	FamilyName    string `json:"family_name"`
	Link          string `json:"link"`
	Picture       string `json:"picture"`
	Gender        string `json:"gender"`
	Locale        string `json:"locale"`
}

var googleOauthConfig = &oauth2.Config{
	RedirectURL:  "http://localhost:5000/api/v1/oauth/google/callback",
	ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
	ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
	Scopes: []string{
		"https://www.googleapis.com/auth/userinfo.profile",
		"https://www.googleapis.com/auth/userinfo.email"},
	Endpoint: google.Endpoint,
}

// OauthHandler is here
func OauthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "<a href='/login'>Log in with Google</a>")
}

// GloginHandler is here
func GloginHandler(w http.ResponseWriter, r *http.Request) {
	oauthStateString := uniuri.New()
	url := googleOauthConfig.AuthCodeURL(oauthStateString)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

// GcallbackHandler is here
func GcallbackHandler(w http.ResponseWriter, r *http.Request) {
	var user *GoogleUser
	code := r.FormValue("code")
	token, _ := googleOauthConfig.Exchange(oauth2.NoContext, code)
	response, _ := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)

	contents, _ := ioutil.ReadAll(response.Body)
	_ = json.Unmarshal(contents, &user)
	fmt.Fprintf(w, "Email: %s\nName: %s\nImage link: %s\n", user.Email, user.Name, user.Picture)
}
