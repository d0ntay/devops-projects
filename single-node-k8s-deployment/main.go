package main

import (
        "fmt"
        "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        fmt.Fprint(w, "{\"message\" : \"Hello from the backend ;D\"}")
}

func main() {
        http.HandleFunc("/api/hello", handler)
        fmt.Println("Server listening on port 8080")
        http.ListenAndServe(":8080", nil)
}
