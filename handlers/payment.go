package handlers

import (
    "encoding/json"
    "net/http"
)

func Payment(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]string{"message": "Pembayaran berhasil dikonfirmasi"})
}
