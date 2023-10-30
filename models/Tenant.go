package models

type Tenant struct {
	TenantID     string `gorm:"column:id_tenant;primaryKey:id_tenant;not null;unique_index" json:"id_tenant"`
	NameTenant   string `gorm:"type:varchar;not null" json:"name_tenant"`
	UrlLogo      string `gorm:"type:varchar;not null" json:"url_logo"`
	EmailContact string `gorm:"type:varchar;not null" json:"email_contact"`
}

func (Tenant) TableName() string {
	return "Tenant"
}