package pattern

/*
	Реализовать паттерн «комманда».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Command_pattern
*/

// Паттерн Command позволяет представить запрос в виде объекта. Из этого следует, что команда - это объект.
// Такие запросы, например, можно ставить в очередь, отменять или возобновлять.
// Пример: В текстовом редакторе каждая команда (например, "вставить текст", "удалить текст") может быть сохранена в стеке,
// что позволяет пользователю отменять или повторять действия.

// + Отложенный запуск
// + Поддержка отмены операций
// + Разделение объектов, инициирующих операции, и объектов, которые их выполняют.
// - Потенциальные проблемы с отладкой
// - Усложнение структуры, много классов

// Command - интерфейс комманды.
type Command interface {
	Execute() string
}

// ToggleOnCommand реализует интерфейс Command.
type ToggleOnCommand struct {
	receiver *Receiver
}

// Комманда выполнения.
func (c *ToggleOnCommand) Execute() string {
	return c.receiver.ToggleOn()
}

// ToggleOffCommand реализует интерфейс Command.
type ToggleOffCommand struct {
	receiver *Receiver
}

// Комманда выполнения.
func (c *ToggleOffCommand) Execute() string {
	return c.receiver.ToggleOff()
}

// Реализация Receiver.
type Receiver struct {
}

// Реализация ToggleOn.
func (r *Receiver) ToggleOn() string {
	return "Toggle On"
}

// Реализация ToggleOff.
func (r *Receiver) ToggleOff() string {
	return "Toggle Off"
}

// Реализация Invoker.
type Invoker struct {
	commands []Command
}

// Store добавляет комманду.
func (i *Invoker) Store(command Command) {
	i.commands = append(i.commands, command)
}

// UnStore удаляет комманду.
func (i *Invoker) UnStore() {
	if len(i.commands) != 0 {
		i.commands = i.commands[:len(i.commands)-1]
	}
}

// Выполнить все команды.
func (i *Invoker) Execute() string {
	var result string
	for _, command := range i.commands {
		result += command.Execute() + "\n"
	}
	return result
}