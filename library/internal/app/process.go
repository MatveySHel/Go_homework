package app

import (
	"fmt"
	"library/internal/objects"
)

func Search_book(lib *objects.Library, title string, author string){
	flag := lib.Search(title, author)
		if !flag{
			fmt.Printf("No such book: %s '%s'\n\n",author, title)
		}else{
			fmt.Printf("The book %s '%s' with id %s is exists\n\n",author, title, lib.CalcId(title, author))
		}
}

func Add_book(lib *objects.Library, title string, author string){
	lib.AddBook(title, author)
	fmt.Printf("The book %s '%s' is successfully added\n\n", author, title)
}

func Remove_book(lib *objects.Library, title string, author string){
	flag := lib.RemoveBook(title,author)
		if !flag{
			fmt.Printf("The book %s '%s' is not found\n\n", author, title)
		}else{
			fmt.Printf("The book %s '%s' is successfully removed\n\n", author, title)
		}
}

func Renew_storage(lib_pointer *objects.Library, storage string){
	lib_pointer.LibStorage = SetUp_storage(storage)
	fmt.Printf("The storage successfully renewed\n\n")
}

func Renew_id_func(lib *objects.Library){
		lib.RenewIdFunc()
		fmt.Printf("The id-function is successfully renewed\n\n")
}

func Show_storage(lib *objects.Library){
	lib.ShowContainer()
}