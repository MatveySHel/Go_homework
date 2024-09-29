package library


// Программа, имитирующая управление библиотекой



import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strings"
)

const MOD = 1_000_000_007


// Описываем структуры книг
type Book struct {
	Id     string
	Title  string
	Author string
	Amount int
}

// Фунция, которая вычитывает id книги(по названию и автору) на основе полиномиального хеша
// Принимает параметр k, который обновляется каждый раз, когда нужно изменить id-функцию
func Calc_id(k *int, str string)string {
	var id int
	runes := []rune(str)
	for i, val := range runes{
		id = (id*(i+(*k)) + int(val) + *k)%MOD
	}
	return fmt.Sprintf("%d", id)	
}

// Интерфейс библиотеки, которому должны удволетворять хранилища
type Library interface {
	Search(k *int, title, author string) (int, string, bool)
	AddBook(k *int, title, author string)
	RemoveBook(k *int, title, author string) bool
	RenewId(k *int)
	ShowContainer()
}

// Тип Storage_map на основе map для хранения книг
type Storage_map map[string]Book

// Тип Storage_slice на основе slice для хранения книг
type Storage_slice []Book

//Метод для *Storage_map для поиска книги в структуре
func (s *Storage_map) Search(k *int, title, author string) (int, string, bool) {
	id := Calc_id(k, fmt.Sprintf("%v_%v", title, author))
	_, key := (*s)[id]
	return -1, id, key
}

//Метод для *Storage_slice для поиска книги в структуре
func (s *Storage_slice) Search(k *int,title, author string) (int, string, bool) {
	id := Calc_id(k, fmt.Sprintf("%v_%v", title, author))
	for index, book := range *s {
		if book.Id == id {
			return index, id, true
		}
	}
	return -1, id, false
}

//Метод для *Storage_map для добавления книги в структуру
func (s *Storage_map) AddBook(k *int,title, author string){

	if _, id, exsists := s.Search(k ,title, author); exsists {
		book := (*s)[id]
		book.Amount++
		(*s)[id] = book
	} else {
		(*s)[id] = Book{Id: id, Title: title, Author: author, Amount: 1}
	}
}

//Метод для *Storage_slice для добавления книги в структуру
func (s *Storage_slice) AddBook(k *int,title, author string){
	if index, id, exsists := s.Search(k, title, author); exsists {
		book := (*s)[index]
		book.Amount++
		(*s)[index] = book
	} else {
		*s = append((*s), Book{Id: id, Title: title, Author: author, Amount: 1})

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

//Метод для *Storage_map, который обновлет id у книг в хранилище
func (s *Storage_map) RenewId(k *int){
	keys := make([]string, 0, len(*s))
	for key := range *s{
		keys = append(keys, key)
	}
	for _ ,key := range keys{
		book := (*s)[key]
		book.Id = Calc_id(k , fmt.Sprintf("%v_%v", book.Title, book.Author))
		(*s)[book.Id] = book
		delete(*s, key)
	}
}

//Метод для *Storage_slice, который обновлет id у книг в хранилище
func (s *Storage_slice) RenewId(k *int){
	for i, v := range *s{
		v.Id = Calc_id(k , fmt.Sprintf("%v_%v", v.Title, v.Author))
		(*s)[i] = v
	}
}

// Метод для *Storage_map для вывода содержимого хранилища на экран в stdout
func (s *Storage_map) ShowContainer(){
	fmt.Print(">library(map): { ")
	for _,v := range *s{
		fmt.Print(v," ")
	}
	fmt.Print("}\n")
}

// Метод для *Storage_map для вывода содержимого хранилища на экран в stdout
func (s *Storage_slice) ShowContainer(){
	fmt.Print(">library(slice): [ ")
	for _,v := range *s{
		fmt.Print(v," ")
	}
	fmt.Print("]\n")
}


// Функция для обнавления хранилища
func SetUp_storage(arg string) Library {
    if arg == "map" {
        return &Storage_map{}
    } else if arg == "slice" {
        return &Storage_slice{}
    }
    return nil
}

// Функция для парсинга запроса
func ParseQuiry(input string) (parsed map[string]string, err error) {
	parsed = make(map[string]string)
	if input == "renew_id_func" || input == "show_storage"{
		parsed["Command"] = input
		return
	}

	splited_by_colon := strings.SplitN(input, ":", 2)
	if match, _ := regexp.MatchString(`^renew_storage: (map|slice)`, input); match{
		parsed["Command"] = splited_by_colon[0]
		parsed["Storage"] = strings.TrimSpace(splited_by_colon[1])
		return
	}
	match, _ := regexp.MatchString(`^(search_book|add_book|remove_book): [A-ZА-ЯЁ][A-ZА-ЯЁa-zа-яё. -]* \'[A-ZА-ЯЁ].*\'$`, input)
	if !match {
		err = fmt.Errorf(">Error: %w", errors.New("invalid command"))
	}else{
		splited_by_q := strings.SplitN(splited_by_colon[1], "'", 2)
		parsed["Command"] = splited_by_colon[0]
		parsed["Author"] = strings.TrimSpace(splited_by_q[0])
		parsed["Title"] = strings.Trim(splited_by_q[1], "' ")
	}
	return
}


// Функция для обработки запроса
func ProccessQuiry(lib_pointer *Library, k *int, parsed map[string]string){
	switch parsed["Command"]{
	case "search_book":
		_,_,flag := (*lib_pointer).Search(k, parsed["Title"], parsed["Author"])
		if !flag{
			fmt.Printf(">No such book: %s %s\n",parsed["Title"], parsed["Author"])
		}else{
			fmt.Printf(">The book %s %s is exists\n",parsed["Title"], parsed["Author"])
		}
	case "add_book":
		(*lib_pointer).AddBook(k, parsed["Title"], parsed["Author"])
		fmt.Println(">The book successfully added")
	case "remove_book":
		flag := (*lib_pointer).RemoveBook(k, parsed["Title"], parsed["Author"])
			if !flag{
				fmt.Println(">The book is not found")
			}else{
				fmt.Println(">The book successfully removed")
			}
	case "renew_storage":
		(*lib_pointer) = SetUp_storage(parsed["Storage"])
	case "renew_id_func":
		*k ++ 
		(*lib_pointer).RenewId(k)
	case "show_storage":
		(*lib_pointer).ShowContainer()

	}
}

// Точка входа в программу
func LibraryLaunch() {
	fmt.Print(">Type following commands:(\n" +
		"  search_book: <author> '<title>',\n" +
		"  add_book: <author> '<title>',\n" +
		"  remove_book: <author> '<title>'\n" +
		"  renew_storage: {slice, map}')\n"  +
		"  renew_id_func\n" +
		"  show_storage )\n" )
	var k int
	lib := SetUp_storage("map")
	lib.ShowContainer()
	for {
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		if err != nil {
			fmt.Println(">Input error:", err)
			return
		}
		if strings.ToLower(strings.TrimSpace(input)) == "exit" {
			fmt.Println(">Input stopped")
			return
		}

		parsed, err2 := ParseQuiry(input)
		if err2 != nil {
			fmt.Println(err2)
			continue
		}
		ProccessQuiry(&lib, &k, parsed)
	}
}

//func main(){
//	LibraryLaunch()
//}
