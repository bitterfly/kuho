package spiderdata

import "time"

// Cinema represents a single cinema (the building)
type Cinema struct {
	Name      string    `json:"name"`
	URL       string    `json:"url"`
	Chain     string    `json:"chain"` // e.g. "Epic Cinemas"
	ShortName string    `json:"short_name"`
	Acquired  time.Time `json:"acquired"`
}
