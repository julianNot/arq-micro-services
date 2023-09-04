package models

type Action struct {
	ActionID    string `gorm:"primaryKey;not null" json:"id_action"`
	NameAction  string `gorm:"not null" json:"name_action"`
	Description string `gorm:"not null" json:"description"`
}
