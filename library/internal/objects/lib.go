package objects

import "fmt"

type Library struct {
	LibStorage Storage
	Id_func Idgenerator
}

func (l *Library) CalcId(title, author string) string {
	return l.Id_func.Id_func(l.Id_func.Key, fmt.Sprintf("%v_%v", title, author))
}
func (l *Library) Search(title, author string) bool {
	id := l.CalcId(title, author)
	_, isExist := l.LibStorage.Search(id)
	return isExist
}

func (l *Library) AddBook(title, author string) {
	id := l.CalcId(title, author)
	l.LibStorage.AddBook(title, author, id)

}

func (l *Library) RemoveBook(title, author string) bool {
	id := l.CalcId(title, author)
	return l.LibStorage.RemoveBook(id)
}

func (l *Library) ShowContainer() {
	l.LibStorage.ShowContainer()
}

func (l *Library) RenewIdFunc() {
	l.Id_func.Key++
	l.LibStorage.RenewId(l.CalcId)

}
