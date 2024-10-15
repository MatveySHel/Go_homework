package utils

import "fmt"

const MOD = 1_000_000_007

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