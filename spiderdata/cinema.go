package spiderdata

import "time"

// Cinema represents a single cinema (the building)
type Cinema struct {
	Name     string    `json:"name"`
	URL      string    `json:"url"`
	Chain    string    `json:"chain"` // e.g. "Epic Cinemas"
	ID       string    `json:"id"`
	Acquired time.Time `json:"acquired"`
}
