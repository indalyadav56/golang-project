package services

import (
	"golang-project/config"
	"golang-project/models"
)


func GetUsers() ([]models.User, error) {
	var users []models.User
	config.DB.Find(&users)
	return users, nil
}

func CreateUser(user models.User) (models.User, error) {
	
	err :=config.DB.Create(&user).Error
  
	if err != nil {
		return user, err
	}

	return user, nil
	
	
  }

func UpdateUser(user models.User) (models.User, error) {
	// Update single field
	// db.Model(&user).Update("Name", "Jack") 
	// Update multiple fields
	// db.Model(&user).Updates(User{Name: "Jack", Age: 30})
	config.DB.Model(&user).Updates(models.User{FirstName: "Jack",})
	return user, nil
}  

func DeleteUser(user models.User) (models.User, error) {
	config.DB.Delete(models.User{}, 1)
	return user, nil
}