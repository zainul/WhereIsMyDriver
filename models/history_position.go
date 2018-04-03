package models

//HistoryPositionTableName ...
const HistoryPositionTableName = "history_positions"

// HistoryPosition ...
type HistoryPosition struct {
	UserID   string  `gorm:"not null" validate:"required"`
	Lat      float32 `gorm:"not null" sql:"type:decimal(9,6);"`
	Lang     float32 `gorm:"not null" sql:"type:decimal(9,6);"`
	Accuracy float32 `gorm:"not null"`
	Base
}
