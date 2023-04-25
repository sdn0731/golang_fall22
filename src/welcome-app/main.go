package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"
	"encoding/json"
)

type Welcome struct {
	Name string
	Time string
}

// Create JsonResponse Struct
type JsonResponse struct {
	Value1 string `json:"key1"`
	Value2 string `json:"key2"`
	JsonNested JsonNested `json:"jsonNested"`
}

// Create JsonNested Struct
type JsonNested struct {
	NestedValue1 string `json:"nestedKey1"`
	NestedValue2 string `json:"nestedKey2"`
}

// Create Contact Struct
type Contact struct {
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
	Address Address `json:"address"`
	ContactInfo ContactInfo `json:"contactInfo"`
}

// Create Nested Address Struct
type Address struct {
	Street string `json:"street"`
	City string `json:"city"`
}

// Create Nested ContactInfo Struct
type ContactInfo struct {
	Email string `json:"email"`
	Phone string `json:"phone"`
}

func main() {
	welcome := Welcome{"Anonymous", time.Now().Format(time.Stamp)}
	templates := template.Must(template.ParseFiles("templates/welcome-template.html"))

	// Create Json Nested Object
	nested := JsonNested{
		NestedValue1: "first nested value",
		NestedValue2: "second nested value",
	}

	// Create Json Response Object
	jsonResp := JsonResponse{
		Value1: "some Data",
		Value2: "other Data",
		JsonNested: nested,
	}

	// Create ContactInfo Nested Object
	nestedContact := ContactInfo{
		Email: "taylorswift@umgstores.com",
		Phone: "2301301978",
	}

	// Create Address Nested Object
	nestedAddress := Address{
		Street: "23 Cornelia Street",
		City: "New York, New York",
	}

	// Create Contact Response Object
	contactResp := Contact{
		FirstName: "Taylor",
		LastName: "Swift",
		Address: nestedAddress,
		ContactInfo: nestedContact,
	}

	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("static"))))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if name := r.FormValue("name"); name != "" {
			welcome.Name = name
		}
		if err := templates.ExecuteTemplate(w, "welcome-template.html", welcome); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	// Create jsonResponse Handler
	http.HandleFunc("/jsonResponse", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(jsonResp)
	})

	// Create contactResponse Handler
	http.HandleFunc("/contactResponse", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(contactResp)
	})

	fmt.Println("Listening")
	fmt.Println(http.ListenAndServe(":8080", nil))
}