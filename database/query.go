package database

import (
	"log"

	pb "github.com/samridhi-sahu/cms/protobuf"
)

func GetAll(list *pb.AllCustomersList) *pb.AllCustomersList{
	query := `select * from customer`
	rows, err := Db.Query(query)
	if err != nil {
		log.Println("error in fetching customers from db ", err)
	}
	for rows.Next() {
		customer := &pb.Customer{}
		err = rows.Scan(&customer.Id, &customer.Name, &customer.City)
		if err != nil {
			log.Println("error in assigning values from row to customer")
		}
		list.Customers = append(list.Customers, customer)
	}
	return list
}

func Get(customer *pb.Customer, req *pb.CustomerRequest) *pb.Customer{
	query := `select * from customer where id = ?`
	row, err := Db.Query(query, req.Id)
	if err != nil {
		log.Println("error in fetching customer form db ", err)
	}
	if row.Next() {
		err = row.Scan(&customer.Id, &customer.Name, &customer.City)
		if err != nil {
			log.Println("error in assigning values from row to customer")
		}
	} 
	//else {
	// 	log.Println("no data fetched with this id")
	// }
	return customer
}

func Create(customer *pb.Customer) *pb.Customer{
	query := `insert into customer(id, name, city) values(?,?,?)`
	_, err := Db.Exec(query, customer.Id, customer.Name, customer.City)
	if err != nil {
		log.Println("error in inserting customer into db", err)
	}
	return customer
}

func Update(req *pb.Customer) *pb.Customer{
	query := `update customer set name = ?, city = ? where id = ?`
	_, err := Db.Exec(query, req.Name, req.City, req.Id)
	if err != nil { 
		log.Println("error in updating customer values ", err)
	}
	return req
}

func Delete(req *pb.CustomerRequest) *pb.Customer{
	query := `delete from customer where id = ?`
	_, err := Db.Exec(query, req.Id)
	if err != nil {
		log.Println("error in deleting customer ", err)
	}
	return &pb.Customer{}
}