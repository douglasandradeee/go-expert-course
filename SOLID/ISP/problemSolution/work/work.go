package work

type WorkTasks interface {
	GetIntoWork()
	StartWorking()
	ContinueToWork()
	LeaveWork()
}
