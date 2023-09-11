package models

type Directory struct {
	DirectoryID  string `gorm:"column:id_directory;primaryKey:id_directory;not null;unique_index" json:"id_directory"`
	UserName     string `gorm:"column:username;not null" json:"username"`
	AccessEmail  string `gorm:"not null" json:"access_email"`
	HashPassword string `gorm:"not null" json:"hash_password"`
	SaltHash     string `gorm:"column:salt_hash" json:"salt_hash"`
	RoleID       string `gorm:"column:id_role;foreignKey:id_role" json:"id_role"`
}

func (Directory) TableName() string {
	return "Directory"
}
