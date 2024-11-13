package handlers

import (
    "encoding/json"
    "net/http"
    "os"
    "hotel-booking/models"
    "fmt"
)

var Users []models.User

func LoadUsers() {
    file, _ := os.Open("data/users.json")
    defer file.Close()
    decoder := json.NewDecoder(file)
    decoder.Decode(&Users)
}

func Login(w http.ResponseWriter, r *http.Request) {
    var credentials models.User
    json.NewDecoder(r.Body).Decode(&credentials)

    for _, user := range Users {
        if user.Username == credentials.Username && user.Password == credentials.Password {
            http.SetCookie(w, &http.Cookie{
                Name:  "userID",
                Value: fmt.Sprint(user.ID),
                Path:  "/",
            })

            if user.Role == "admin" {
                http.Redirect(w, r, "/admin/dashboard", http.StatusSeeOther)
            } else {
                http.Redirect(w, r, "/hotels", http.StatusSeeOther)
            }
            return
        }
    }
    http.Error(w, "Invalid credentials", http.StatusUnauthorized)
}
