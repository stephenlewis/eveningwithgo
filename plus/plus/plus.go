package hello

import (
	"appengine"
	"appengine/urlfetch"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func init() {
	http.HandleFunc("/", handler)
}

const user = "<user_id>"
const key = "<key>"

type Person struct {
	DisplayName string
}

func handler(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	client := urlfetch.Client(c)

	url := fmt.Sprintf("https://www.googleapis.com/plus/v1/people/%s?key=%s", user, key)
	resp, err := client.Get(url)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	person := Person{}
	if err := json.Unmarshal(body, &person); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "text/plain")
	fmt.Fprintln(w, person.DisplayName)
}
