package main

/*
=== Утилита grep ===

Реализовать утилиту фильтрации (man grep)

Поддержать флаги:
-A - "after" печатать +N строк после совпадения
-B - "before" печатать +N строк до совпадения
-C - "context" (A+B) печатать ±N строк вокруг совпадения
-c - "count" (количество строк)
-i - "ignore-case" (игнорировать регистр)
-v - "invert" (вместо совпадения, исключать)
-F - "fixed", точное совпадение со строкой, не паттерн
-n - "line num", печатать номер строки

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

// GrepConfig для хранения флагов
type GrepConfig struct {
	After      int
	Before     int
	Context    int
	Count      bool
	IgnoreCase bool
	Inverted   bool
	Fixed      bool
	LineNum    bool
	Pattern    string
	Filename   string
}

// Функция для парсинга аргументов
func parseFlags(args []string) GrepConfig {
	fs := flag.NewFlagSet("grep", flag.ExitOnError)
	config := GrepConfig{}
	fs.IntVar(&config.After, "A", 0, "Вывести +N строк после совпадения")
	fs.IntVar(&config.Before, "B", 0, "Вывести +N строк перед совпадением")
	fs.IntVar(&config.Context, "C", 0, "Вывести +-N строк вокруг совпадения")
	fs.BoolVar(&config.Count, "c", false, "Вывести количество строк с совпадением")
	fs.BoolVar(&config.IgnoreCase, "i", false, "Игнорировать регистр символов")
	fs.BoolVar(&config.Inverted, "v", false, "Исключить совпадающие строки")
	fs.BoolVar(&config.Fixed, "F", false, "Точное совпадение со строкой, а не с шаблоном")
	fs.BoolVar(&config.LineNum, "n", false, "Вывести номера строк")

	fs.Parse(args)

	nonFlags := fs.Args()
	if len(nonFlags) != 2 {
		log.Panic("Требуется два аргумента: строка для поиска и имя файла")
	}
	config.Pattern = nonFlags[0]
	config.Filename = nonFlags[1]

	return config
}

// Функция match проверяет, соответствует ли строка заданному шаблону с учетом опций
func match(line, pattern string, fixed, ignoreCase bool) bool {
	if ignoreCase {
		line = strings.ToLower(line)
		pattern = strings.ToLower(pattern)
	}
	if fixed {
		return line == pattern
	}
	return strings.Contains(line, pattern)
}

// Функция processFile обрабатывает файл с заданными параметрами
func grep(config GrepConfig) {
	file, err := os.Open(config.Filename)
	if err != nil {
		log.Panic(err)
	}
	defer file.Close()

	var matches []int
	var lines []string
	scanner := bufio.NewScanner(file)
	for i := 0; scanner.Scan(); i++ {
		str := scanner.Text()
		if match(str, config.Pattern, config.Fixed, config.IgnoreCase) != config.Inverted {
			matches = append(matches, i)
		}
		lines = append(lines, str)
	}

	if err = scanner.Err(); err != nil {
		log.Panic(err)
	}

	if config.Count {
		fmt.Println(len(matches))
		return
	}

	if config.Context > 0 {
		config.Before = config.Context
		config.After = config.Context
	}

	printed := make(map[int]struct{})
	for _, num := range matches {
		start := num - config.Before
		if start < 0 {
			start = 0
		}
		end := num + config.After
		if end >= len(lines) {
			end = len(lines) - 1
		}

		for i := start; i <= end; i++ {
			if _, ok := printed[i]; !ok {
				if config.LineNum {
					fmt.Print(i+1, " ")
				}
				fmt.Println(lines[i])
				printed[i] = struct{}{}
			}
		}
	}
}

func main() {
	config := parseFlags(os.Args[1:])
	grep(config)
}
