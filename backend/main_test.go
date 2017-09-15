package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDBConnection(t *testing.T) {
	t.Log("Testing database connection")
	db := NewDB()
	session, err := db.Connect()
	if err != nil {
		t.Fatal(err)
	}
	session.Close()
}

func TestListDownloads(t *testing.T) {
	req, err := http.NewRequest("GET", "/downloads?startDate=2017-09-01T12:00:00-03:00&endDate=2017-09-04T23:59:59-03:00", nil)
	if err != nil {
		t.Fatal(err)
	}
	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(ListDownloads)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestDownloadsByCountry(t *testing.T) {
	req, err := http.NewRequest("GET", "/downloads/by-country?startDate=2017-09-01T12:00:00-03:00&endDate=2017-09-04T23:59:59-03:00", nil)
	if err != nil {
		t.Fatal(err)
	}
	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(DownloadsByCountry)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestDownloadsByTime(t *testing.T) {
	req, err := http.NewRequest("GET", "/downloads/by-time?startDate=2017-09-01T12:00:00-03:00&endDate=2017-09-04T23:59:59-03:00", nil)
	if err != nil {
		t.Fatal(err)
	}
	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(DownloadsByTime)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestDownloadsByAppID(t *testing.T) {
	req, err := http.NewRequest("GET", "/downloads/by-time?startDate=2017-09-01T12:00:00-03:00&endDate=2017-09-04T23:59:59-03:00", nil)
	if err != nil {
		t.Fatal(err)
	}
	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(DownloadsByAppID)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}
