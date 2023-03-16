package main

import (
	"html/template"
	"net/http"
	"strconv"
)

type Contact struct {
	Id        string
	FirstName string
	LastName  string
	Email     string
	Phone     string
	Address   string
	City      string
}

type ContactsListData struct {
	Contacts []Contact
}

type FormData struct {
	FormName string
}

func main() {
	contactList := make([]Contact, 0)
	id := 100

	list_template := template.Must(template.ParseFiles("./views/list.html"))
	form_template := template.Must(template.ParseFiles("./views/form.html"))
	details_template := template.Must(template.ParseFiles("./views/details.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			contact := Contact{
				Id:        strconv.FormatInt(int64(id), 10),
				FirstName: r.FormValue("firstname"),
				LastName:  r.FormValue("lastname"),
				Email:     r.FormValue("email"),
				Phone:     r.FormValue("phone"),
				Address:   r.FormValue("address"),
				City:      r.FormValue("city"),
			}

			id++
			contactList = append(contactList, contact)
		}

		data := ContactsListData{Contacts: contactList}

		list_template.Execute(w, data)
	})

	http.HandleFunc("/form", func(w http.ResponseWriter, r *http.Request) {
		data := FormData{FormName: "Please enter a new contact"}
		form_template.Execute(w, data)
	})

	http.HandleFunc("/contact/", func(w http.ResponseWriter, r *http.Request) {
		var data Contact
		temp_string := r.URL.String()
		temp_id := temp_string[len(temp_string)-3:] // ottaa kolme viimeistä merkkiä, tämä kannattaisi hanskata vaikka regexpillä jos oikeasti tekisi, nyt ei toimi jos lopussa vaikka kysymysmerkki
		for _, contact := range contactList {
			if contact.Id == temp_id {
				data = contact
			}

			details_template.Execute(w, data)
		}

	})

	http.ListenAndServe(":3000", nil)
}
