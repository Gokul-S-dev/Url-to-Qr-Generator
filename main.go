package main

import (
	"log"
	"net/http"

	"github.com/skip2/go-qrcode"
)

func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}
		
		next.ServeHTTP(w, r)
	})
}


func generateQrHandler(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodGet {
		http.Error(w,"Send get method",http.StatusBadRequest)
	}

	url := r.URL.Query().Get("url")
	if url == "" {
	       http.Error(w,"No url passed",http.StatusBadRequest)
	}
	png , err := qrcode.Encode(url,qrcode.Medium,256)

	if err != nil {
		http.Error(w,"Error occured while generating QR code",http.StatusBadRequest)
	}
	w.Header().Set("Content-Type","image/png")
	w.WriteHeader(http.StatusOK)
	w.Write(png)
}
func main() {
	http.Handle("/qr", enableCORS(http.HandlerFunc(generateQrHandler)))

	log.Println("Application is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080",nil))

}