package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
    "encoding/json"
	"io/ioutil"
)

func TestRouter(t *testing.T) {
	r := Router("", "", "")
	ts := httptest.NewServer(r)
	defer ts.Close()

	res, err := http.Get(ts.URL + "/home")
	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != http.StatusOK {
		t.Errorf("Status code for /home is wrong. Have: %d, want: %d.", res.StatusCode, http.StatusOK)
	}

	res, err = http.Post(ts.URL+"/home", "text/plain", nil)
	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != http.StatusMethodNotAllowed {
		t.Errorf("Status code for /home is wrong. Have: %d, want: %d.", res.StatusCode, http.StatusMethodNotAllowed)
	}

	res, err = http.Get(ts.URL + "/not-exists")
	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != http.StatusNotFound {
		t.Errorf("Status code for /home is wrong. Have: %d, want: %d.", res.StatusCode, http.StatusNotFound)
	}
}

func TestHomeHandler(t *testing.T) {
	w := httptest.NewRecorder()
    version := "0.1.0"
	commit := "some test hash"
	repo := "https://github.com/andrewozhegov/k8s-introduction"
	h := home(version, commit, repo)
	h(w, nil)

	resp := w.Result()
	if have, want := resp.StatusCode, http.StatusOK; have != want {
		t.Errorf("Status code is wrong. Have: %d, want: %d.", have, want)
	}

	greeting, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		t.Fatal(err)
	}
	info := struct {
			Version   string `json:"version"`
			Commit    string `json:"commit"`
			Repo      string `json:"repo"`
    }{}
	err = json.Unmarshal(greeting, &info)
	if err != nil {
		t.Fatal(err)
	}
	if info.Version != version {
		t.Errorf("Release version is wrong. Have: %s, want: %s", info.Version, version)
	}
	if info.Repo != repo {
		t.Errorf("Repo name is wrong. Have: %s, want: %s", info.Repo, repo)
	}
	if info.Commit != commit {
		t.Errorf("Commit is wrong. Have: %s, want: %s", info.Commit, commit)
	}
}

func TestHealthHandler(t *testing.T) {}
func TestReadyHandler(t *testing.T) {}

