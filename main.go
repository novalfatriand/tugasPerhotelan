package main

import (
    "net/http"
    "hotel-booking/handlers"
)

func main() {
    // Memuat data dari JSON ke dalam aplikasi
    handlers.LoadHotels()
    handlers.LoadUsers()
    handlers.LoadBookings()

    // Routing untuk halaman utama login
    http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, "static/login.html")
    })

    // Routing untuk halaman daftar hotel
    http.HandleFunc("/hotels", func(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, "static/hotels.html")
    })

    // Routing untuk halaman detail hotel
    http.HandleFunc("/hotel/", func(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, "static/hotel_detail.html")
    })

    // Routing untuk halaman booking
    http.HandleFunc("/booking", func(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, "static/booking.html")
    })

    // Routing untuk halaman payment
    http.HandleFunc("/payment", func(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, "static/payment.html")
    })

    // Routing untuk dashboard admin
    http.HandleFunc("/admin/dashboard", func(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, "static/admin_dashboard.html")
    })

    // API untuk autentikasi
    http.HandleFunc("/auth/login", handlers.Login)

    // API untuk CRUD hotel dan detail hotel
    http.HandleFunc("/api/hotels", handlers.GetHotels)
    http.HandleFunc("/api/hotel/", handlers.GetHotelDetails)

    // API untuk booking hotel
    http.HandleFunc("/api/booking", handlers.Booking)

    // API untuk payment
    http.HandleFunc("/api/payment", handlers.Payment)

    // API untuk admin dashboard
    http.HandleFunc("/api/admin/dashboard", handlers.AdminDashboard)
    http.HandleFunc("/api/admin/confirm_payment", handlers.ConfirmPayment)

    // Menjalankan server pada port 8080
    http.ListenAndServe(":8080", nil)
}
