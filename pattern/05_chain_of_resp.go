package pattern

/*
	Реализовать паттерн «цепочка вызовов».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Chain-of-responsibility_pattern
*/

// Паттерн Chain Of Responsibility позволяет избежать привязки объекта-отправителя запроса к объекту-получателю запроса,
// при этом давая шанс обработать этот запрос нескольким объектам. Получатели связываются в цепочку, и запрос передается
// по цепочке, пока не будет обработан каким-то объектом.
//
// Пример: Когда система должна обрабатывать различные типы запросов, но заранее неизвестно, какие именно запросы будут поступать.
// В веб-приложении для обработки HTTP-запросов, где каждый обработчик (например, аутентификация, авторизация, логирование)
// может проверять и обрабатывать запрос, передавая его дальше по цепочке, если не может его обработать.

// + Каждый обработчик отвечает только за свою часть логики обработки запроса
// + Легкость в добавлении новых обработчиков в цепь без изменения существующего кода
// + Паттерн уменьшает зависимость между отправителем и получателем запроса, что позволяет избежать жесткой привязки между ними
// - Возможны необработанные запросы
// - Сложно отследить, какой обработчик в цепочке обработал запрос


// Handler - интерфейс обработчика.
type Handler interface {
	SendRequest(message int) string
	SetNext(handler Handler)
}

// BaseHandler базовый обработчик, реализующий общую логику установки следующего обработчика
type BaseHandler struct {
	next Handler
}

func (h *BaseHandler) SendRequest(message int) {
	if h.next != nil {
		h.next.SendRequest(message)
	}
}

func (h *BaseHandler) SetNext(handler Handler) {
	h.next = handler
}


// ConcreteHandlerA реализует обработчик "A".
type ConcreteHandlerA struct {
	BaseHandler
}

// Реализация SendRequest.
func (h *ConcreteHandlerA) SendRequest(message int) (result string) {
	if message == 1 {
		result = "Im handler A"
	} else if h.next != nil {
		result = h.next.SendRequest(message)
	}
	return
}

// ConcreteHandlerB реализует обработчик "B".
type ConcreteHandlerB struct {
	BaseHandler
}

// Реализация SendRequest.
func (h *ConcreteHandlerB) SendRequest(message int) (result string) {
	if message == 2 {
		result = "Im handler B"
	} else if h.next != nil {
		result = h.next.SendRequest(message)
	}
	return
}

// ConcreteHandlerC реализует обработчик "C".
type ConcreteHandlerC struct {
	BaseHandler
}

// Реализация SendRequest.
func (h *ConcreteHandlerC) SendRequest(message int) (result string) {
	if message == 3 {
		result = "Im handler C"
	} else if h.next != nil {
		result = h.next.SendRequest(message)
	}
	return
}