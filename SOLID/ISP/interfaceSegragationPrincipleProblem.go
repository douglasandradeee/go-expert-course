package isp

type Work interface {
	GetIntoWork()
	StartWorking()
	CoffeBreak()
	CheckOil()
	ToLunch()
	ChargeTheBattery()
	ContinueToWork()
	LeaveWork()
}

type Robot struct{}
type Human struct{}

func (h *Human) GetIntoWork() {
	//humano entrando no trabalho.
}

func (h *Human) StartWorking() {
	// humano começando a trabalhar.
}

func (h *Human) CoffeBreak() {
	// humano fazendo pausa para o café.

}

func (h *Human) CheckOil() {
	// humano faz checagem de Óleo?
}

func (h *Human) ToLunch() {
	// humano fazendo pausa para o almoço.

}

func (h *Human) ChargeTheBattery() {
	// humano carrega a sua bateria?
	// humano possui bateria?
}

func (h *Human) ContinueToWork() {
	// humano contiuando a trabalhar.
}

func (h *Human) LeaveWork() {
	// humando saindo/deixando o trabalho.
}

func (r *Robot) GetIntoWork() {
	//robô entrando no trabalho.
}

func (r *Robot) StartWorking() {
	// robô começando a trabalhar.
}

func (r *Robot) CoffeBreak() {
	// robô faz pausa para o café?

}

func (r *Robot) CheckOil() {
	// robô fazendo checagem de Óleo.
}

func (r *Robot) ToLunch() {
	// robô  faz pausa para o almoço?
}

func (r *Robot) ChargeTheBattery() {
	// robô  carregando a sua bateria.
}

func (r *Robot) ContinueToWork() {
	// robô contiuando a trabalhar.
}

func (r *Robot) LeaveWork() {
	// robô saindo/deixando o trabalho.
}
