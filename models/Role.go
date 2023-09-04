package models

type Role struct {
	RoleID      string      `gorm:"primaryKey;not null;unique_index" json:"id_role"`
	NameRole    string      `gorm:"not null" json:"name_role"`
	Description string      `json:"description"`
	Actions     []Action    `gorm:"many2many:role_actions;" json:"actions"`
	Directories []Directory `gorm:"has_many:directories;"`
}
