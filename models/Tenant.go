package models

type Tenant struct {
	TenantID     string `gorm:"primaryKey;not null;unique_index" json:"id_tenant"`
	Name         string `gorm:"type:varchar;not null" json:"name_tenant"`
	UrlLogo      string `gorm:"type:varchar;not null" json:"url_logo"`
	EmailContact string `gorm:"type:varchar;not null" json:"email_contact"`
}
