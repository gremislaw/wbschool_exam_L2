package pattern

/*
	Реализовать паттерн «фабричный метод».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Some_method_pattern
*/

// Паттерн Factory Method применяется для создания объектов с определенным интерфейсом, реализации которого предоставляются потомками.
// Другими словами, есть базовый абстрактный класс фабрики, который говорит, что каждая его наследующая фабрика должна реализовать
// такой-то метод для создания своих продуктов.
//
// Пример: Создание класса PaymentFactory, который имеет метод createPaymentMethod().
// Подклассы реализуют создание конкретных методов оплаты, таких как CreditCardPayment и PayPalPayment

// + Фабричный метод позволяет создавать объекты, не привязываясь к конкретным классам
// + Логика создания в одном месте
// + Легко добавлять новые классы продуктов, не изменяя существующий код.
// - Добавление новых классов и методов может усложнить код


// action помогает клиентам найти доступные действия.
type action string

const (
	A action = "A"
	B action = "B"
)

// Creator - интерфейс фабрики.
type Creator interface {
	CreateSomeProduct(action action) SomeProduct // Фабричный метод
}

// SomeProduct - интерфейс продукта.
// Все продукты возвращенные фабрикой должны иметь один интерфейс.
type SomeProduct interface {
	Use() string // Каждый продукт должен быть доступным для использования
}

// ConcreteCreator реализует интерфейс Creator.
type ConcreteCreator struct{}

// NewCreator - конструктор создания ConcreteCreator.
func NewCreator() Creator {
	return &ConcreteCreator{}
}

// CreateSomeProduct - фабричный метод.
func (p *ConcreteCreator) CreateSomeProduct(action action) SomeProduct {
	var SomeProduct SomeProduct

	switch action {
	case A:
		SomeProduct = &ConcreteSomeProductA{string(action)}
	case B:
		SomeProduct = &ConcreteSomeProductB{string(action)}
	default:
		SomeProduct = nil
	}

	return SomeProduct
}

// ConcreteSomeProductA implements Someproduct "A".
type ConcreteSomeProductA struct {
	action string
}

// Use returns Someproduct action.
func (p *ConcreteSomeProductA) Use() string {
	return p.action
}

// ConcreteSomeProductB implements Someproduct "B".
type ConcreteSomeProductB struct {
	action string
}

// Use returns Someproduct action.
func (p *ConcreteSomeProductB) Use() string {
	return p.action
}