package events

import (
	"errors"
	"sync"
)

var ErrorHandlerAlreadyRegistered = errors.New("handler already registered")

//Implementando nosso dispatcher/disparador de eventos.

/*
EventDispatcher - precisará ter os nossos handlers registrados.
Ele precisará ter o nome do evento e os nossos handlers de acordo com o evento.
Por isso essa struct terá um map de handlers onde o nome do evento será a chave,
e os handlers do evento serão o valor, lembrando que um evento pode ter mais de um handler
*/
type EventDispatcher struct {
	//mapa de handlers.
	handlers map[string][]EventHandlerInterface
}

// NewEventDispatcher - Criando um novo dispatcher/disparador de eventos.
func NewEventDispatcher() *EventDispatcher {
	return &EventDispatcher{
		handlers: make(map[string][]EventHandlerInterface),
	}
}

// Dispatch - disparando o evento.
func (ed *EventDispatcher) Dispatch(event EventInterface) error {
	//verificando se existe handler para o evento name passado.
	//Caso exista, percorra os handlers do evento e chame o handler do evento.
	if handlers, ok := ed.handlers[event.GetName()]; ok {
		//percorrendo os handlers do evento.
		wg := &sync.WaitGroup{}
		for _, handler := range handlers {
			wg.Add(1)
			//chamando o handler do evento.
			go handler.Handle(event, wg)
		}
		wg.Wait()
	}
	return nil
}

func (ed *EventDispatcher) Register(eventName string, handler EventHandlerInterface) error {
	//verificando se o evento já existe.\

	/*	verificando se o evento solicitado para ser registrado já existe dentro do mapa de handlers.
		Vai passar o nome do evento como chave de busca, quando o evento for encontrador iremos percorrer
		o mapa de handlers do evento e verificar se o handler já está registrado, se estiver registrado
		retornamos um erro, se não estiver registrado, registramos o handler.
	*/
	if _, ok := ed.handlers[eventName]; ok {
		for _, h := range ed.handlers[eventName] {
			if h == handler {
				return ErrorHandlerAlreadyRegistered
			}
		}
	}
	//registrando o handler no map.
	ed.handlers[eventName] = append(ed.handlers[eventName], handler)
	return nil
}

func (ed *EventDispatcher) Remove(eventName string, handler EventHandlerInterface) {
	//verificando se o evento existe.
	// se existir irei percorrer o map de handlers do evento e verificar se o handler já está registrado.
	if _, ok := ed.handlers[eventName]; ok {
		//percorrendo os handlers do evento.
		for i, h := range ed.handlers[eventName] {
			if h == handler {
				//removendo o handler do evento.
				ed.handlers[eventName] = append(ed.handlers[eventName][:i], ed.handlers[eventName][i+1:]...)
			}
		}
	}
}

// Clear - limpando o mapa de handlers.
func (ed *EventDispatcher) Clear() {
	//Zerando o mapa de handlers.
	ed.handlers = make(map[string][]EventHandlerInterface)
}

func (ed *EventDispatcher) Has(eventName string, handler EventHandlerInterface) bool {
	//verificando se o evento existe.
	// se existir irei percorrer o map de handlers do evento e verificar se o handler já está registrado.
	if _, ok := ed.handlers[eventName]; ok {
		//percorrendo os handlers do evento.
		for _, h := range ed.handlers[eventName] {
			if h == handler {
				return true
			}
		}
	}
	return false
}
