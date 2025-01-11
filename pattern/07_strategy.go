package pattern

/*
	Реализовать паттерн «стратегия».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Strategy_pattern
*/

// Паттерн Strategy определяет набор алгоритмов схожих по роду деятельности, инкапсулирует их в отдельный класс и делает их подменяемыми.
// Паттерн Strategy позволяет подменять алгоритмы без участия клиентов, которые используют эти алгоритмы.
//
// Пример: Создание интерфейса PaymentStrategy, который определяет метод pay().
// Подклассы, такие как CreditCardPayment и PayPalPayment, реализуют конкретные алгоритмы обработки платежей.

// + Алгоритмы могут быть заменены во время выполнения
// + Алгоритмы изолированы в собственных классах
// + Новые стратегии могут быть добавлены без изменения существующего кода
// - Необходимость знания различий между стратегиями
// - Увеличение количества классов и объектов

// StrategySort - интерфейс для сортирующих алгоритмов.
type StrategySort interface {
	Sort([]int)
}

// BubbleSort реализует алгоритм "Пузырьковая сортировка".
type BubbleSort struct {
}

// Sort сортирует массив чисел.
func (s *BubbleSort) Sort(a []int) {
	size := len(a)
	if size < 2 {
		return
	}
	for i := 0; i < size; i++ {
		for j := size - 1; j >= i+1; j-- {
			if a[j] < a[j-1] {
				a[j], a[j-1] = a[j-1], a[j]
			}
		}
	}
}

// InsertionSort реализует алгоритм "Сортировка вставками".
type InsertionSort struct {
}

// Sort сортирует массив чисел.
func (s *InsertionSort) Sort(a []int) {
	size := len(a)
	if size < 2 {
		return
	}
	for i := 1; i < size; i++ {
		var j int
		var buff = a[i]
		for j = i - 1; j >= 0; j-- {
			if a[j] < buff {
				break
			}
			a[j+1] = a[j]
		}
		a[j+1] = buff
	}
}

// Context контекст, исполняющий стратегию
type Context struct {
	strategy StrategySort
}

// Algorithm меняет стратегию.
func (c *Context) Algorithm(a StrategySort) {
	c.strategy = a
}

// Sort сортирует массив чисел по выбранной стратегии.
func (c *Context) Sort(s []int) {
	c.strategy.Sort(s)
}