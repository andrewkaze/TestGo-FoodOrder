package Models

type User struct {
	Id      uint   `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
	Password string `json:"password"`
	Role string `json:"role"`
}
type Menu struct {
	Id      uint   `json:"id"`
	Name    string `json:"name"`
	Price   string `json:"price"`
	Description   string `json:"desc"`

}

type Order struct {
	Id      uint   `json:"id"`
	IdUser    string `json:"idUser"`
	IdMenu   string `json:"idMenu"`
	DeliveryAddress   string `json:"deliveryAddress"`
	Status   string `json:"status"`

}
func (b *User) TableName() string {
	return "user"
}

func (b *Menu) TableName() string {
	return "menu"
}

func (b *Order) TableName() string {
	return "order"
}

//func (u *User)BeforeCreate(tx *gorm.DB)  (err error){
//	u.Name = Service.Encrypt(u.Name)
//	return
//}
//
//func (u *User)BeforeUpdate(tx *gorm.DB)  (err error){
//	u.Name = Service.Encrypt(u.Name)
//	return
//}
//
//func (u *User) AfterFind(tx *gorm.DB) (err error) {
//	u.Name = Service.Decrypt(u.Name)
//	return
//}
