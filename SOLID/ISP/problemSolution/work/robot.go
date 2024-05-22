package work

type Robot struct{}

type RobotTasks interface {
	WorkTasks
	CheckOil()
	ChargeTheBattery()
}

func (r *Robot) GetIntoWork() {
	// robô entrando no trabalho.
}

func (r *Robot) StartWorking() {
	// robô começando a trabalhar.
}

func (r *Robot) CheckOil() {
	// robô fazendo checagem de Óleo.
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
