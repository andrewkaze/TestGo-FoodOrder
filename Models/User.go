package Models

import (
	//"github.com/getsentry/sentry-go"
	"foodorder/Config"
)

func GetAllUser(user *[]User) (err error)  {
	if err = Config.DB.Find(user).Error; err != nil{
		return err
	}
	return nil
}

func GetAllMenu(menu *[]Menu) (err error)  {
	if err = Config.DB.Table("menus").Find(menu).Error; err != nil{
		return err
	}
	return nil

}

func CreateMenu(menu *Menu) (err error)  {
	if err = Config.DB.Table("menus").Create(menu).Error; err != nil{
		return err
	}
	//sentry.CaptureMessage("User created")
	return nil
}

func CreateUser(user *User) (err error)  {
	if err = Config.DB.Create(user).Error; err != nil{
		return err
	}
	//sentry.CaptureMessage("User created")
	return nil
}

func DeleteMenu(menu *Menu, id string)(err error)  {
	Config.DB.Table("menus").Where("id=?", id).Delete(menu)
	return nil
}

func GetMenuByID(menu *Menu, id string)(err error)  {
	if err = Config.DB.Table("menus").Where("id=?", id).First(menu).Error; err != nil{
		return err
	}
	return nil
}

func UpdateMenu(menu *Menu, id string) (err error) {
	Config.DB.Table("menus").Save(menu)
	return nil
}

func UpdateOrder(order *Order, id string) (err error) {
	Config.DB.Table("orders").Save(order)
	return nil
}

func CreateOrder(order *Order) (err error)  {
	if err = Config.DB.Table("orders").Create(order).Error; err != nil{
		return err
	}
	//sentry.CaptureMessage("User created")
	return nil
}

func GetOrderByID(order *Order, id string)(err error)  {
	if err = Config.DB.Table("orders").Where("id=?", id).First(order).Error; err != nil{
		return err
	}
	return nil
}

func GetAllOrder(order *[]Order) (err error)  {
	if err = Config.DB.Table("orders").Find(order).Error; err != nil{
		return err
	}
	return nil
}

func GetAllOrderByIDUser(order *[]Order, id string)(err error)  {
	if err = Config.DB.Table("orders").Where("id_user=?", id).Find(order).Error; err != nil{
		return err
	}
	return nil
}

func GetUserByID(user *User, id string)(err error)  {
	if err = Config.DB.Where("id=?", id).First(user).Error; err != nil{
		return err
	}
	return nil
}

func GetUserByEmail(user *User, email string)(err error)  {
	if err = Config.DB.Model(&User{}).Where("email=?", email).First(user).Error; err != nil{
		return err
	}
	return nil
}

func GetUserByEmailPassword(user *User, email string, password string)(err error)  {
	if err = Config.DB.Table("user").Select("id,name,email,address,password,role").Where("email=?", email).Where("password=?",password).First(user).Error; err != nil{
		return err
	}
	return nil
}

func UpdateUser(user *User, id string) (err error) {
	Config.DB.Save(user)
	return nil
}

func DeleteUser(user *User, id string)(err error)  {
	Config.DB.Where("id=?", id).Delete(user)
	return nil
}

