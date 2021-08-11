package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetupUserRoutes(t *testing.T) {
	client := SetupTestRouter(t)
	defer client.Close()

	router := SetupRouter(client)
	requestTester := GetJSONRequestTester(router)

	// Invalid payload: Invalid JSON
	w := requestTester("POST", "/user/", `{"user":"Frodo,"email":"frodo@shire.me","password":"myprecious"}`)

	assert.Equal(t, 400, w.Code)
	assert.Equal(t, `{"error":"invalid character 'e' after object key:value pair"}`, w.Body.String())

	// Valid user registering
	w = requestTester("POST", "/user/", `{"username":"Frodo","email":"frodo@shire.me","password":"myprecious"}`)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, `{"username":"Frodo","email":"frodo@shire.me"}`, w.Body.String())

	// Same user trying to register again
	w = requestTester("POST", "/user/", `{"username":"Frodo","email":"frodo@shire.me","password":"myprecious"}`)

	assert.Equal(t, 409, w.Code)
	assert.Equal(t, `{"error":"User or email already taken"}`, w.Body.String())

	// Valid user logging in
	w = requestTester("POST", "/login", `{"username":"Frodo","password":"myprecious"}`)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, `{"username":"Frodo","email":"frodo@shire.me"}`, w.Body.String())

	// Wrong password
	w = requestTester("POST", "/login", `{"username":"Frodo","password":"sam4ever"}`)

	assert.Equal(t, 401, w.Code)
	assert.Equal(t, `{"error":"Invalid username/password"}`, w.Body.String())

	// Invalid payload: Invalid JSON
	w = requestTester("POST", "/login", `{"user":"Frodo,"password":"myprecious"}`)

	assert.Equal(t, 400, w.Code)
	assert.Equal(t, `{"error":"invalid character 'p' after object key:value pair"}`, w.Body.String())

}
