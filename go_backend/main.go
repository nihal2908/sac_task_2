package main

import (
    "encoding/json"
    "io/ioutil"
    "net/http"
)

var storedString = "Hello, world!" // Initial value

func getData(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]string{"data": storedString})
}

func postData(w http.ResponseWriter, r *http.Request) {
    var input map[string]string
    body, _ := ioutil.ReadAll(r.Body)
    json.Unmarshal(body, &input)
    if value, ok := input["data"]; ok {
        storedString = value
        w.WriteHeader(http.StatusOK)
        w.Write([]byte("Data updated successfully"))
    } else {
        http.Error(w, "Invalid input", http.StatusBadRequest)
    }
}

func getImage(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "static/image.png") // Ensure image is stored in `static/` folder
}

func main() {
    http.HandleFunc("/data", func(w http.ResponseWriter, r *http.Request) {
        if r.Method == http.MethodGet {
            getData(w, r)
        } else if r.Method == http.MethodPost {
            postData(w, r)
        } else {
            http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        }
    })

    http.HandleFunc("/image", getImage)


    http.ListenAndServe(":4000", nil)
}