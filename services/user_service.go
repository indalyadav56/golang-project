package services

import (
	"golang-project/database"
	"golang-project/models"
)


func GetUsers() ([]models.User, error) {
	var users []models.User
	database.DB.Find(&users)
	return users, nil
}

func CreateUser(u models.User) (models.User, error) {
	database.DB.Save(&u)
	return u, nil
  }

func UpdateUser(user models.User) (models.User, error) {
	// Update single field
	// db.Model(&user).Update("Name", "Jack") 
	// Update multiple fields
	// db.Model(&user).Updates(User{Name: "Jack", Age: 30})
	database.DB.Model(&user).Updates(models.User{FirstName: "Jack",})
	return user, nil
}  

func DeleteUser(user models.User) (models.User, error) {
	database.DB.Delete(models.User{}, 1)
	return user, nil
}