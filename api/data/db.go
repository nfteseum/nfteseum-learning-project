package data

import "time"

var ()

func init() {
	// Set UTC timezone for all our data models & ignore TZ coming from OS env.
	time.Local = time.UTC
}

func PrepareDB() {}

// something
