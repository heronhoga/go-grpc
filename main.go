package main

import (
	"fmt"
	pb "go-grpc/pb"
	"log"

	"google.golang.org/protobuf/proto"
)

func main() {

	products := &pb.Products{
		Data: []*pb.Product{
			{
				Id: 1,
				Name: "Nike Air Tees",
				Price: 100000,
				Stock: 5,
				Category: &pb.Category{
					Id: 1,
					Name: "shirt",
				},
			},
			{
				Id: 2,
				Name: "Nike Training short",
				Price: 120000,
				Stock: 10,
				Category: &pb.Category{
					Id: 2,
					Name: "pant",
				},
			},
		},
	}
	
	data, err := proto.Marshal(products)
	if err != nil {
		log.Fatal(err)
	}
	//COMPACT BINARY WIRE FORMAT
	fmt.Println(data)
}