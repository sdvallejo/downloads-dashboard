package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

// ListDownloads List the downloads between startDate and endDate
func ListDownloads(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	startDate, err := time.Parse(time.RFC3339, r.URL.Query().Get("startDate"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	endDate, err := time.Parse(time.RFC3339, r.URL.Query().Get("endDate"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	dm := DownloadManager{}
	downloads, err := dm.ListDownloads(startDate, endDate)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Panic(err.Error())
		return
	}

	data, err := json.Marshal(downloads)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Panic(err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

// DownloadsByCountry List the ammount of downloads by country between startDate and
// endDate
func DownloadsByCountry(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	startDate, err := time.Parse(time.RFC3339, r.URL.Query().Get("startDate"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	endDate, err := time.Parse(time.RFC3339, r.URL.Query().Get("endDate"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	dm := DownloadManager{}
	countries, err := dm.TotalsByCountry(startDate, endDate)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Panic(err.Error())
		return
	}

	data, err := json.Marshal(countries)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Panic(err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

// DownloadsByTime List the ammount of downloads by time of the day between startDate and
// endDate
func DownloadsByTime(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	startDate, err := time.Parse(time.RFC3339, r.URL.Query().Get("startDate"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	endDate, err := time.Parse(time.RFC3339, r.URL.Query().Get("endDate"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	dm := DownloadManager{}
	groups, err := dm.TotalsByTime(startDate, endDate)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Panic(err.Error())
		return
	}

	data, err := json.Marshal(groups)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Panic(err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

// DownloadsByAppID List the ammount of downloads by app_id between startDate and
// endDate
func DownloadsByAppID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	startDate, _ := time.Parse(time.RFC3339, r.URL.Query().Get("startDate"))
	endDate, _ := time.Parse(time.RFC3339, r.URL.Query().Get("endDate"))

	dm := DownloadManager{}
	groups, err := dm.TotalsByAppID(startDate, endDate)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Panic(err.Error())
		return
	}

	data, err := json.Marshal(groups)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Panic(err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}
