package routing_test

import "testing"
import "net/http"
import "net/http/httptest"

import "github.com/hinrikg/gophercon/pkg/routing"

func TestCreateRouter(t *testing.T) {
	router := routing.CreateRouter()

	testServer := httptest.NewServer(router)
	defer testServer.Close()

	response, error := http.Get(testServer.URL + "/home")
	if error != nil {
		t.Fatal(error)
	}

	if response.StatusCode != http.StatusOK {
		t.Errorf("Wrong status code %d (%d expected)", response.StatusCode, http.StatusOK)
	}
}
