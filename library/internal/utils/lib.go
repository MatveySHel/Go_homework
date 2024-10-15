package utils

// Интерфейс библиотеки, которому должны удволетворять хранилища
type Library interface {
	Search(k *int, title, author string) (int, string, bool)
	AddBook(k *int, title, author string)
	RemoveBook(k *int, title, author string) bool
	RenewId(k *int)
	ShowContainer()
}