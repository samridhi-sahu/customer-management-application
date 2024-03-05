package main

import (
	"context"
	"log"
	"reflect"
	"testing"

	pb "github.com/samridhi-sahu/cms/protobuf"
)

func TestGetCustomer(t *testing.T) {
	s := &CustomerServiceServer{}
	tests := map[string]struct {
		in  *pb.CustomerRequest
		out *pb.Customer
	}{
		"Success": {
			in: &pb.CustomerRequest{Id: "4b108556-6b46-4ef3-9ef8-22366ff63498"},
			out: &pb.Customer{
				Id:   "6961a188-efee-4aac-855d-df7acbc6f0",
				Name: "John Doe",
				City: "New York",
			},
		},
		"Not_Found": {
			in: &pb.CustomerRequest{Id: "5b108556-6b46-4ef3-9ef8-22366ff63498"},
			out: &pb.Customer{},
		},
		"Empty": {
			in: &pb.CustomerRequest{Id: ""},
			out: &pb.Customer{},
		},
		"Invalid": {
			in: &pb.CustomerRequest{Id: "test"},
			out: &pb.Customer{},
		},
	}
	for key, value := range tests {
		t.Run(key, func(t *testing.T) {
			response, err := s.GetCustomer(context.Background(), value.in)
			if err != nil {
				log.Println("get customer response err ", err)
			}
			reflect.DeepEqual(value.out, response)
		})
	}
}
