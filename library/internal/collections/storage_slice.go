package collections

import (
	"fmt"
	"library/internal/objects"
)
type Storage_slice []objects.Book

//Метод для *Storage_slice для поиска книги в структуре
func (s *Storage_slice) Search(id string) (int, bool) {

	for index, book := range (*s) {
		if book.Id == id {
			return index, true
		}
	}
	return -1, false
}

//Метод для *Storage_slice для добавления книги в структуру
func (s *Storage_slice) AddBook(title, author, id string){
	index := 0
	for index < len(*s){
		if (*s)[index].Id == id {
			break
		}else{
			index++
		}
	}
	if index != len(*s){
		book := (*s)[index]
		book.Amount++
		(*s)[index] = book
	} else {
		(*s) = append(*s, objects.Book{Id: id, Title: title, Author: author, Amount: 1})

	}
}

//Метод для *Storage_slice для удаления книги из хранилища
func (s *Storage_slice) RemoveBook(id string) bool {
	index := 0
	for index < len(*s){
		if (*s)[index].Id == id {
			break
		}else{
			index++
		}
	}
	if index != len(*s) {
		book := (*s)[index]
		book.Amount--
		(*s)[index] = book
		if book.Amount == 0 {
			(*s) = append((*s)[:index], (*s)[index+1:]...)
		}
		return true
	}
	return false
}

//Метод для *Storage_slice, который обновлет id у книг в хранилище
func (s *Storage_slice) RenewId(id_func func(string,string)string){
	for i, book := range *s{
		book.Id = id_func(book.Title, book.Author)
		(*s)[i] = book
	}
}


// Метод для *Storage_map для вывода содержимого хранилища на экран в stdout
func (s *Storage_slice) ShowContainer(){
	fmt.Print("library(slice): [ ")
	for _,v := range (*s){
		fmt.Print(v," ")
	}
	fmt.Print("]\n\n")
}