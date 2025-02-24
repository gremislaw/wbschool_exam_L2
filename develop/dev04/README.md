# Утилита для поиска анаграмм
Программа находит группы анаграмм в списке слов.

## Установка
```sh
go build -o anagrams ./anagrams
```

**Пример использования:**
```sh
./anagrams wordlist.txt
```

Вывод:
```
[listen silent enlist]
[evil vile live]
```
---

## Тестирование
Запуск тестов:
```sh
go test ./...
```

