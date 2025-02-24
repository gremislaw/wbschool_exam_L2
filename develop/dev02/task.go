package main

/*
=== Задача на распаковку ===

Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:
	- "a4bc2d5e" => "aaaabccddddde"
	- "abcd" => "abcd"
	- "45" => "" (некорректная строка)
	- "" => ""
Дополнительное задание: поддержка escape - последовательностей
	- qwe\4\5 => qwe45 (*)
	- qwe\45 => qwe44444 (*)
	- qwe\\5 => qwe\\\\\ (*)

В случае если была передана некорректная строка функция должна возвращать ошибку. Написать unit-тесты.

Функция должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/


import (
	"errors"
	"strconv"
	"unicode"
)

// Unpack распаковывает строку с учетом повторений и escape-последовательностей
func Unpack(input string) (string, error) {
	if len(input) == 0 {
		return "", nil
	}

	var result []rune
	var lastChar rune  // предыдущий символ
	var isEscaped bool // флаг для обработки escape-последовательностей

	for _, char := range input {
		switch {
		case isEscaped:
			// Обрабатываем символ после '\'
			if err := handleEscape(&result, char); err != nil {
				return "", err
			}
			isEscaped = false
			lastChar = char
		case char == '\\':
			// Обнаружен символ '\', включаем режим escape
			isEscaped = true
		case unicode.IsLetter(char):
			// Если буква, добавляем в результат
			lastChar = addChar(&result, char)
		case unicode.IsDigit(char):
			// Если цифра, повторяем предыдущий символ
			if err := repeatLastChar(&result, lastChar, char); err != nil {
				return "", err
			}
		default:
			// Встречен недопустимый символ
			return "", errors.New("invalid string: contains non-alphanumeric characters")
		}
	}

	// Если строка оканчивается на '\', это ошибка
	if isEscaped {
		return "", errors.New("invalid escape sequence")
	}

	return string(result), nil
}

// handleEscape добавляет символ после '\' в результат
func handleEscape(result *[]rune, char rune) error {
	if !unicode.IsDigit(char) && char != '\\' {
		return errors.New("invalid escape sequence")
	}
	*result = append(*result, char)
	return nil
}

// repeatLastChar повторяет предыдущий символ указанное число раз
func repeatLastChar(result *[]rune, lastChar rune, char rune) error {
	if lastChar == 0 {
		return errors.New("invalid string")
	}
	count, _ := strconv.Atoi(string(char))
	for j := 1; j < count; j++ {
		*result = append(*result, lastChar)
	}
	return nil
}

// addChar добавляет символ в результат и возвращает его как последний обработанный
func addChar(result *[]rune, char rune) rune {
	*result = append(*result, char)
	return char
}
