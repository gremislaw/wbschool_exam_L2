#!/bin/bash

for i in $(seq 1 11); do
    # Форматируем номер с ведущим нулем
    formatted_number=$(printf "%.2d" $i)
    
    # Создаем папку с префиксом 'l1.'
    mkdir "dev$formatted_number"
    
    # Создаем файл внутри каждой папки
    file="dev$formatted_number/task.go"
    touch "$file"

    # Записываем содержимое в файл
    echo 'package main' > "$file"
    echo '' >> "$file"  # Пустая строка для разделения
    echo 'import "fmt"' >> "$file"
    echo '' >> "$file"  # Пустая строка для разделения
    echo '' >> "$file"  # Пустая строка для разделения
    echo 'func main() {' >> "$file"
    echo '' >> "$file"  # Пустая строка для разделения
    echo '}' >> "$file"

    # Создаем тест внутри каждой папки
    file="dev$formatted_number/task_test.go"
    touch "$file"

    # Записываем содержимое в файл
    echo 'package main' > "$file"
    echo '' >> "$file"  # Пустая строка для разделения
    echo 'import "fmt"' >> "$file"
    echo '' >> "$file"  # Пустая строка для разделения
    echo '' >> "$file"  # Пустая строка для разделения
    echo 'func main() {' >> "$file"
    echo '' >> "$file"  # Пустая строка для разделения
    echo '}' >> "$file"

done