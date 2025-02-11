package models

type JwtVerify struct {
	Token        string `gorm:"varchar(255)" json:"token,omitempty"`
	RefreshToken string `gorm:"varchar(255)" json:"refresh_token,omitempty"`
}
