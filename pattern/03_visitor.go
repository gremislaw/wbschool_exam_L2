package pattern

/*
	Реализовать паттерн «посетитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Visitor_pattern
*/

// Паттерн Visitor позволяет обойти набор элементов (объектов) с разнородными интерфейсами,
// а также позволяет добавить новый метод в класс объекта, при этом, не изменяя сам класс этого объекта.
//
// Пример: В системе управления заказами можно добавить новый метод для расчета налогов на основе
// различных типов заказов (физические товары, цифровые продукты) без изменения классов самих заказов.

// + Упрощение добавления новых операций
// + Объединение родственных операций
// + Инкапсуляция поведения
// - Сложность добавления новых классов
// - Усложнение структуры кода
// - Может нарушить инкапсуляцию объектов

// Visitor - интерфейс посетителя
type Visitor interface {
	VisitShaurma(p *Shaurma) string
	VisitPizzeria(p *Pizzeria) string
	VisitBurgerKing(p *BurgerKing) string
}

// Place - интерфейс места, которое посетитель должен посетить
type Place interface {
	Accept(v Visitor) string
}

// People реализует интерфейс Visitor
type People struct {
}

// VisitShaurma реализует посещение в Shaurma
func (v *People) VisitShaurma(p *Shaurma) string {
	return p.BuyShaurma()
}

// VisitPizzeria реализует посещение в Pizzeria
func (v *People) VisitPizzeria(p *Pizzeria) string {
	return p.BuyPizza()
}

// VisitBurgerKing реализует посещение в BurgerKing
func (v *People) VisitBurgerKing(p *BurgerKing) string {
	return p.BuyBurger()
}

// City реализует коллекцию мест, которые можно посетить
type City struct {
	places []Place
}

// Add добавляет Place в коллекцию.
func (c *City) Add(p Place) {
	c.places = append(c.places, p)
}

// Accept реализует посещение во все места в городе
func (c *City) Accept(v Visitor) string {
	var result string
	for _, p := range c.places {
		result += p.Accept(v)
	}
	return result
}

// Shaurma реализует интерфейс Place
type Shaurma struct {
}

// Реализация Accept
func (s *Shaurma) Accept(v Visitor) string {
	return v.VisitShaurma(s)
}

// Реализация BuyShaurma
func (s *Shaurma) BuyShaurma() string {
	return "Buy Shaurma..."
}

// Pizzeria реализует интерфейс Place
type Pizzeria struct {
}

// Реализация Accept
func (p *Pizzeria) Accept(v Visitor) string {
	return v.VisitPizzeria(p)
}

// Реализация BuyPizza
func (p *Pizzeria) BuyPizza() string {
	return "Buy pizza..."
}

// BurgerKing реализует интерфейс Place
type BurgerKing struct {
}

// Реализация Accept
func (b *BurgerKing) Accept(v Visitor) string {
	return v.VisitBurgerKing(b)
}

// Реализация BuyBurger
func (b *BurgerKing) BuyBurger() string {
	return "Buy burger..."
}