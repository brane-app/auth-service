package main

import (
	"github.com/imonke/monkebase"
	"github.com/imonke/monketype"

	"encoding/json"
	"io/ioutil"
	"net/http"
)

var (
	badRequest map[string]interface{} = map[string]interface{}{"error": "bad_request"}
	badAuth    map[string]interface{} = map[string]interface{}{"error": "bad_auth"}
)

type AuthBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Secret   string `json:"secret"`
}

func getBody(request *http.Request) (body AuthBody, err error) {
	if request.Close {
		defer request.Body.Close()
	}

	if request.Body == nil {
		return
	}

	var bytes []byte
	if bytes, err = ioutil.ReadAll(request.Body); err != nil {
		return
	}

	json.Unmarshal(bytes, &body)
	return
}

func rotatedTokens(id string) (r_map map[string]interface{}, err error) {
	var token string
	var expires int64
	if token, expires, err = monkebase.CreateToken(id); err != nil {
		return
	}

	var secret string
	if secret, err = monkebase.CreateSecret(id); err != nil {
		return
	}

	r_map = map[string]interface{}{
		"auth": map[string]interface{}{
			"token":   token,
			"expires": expires,
			"secret":  secret,
		},
	}

	return
}

func authenticate(body AuthBody) (code int, r_map map[string]interface{}, err error) {
	var who monketype.User
	var exists bool
	if who, exists, err = monkebase.ReadSingleUserEmail(body.Email); err != nil {
		return
	}

	var authed bool = false

	switch {
	case !exists:
	case body.Secret != "":
		authed, err = monkebase.CheckSecret(who.ID, body.Secret)
	case body.Password != "":
		authed, err = monkebase.CheckPassword(who.ID, body.Password)
	}

	if authed && err == nil {
		r_map, err = rotatedTokens(who.ID)
		code = 200
		return
	}

	code = 401
	r_map = badAuth
	return
}

func postAuth(request *http.Request) (code int, r_map map[string]interface{}, err error) {
	var body AuthBody
	body, err = getBody(request)

	switch {
	case err != nil:
	case body.Password == "" && body.Secret == "", body.Email == "":
		code = 400
		r_map = badRequest
	default:
		code, r_map, err = authenticate(body)
	}

	return
}
