package pdl

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func beforeTest() {
	muxAPI := http.NewServeMux()
	testAPIServer := httptest.NewServer(muxAPI)
	defer testAPIServer.Close()

	muxAPI.HandleFunc("/users/1", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "testdata/users-1.json")
	})
}

func TestNormal(t *testing.T) {
	ts := httptest.NewServer(sampleHandler)
	defer ts.Close()

	// リクエストの送信先はテストサーバのURLへ。
	r, err := http.Get(ts.URL)
	if err != nil {
		t.Fatalf("Error by http.Get(). %v", err)
	}

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		t.Fatalf("Error by ioutil.ReadAll(). %v", err)
	}

	if "Hello HTTP Test" != string(data) {
		t.Fatalf("Data Error. %v", string(data))
	}
}
