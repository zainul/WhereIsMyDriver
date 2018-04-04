package api

// GetDriver struct for get driver response
type GetDriver struct {
	Distance  float64 `json:"distance"`
	ID        int     `json:"id"`
	Latitude  float32 `json:"latitude" gorm:"column:current_latitude"`
	Longitude float32 `json:"longitude" gorm:"column:current_longitude"`
}

// UpdateLocation struct for update driver location
type UpdateLocation struct {
	Latitude  float32 `json:"latitude"`
	Longitude float32 `json:"longitude"`
	Accuracy  float32 `json:"accuracy"`
}
