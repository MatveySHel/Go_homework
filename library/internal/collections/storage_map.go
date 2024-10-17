package collections

import (
	"fmt"
	"library/internal/objects"
)


type Storage_map map[string]objects.Book

//Метод для *Storage_map для поиска книги в структуре
func (s *Storage_map) Search(id string) (int,  bool) {
	_, key := (*s)[id]
	return -1, key
}

//Метод для *Storage_map для добавления книги в структуру
func (s *Storage_map) AddBook(title, author, id string){

	if _, key := (*s)[id]; key {
		book := (*s)[id]
		book.Amount++
		(*s)[id] = book
	} else {
		(*s)[id] = objects.Book{Id: id, Title: title, Author: author, Amount: 1}
	}
}

//Метод для *Storage_map для удаления книги из хранилища
func (s *Storage_map) RemoveBook(id string) bool {
	if _, key := (*s)[id]; key {
		book := (*s)[id]
		book.Amount--
		(*s)[id] = book
		if book.Amount == 0 {
			delete((*s), id)
		}
		return true
	}
	return false
}

//Метод для *Storage_map, который обновлет id у книг в хранилище
func (s *Storage_map) RenewId(id_func func(string,string)string){
	keys := make([]string, 0, len(*s))
	for key := range *s{
		keys = append(keys, key)
	}
	for _ ,key := range keys{
		book := (*s)[key]
		book.Id = id_func(book.Title, book.Author)
		(*s)[book.Id] = book
		delete(*s, key)
	}
}

// Метод для *Storage_map для вывода содержимого хранилища на экран в stdout
func (s *Storage_map) ShowContainer(){
	fmt.Print("library(map): { ")
	for _,v := range (*s){
		fmt.Print(v," ")
	}
	fmt.Print("}\n\n")
}