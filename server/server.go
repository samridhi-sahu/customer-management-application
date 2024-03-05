package main

import (
	"context"
	"log"
	"net"

	"github.com/google/uuid"
	"github.com/samridhi-sahu/cms/database"
	pb "github.com/samridhi-sahu/cms/protobuf"
	"google.golang.org/grpc"
)

type CustomerServiceServer struct {
	pb.UnimplementedCustomerServicesServer
}
func init(){
	database.Setup()
}
func main() {
	
	database.CreateTable()
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("grpc server lis error ", err)
	}
	grpcserver := grpc.NewServer()
	pb.RegisterCustomerServicesServer(grpcserver, &CustomerServiceServer{})
	grpcserver.Serve(lis)
}

func (s *CustomerServiceServer) GetAllCustomers(ctx context.Context, req *pb.Request) (*pb.AllCustomersList, error) {
	list := &pb.AllCustomersList{}
	return database.GetAll(list), nil
}

func (s *CustomerServiceServer) GetCustomer(ctx context.Context, req *pb.CustomerRequest) (*pb.Customer, error) {
	customer := &pb.Customer{}
	return database.Get(customer, req), nil
}

func (s *CustomerServiceServer) CreateCustomer(ctx context.Context, req *pb.NewCustomer) (*pb.Customer, error) {
	customer := &pb.Customer{}
	customer.Id = uuid.New().String()
	customer.Name = req.Name
	customer.City = req.City
	return database.Create(customer), nil
}

func (s *CustomerServiceServer) UpdateCustomer(ctx context.Context, req *pb.Customer) (*pb.Customer, error) {
	return database.Update(req), nil
}

func (s *CustomerServiceServer) DeleteCustomer(ctx context.Context, req *pb.CustomerRequest) (*pb.Customer, error) {
	return database.Delete(req), nil
}
