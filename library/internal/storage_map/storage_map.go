package storage_map

import (
	"library/internal/utils"
	"fmt"
)

type Storage_map map[string]utils.Book

//Метод для *Storage_map для поиска книги в структуре
func (s *Storage_map) Search(k *int, title, author string) (int, string, bool) {
	id := utils.Calc_id(k, fmt.Sprintf("%v_%v", title, author))
	_, key := (*s)[id]
	return -1, id, key
}

//Метод для *Storage_map для добавления книги в структуру
func (s *Storage_map) AddBook(k *int,title, author string){

	if _, id, exsists := s.Search(k ,title, author); exsists {
		book := (*s)[id]
		book.Amount++
		(*s)[id] = book
	} else {
		(*s)[id] = utils.Book{Id: id, Title: title, Author: author, Amount: 1}
	}
}

//Метод для *Storage_map для удаления книги из хранилища
func (s *Storage_map) RemoveBook(k *int,title, author string) bool {
	if _, id, exsists := s.Search(k, title, author); exsists {
		book := (*s)[id]
		book.Amount--
		(*s)[id] = book
		if book.Amount == 0 {
			delete(*s, id)
		}
		return true
	}
	return false
}

//Метод для *Storage_map, который обновлет id у книг в хранилище
func (s *Storage_map) RenewId(k *int){
	keys := make([]string, 0, len(*s))
	for key := range *s{
		keys = append(keys, key)
	}
	for _ ,key := range keys{
		book := (*s)[key]
		book.Id = utils.Calc_id(k , fmt.Sprintf("%v_%v", book.Title, book.Author))
		(*s)[book.Id] = book
		delete(*s, key)
	}
}

// Метод для *Storage_map для вывода содержимого хранилища на экран в stdout
func (s *Storage_map) ShowContainer(){
	fmt.Print("library(map): { ")
	for _,v := range *s{
		fmt.Print(v," ")
	}
	fmt.Print("}\n\n")
}