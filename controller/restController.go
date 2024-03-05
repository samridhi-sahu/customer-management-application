package controller

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/samridhi-sahu/cms/database"
	pb "github.com/samridhi-sahu/cms/protobuf"
)

func Logging(ctx *gin.Context) {
	ctx.Next()
	startTime := time.Now()
	duration := time.Since(startTime)
    log.Printf("[%s] %s %s - %v", ctx.Request.Method, ctx.Request.URL.Path, ctx.ClientIP(), duration)
}

func GetCustomers(ctx *gin.Context) {
	customersListfetch, err := database.Client.GetAllCustomers(context.Background(), &pb.Request{})
	if err != nil {
		log.Println("error in fetching list of customer, client.getallcustomer error ", err)
	}
	var customersList = customersListfetch.Customers
	ctx.JSON(http.StatusOK, customersList)
}
func GetACustomer(ctx *gin.Context) {
	id := ctx.Query("id")
	customer, err := database.Client.GetCustomer(context.Background(), &pb.CustomerRequest{Id: id})
	if err != nil {
		log.Println("error in fetching customer, client.getcustomer error ", err)
	}
	ctx.JSON(http.StatusOK, customer)
}
func AddCustomer(ctx *gin.Context) {
	var newcustomer *pb.NewCustomer
	err := ctx.BindJSON(&newcustomer)
	if err != nil {
		log.Println("error in parsing body ", err)
	}
	customer, err := database.Client.CreateCustomer(context.Background(), newcustomer)
	if err != nil {
		log.Println("error in creating customer, client.createcustomer error ", err)
	}
	ctx.JSON(http.StatusOK, customer)
}
func UpdateACustomer(ctx *gin.Context) {
	id := ctx.Query("id")
	updatecustomer := &pb.NewCustomer{}
	err := ctx.BindJSON(updatecustomer)
	if err != nil {
		log.Println("error in parsing request body ", err)
	}
	customer, err := database.Client.UpdateCustomer(context.Background(),
		&pb.Customer{Id: id, Name: updatecustomer.Name, City: updatecustomer.City})
	if err != nil {
		log.Println("error in updating customer, client.updatecustomer error ", err)
	}
	ctx.JSON(http.StatusOK, customer)
}
func DeleteACustomer(ctx *gin.Context) {
	id := ctx.Query("id")
	customer, err := database.Client.DeleteCustomer(context.Background(), &pb.CustomerRequest{Id: id})
	if err != nil {
		log.Println("error in deleting customer, client.deletecustomer error ", err)
	}
	ctx.JSON(http.StatusOK, customer)
}
