package app

import (
	"library/internal/utils"
	"fmt"
)

func Search_book(lib utils.Library, k *int, title string, author string){
	_,id,flag := lib.Search(k, title, author)
		if !flag{
			fmt.Printf("No such book: %s '%s'\n\n",author, title)
		}else{
			fmt.Printf("The book %s '%s' with id %s is exists\n\n",author, title, id)
		}
}

func Add_book(lib utils.Library, k *int, title string, author string){
	lib.AddBook(k, title, author)
	fmt.Printf("The book %s '%s' is successfully added\n\n", author, title)
}

func Remove_book(lib utils.Library, k *int, title string, author string){
	flag := lib.RemoveBook(k, title,author)
		if !flag{
			fmt.Printf("The book %s '%s' is not found\n\n", author, title)
		}else{
			fmt.Printf("The book %s '%s' is successfully removed\n\n", author, title)
		}
}

func Renew_storage(lib_pointer *utils.Library, storage string){
	(*lib_pointer) = SetUp_storage(storage)
	fmt.Printf("The storage successfully renewed\n\n")
}

func Renew_id_func(lib utils.Library, k *int){
		*k ++ 
		lib.RenewId(k)
		fmt.Printf("The id-function is successfully renewed\n\n")
}

func Show_storage(lib utils.Library){
	(lib).ShowContainer()
}