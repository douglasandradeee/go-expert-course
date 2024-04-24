package events

import (
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type TestEvent struct {
	Name    string
	Payload interface{}
}

// GetName - função de teste para obter o nome do evento.
func (te *TestEvent) GetName() string {
	return te.Name
}

// GetPayload - função de teste para obter o payload do evento.
func (te *TestEvent) GetPayload() interface{} {
	return te.Payload
}

// GetDateTime - função de teste para obter a data e hora do evento.
func (te *TestEvent) GetDateTime() time.Time {
	return time.Now()
}

// TestEventHandler - Criando o eventHandler de teste.
type TestEventHandler struct {
	ID int
}

// Handle - função handler de teste.
func (teh *TestEventHandler) Handle(event EventInterface, wg *sync.WaitGroup) {
}

// EventDispatcherTestSuite - Criando o suite de teste para o eventDispatcher.
type EventDispatcherTestSuite struct {
	suite.Suite
	event            TestEvent
	event2           TestEvent
	handler          TestEventHandler
	handler2         TestEventHandler
	handler3         TestEventHandler
	eventDispatacher *EventDispatcher
}

// SetupTest - Prepara os dados para o teste sem repetir código.
func (suite *EventDispatcherTestSuite) SetupTest() {
	suite.eventDispatacher = NewEventDispatcher()
	suite.handler = TestEventHandler{
		ID: 1,
	}

	suite.handler2 = TestEventHandler{
		ID: 2,
	}

	suite.handler3 = TestEventHandler{
		ID: 3,
	}
	suite.event = TestEvent{Name: "Teste Evento1", Payload: "Teste Payload Evento1"}
	suite.event = TestEvent{Name: "Teste Evento2", Payload: "Teste Payload Evento2"}
}

func (suite *EventDispatcherTestSuite) TestEventDispatcher_RegisterSucess() {
	//Registrando o evento 1 e verificando se  o handler 1 foi registrado sem errors.
	err := suite.eventDispatacher.Register(suite.event.GetName(), &suite.handler)
	suite.Nil(err)
	//verificando se o handler que passamos para ser registrado foi registrado e se a quantidade de handlers é 1.
	suite.Equal(1, len(suite.eventDispatacher.handlers[suite.event.GetName()]))

	//Registrando o evento 2 e verificando se  o handler 2 foi registrado sem errors.
	err = suite.eventDispatacher.Register(suite.event.GetName(), &suite.handler2)
	suite.Nil(err)
	//verificando se o handler que passamos para ser registrado foi registrado e se a quantidade de handlers é 2.
	suite.Equal(2, len(suite.eventDispatacher.handlers[suite.event.GetName()]))

	//Verificando se o handler Registrado é o mesmo que passamos para ser registrado.
	assert.Equal(suite.T(), &suite.handler, suite.eventDispatacher.handlers[suite.event.GetName()][0])
	assert.Equal(suite.T(), &suite.handler2, suite.eventDispatacher.handlers[suite.event.GetName()][1])
}

func (suite *EventDispatcherTestSuite) TestEventDispatcher_RegisterError_WithSameHandler() {
	//Registrando o evento 1 e verificando se  o handler 1 foi registrado sem errors.
	err := suite.eventDispatacher.Register(suite.event.GetName(), &suite.handler)
	suite.Nil(err)
	//verificando se o handler que passamos para ser registrado foi registrado e se a quantidade de handlers é 1.
	suite.Equal(1, len(suite.eventDispatacher.handlers[suite.event.GetName()]))

	//Registrando o evento o mesmo evento que já foi registrado anteriormente a fim de capturar o error.
	err = suite.eventDispatacher.Register(suite.event.GetName(), &suite.handler)
	suite.Equal(ErrorHandlerAlreadyRegistered, err)
	//verificando se o handler que passamos para ser registrado foi registrado e se a quantidade de handlers é 1 ou 2.
	// em tese tem de ser 1 pois o handler já foi registrado anteriormente.
	suite.Equal(1, len(suite.eventDispatacher.handlers[suite.event.GetName()]))

}

func (suite *EventDispatcherTestSuite) TestEventDispatcher_Clear() {
	//Event 1
	//Registando primeiro handler.
	err := suite.eventDispatacher.Register(suite.event.GetName(), &suite.handler)
	suite.Nil(err)
	suite.Equal(1, len(suite.eventDispatacher.handlers[suite.event.GetName()]))

	//Registrando segundo handler.
	err = suite.eventDispatacher.Register(suite.event.GetName(), &suite.handler2)
	suite.Nil(err)
	suite.Equal(2, len(suite.eventDispatacher.handlers[suite.event.GetName()]))

	//Event 2
	err = suite.eventDispatacher.Register(suite.event2.GetName(), &suite.handler3)
	suite.Nil(err)
	suite.Equal(1, len(suite.eventDispatacher.handlers[suite.event2.GetName()]))

	//chamando o clear para limpar os handlers e verificando se realmente foram limpos.
	suite.eventDispatacher.Clear()
	suite.Equal(0, len(suite.eventDispatacher.handlers[suite.event.GetName()]))
}

func (suite *EventDispatcherTestSuite) TestEventDispatcher_Has() {
	//Registando primeiro handler.
	err := suite.eventDispatacher.Register(suite.event.GetName(), &suite.handler)
	suite.Nil(err)
	suite.Equal(1, len(suite.eventDispatacher.handlers[suite.event.GetName()]))

	//Registrando segundo handler.
	err = suite.eventDispatacher.Register(suite.event.GetName(), &suite.handler2)
	suite.Nil(err)
	suite.Equal(2, len(suite.eventDispatacher.handlers[suite.event.GetName()]))

	assert.True(suite.T(), suite.eventDispatacher.Has(suite.event.GetName(), &suite.handler))
	assert.True(suite.T(), suite.eventDispatacher.Has(suite.event.GetName(), &suite.handler2))
	assert.False(suite.T(), suite.eventDispatacher.Has(suite.event.GetName(), &suite.handler3))
}

func (suite *EventDispatcherTestSuite) TestEventDispatcher_Remove() {
	// Event 1
	// Registrando primeiro handler.
	err := suite.eventDispatacher.Register(suite.event.GetName(), &suite.handler)
	suite.Nil(err)
	suite.Equal(1, len(suite.eventDispatacher.handlers[suite.event.GetName()]))
	// Registrando segundo handler.
	err = suite.eventDispatacher.Register(suite.event.GetName(), &suite.handler2)
	suite.Nil(err)
	suite.Equal(2, len(suite.eventDispatacher.handlers[suite.event.GetName()]))

	// Event 2
	err = suite.eventDispatacher.Register(suite.event2.GetName(), &suite.handler3)
	suite.Nil(err)
	suite.Equal(1, len(suite.eventDispatacher.handlers[suite.event2.GetName()]))

	// Removendo o handler 1 do evento 1.
	suite.eventDispatacher.Remove(suite.event.GetName(), &suite.handler)
	suite.Equal(1, len(suite.eventDispatacher.handlers[suite.event.GetName()]))
	assert.Equal(suite.T(), &suite.handler2, suite.eventDispatacher.handlers[suite.event.GetName()][0])

	// Removendo o handler 2 do evento 1.
	suite.eventDispatacher.Remove(suite.event.GetName(), &suite.handler2)
	suite.Equal(0, len(suite.eventDispatacher.handlers[suite.event.GetName()]))

	// Removendo o handler 3 do evento 2.
	suite.eventDispatacher.Remove(suite.event2.GetName(), &suite.handler3)
	suite.Equal(0, len(suite.eventDispatacher.handlers[suite.event2.GetName()]))

}

type MockHandler struct {
	mock.Mock
}

func (m *MockHandler) Handle(event EventInterface, wg *sync.WaitGroup) {
	m.Called(event)
	wg.Done()
}

func (suite *EventDispatcherTestSuite) TestEventDispatch_Dispatch() {
	eh := &MockHandler{}
	eh.On("Handle", &suite.event)

	eh2 := &MockHandler{}
	eh2.On("Handle", &suite.event)

	suite.eventDispatacher.Register(suite.event.GetName(), eh)
	suite.eventDispatacher.Register(suite.event.GetName(), eh2)

	suite.eventDispatacher.Dispatch(&suite.event)
	eh.AssertExpectations(suite.T())
	eh2.AssertExpectations(suite.T())
	eh.AssertNumberOfCalls(suite.T(), "Handle", 1)
	eh2.AssertNumberOfCalls(suite.T(), "Handle", 1)
}

// TestSuite - função de teste para o suite de teste do eventDispatcher.
func TestSuite(t *testing.T) {
	suite.Run(t, new(EventDispatcherTestSuite))
}
