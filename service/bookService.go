package service

import (
	//"github.com/mashingan/smapping"
	"golang_api/dto"
	"golang_api/entity"
	"golang_api/respository"
	//"log"
)

type BookService interface {
	Insert(book dto.BookCreateDTO) entity.Book
	Update(book dto.BookUpdateDTO) entity.Book
	Delete(book entity.Book)
	List() []entity.Book
	FindById(bookID uint) entity.Book
	IsAllowedToEdit(userId uint,bookID uint) bool
}

type bookService struct {
	bookRepository respository.BookRepository
}

func (b bookService) Insert(book dto.BookCreateDTO) entity.Book {
	bookToCreate := entity.Book{}
	//err := smapping.FillStruct(&bookToCreate, smapping.MapFields(&book))
	//if err != nil {
	//	log.Fatalf("Failed map %v", err)
	//}
	bookToCreate.UserId =book.UserID
	bookToCreate.Title =book.Title
	bookToCreate.Description =book.Description

	res := b.bookRepository.InsertBook(bookToCreate)

	return res
}

func (b bookService) Update(book dto.BookUpdateDTO) entity.Book {
	bookToUpdate := entity.Book{}

	bookToUpdate.ID = uint(book.ID)
	bookToUpdate.Title = book.Title
	bookToUpdate.Description = book.Description
	bookToUpdate.UserId = book.UserID

	res := b.bookRepository.UpdateBook(bookToUpdate)
	return res
}

func (b bookService) Delete(book entity.Book) {
	b.bookRepository.DeleteBook(book)
}

func (b bookService) List() []entity.Book {
	return b.bookRepository.BookList()

}

func (b bookService) FindById(bookID uint) entity.Book {
	return b.bookRepository.FindBookByID(bookID)
}

func (b bookService) IsAllowedToEdit(userId uint, bookID uint) bool {
	book := b.bookRepository.FindBookByID(bookID)
	return book.UserId == uint64(userId)
}

func NewBookService(bookRep respository.BookRepository) BookService {
	return &bookService{
		bookRepository: bookRep,
	}
}