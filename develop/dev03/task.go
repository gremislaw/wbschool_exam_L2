package main

/*
=== Утилита sort ===

Отсортировать строки (man sort)
Основное

Поддержать ключи

-k — указание колонки для сортировки
-n — сортировать по числовому значению
-r — сортировать в обратном порядке
-u — не выводить повторяющиеся строки

Дополнительное

Поддержать ключи

-M — сортировать по названию месяца
-b — игнорировать хвостовые пробелы
-c — проверять отсортированы ли данные
-h — сортировать по числовому значению с учётом суффиксов

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
	"unicode"
)

// customSort хранит параметры сортировки и реализует интерфейс sort.Interface
type customSort struct {
	Lines             []string
	Column            int
	Numeric           bool
	Reverse           bool
	Unique            bool
	Month             bool
	IgnoreSpaces      bool
	CheckSorted       bool
	NumericWithSuffix bool
}

func (cs *customSort) Len() int {
	return len(cs.Lines)
}

func (cs *customSort) Swap(i, j int) {
	cs.Lines[i], cs.Lines[j] = cs.Lines[j], cs.Lines[i]
}

func (cs *customSort) Less(i, j int) bool {
	fieldsI := strings.Fields(cs.Lines[i])
	fieldsJ := strings.Fields(cs.Lines[j])

	if cs.Column > len(fieldsI) || cs.Column > len(fieldsJ) {
		return false
	}

	fieldI := fieldsI[cs.Column-1]
	fieldJ := fieldsJ[cs.Column-1]

	if cs.Numeric {
		numI, errI := strconv.ParseFloat(fieldI, 64)
		numJ, errJ := strconv.ParseFloat(fieldJ, 64)
		if errI != nil || errJ != nil {
			log.Fatal("Ошибка преобразования в число")
		}
		return (numI < numJ) != cs.Reverse
	}

	if cs.Month {
		timeI, errI := time.Parse("Jan", fieldI)
		timeJ, errJ := time.Parse("Jan", fieldJ)
		if errI != nil || errJ != nil {
			log.Fatal("Ошибка обработки месяца")
		}
		return (timeI.Before(timeJ)) != cs.Reverse
	}

	if cs.NumericWithSuffix {
		numI := parseWithSuffix(fieldI)
		numJ := parseWithSuffix(fieldJ)
		return (numI < numJ) != cs.Reverse
	}

	return (fieldI < fieldJ) != cs.Reverse
}

// parseWithSuffix конвертирует числа с суффиксами (K, M, G и т. д.)
func parseWithSuffix(s string) float64 {
	s = strings.TrimSpace(s)
	var i int
	for i = 0; i < len(s); i++ {
		if !unicode.IsDigit(rune(s[i])) && s[i] != '.' {
			break
		}
	}
	numPart := s[:i]
	suffixPart := s[i:]
	num, err := strconv.ParseFloat(numPart, 64)
	if err != nil {
		log.Fatal("Ошибка обработки числа с суффиксом")
	}

	switch strings.ToUpper(suffixPart) {
	case "K":
		num *= 1e3
	case "M":
		num *= 1e6
	case "G":
		num *= 1e9
	case "T":
		num *= 1e12
	case "P":
		num *= 1e15
	case "E":
		num *= 1e18
	}

	return num
}

// removeDuplicates удаляет дубликаты из списка строк
func removeDuplicates(data *customSort) {
	lines := make(map[string]struct{})
	var uniqueLines []string
	for _, line := range data.Lines {
		if _, exists := lines[line]; !exists {
			lines[line] = struct{}{}
			uniqueLines = append(uniqueLines, line)
		}
	}
	data.Lines = uniqueLines
}

// parseFlags считывает аргументы командной строки и возвращает структуру с настройками
func parseFlags() *customSort {
	data := &customSort{}
	flag.IntVar(&data.Column, "k", 1, "Столбец для сортировки (по умолчанию 1)")
	flag.BoolVar(&data.Numeric, "n", false, "Сортировка по числовому значению")
	flag.BoolVar(&data.Reverse, "r", false, "Обратная сортировка")
	flag.BoolVar(&data.Unique, "u", false, "Вывод только уникальных строк")
	flag.BoolVar(&data.Month, "M", false, "Сортировка по названию месяца")
	flag.BoolVar(&data.IgnoreSpaces, "b", false, "Игнорировать завершающие пробелы")
	flag.BoolVar(&data.CheckSorted, "c", false, "Проверить, отсортирован ли файл")
	flag.BoolVar(&data.NumericWithSuffix, "h", false, "Числовая сортировка с учетом суффиксов (K, M, G и т. д.)")

	flag.Parse()
	return data
}

// readLines читает строки из файла и возвращает их
func readLines(filename string, ignoreSpaces bool) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("ошибка открытия файла: %w", err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if ignoreSpaces {
			line = strings.TrimRight(line, " ")
		}
		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("ошибка чтения файла: %w", err)
	}

	return lines, nil
}

// checkIfSorted проверяет, отсортирован ли файл
func checkIfSorted(data *customSort) {
	if sort.IsSorted(data) {
		fmt.Println("Файл отсортирован")
	} else {
		fmt.Println("Файл не отсортирован")
	}
}

// sortAndPrint выполняет сортировку и выводит результат
func sortAndPrint(data *customSort) {
	sort.Sort(data)

	if data.Unique {
		removeDuplicates(data)
	}

	for _, line := range data.Lines {
		fmt.Println(line)
	}
}

func mySort(filename string) {
	data := parseFlags()

	lines, err := readLines(filename, data.IgnoreSpaces)
	if err != nil {
		log.Fatal(err)
	}
	data.Lines = lines

	if data.CheckSorted {
		checkIfSorted(data)
		return
	}

	sortAndPrint(data)
}

func main() {
	mySort("test1") // числа
	//mySort("test2") // месяцы
}
