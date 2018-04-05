package models

import "time"

//HistoryPositionTableName ...
const HistoryPositionTableName = "history_positions"

// HistoryPosition ...
type HistoryPosition struct {
	UserID    int     `gorm:"not null" validate:"required"`
	Latitude  float32 `gorm:"not null" sql:"type:decimal(9,6);"`
	Longitude float32 `gorm:"not null" sql:"type:decimal(9,6);"`
	Accuracy  float32 `gorm:"not null"`
	Base
}

// SetDefault use for set default value if created and updated time
func (u *HistoryPosition) SetDefault() {
	u.Base.CreatedAt = time.Now()
	u.Base.UpdatedAt = time.Now()
}

// TableName ...
func (u *HistoryPosition) TableName() string {
	return HistoryPositionTableName
}
