package app

import 	"library/internal/objects"

func RunScenario(){
	generator := objects.Idgenerator{Id_func: objects.Calc_id, Key: 0}
	lib := &objects.Library{LibStorage: SetUp_storage("map"), Id_func: generator}
	Show_storage(lib)
	Add_book(lib, "Дубровский", "Пушкин А. С.")
	Add_book(lib, "Война и мир", "Толстой Л. Н.")
	Add_book(lib, "Дубровский", "Пушкин А. С.")
	Add_book(lib, "Герой нашего времени", "Лермонтов М. Ю.")
	Show_storage(lib)
	Search_book(lib, "Война и мир", "Толстой Л. Н.")
	Search_book(lib, "Мцыри", "Лермонтов М. Ю.")
	Remove_book(lib, "Герой нашего времени", "Лермонтов М. Ю.")
	Show_storage(lib)
	Renew_storage(lib, "slice")
	Show_storage(lib)
	Add_book(lib, "Дубровский", "Пушкин А. С.")
	Add_book(lib, "Война и мир", "Толстой Л. Н.")
	Show_storage(lib)
	Search_book(lib, "Война и мир", "Толстой Л. Н.")
	Search_book(lib,"Преступление и наказание", "Достоевский Ф. М.")
	Renew_id_func(lib)
	Show_storage(lib)
	Renew_id_func(lib)
	Show_storage(lib)
	Renew_id_func(lib)
	Show_storage(lib)
	Search_book(lib, "Война и мир", "Толстой Л. Н.")
}