package service

import (
	"fmt"
	"log"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func IsMorze(str string) (bool, error) {
	runes := []rune(str)
	if len(runes) == 0 {
		return false, fmt.Errorf("empty string")
	}
	for _, v := range runes {
		if v != '.' && v != '-' && v != ' ' {
			return false, nil
		}
	}
	return true, nil
}

func Translate(str string) string {
	res, err := IsMorze(str)
	if err != nil {
		log.Println(err)
		return ""
	}
	if res {
		return morse.ToText(str)
	}
	return morse.ToMorse(str)
}
