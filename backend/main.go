package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/downloads", ListDownloads)
	http.HandleFunc("/downloads/by-country", DownloadsByCountry)
	http.HandleFunc("/downloads/by-time", DownloadsByTime)
	http.HandleFunc("/downloads/by-appid", DownloadsByAppID)
	http.ListenAndServe(":8000", nil)
}
