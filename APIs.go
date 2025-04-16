package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"unicode"
)

// isNumeric checks if a string is a valid numeric value
func isNumeric(s string) bool {
	for _, char := range s {
		if !unicode.IsDigit(char) {
			return false
		}
	}
	return true
}

// GreetHandler handles GET requests with a query parameter
func GreetHandler(w http.ResponseWriter, r *http.Request) {

	/*
		fmt.Println(r)

		fmt.Printf("\n")

		fmt.Println(w)

		r.ParseForm()       // parse arguments, you have to call this by yourself
		fmt.Println(r.Form) // print form information in server side
		fmt.Println("path", r.URL.Path)
		fmt.Println("scheme", r.URL.Scheme)
		fmt.Println(r.Form["url_long"])
		for k, v := range r.Form {
			fmt.Println("key:", k)
			fmt.Println("val:", strings.Join(v, ""))
		}
	*/

	if r.Method != http.MethodGet {
		http.Error(w, "Only GET method is allowed", http.StatusMethodNotAllowed)
		return
	}

	// Read query parameters
	name := r.URL.Query().Get("name")
	age := r.URL.Query().Get("age")

	// Scenario 1: No parameters provided
	if name == "" && age == "" {
		http.Error(w, "You provided nothing.", http.StatusBadRequest)
		return
	}

	// Scenario 2: Name provided but age is missing
	if name != "" && age == "" {
		http.Error(w, "Age is missing.", http.StatusBadRequest)
		return
	}

	// Scenario 3: Both name and age are provided but age not in correct format
	if name != "" && age != "" && !isNumeric(age) {
		http.Error(w, "Invalid age provided. Age must be a number.", http.StatusBadRequest)
		return
	}

	// Scenario 4: Both name and age are provided
	if name != "" && isNumeric(age) {

		// Convert age to an integer
		age_int, err := strconv.Atoi(age)
		if err != nil {
			http.Error(w, "Error converting age to number.", http.StatusInternalServerError)
			return
		}

		if age_int > 18 {

			// Set headers before WriteHeader then WriteHeader
			w.Header().Set("Content-Type", "application/json")
			w.Header().Set("Custom-Header", "Sadat") // no space in key

			w.WriteHeader(201) //w.WriteHeader(http.StatusOK)

			response := map[string]string{
				"message": fmt.Sprintf("Hello, %s! Your age is %s. You are an adult.", name, age),
			}
			// Writting json file to response
			json.NewEncoder(w).Encode(response)

		} else {
			response := fmt.Sprintf("Hello, %s!\nYour age is %s\nYou are a minor.\n", name, age)
			fmt.Fprintln(w, response)
		}

		return
	}

	response := fmt.Sprintf("You didn't get caught.\n")
	fmt.Fprintln(w, response)

	return

}

// EchoHandler handles POST requests with JSON input
func EchoHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	// Define a struct to parse JSON input
	var requestData struct {
		Message string `json:"message"`
	}

	// Decode JSON from request body
	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Send JSON response
	responseData := map[string]string{"echo": requestData.Message}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(responseData)

}

func main() {
	http.HandleFunc("/get", GreetHandler)
	http.HandleFunc("/post", EchoHandler)

	port := "8080"
	fmt.Printf("Server running on port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))

}
