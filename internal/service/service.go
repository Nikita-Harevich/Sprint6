package service

import (
	"errors"
	"log"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func ConvertSrting(s string) (string, error) {

	if len(s) == 0 {
		log.Print("Ошибка нулевая строка")
		return "", errors.New("длинна строки 0")
	}

	if string(s[0]) != "." || string(s[0]) != "-" {
		return morse.ToMorse(s), nil
	}

	return morse.ToText(s), nil
}
