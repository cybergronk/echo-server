package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	// Look up PORT environment variable
	var port string
	{
		if p, ok := os.LookupEnv("PORT"); !ok {
			fmt.Println("No port specified, using app default...")
			port = "36110"
		} else {
			port = p
		}
	}

	// Register routes
	http.HandleFunc("/", returnAliveStatus)
	http.HandleFunc("/echo", echoBody)

	// Begin listening
	fmt.Printf("Now listening on port %s...\n", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Printf("Could not start server: %s\n", err.Error())
		os.Exit(1)
	}
}

// returnAliveStatus returns whether the server is running
func returnAliveStatus(w http.ResponseWriter, r *http.Request) {
	writeMessage(w, http.StatusOK, "I live!")
}

// echoBody returns exactly what it was given in the request body
func echoBody(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		writeMessage(w, http.StatusInternalServerError, fmt.Sprintf("Could not read request body: %s", err.Error()))
		return
	}

	writeBytes(w, http.StatusOK, body)
}

// writeBytes sets Content-Type to JSON and responds with given status code and byte slice
func writeBytes(writer http.ResponseWriter, statusCode int, body []byte) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(statusCode)
	writer.Write(body)
}

// writeMessage sets Content-Type to JSON and responds with given status code and string
// wrapped around a JSON "message" property
func writeMessage(writer http.ResponseWriter, statusCode int, message string) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(statusCode)
	io.WriteString(writer, fmt.Sprintf(`{"message":"%s"}`, message))
}
