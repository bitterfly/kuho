package spiderdata

import "time"

// Film represents a film as seen in the cinema (different versions, e.g. 3D,
// IMAX, 70mm etc are separate films)
type Film struct {
	ID              string       `json:"id"`
	ImdbID          *int         `json:"imdb_id"`
	ImdbIDCertainty float32      `json:"imdb_id_certainty"`
	Title           string       `json:"title"`
	Year            int          `json:"year"`
	URL             string       `json:"url"`
	Screenings      []*Screening `json:"screenings"`
	Rating          string       `json:"rating"`
	Acquired        time.Time    `json:"acquired"`
}

// Screening is a performance of a film in a particular cinema at a particular time
type Screening struct {
	ID              string        `json:"id"`
	CinemaShortName string        `json:"cinema_short_name"`
	Hall            string        `json:"hall"`
	Time            time.Time     `json:"time"`
	Tickets         []*Ticket     `json:"tickets"`
	Active          bool          `json:"active"`
	IsSubtitled     bool          `json:"is_subtitled"`
	IsDubbed        bool          `json:"is_dubbed"`
	Is3D            bool          `json:"is_3d"`
	Is4D            bool          `json:"is_4d"`
	IsImax          bool          `json:"is_imax"`
	Language        string        `json:"language"`
	Variant         string        `json:"variant"`
	Duration        time.Duration `json:"duration"`
	Acquired        time.Time     `json:"acquired"`
}

// Ticket is a ticket for a single performance
type Ticket struct {
	BookingURL string    `json:"booking_url"`
	Price      float32   `json:"price"`
	Currency   string    `json:"currency"`
	Type       string    `json:"type"` // e.g. "elderly", "student", "regular"
	Acquired   time.Time `json:"acquired"`
}
