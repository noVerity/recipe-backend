package main

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetupUserRoutes(t *testing.T) {
	client, requestTester := SetupTestORM(t)
	defer client.Close()

	userRoute := "/user"
	loginRoute := "/login"

	// Invalid payload: Invalid JSON
	w := requestTester(
		http.MethodPost,
		userRoute,
		false,
		`{
			"user""Frodo",
			"email":"frodo@shire.me",
			"password":"myprecious"
		}`,
	)

	assert.Equal(t, http.StatusBadRequest, w.Code)

	// Valid user registering
	w = requestTester(
		http.MethodPost,
		userRoute,
		false,
		`{
			"username":"Frodo",
			"email":"frodo@shire.me",
			"password":"myprecious"
		}`,
	)

	assert.Equal(t, http.StatusOK, w.Code)

	// Valid user getting user data
	w = requestTester(
		http.MethodGet,
		userRoute,
		true,
		``,
	)

	assert.Equal(t, http.StatusOK, w.Code)

	// Same user trying to register again
	w = requestTester(
		http.MethodPost,
		userRoute,
		false,
		`{
			"username":"Frodo",
			"email":"frodo@shire.me",
			"password":"myprecious"
		}`,
	)

	assert.Equal(t, http.StatusConflict, w.Code)
	assert.Equal(t, `{"error":"User or email already taken"}`, w.Body.String())

	// Valid user logging in
	w = requestTester(
		http.MethodPost,
		loginRoute,
		false,
		`{
			"username":"Frodo",
			"password":"myprecious"
		}`,
	)

	assert.Equal(t, http.StatusOK, w.Code)

	// Wrong password
	w = requestTester(
		http.MethodPost,
		loginRoute,
		false,
		`{
			"username":"Frodo",
			"password":"sam4ever"
		}`,
	)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
	assert.Equal(t, `{"error":"Invalid username/password"}`, w.Body.String())

	// Invalid payload: Invalid JSON
	w = requestTester(
		http.MethodPost,
		loginRoute,
		false,
		`{
			"user""Frodo",
			"password":"myprecious"
		}`,
	)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, "{\"error\":\"invalid character '\\\"' after object key\"}", w.Body.String())

}
