package models

type Society struct{
	Id int64 `json:"id"`
	Name string `json:"name"`
	Address string `json:"address"`
	RegistrationNo string `json:"registration_no"`

}