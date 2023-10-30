package models

type Role struct {
	IdRole      string `gorm:"column:id_role;not null" json:"id_role"`
	NameRole    string `gorm:"column:name_role;not null" json:"name_role"`
	Description string `gorm:"column:description" json:"description"`
	// Actions     []Action    `gorm:"many2many:role_actions;" json:"actions"`
	// Directories []Directory `gorm:"has_many:directories;"`
}
