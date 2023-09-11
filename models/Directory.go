package models

type Directory struct {
	IdDirectory  string `gorm:"primaryKey:directory_id;unique_index" json:"directory_id"`
	UserName     string `gorm:"not null" json:"user_name"`
	AccessEmail  string `gorm:"not null" json:"access_email"`
	HashPassword string `gorm:"not null" json:"hash_password"`
	SaltHash     string `json:"salt_hash"`
	RoleID       string `gorm:"foreignKey:role_id" json:"role_id"`
}
