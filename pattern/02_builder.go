package pattern

/*
	Реализовать паттерн «строитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Builder_pattern
*/

// Паттерн Builder определяет процесс поэтапного построения сложного продукта. 
// После того как будет построена последняя его часть, продукт можно использовать.
//
// Пример: Создание объекта Pizza, где можно поочередно добавлять ингредиенты (сыр, грибы, перец и т.д.) и настраивать его размер и тип теста.

// + Позволяет поэтапно создавать сложные объекты
// + Улучшает читаемость кода, так как каждый метод строителя отвечает за установку одного параметра
// + Гибкость
// - Привязанность к конкретным классам строителей
// - Может усложнять код, если объект не требует поэтапной сборки
// - Увеличение количества классов

// Builder определяет интерфейс строителя
type Builder interface {
	MakeHeader(str string)
	MakeBody(str string)
	MakeFooter(str string)
}

// Director реализует управляющего строителем
type Director struct {
	builder Builder
}

// Construct - управление строителем, т.е. что и в каком порядке делать строителю.
func (d *Director) Construct() {
	d.builder.MakeHeader("Header")
	d.builder.MakeBody("Body")
	d.builder.MakeFooter("Footer")
}

// ConcreteBuilder реализует интерфейс строителя
type ConcreteBuilder struct {
	product *Product
}

// MakeHeader строит хедер документа
func (b *ConcreteBuilder) MakeHeader(str string) {
	b.product.Content += "<header>" + str + "</header>"
}

// MakeBody строит тело документа.
func (b *ConcreteBuilder) MakeBody(str string) {
	b.product.Content += "<article>" + str + "</article>"
}

// MakeFooter строит футер документа.
func (b *ConcreteBuilder) MakeFooter(str string) {
	b.product.Content += "<footer>" + str + "</footer>"
}

// Реализация продукта.
type Product struct {
	Content string
}

// Show возвращает содержание продукта.
func (p *Product) Show() string {
	return p.Content
}