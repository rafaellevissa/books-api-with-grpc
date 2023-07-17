package grpc_go

import (
	"context"
	"crud/internal/database"
	"crud/internal/models"
	"crud/proto/pb"
	"errors"
	"strconv"

	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedBookServiceServer
}

func (service Server) CreateBook(ctx context.Context, req *pb.Book) (*pb.BookResponse, error) {
	db := database.GetDatabase()

	var book models.Book

	book.Name = req.Name
	book.Description = req.Description
	book.MediumPrice = req.MediumPrice
	book.Author = req.Author
	book.ImageURL = req.ImgUrl

	//Criando
	err := db.Create(&book).Error
	if err != nil {
		return &pb.BookResponse{Status: 500}, errors.New("erro ao criar o livro")
	}

	response := &pb.BookResponse{
		Book: &pb.Book{
			Id:          strconv.FormatUint(uint64(book.ID), 10),
			Name:        book.Name,
			Description: book.Description,
			MediumPrice: book.MediumPrice,
			Author:      book.Author,
			ImgUrl:      book.ImageURL,
		},
		Status: 200,
	}

	return response, nil
}

func (service Server) FindAllBooks(ctx context.Context, in *pb.Void) (*pb.BookList, error) {
	db := database.GetDatabase()

	var books []models.Book

	err := db.Find(&books).Error
	if err != nil {
		return &pb.BookList{Books: []*pb.Book{{}}, Status: &pb.BookResponse{Status: 404}}, errors.New("erro ao listar os livros")
	}

	sliceBook := []*pb.Book{}

	for _, item := range books {
		book := pb.Book{
			Id:          strconv.FormatUint(uint64(item.ID), 10),
			Name:        item.Name,
			Description: item.Description,
			MediumPrice: item.MediumPrice,
			Author:      item.Author,
			ImgUrl:      item.ImageURL,
		}

		sliceBook = append(sliceBook, &book)
	}

	response := &pb.BookList{
		Books: sliceBook,
		Status: &pb.BookResponse{
			Status: 200,
		},
	}

	return response, nil
}

func (service Server) FindBook(ctx context.Context, req *pb.BookId) (*pb.BookResponse, error) {
	bookId := req.Id

	//Transformando para inteiro
	newBookId, err := strconv.Atoi(bookId)
	if err != nil {
		return &pb.BookResponse{Book: &pb.Book{}, Status: 400}, errors.New("iD has to be integer")
	}

	db := database.GetDatabase()

	var book models.Book

	err = db.First(&book, newBookId).Error
	if err != nil {
		return &pb.BookResponse{Book: &pb.Book{}, Status: 404}, errors.New("livro não encontrado")
	}

	response := &pb.BookResponse{
		Book: &pb.Book{
			Id:          strconv.FormatUint(uint64(book.ID), 10),
			Name:        book.Name,
			Description: book.Description,
			MediumPrice: book.MediumPrice,
			Author:      book.Author,
			ImgUrl:      book.ImageURL,
		},
		Status: 200,
	}

	return response, nil
}

func (service Server) EditBook(ctx context.Context, req *pb.Book) (*pb.BookResponse, error) {
	bookId := req.Id

	//Transformando para inteiro
	newBookId, err := strconv.Atoi(bookId)
	if err != nil {
		return &pb.BookResponse{Book: &pb.Book{}, Status: 400}, errors.New("iD has to be integer")
	}

	db := database.GetDatabase()

	var book models.Book

	err = db.First(&book, newBookId).Error
	if err != nil {
		return &pb.BookResponse{Book: &pb.Book{}, Status: 404}, errors.New("livro não encontrado")
	}

	book.Name = req.Name
	book.Description = req.Description
	book.MediumPrice = req.MediumPrice
	book.Author = req.Author
	book.ImageURL = req.ImgUrl

	//Salvando
	err = db.Save(&book).Error
	if err != nil {
		return &pb.BookResponse{Status: 500}, errors.New("erro ao atualizar livro")
	}

	response := &pb.BookResponse{
		Book: &pb.Book{
			Id:          strconv.FormatUint(uint64(book.ID), 10),
			Name:        book.Name,
			Description: book.Description,
			MediumPrice: book.MediumPrice,
			Author:      book.Author,
			ImgUrl:      book.ImageURL,
		},
		Status: 201,
	}

	return response, nil
}

func (service Server) DeleteBook(ctx context.Context, req *pb.BookId) (*pb.BookResponse, error) {
	bookId := req.Id

	//Transformando para inteiro
	newBookId, err := strconv.Atoi(bookId)
	if err != nil {
		return &pb.BookResponse{Book: &pb.Book{}, Status: 400}, errors.New("iD has to be integer")
	}

	db := database.GetDatabase()

	var book models.Book

	err = db.First(&book, newBookId).Error
	if err != nil {
		return &pb.BookResponse{Status: 404}, errors.New("livro não encontrado")
	}

	err = db.Delete(&book, req.Id).Error
	if err != nil {
		return &pb.BookResponse{Status: 400}, errors.New("erro ao deletar livro")
	}

	response := &pb.BookResponse{
		Status: 204,
	}

	return response, nil
}

func RegisterService(grpc_server grpc.Server) {
	pb.RegisterBookServiceServer(&grpc_server, &Server{})
}
