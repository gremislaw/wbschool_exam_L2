package pattern

/*
	Реализовать паттерн «состояние».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/State_pattern
*/

// Паттерн должен применяться:
// * когда поведение объекта зависит от его состояния
// * поведение объекта должно изменяться во время выполнения программы
// * состояний достаточно много и использовать для этого условные операторы, разбросанные по коду, достаточно затруднительно
//
// пример: корзина может находиться в состояниях "Пустая", "Наполненная" и "Оплаченная", где доступные действия зависят от состояния корзины.

// + Читаемость кода
// + Код, связанный с определенным состоянием, сосредоточен в одном месте
// + Гибкость
// + Чистота архитектуры
// - Много классов


// MobileAlertStater - интерфейс для состояний.
type MobileAlertStater interface {
	Alert() string
}

// MobileAlert воспроизводит сигнал исходя из текущего состояния
type MobileAlert struct {
	state MobileAlertStater
}

// Alert возвращает строку сигнал
func (a *MobileAlert) Alert() string {
	return a.state.Alert()
}

// SetState меняет состояние
func (a *MobileAlert) SetState(state MobileAlertStater) {
	a.state = state
}

// NewMobileAlert - конструктор создания MobileAlert.
func NewMobileAlert() *MobileAlert {
	return &MobileAlert{state: &MobileAlertVibration{}}
}

// MobileAlertVibration реализует сигнал вибрации
type MobileAlertVibration struct {
}

// Alert возвращает строку сигнал
func (a *MobileAlertVibration) Alert() string {
	return "Brrr... Brrr..."
}

// MobileAlertSong реализует сигнал бип
type MobileAlertSong struct {
}

// Alert возвращает строку
func (a *MobileAlertSong) Alert() string {
	return "Сигма-сигма бой, сигма бой, сигма бо-о-ой..."
}