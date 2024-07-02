package book

import (
    "context"

    "google.golang.org/grpc/codes"
    "google.golang.org/grpc/status"
)

type Book struct {
    ID     string
    Title  string
    Author string
}

var books = make(map[string]Book)

type BookServiceServer struct {
    UnimplementedBookServiceServer
}

func (s *BookServiceServer) AddBook(ctx context.Context, req *AddBookRequest) (*AddBookResponse, error) {
    book := Book{
        ID:     "1", 
        Title:  req.GetTitle(),
        Author: req.GetAuthor(),
    }
    books[book.ID] = book
    return &AddBookResponse{BookId: book.ID}, nil
}

func (s *BookServiceServer) GetBook(ctx context.Context, req *GetBookRequest) (*Book, error) {
    book, exists := books[req.GetBookId()]
    if !exists {
        return nil, status.Errorf(codes.NotFound, "Book not found")
    }
    return &Book{
        BookId: book.ID,
        Title:  book.Title,
        Author: book.Author,
    }, nil
}
