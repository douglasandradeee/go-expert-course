package events

import (
	"sync"
	"time"
)

// Definindo a interface de eventos ou o evento em si.
type EventInterface interface {
	//obtêm o nome do evento.
	GetName() string
	//obtêm a data e hora do evento.
	GetDateTime() time.Time
	//obtêm o payload do evento.
	GetPayload() interface{}
}

// Operação que serão executadas quando um evento for chamado.
type EventHandlerInterface interface {
	Handle(event EventInterface, wg *sync.WaitGroup)
}

// Interface para o disparar as operações de um dado evento.
// Gerenciador de eventos/operações.
// Responsável por registrar e disparar os eventos.
type EventDispatcherInterface interface {
	//Registra um evento.
	Register(eventName string, handler EventHandlerInterface) error
	//Dispara um evento.
	Dispatach(event EventInterface) error
	//Remove um evento.
	Remove(eventName string, handler EventHandlerInterface) error
	//Verifica se um evento existe.
	Has(eventName string, handler EventHandlerInterface) bool
	//Limpa todos os eventos.
	Clear() error
}
