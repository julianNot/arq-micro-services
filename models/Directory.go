package models

type Directory struct {
	DirectoryID  string `gorm:"primaryKey;not null;unique_index" json:"id_directory"`
	UserName     string `gorm:"not null" json:"username"`
	AccessEmail  string `gorm:"not null" json:"access_email"`
	HashPassword string `gorm:"not null" json:"hash_password"`
	SaltHash     string `json:"salt_hash"`
	RoleID       string `gorm:"foreignKey:role_id" json:"role_id"`
}
