package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"payment/settings/database"

	res "payment/utils/response"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

type Employee struct {
	Name  string `json:"name,omitempty"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

func Test(c *gin.Context) {
	c.JSON(200, "Testing...")
}

func AddOneEmp(c *gin.Context) {

	//db connection
	collection, _ := database.MongoConnection()

	fmt.Println("Adding emp")
	var emp Employee

	jsonBytes, _ := ioutil.ReadAll(c.Request.Body)
	json.Unmarshal(jsonBytes, &emp)
	fmt.Println("emp::::", emp)

	msg, err := collection.InsertOne(context.Background(), emp)
	if err != nil {
		log.Print(err)
	} else {
		log.Print(":::", msg)
	}
	c.JSON(200, msg)
}

func GetEmployees(c *gin.Context) {
	//db connection
	collection, _ := database.MongoConnection()

	// getting emp values
	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		fmt.Println("Error while retrieving employees:", err)
		return
	}
	var employees []Employee
	for cursor.Next(context.Background()) {
		var emp Employee
		if err := cursor.Decode(&emp); err != nil {
			fmt.Println("error while decoding:", err)
			return
		}
		employees = append(employees, emp)
	}
	c.JSON(200, employees)
}
func GetOneEmployee(c *gin.Context) {
	//connect mongodb
	collection, _ := database.MongoConnection()
	name := c.Param("name")
	fmt.Println("name:::", name)
	filterbyname := bson.M{"name": name}
	var emp Employee
	err := collection.FindOne(context.Background(), filterbyname).Decode(&emp)
	if err != nil {
		fmt.Println("error while finding data:::", err)
	}

	c.JSON(200, emp)
}

func UpdateEmployee(c *gin.Context) {
	//db connection
	collection, _ := database.MongoConnection()
	var emp Employee
	err := c.ShouldBindJSON(&emp)
	if err != nil {
		fmt.Println("Error while reading and unmarshaling data:", err)
	}
	name := emp.Name
	filterbyname := bson.M{"name": name}

	fmt.Println("emppp", emp)

	update := bson.M{
		"$set": bson.M{
			"name":  emp.Name,
			"email": emp.Email,
			"age":   emp.Age,
		},
	}
	fmt.Println("update:::", update)
	res, err := collection.UpdateOne(context.Background(), filterbyname, update)
	if err != nil {
		log.Print("Error while updating data", err)
	}
	c.JSON(200, res)
}

func DeleteOneEmployee(c *gin.Context) {
	//connect mongodb
	collection, _ := database.MongoConnection()
	name := c.Param("name")
	fmt.Println("name:::", name)
	filterbyname := bson.M{"name": name}

	// res, err := collection.DeleteOne(context.Background(), filterbyname)
	result, err := collection.DeleteMany(context.Background(), filterbyname)
	if err != nil {
		fmt.Println("error while finding data:::", err)
	}
	res.Response(c, 200, result, "Deleted one record")
}
