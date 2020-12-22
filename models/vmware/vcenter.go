package vmware

import (
	"encoding/base64"
	"github.com/jinzhu/gorm"
)

type VCenter struct {
	gorm.Model
	Name		string	`json:"name" gorm:"not null"`
	Description	string	`json:"description"`
	HostnameOrIp	string	`json:"hostnameOrIp" gorm:"not null"`
	Username	string	`json:"username" gorm:"not null"`
	Password	string	`json:"password" gorm:"not null"`
	SslThumbprint	string	`json:"sslThumbprint" gorm: "not null"`
	ProductVersion	string	`json:"productVersion"`
	ProductType	string	`json:"productType"`
}

func (vc *VCenter) BeforeSave(tx *gorm.DB) (err error) {
	encodedPassword := base64.StdEncoding.EncodeToString([]byte(vc.Password))
	vc.Password = encodedPassword
	return
}

func (vc *VCenter) AfterFind(tx *gorm.DB) (err error) {
	decodedPassword, _ := base64.StdEncoding.DecodeString(vc.Password)
	vc.Password = string(decodedPassword)
	return
}
