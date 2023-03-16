package main

import (
	"html/template"
	"net/http"
)

type Item struct {
	Type  string
	Count int
	Price int
}

type ShoppingData struct {
	ShoppingList string
	Items        []Item
}

func main() {
	t := template.Must(template.ParseFiles("./views/layout.html"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := ShoppingData{
			ShoppingList: "MyList",
			Items:        []Item{{Type: "Banana", Count: 5, Price: 3}, {Type: "Apple", Count: 10, Price: 4}, {Type: "Beer", Count: 100, Price: 6}},
		}

		t.Execute(w, data)
	})
	http.ListenAndServe(":3000", nil)
}
