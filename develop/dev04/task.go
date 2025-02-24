package main

/*
=== Поиск анаграмм по словарю ===

Напишите функцию поиска всех множеств анаграмм по словарю.
Например:
'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
'листок', 'слиток' и 'столик' - другому.

Входные данные для функции: ссылка на массив - каждый элемент которого - слово на русском языке в кодировке utf8.
Выходные данные: Ссылка на мапу множеств анаграмм.
Ключ - первое встретившееся в словаре слово из множества
Значение - ссылка на массив, каждый элемент которого, слово из множества. Массив должен быть отсортирован по возрастанию.
Множества из одного элемента не должны попасть в результат.
Все слова должны быть приведены к нижнему регистру.
В результате каждое слово должно встречаться только один раз.

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

import (
	"fmt"
	"sort"
	"strings"
)

// findAnagrams ищет все группы анаграмм в списке слов
func findAnagrams(words []string) map[string][]string {
	anagramGroups := make(map[string][]string)
	seenWords := make(map[string]bool)

	for _, word := range words {
		lowerWord := strings.ToLower(word)
		sortedWord := sortString(lowerWord)

		if !seenWords[lowerWord] {
			anagramGroups[sortedWord] = append(anagramGroups[sortedWord], lowerWord)
			seenWords[lowerWord] = true
		}
	}

	return filterAnagramGroups(anagramGroups)
}

// sortString сортирует буквы в слове в алфавитном порядке
func sortString(s string) string {
	runes := []rune(s)
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})
	return string(runes)
}

// filterAnagramGroups удаляет группы, где только одно слово, и сортирует группы
func filterAnagramGroups(groups map[string][]string) map[string][]string {
	result := make(map[string][]string)
	for _, group := range groups {
		if len(group) > 1 {
			sort.Strings(group)
			result[group[0]] = group
		}
	}
	return result
}

// printAnagrams выводит результат
func printAnagrams(anagrams map[string][]string) {
	for key, group := range anagrams {
		fmt.Printf("%s: %v\n", key, group)
	}
}

func main() {
	words := []string{"пятак", "пятка", "пятка", "тяпка", "листок", "слиток", "столик", "слово"}
	anagrams := findAnagrams(words)
	printAnagrams(anagrams)
}
