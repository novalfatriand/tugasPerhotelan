package handlers

import (
    "encoding/json"
    "net/http"
    "hotel-booking/models"
    "strconv"
)

func AdminDashboard(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(Bookings)
}

func ConfirmPayment(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    bookingIDStr := r.URL.Query().Get("booking_id")
    bookingID, err := strconv.Atoi(bookingIDStr)
    if err != nil {
        http.Error(w, "Invalid booking ID", http.StatusBadRequest)
        return
    }

    for i, booking := range Bookings {
        if booking.ID == bookingID {
            Bookings[i].Confirmed = true
            saveBookings() // Simpan perubahan ke JSON
            json.NewEncoder(w).Encode(map[string]string{"message": "Pembayaran dikonfirmasi"})
            return
        }
    }

    http.Error(w, "Booking not found", http.StatusNotFound)
}
