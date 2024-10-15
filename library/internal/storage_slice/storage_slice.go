package storage_slice

import (
	"library/internal/utils"
	"fmt"
)
type Storage_slice []utils.Book

//Метод для *Storage_slice для поиска книги в структуре
func (s *Storage_slice) Search(k *int,title, author string) (int, string, bool) {
	id := utils.Calc_id(k, fmt.Sprintf("%v_%v", title, author))
	for index, book := range *s {
		if book.Id == id {
			return index, id, true
		}
	}
	return -1, id, false
}

//Метод для *Storage_slice для добавления книги в структуру
func (s *Storage_slice) AddBook(k *int,title, author string){
	if index, id, exsists := s.Search(k, title, author); exsists {
		book := (*s)[index]
		book.Amount++
		(*s)[index] = book
	} else {
		*s = append((*s), utils.Book{Id: id, Title: title, Author: author, Amount: 1})

	}
}

//Метод для *Storage_slice для удаления книги из хранилища
func (s *Storage_slice) RemoveBook(k *int,title, author string) bool {
	if index, _, exsists := s.Search(k, title, author); exsists {
		book := (*s)[index]
		book.Amount--
		(*s)[index] = book
		if book.Amount == 0 {
			*s = append((*s)[:index], (*s)[index+1:]...)
		}
		return true
	}
	return false
}

//Метод для *Storage_slice, который обновлет id у книг в хранилище
func (s *Storage_slice) RenewId(k *int){
	for i, v := range *s{
		v.Id = utils.Calc_id(k , fmt.Sprintf("%v_%v", v.Title, v.Author))
		(*s)[i] = v
	}
}


// Метод для *Storage_map для вывода содержимого хранилища на экран в stdout
func (s *Storage_slice) ShowContainer(){
	fmt.Print("library(slice): [ ")
	for _,v := range *s{
		fmt.Print(v," ")
	}
	fmt.Print("]\n\n")
}