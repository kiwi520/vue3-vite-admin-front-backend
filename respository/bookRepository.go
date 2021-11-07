package respository

import (
	"golang_api/entity"
	"gorm.io/gorm"
)

type BookRepository interface {
	InsertBook(book entity.Book) entity.Book
	UpdateBook(book entity.Book) entity.Book
	DeleteBook(book entity.Book)
	BookList() []entity.Book
	FindBookByID(bookID uint) entity.Book

}

type bookConnection struct {
	bookConnection *gorm.DB
}

func (b bookConnection) InsertBook(book entity.Book) entity.Book {
	b.bookConnection.Save(&book)
	b.bookConnection.Preload("User").Find(&book)
	return book
}

func (b bookConnection) UpdateBook(book entity.Book) entity.Book {
	b.bookConnection.Model(&book).Select("*").Omit("id","CreatedAt").Updates(map[string]interface{}{"title":book.Title,"description":book.Description})
	b.bookConnection.Preload("User").Find(&book)
	return book
}

func (b bookConnection) DeleteBook(book entity.Book) {
	b.bookConnection.Delete(&entity.Book{},book.ID)
}

func (b bookConnection) BookList() []entity.Book {
	var bookList = []entity.Book{}
	b.bookConnection.Preload("User").Find(&bookList)
	return bookList
}

func (b bookConnection) FindBookByID(bookID uint) entity.Book {
	var book = entity.Book{}
	b.bookConnection.Preload("User").Find(&book,bookID)
	return book
}

func NewBookRepository(db *gorm.DB) BookRepository {
	return &bookConnection{
		bookConnection: db,
	}
}
