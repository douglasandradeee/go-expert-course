package services

import (
	"math/rand"
	"strconv"

	"github.com/douglasandradeee/go-expert-course/SOLID/DIP/problemSolution/repository"
)

type UserServices struct {
	Repo repository.UserRepositoryInterface
}

func (s *UserServices) GetUserByIDServices(user_id string) *repository.User {
	return s.Repo.GetUserByID(user_id)
}

func (s *UserServices) InsertUserServices(name string) *repository.User {
	user_id := strconv.Itoa(rand.Intn(100000))
	return s.Repo.InsertUser(user_id, name)
}
