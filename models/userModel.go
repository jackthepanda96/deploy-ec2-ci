package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `json:"email"`
	Gender   string `json:"gender"`
	Password string `json:"password"`
	Token    string `json:"token"`
	// WritedBook Book
}

type UserDBModel struct {
	db *gorm.DB
}

func NewUserModel(db *gorm.DB) *UserDBModel {
	return &UserDBModel{db: db}
}

type UserModel interface {
	GetAll() ([]User, error)
	GetByName(name string) (User, error)
	GetByID(id int) (User, error)
	Insert(newUser User) (User, error)
	Update(newData User, userID int) (User, error)
	Delete(userID int) (User, error)
	GetByEmailAndPassword(email, password string) (User, error)
}

func (u *UserDBModel) GetAll() ([]User, error) {
	user := []User{}
	if err := u.db.Find(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserDBModel) GetByName(name string) (User, error) {
	user := User{}
	if err := u.db.Where("name = ?", name).Find(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (u *UserDBModel) GetByID(id int) (User, error) {
	user := User{}
	if err := u.db.Find(&user, id).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (u *UserDBModel) Insert(newUser User) (User, error) {
	if err := u.db.Save(&newUser).Error; err != nil {
		return newUser, err
	}
	return newUser, nil
}

func (u *UserDBModel) Update(newData User, userID int) (User, error) {
	user := User{}
	err := u.db.First(&user, userID).Error

	if err != nil {
		return user, err
	}

	user.Name = newData.Name
	user.Email = newData.Email
	user.Gender = newData.Gender
	user.Password = newData.Password
	user.Token = newData.Token

	err = u.db.Save(&user).Error
	return user, err
}

func (u *UserDBModel) Delete(userID int) (User, error) {
	user := User{}
	err := u.db.First(&user, userID).Error

	if err != nil {
		return user, err
	}

	if err := u.db.Delete(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (u *UserDBModel) GetByEmailAndPassword(email string, password string) (User, error) {
	user := User{}
	err := u.db.Where("email = ? AND password = ?", email, password).First(&user).Error
	return user, err
}
