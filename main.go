package main

import (
	"log"
	"net/http"

	"github.com/skip2/go-qrcode"
)
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
	http.HandleFunc("/qr",generateQrHandler)

	log.Println("Application is running on http://localhost:8000")
	log.Fatal(http.ListenAndServe(":8080",nil))

}