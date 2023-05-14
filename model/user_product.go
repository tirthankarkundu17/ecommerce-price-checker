package model

import "gorm.io/gorm"

type UserProduct struct {
	UserId string `json:"userId"`
	URL    string `json:"url"`
}

func (u *UserProduct) SaveUserProduct(db *gorm.DB) (*UserProduct, error) {

	err := db.Debug().Create(&u).Error
	if err != nil {
		return &UserProduct{}, err
	}
	return u, nil
}

func (u *UserProduct) GetUserProduct(db *gorm.DB, userId string) (*[]UserProduct, error) {

	var err error
	products := []UserProduct{}
	err = db.Debug().Model(&UserProduct{}).Limit(100).Find(&products, userId).Error
	if err != nil {
		return &[]UserProduct{}, err
	}
	return &products, err
}
