package main

import (
	"log"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/gin-gonic/gin"
	"github.com/samridhi-sahu/cms/database"
	"github.com/samridhi-sahu/cms/graph"
	pb "github.com/samridhi-sahu/cms/protobuf"
	"github.com/samridhi-sahu/cms/router"
	"go.appointy.com/jaal"
	"go.appointy.com/jaal/schemabuilder"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	//conn, err := grpc.Dial(":8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.Dial("server:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatal("grpc client connection error", err)
	}
	defer conn.Close()
	database.Client = pb.NewCustomerServicesClient(conn)
	r := gin.Default()
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))
	router.GraphQLRoutes(srv, r)
	router.RestRoutes(r)
	sb := schemabuilder.NewSchema()
	pb.RegisterInputAllCustomersList(sb)
	pb.RegisterInputCreateCustomerInput(sb)
	pb.RegisterInputCustomer(sb)
	pb.RegisterInputCustomerRequest(sb)
	pb.RegisterInputDeleteCustomerInput(sb)
	pb.RegisterInputNewCustomer(sb)
	pb.RegisterInputRequest(sb)
	pb.RegisterInputUpdateCustomerInput(sb)
	pb.RegisterPayloadAllCustomersList(sb)
	pb.RegisterPayloadCreateCustomerPayload(sb)
	pb.RegisterPayloadCustomer(sb)
	pb.RegisterPayloadCustomerRequest(sb)
	pb.RegisterPayloadDeleteCustomerPayload(sb)
	pb.RegisterPayloadNewCustomer(sb)
	pb.RegisterPayloadRequest(sb)
	pb.RegisterPayloadUpdateCustomerPayload(sb)
	pb.RegisterCustomerServicesOperations(sb, database.Client)
	schema, err := sb.Build()
	if err != nil {
		log.Fatal("schema build error ", err)
	}
	r.POST("/jaal", gin.WrapH(jaal.HTTPHandler(schema)))
	err = r.Run(":4000")
	if err != nil {
		log.Fatal("http listen error ", err)
	}
}
