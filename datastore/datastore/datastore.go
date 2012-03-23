package datastore

import (
	"appengine"
	"appengine/datastore"
	"fmt"
	"net/http"
)

type Person struct {
	FirstName string
	Surname   string
}

func init() {
	http.HandleFunc("/write", write)
	http.HandleFunc("/query", query)
}

func write(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

	p := Person{
		FirstName: r.FormValue("first_name"),
		Surname:   r.FormValue("surname"),
	}

	key, err := datastore.Put(c, datastore.NewIncompleteKey(c, "person", nil), &p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Written with key %v\n", key)
}

func query(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	q := datastore.NewQuery("person")

	w.Header().Add("Content-Type", "text/plain")

	for t := q.Run(c); ; {
		var p Person
		key, err := t.Next(&p)
		if err == datastore.Done {
			break
		}
		fmt.Fprintf(w, "%v: %s %s\n", key, p.FirstName, p.Surname)
	}
}
