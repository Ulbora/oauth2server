package oauthclient

import (
	"fmt"
	"net/http"
	"testing"

	cp "github.com/Ulbora/GoAuth2/compresstoken"
	m "github.com/Ulbora/GoAuth2/managers"
	lg "github.com/Ulbora/Level_Logger"
)

func TestOauthClient_Authorize(t *testing.T) {
	var man m.MockManager
	man.MockValidateAccessTokenSuccess = true

	var oc OauthClient
	oc.Manager = &man
	var l lg.Logger
	oc.Log = &l
	//oc.TokenCompressed = true
	c := oc.GetNewClient()
	var cl Claim
	cl.Role = "testRole"
	cl.URL = "testURL"
	cl.Scope = "web"
	r, _ := http.NewRequest("GET", "/testurl", nil)
	r.Header.Set("Authorization", "Bearer jdljdfldjslkjdslkldksldfks")
	r.Header.Set("hashed", "true")
	r.Header.Set("clientId", "22")
	r.Header.Set("userId", "lfo")

	suc := c.Authorize(r, &cl)
	fmt.Println("suc", suc)
	if !suc {
		t.Fail()
	}
}

func TestOauthClient_AuthorizeCompressed(t *testing.T) {
	var man m.MockManager
	man.MockValidateAccessTokenSuccess = true

	var oc OauthClient
	oc.Manager = &man
	oc.TokenCompressed = true
	var l lg.Logger
	oc.Log = &l
	c := oc.GetNewClient()
	var cl Claim
	cl.Role = "testRole"
	cl.URL = "testURL"
	cl.Scope = "web"
	var token = "jdljdfldjslkjdsdfgdfgdffgdfgfdfgdfgdfgdfgdfdfdfdfdfdfdfdfgdgdfgdffgdfgdfdfgfdfgdfdfgddddgdgdgdgdgdgdgddggdgdgdgdggdfgdfgdfgdgflkldksldfks"
	fmt.Println("len of token: ", len(token))
	var jc cp.JwtCompress
	tkn := jc.CompressJwt(token)
	fmt.Println("tkn", tkn)
	fmt.Println("len of compressed token: ", len(tkn))
	r, _ := http.NewRequest("GET", "/testurl", nil)
	r.Header.Set("Authorization", "Bearer "+tkn)
	r.Header.Set("hashed", "true")
	r.Header.Set("clientId", "22")
	r.Header.Set("userId", "lfo")

	suc := c.Authorize(r, &cl)
	fmt.Println("suc", suc)
	if !suc {
		t.Fail()
	}
}
