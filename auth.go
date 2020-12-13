package main

import (
	"github.com/gastrodon/groudon"
	"git.gastrodon.io/imonke/monkebase"
	"git.gastrodon.io/imonke/monketype"

	"net/http"
)

var (
	badAuth map[string]interface{} = map[string]interface{}{"error": "bad_auth"}
)

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

func postAuth(request *http.Request) (code int, r_map map[string]interface{}, err error) {
	var body AuthBody
	var external error
	if err, external = groudon.SerializeBody(request.Body, &body); err != nil || external != nil || (body.Secret == nil && body.Password == nil) {
		code = 400
		return
	}

	var who monketype.User
	var exists bool
	if who, exists, err = monkebase.ReadSingleUserEmail(body.Email); err != nil {
		return
	}

	var authed bool
	switch {
	case !exists:
	case body.Secret != nil && *body.Secret != "":
		authed, err = monkebase.CheckSecret(who.ID, *body.Secret)
	case body.Password != nil && *body.Password != "":
		authed, err = monkebase.CheckPassword(who.ID, *body.Password)
	}

	if authed && err == nil {
		code = 200
		r_map, err = rotatedTokens(who.ID)
		return
	}

	code = 401
	return
}
