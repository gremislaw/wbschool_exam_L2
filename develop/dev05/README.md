# Grep


### Установка
```sh
go build -o grep ./grep
```


### 📌 Описание
Утилита `grep` позволяет искать строки в файле по заданному шаблону с поддержкой дополнительных флагов:

### 🔧 Поддерживаемые флаги
- `-n` — показать номера строк с совпадениями
- `-i` — регистронезависимый поиск
- `-v` — инвертировать поиск (показать строки без совпадений)
- `-c` — вывести количество найденных строк
- `-F` — точное совпадение строки
- `-A <N>` — вывести N строк после совпадения
- `-B <N>` — вывести N строк перед совпадением
- `-C <N>` — вывести N строк до и после совпадения

### 🚀 Запуск
```sh
./grep [ФЛАГИ] "шаблон" файл.txt
```

### ✅ Примеры использования
```sh
# Найти "error" в файле log.txt с номерами строк
./grep -n "error" log.txt

# Вывести количество строк с "failed"
./grep -c "failed" log.txt
```

### 🛠 Тесты
Тесты находятся в файле `task_test.go`.
Запуск тестов:
```sh
 go test -v
```
Тесты проверяют:
- Поиск строки в файле
- Работу флагов (-n, -i, -v, -c, -F, -A, -B, -C)
- Обработку ошибок и корректность вывода
- Отсутствие паник с `recover()`