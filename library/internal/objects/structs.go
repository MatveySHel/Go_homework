package objects

// Описываем структуры книг
type Book struct {
	Id     string
	Title  string
	Author string
	Amount uint
}

// Интерфейс хранилищ, которому должны удволетворять хранилища
type Storage interface {
	Search(string) (int, bool)
	AddBook(string, string, string)
	RemoveBook(string) bool
	RenewId(func(string,string)string)
	ShowContainer()
}


type Idgenerator struct {
	Id_func func(int, string) string
	Key int
}