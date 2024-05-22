package dip

import (
	"math/rand"
	"strconv"
)

type User struct {
	Id   string
	Name string
}

var (
	inMemoUsers map[string]User
)

func init() {
	inMemoUsers = make(map[string]User)
}

func main() {
	// name:= "Douglas"

}

// Baixo Nível - Dependendo de detalhes - Inserindo usuário - Repositório
func InsertUserIntoMySQLRepo(user_id string, name string) *User {
	user := User{Id: user_id, Name: name}
	inMemoUsers[user_id] = user
	return &user
}

// Baixo Nível - Dependendo de detalhes - Buscando usuário - Repositório
func GetUserByIDFromSQLRepo(user_id string) *User {
	user, ok := inMemoUsers[user_id]
	if !ok {
		return nil
	}

	return &user
}

// Alto Nível - Dependendo de abstrações - Buscando usuário - Services
func GetUserByIDServices(user_id string) *User {
	return GetUserByIDFromSQLRepo(user_id)
}

// Alto Nível - Dependendo de abstrações - Inserindo usuário - Services
func InsertUserServices(name string) *User {
	user_id := strconv.Itoa(rand.Intn(100000))
	return InsertUserIntoMySQLRepo(user_id, name)
}
