package handlers

import (
    "encoding/json"
    "net/http"
    "os"
    "strconv"
    "hotel-booking/models"
)

var Bookings []models.Booking

// Fungsi untuk memuat data booking dari JSON
func LoadBookings() {
    file, _ := os.Open("data/bookings.json")
    defer file.Close()
    decoder := json.NewDecoder(file)
    decoder.Decode(&Bookings)
}

// Handler untuk booking hotel
func Booking(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    var newBooking models.Booking

    if err := json.NewDecoder(r.Body).Decode(&newBooking); err != nil {
        http.Error(w, "Invalid booking data", http.StatusBadRequest)
        return
    }

    newBooking.ID = len(Bookings) + 1
    totalPrice := calculateTotalPrice(newBooking.HotelID, newBooking.CheckIn, newBooking.CheckOut)
    newBooking.TotalPrice = totalPrice
    newBooking.Confirmed = false

    Bookings = append(Bookings, newBooking)

    // Simpan booking ke file JSON
    saveBookings()

    json.NewEncoder(w).Encode(newBooking)
}

// Fungsi untuk menghitung total harga berdasarkan hotel dan tanggal
func calculateTotalPrice(hotelID int, checkIn, checkOut string) int {
    for _, hotel := range Hotels {
        if hotel.ID == hotelID {
            // Harga per malam * jumlah malam (dianggap 1 malam untuk contoh ini)
            return hotel.PricePerNight
        }
    }
    return 0
}

// Fungsi untuk menyimpan booking ke file JSON
func saveBookings() {
    file, _ := os.Create("data/bookings.json")
    defer file.Close()
    encoder := json.NewEncoder(file)
    encoder.Encode(Bookings)
}
