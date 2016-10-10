package spiderdata

// Request contains all data given in a single request from a spider to the
// filling server
type Request struct {
	Cinemas []*Cinema `json:"cinemas"`
	Films   []*Film   `json:"films"`
}
