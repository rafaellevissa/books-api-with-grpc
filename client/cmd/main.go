package main

import (
	"context"
	pb "crud/client/proto/pb"
	"log"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:5000", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	routehttp := "/FindAllBooks"
	client := pb.NewBookServiceClient(conn)

	if routehttp == "/CreateBook" {
		req := &pb.Book{
			Name:        "Livro Teste 3",
			Description: "Um livro feito para teste '3'",
			MediumPrice: 11.99,
			Author:      "Carlos da Pamonha da Silva Junior",
			ImgUrl:      "acarajé.com.br",
		}

		res, err := client.CreateBook(context.Background(), req)
		if err != nil {
			log.Fatal(err)
		}

		log.Print(res)
	} else if routehttp == "/FindAllBooks" {
		req := &pb.Void{}

		res, err := client.FindAllBooks(context.Background(), req)
		if err != nil {
			log.Fatal(err)
		}

		log.Print(res)
	} else if routehttp == "/FindBook" {
		req := &pb.BookId{Id: "10"}

		res, err := client.FindBook(context.Background(), req)
		if err != nil {
			log.Fatal(err)
		}

		log.Print(res)
	} else if routehttp == "/EditBook" {
		req := &pb.Book{
			Id:          "9",
			Name:        "Livro teste 9 alterado",
			Description: "Fiz a alteração agora",
			MediumPrice: 835.62,
			Author:      "Sergio",
			ImgUrl:      "alteracao.com.br",
		}

		res, err := client.EditBook(context.Background(), req)
		if err != nil {
			log.Fatal(err)
		}

		log.Print(res)
	} else if routehttp == "/DeleteBook" {
		req := &pb.BookId{
			Id: "10",
		}
		res, err := client.DeleteBook(context.Background(), req)
		if err != nil {
			log.Fatal(err)
		}

		log.Print(res)
	}

}
