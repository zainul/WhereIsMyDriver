package api

// GetDriver struct for get driver response
type GetDriver struct {
	Distance  int     `json:"distance"`
	ID        int     `json:"id"`
	Latitude  float32 `json:"latitude"`
	Longitude float32 `json:"longitude"`
}

// UpdateLocation struct for update driver location
type UpdateLocation struct {
	Latitude  float32 `json:"latitude"`
	Longitude float32 `json:"longitude"`
	Accuracy  float32 `json:"accuracy"`
}
