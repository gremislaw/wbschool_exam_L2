package main

import (
	"os"
	"strings"
	"testing"
)

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

// TestSortNumbersWithoutSuffix проверяет обычную числовую сортировку (-n)
func TestSortNumbersWithoutSuffix(t *testing.T) {
	testContent := `3 apples
10 bananas
2 oranges
1 kiwi
5 grapes`

	fileName := createTestFile(t, testContent)
	defer os.Remove(fileName)

	tests := []struct {
		args     []string
		expected []string
	}{
		// Числовая сортировка (-n), игнорируя слова
		{[]string{"-n", fileName}, []string{
			"1 kiwi",
			"2 oranges",
			"3 apples",
			"5 grapes",
			"10 bananas",
		}},
		// Обратная числовая сортировка (-n -r)
		{[]string{"-n", "-r", fileName}, []string{
			"10 bananas",
			"5 grapes",
			"3 apples",
			"2 oranges",
			"1 kiwi",
		}},
	}

	runTests(t, tests)
}

// TestSortNumbers проверяет сортировку чисел, включая суффиксы (K, M, G, T)
func TestSortNumbers(t *testing.T) {
	testContent := `3 apples
10 bananas
2 oranges
1 kiwi
5 grapes
100K units
1M views
500K likes
2G data
1T storage`

	fileName := createTestFile(t, testContent)
	defer os.Remove(fileName)

	tests := []struct {
		args     []string
		expected []string
	}{
		// Числовая сортировка с суффиксами (-h)
		{[]string{"-h", fileName}, []string{
			"1 kiwi",
			"2 oranges",
			"3 apples",
			"5 grapes",
			"10 bananas",
			"100K units",
			"500K likes",
			"1M views",
			"2G data",
			"1T storage",
		}},
		// Обратная числовая сортировка (-h -r)
		{[]string{"-h", "-r", fileName}, []string{
			"1T storage",
			"2G data",
			"1M views",
			"500K likes",
			"100K units",
			"10 bananas",
			"5 grapes",
			"3 apples",
			"2 oranges",
			"1 kiwi",
		}},
	}

	runTests(t, tests)
}

// TestSortMonths проверяет сортировку по названиям месяцев и временам года
func TestSortMonths(t *testing.T) {
	testContent := `Jan winter
Mar spring
Feb winter
Dec winter`

	fileName := createTestFile(t, testContent)
	defer os.Remove(fileName)

	tests := []struct {
		args     []string
		expected []string
	}{
		// Сортировка по временам года
		{[]string{fileName}, []string{
			"Dec winter",
			"Feb winter",
			"Jan winter",
			"Mar spring",
		}},
		// Сортировка по месяцам (-M)
		{[]string{"-M", fileName}, []string{
			"Jan winter",
			"Feb winter",
			"Mar spring",
			"Dec winter",
		}},
		// Обратная сортировка по месяцам (-M -r)
		{[]string{"-M", "-r", fileName}, []string{
			"Dec winter",
			"Mar spring",
			"Feb winter",
			"Jan winter",
		}},
	}

	runTests(t, tests)
}

// runTests — универсальная функция для тестирования сортировки
func runTests(t *testing.T, tests []struct {
	args     []string
	expected []string
}) {
	for _, tt := range tests {
		t.Run(strings.Join(tt.args, " "), func(t *testing.T) {
			// Перехватываем вывод программы
			old := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w

			// Запускаем sort с тестовыми аргументами
			os.Args = append([]string{"sort"}, tt.args...)
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
				if outLines[i] != tt.expected[i] {
					t.Errorf("Ожидалось: %q, Получено: %q", tt.expected[i], outLines[i])
				}
			}
		})
	}
}
