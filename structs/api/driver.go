package api

// GetDriver struct for get driver response
type GetDriver struct {
	Distance  int     `json:"distance"`
	ID        int     `json:"id"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}
