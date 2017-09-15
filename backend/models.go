package main

import (
	"errors"
	"time"

	r "gopkg.in/gorethink/gorethink.v3"
)

// Downloads Downloads class
type Downloads struct {
	ID           string    `json:"id" gorethink:"id,omitempty"`
	Latitude     float32   `json:"latitude" gorethink:"latitude"`
	Longitude    float32   `json:"longitude" gorethink:"longitude"`
	AppID        string    `json:"app_id" gorethink:"app_id"`
	DownloadedAt time.Time `json:"downloaded_at" gorethink:"downloaded_at"`
	Country      string    `json:"country" gorethink:"country"`
}

// DownloadManager DownloadsManager class, uses Downloads
type DownloadManager struct{}

// SaveDownload Saves a Download object on the database
func (d *DownloadManager) SaveDownload(*Downloads) error {
	return errors.New("Not implemented")
}

// ListDownloads List the downloads between startDate and endDate
func (d *DownloadManager) ListDownloads(startDate time.Time, endDate time.Time) (*[]*Downloads, error) {
	db := NewDB()
	session, err := db.Connect()
	if err != nil {
		return nil, err
	}

	defer session.Close()

	cursor, err := r.Table("Downloads").Filter(func(row r.Term) r.Term {
		return row.Field("downloaded_at").During(startDate, endDate)
	}).Run(session)

	defer cursor.Close()
	if err != nil {
		return nil, err
	}

	var downloads []*Downloads
	err = cursor.All(&downloads)
	return &downloads, err
}

// TotalsByCountry List the ammount of downloads by country between startDate and
// endDate
func (d *DownloadManager) TotalsByCountry(startDate time.Time, endDate time.Time) ([]interface{}, error) {
	db := NewDB()
	session, err := db.Connect()
	if err != nil {
		return nil, err
	}

	defer session.Close()

	cursor, err := r.Table("Downloads").Filter(func(row r.Term) r.Term {
		return row.Field("downloaded_at").During(startDate, endDate)
	}).Group("country").Ungroup().Map(func(group r.Term) interface{} {
		return map[string]interface{}{
			"country": group.Field("group"),
			"total":   group.Field("reduction").Count(),
		}
	}).Run(session)

	defer cursor.Close()
	if err != nil {
		return nil, err
	}

	var countries []interface{}
	err = cursor.All(&countries)

	return countries, err
}

// TotalsByTime List the ammount of downloads by time of the day between startDate and
// endDate
func (d *DownloadManager) TotalsByTime(startDate time.Time, endDate time.Time) ([]interface{}, error) {
	db := NewDB()
	session, err := db.Connect()
	if err != nil {
		return nil, err
	}

	defer session.Close()

	cursor, err := r.Table("Downloads").Filter(func(row r.Term) r.Term {
		return row.Field("downloaded_at").During(startDate, endDate)
	}).Group(func(row r.Term) r.Term {
		time := row.Field("downloaded_at").Hours()
		return r.Branch(
			time.Ge(5).And(time.Lt(12)),
			"Morning",
			time.Ge(12).And(time.Lt(17)),
			"Afternoon",
			time.Ge(17).And(time.Lt(21)),
			"Evening",
			time.Ge(21).Or(time.Lt(5)),
			"Night",
			nil)
	}).Ungroup().Map(func(group r.Term) interface{} {
		return map[string]interface{}{
			"time":  group.Field("group"),
			"total": group.Field("reduction").Count(),
		}
	}).Run(session)

	defer cursor.Close()
	if err != nil {
		return nil, err
	}

	var paises []interface{}
	err = cursor.All(&paises)

	return paises, err
}

// TotalsByAppID List the ammount of downloads by app_id between startDate and
// endDate
func (d *DownloadManager) TotalsByAppID(startDate time.Time, endDate time.Time) ([]interface{}, error) {
	db := NewDB()
	session, err := db.Connect()
	if err != nil {
		return nil, err
	}

	defer session.Close()

	cursor, err := r.Table("Downloads").Filter(func(row r.Term) r.Term {
		return row.Field("downloaded_at").During(startDate, endDate)
	}).Group("app_id").Ungroup().Map(func(group r.Term) interface{} {
		return map[string]interface{}{
			"app_id": group.Field("group"),
			"total":  group.Field("reduction").Count(),
		}
	}).Run(session)

	defer cursor.Close()
	if err != nil {
		return nil, err
	}

	var apps []interface{}
	err = cursor.All(&apps)

	return apps, err
}
