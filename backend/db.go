package main

import (
	r "gopkg.in/gorethink/gorethink.v3"
)

// DB DB class
type DB struct {
	ConnectOpts r.ConnectOpts
}

// NewDB DB constructor
func NewDB() *DB {
	connectOpts := r.ConnectOpts{Address: "localhost:28015", Database: "Downloads"}
	return &DB{
		ConnectOpts: connectOpts,
	}
}

// Connect Connects to the database
func (db *DB) Connect() (*r.Session, error) {
	session, err := r.Connect(r.ConnectOpts{
		Address:  "localhost:28015",
		Database: "Downloads",
	})

	return session, err
}
