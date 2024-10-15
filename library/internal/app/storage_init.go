package app

import (
	"library/internal/utils"
	"library/internal/storage_map"
	"library/internal/storage_slice"

)

// Функция для обнавления хранилища
func SetUp_storage(arg string) utils.Library {
    if arg == "map" {
        return &storage_map.Storage_map{}
    } else if arg == "slice" {
        return &storage_slice.Storage_slice{}
    }
    return nil
}