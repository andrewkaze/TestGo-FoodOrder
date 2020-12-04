package Controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/streadway/amqp"
	"log"
	"net/http"
	"foodorder/Models"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func bodyFrom(order *Models.Order) string {

	var s string

	s=`{ "id" : `+ `"`+string(order.Id)+`"` +`"idUser" : `+ `"`+order.IdUser+`"` + `,"idMenu" : ` + `"` + order.IdMenu +`"` + `,"status" : ` +`"received" }`

	return s
}

func SendMq (order *Models.Order,operation string) error{
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"task_queue", // name
		true,         // durable
		false,        // delete when unused
		false,        // exclusive
		false,        // no-wait
		nil,          // arguments
	)
	failOnError(err, "Failed to declare a queue")

	body := bodyFrom(order)
	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "application/json",
			Body:         []byte(body),
		})
	failOnError(err, "Failed to publish a message")
	log.Printf(" [x] Sent %s", body)
	return err
}

func GetUsers(c *gin.Context)  {
	var user[]Models.User
	err := Models.GetAllUser(&user)
	if err != nil{
		c.AbortWithStatus(http.StatusNotFound)
	}else{
		format := c.DefaultQuery("format", "json")
		if format == "json"{
			c.JSON(http.StatusOK, user)
		}else{
			c.XML(http.StatusOK, user)
		}
	}
}

func GetMenu(c *gin.Context)  {
	var menu[]Models.Menu
	err := Models.GetAllMenu(&menu)
	if err != nil{
		c.AbortWithStatus(http.StatusNotFound)
	}else{
		format := c.DefaultQuery("format", "json")
		if format == "json"{
			c.JSON(http.StatusOK, menu)
		}else{
			c.XML(http.StatusOK, menu)
		}
	}
}

func CreateMenu(c *gin.Context){
	var menu Models.Menu
	c.BindJSON(&menu)
	//err := SendMq(&user,"createUser")
	err := Models.CreateMenu(&menu)
	if err != nil{
		c.AbortWithStatus(http.StatusNotFound)
	}else{
		c.JSON(http.StatusOK, menu)
	}
}

func UpdateMenu(c *gin.Context)  {
	var menu Models.Menu
	id := c.Params.ByName("id")

	err := Models.GetMenuByID(&menu, id)
	if err != nil{
		c.JSON(http.StatusNotFound, menu)
	}

	c.BindJSON(&menu)
	//err = SendMq(&user,"updateUser")
	err = Models.UpdateMenu(&menu, id)
	if err != nil{
		c.AbortWithStatus(http.StatusNotFound)
	}else{
		c.JSON(http.StatusOK, menu)
	}
}

func UpdateOrder(c *gin.Context)  {
	var order Models.Order

	id := c.Params.ByName("id")

	err := Models.GetOrderByID(&order, id)
	if err != nil{
		c.JSON(http.StatusNotFound, order)
	}

	c.BindJSON(&order)
	//err = SendMq(&user,"updateUser")
	order.Status="Processed"
	err = Models.UpdateOrder(&order, id)
	if err != nil{
		c.AbortWithStatus(http.StatusNotFound)
	}else{
		c.JSON(http.StatusOK, order)
	}
}


func DeleteMenu(c *gin.Context)  {
	var menu Models.Menu
	id := c.Params.ByName("id")
	//err := SendMq(&user,"deleteUser")
	err := Models.DeleteMenu(&menu, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id" + id: "is deleted"})
	}
}

func CreateOrder(c *gin.Context){
	var order Models.Order
	order.Status = "Received"
	c.BindJSON(&order)

	//err := SendMq(&user,"createUser")
	err := Models.CreateOrder(&order)
	if err != nil{
		c.AbortWithStatus(http.StatusNotFound)
	}else{
		c.JSON(http.StatusOK, order)
	}
}

func GetOrders(c *gin.Context)  {
	var order[]Models.Order
	err := Models.GetAllOrder(&order)
	if err != nil{
		c.AbortWithStatus(http.StatusNotFound)
	}else{
		format := c.DefaultQuery("format", "json")
		if format == "json"{
			c.JSON(http.StatusOK, order)
		}else{
			c.XML(http.StatusOK, order)
		}
	}
}



func GetAllOrderByIDUser(c *gin.Context)  {

	var order[]Models.Order
	id := c.Params.ByName("id")
	err := Models.GetAllOrderByIDUser(&order, id)
	if err != nil{
		c.AbortWithStatus(http.StatusNotFound)
	}else{
		c.JSON(http.StatusOK, order)
	}
}

//func CreateOrder(c *gin.Context){
//	var user Models.User
//	c.BindJSON(&user)
//	err := SendMq(&user,"createOrder")
//	//err := Models.CreateUser(&user)
//	if err != nil{
//		c.AbortWithStatus(http.StatusNotFound)
//	}else{
//		c.JSON(http.StatusOK, user)
//	}
//}



func CreateUser(c *gin.Context){
	var user Models.User
	c.BindJSON(&user)
	//err := SendMq(&user,"createUser")
	err := Models.CreateUser(&user)
	if err != nil{
		c.AbortWithStatus(http.StatusNotFound)
	}else{
		c.JSON(http.StatusOK, user)
	}
}

func GetUserByID(c *gin.Context)  {
	var user Models.User
	id := c.Params.ByName("id")
	err := Models.GetUserByID(&user, id)
	if err != nil{
		c.AbortWithStatus(http.StatusNotFound)
	}else{
		c.JSON(http.StatusOK, user)
	}
}

func UpdateUser(c *gin.Context)  {
	var user Models.User
	id := c.Params.ByName("id")

	err := Models.GetUserByID(&user, id)
	if err != nil{
		c.JSON(http.StatusNotFound, user)
	}

	c.BindJSON(&user)
	//err = SendMq(&user,"updateUser")
	err = Models.UpdateUser(&user, id)
	if err != nil{
		c.AbortWithStatus(http.StatusNotFound)
	}else{
		c.JSON(http.StatusOK, user)
	}
}

func DeleteUser(c *gin.Context)  {
	var user Models.User
	id := c.Params.ByName("id")
	//err := SendMq(&user,"deleteUser")
	err := Models.DeleteUser(&user, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id" + id: "is deleted"})
	}
}