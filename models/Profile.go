package models

import (
	"github.com/google/uuid"
	"time"
	_"github.com/jinzhu/gorm"
)

type Profile struct {
	ID 		 	uuid.UUID   	`json:"id" gorm:"default:uuid_generate_v3()"`
	DOB      	string   		`json:"dob"`
	Address   	Address  		`json:"address"`
	UserID    	uuid.UUID   	`json:"user_id"`
	CreatedAt	time.Time		`json:"created_at"`
	UpdatedAt	time.Time		`json:"updated_at"`
}

type Address struct {
	Street  string `json:"street"`
	City    string `json:"city"`
	State   string `json:"state"`
	Country string `json:"country"`
}

func (p *Profile) CreateProfile(profile *Profile) (*Profile, error) {

	err := DB.Model(Profile{}).Create(profile).Error
	if err != nil {
		return nil, err
	}
	
	return profile, nil
}

func (p *Profile) ProfileDetail(user_id string) (*Profile, error)  {
	var profile *Profile
	err := DB.Model(Profile{}).Where("user_id = ?", user_id).First(profile).Error
	if err != nil {
		return nil, err
	}
	
	return profile, nil
}

func (p *Profile) DeleteProfile(user_id string) (error){
	err := DB.Model(Profile{}).Where("user_id = ?", user_id).Delete(p).Error
	if err != nil{
		return err
	}
	
	return nil
}

func (p *Profile) UpdateProfile(user_id string) (*Profile, error) {
	err := DB.Model(Profile{}).Where("user_id = ?", user_id).Update(p).Error
	if err != nil {
		return nil, err
	}
	
	return p, nil
}