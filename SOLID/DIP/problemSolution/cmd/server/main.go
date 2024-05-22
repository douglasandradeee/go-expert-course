package main

import (
	"fmt"

	"github.com/douglasandradeee/go-expert-course/SOLID/DIP/problemSolution/repository"
	"github.com/douglasandradeee/go-expert-course/SOLID/DIP/problemSolution/services"
)

func main() {
	// Instanciando o repositório de usuários. - Injeção de dependência.
	mysqlRepositository := &repository.MysqlRepositository{}
	// Instanciando o serviço de usuários. - Injeção de dependência.
	userServices := &services.UserServices{Repo: mysqlRepositository}

	name := "Douglas"
	userInserted := userServices.InsertUserServices(name)
	fmt.Println("Usuário inserido: ", userInserted)

	userReturned := userServices.GetUserByIDServices(userInserted.ID)
	fmt.Println("Usuário retornado: ", userReturned)

}
