package app

import (
	"library/internal/objects"
	"library/internal/collections"

)

// Функция для обнавления хранилища
func SetUp_storage(arg string) objects.Storage {
    if arg == "map" {
        return &collections.Storage_map{}
    } else if arg == "slice" {
        return &collections.Storage_slice{}
    }
    return nil
}