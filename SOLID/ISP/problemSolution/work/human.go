package work

type Human struct{}

type HumanTasks interface {
	WorkTasks
	CoffeBreak()
	ToLunch()
}

func (h *Human) GetIntoWork() {
	//humano entrando no trabalho.
}

func (h *Human) StartWorking() {
	// humano começando a trabalhar.
}

func (h *Human) CoffeBreak() {
	// humano fazendo pausa para o café.
}

func (h *Human) ToLunch() {
	// humano fazendo pausa para o almoço.

}

func (h *Human) ContinueToWork() {
	// humano contiuando a trabalhar.
}

func (h *Human) LeaveWork() {
	// humando saindo/deixando o trabalho.
}
