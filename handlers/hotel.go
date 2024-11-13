package handlers

import (
    "encoding/json"
    "net/http"
    "os"
    "strconv"
    "strings"
    "hotel-booking/models"
)

var Hotels []models.Hotel

// Fungsi untuk memuat data hotel dari JSON
func LoadHotels() {
    file, _ := os.Open("data/hotels.json")
    defer file.Close()
    decoder := json.NewDecoder(file)
    decoder.Decode(&Hotels)
}

// Handler untuk mendapatkan daftar hotel
func GetHotels(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(Hotels)
}

// Handler untuk mendapatkan detail hotel berdasarkan ID
func GetHotelDetails(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    idStr := strings.TrimPrefix(r.URL.Path, "/api/hotel/")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        http.Error(w, "Invalid hotel ID", http.StatusBadRequest)
        return
    }

    for _, hotel := range Hotels {
        if hotel.ID == id {
            json.NewEncoder(w).Encode(hotel)
            return
        }
    }

    http.Error(w, "Hotel not found", http.StatusNotFound)
}
