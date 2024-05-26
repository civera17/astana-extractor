package client

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAllProjectsSuccess(t *testing.T) {
	testProjects := `{"data":[{"gid":"1207408218652434","name":"test-proj21","resource_type":"project"},{"gid":"1207408218651330","name":"project-test","resource_type":"project"}]}`
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(testProjects))
	}))

	c := NewAstanaClient(server.URL)

	projs, err := c.GetAllProjects("10")
	if err != nil {
		t.Error()
	}
	if len(projs) == 0 {
		t.Error()
	}
}

func TestGetAllProjectsFail(t *testing.T) {
	testProjects := `{"data":[{"gid":"1207408218652434","name":"test-proj21","resource_type":"project"},{"gid":"1207408218651330","name":"project-test","resource_type":"project"}]}`
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(testProjects))
	}))

	c := NewAstanaClient(server.URL)

	_, err := c.GetAllProjects("10")
	if err == nil {
		t.Error()
	}
	assert.ErrorContains(t, err, "request did not respond with 200")
}

func TestGetAllUsersSuccess(t *testing.T) {
	testUsers := `{"data":[{"gid":"1207408218652434","email":"test-proj21","resource_type":"user"},{"gid":"1207408218651330","email":"project-test","resource_type":"user"}]}`
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(testUsers))
	}))

	c := NewAstanaClient(server.URL)

	projs, err := c.GetAllProjects("10")
	assert.NoError(t, err)
	assert.Len(t, projs, 2)
}

func TestGetAllUsersFail(t *testing.T) {
	testUsers := `{"data":[{"gid":"1207408218652434","email":"test-proj21","resource_type":"user"},{"gid":"1207408218651330","email":"project-test","resource_type":"user"}]}`
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(testUsers))
	}))

	c := NewAstanaClient(server.URL)

	_, err := c.GetAllProjects("10")
	if err == nil {
		t.Error()
	}
	assert.ErrorContains(t, err, "request did not respond with 200")
}
