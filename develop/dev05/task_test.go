package main

import (
	"os"
	"strings"
	"testing"
)

// recoverFromPanic обрабатывает возможные паники в тестах
func recoverFromPanic(t *testing.T) {
	if r := recover(); r != nil {
		t.Errorf("Тест завершился паникой: %v", r)
	}
}

// createTestFile создает временный файл с содержимым для тестов
func createTestFile(t *testing.T, content string) string {
	tempFile, err := os.CreateTemp("", "testfile_*.txt")
	if err != nil {
		t.Fatal(err)
	}
	tempFile.WriteString(content)
	tempFile.Close()
	return tempFile.Name()
}

// TestGrep проверяет различные флаги утилиты grep
func TestGrep(t *testing.T) {
	testContent := 
`Hello world
This is a test line
ERROR: Something went wrong
Warning: Check your input
Another error occurred
Hello again`

	fileName := createTestFile(t, testContent)
	defer os.Remove(fileName) // Удаляем временный файл после теста

	tests := []struct {
		name     string
		args     []string
		expected []string
	}{
		{"Простое совпадение", []string{"ERROR", fileName}, []string{"ERROR: Something went wrong"}},
		{"Регистронезависимый поиск", []string{"-i", "hello", fileName}, []string{"Hello world", "Hello again"}},
		{"Точное совпадение", []string{"-F", "Warning: Check your input", fileName}, []string{"Warning: Check your input"}},
		{"Инверсия поиска", []string{"-v", "-i", "error", fileName}, []string{"Hello world", "This is a test line", "Warning: Check your input", "Hello again"}},
		{"Вывод номеров строк", []string{"-n", "-i", "error", fileName}, []string{"3 ERROR: Something went wrong", "5 Another error occurred"}},
		{"Подсчет количества строк", []string{"-c", "error", fileName}, []string{"1"}},
		{"Вывод N строк после совпадения", []string{"-A", "1", "ERROR", fileName}, []string{"ERROR: Something went wrong", "Warning: Check your input"}},
		{"Вывод N строк перед совпадением", []string{"-B", "1", "ERROR", fileName}, []string{"This is a test line", "ERROR: Something went wrong"}},
		{"Вывод N строк до и после совпадения", []string{"-C", "1", "ERROR", fileName}, []string{"This is a test line", "ERROR: Something went wrong", "Warning: Check your input"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer recoverFromPanic(t) // Обрабатываем панику

			// Перехватываем вывод программы
			old := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w

			// Запускаем grep с тестовыми аргументами
			os.Args = append([]string{"grep"}, tt.args...)
			main()

			// Читаем вывод
			w.Close()
			os.Stdout = old
			var output strings.Builder
			buf := make([]byte, 1024)
			for {
				n, _ := r.Read(buf)
				if n == 0 {
					break
				}
				output.Write(buf[:n])
			}

			// Проверяем вывод
			outLines := strings.Split(strings.TrimSpace(output.String()), "\n")
			if len(outLines) != len(tt.expected) {
				t.Errorf("Ожидалось %d строк, получено %d", len(tt.expected), len(outLines))
			}
			for i := range tt.expected {
				if i < len(outLines) && strings.TrimSpace(outLines[i]) != tt.expected[i] {
					t.Errorf("Ожидалось: %q, Получено: %q", tt.expected[i], strings.TrimSpace(outLines[i]))
				}
			}
		})
	}
}
