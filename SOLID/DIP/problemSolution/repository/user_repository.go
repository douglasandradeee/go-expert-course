package repository

type User struct {
	ID   string
	Name string
}

// lista de usuários em memória.
var (
	inMemoUsers map[string]User
)

type MysqlRepositository struct{}

type MongoDbRepositository struct{}

type UserRepositoryInterface interface {
	InsertUser(user_id string, name string) *User
	GetUserByID(user_id string) *User
}

// Inicializa a lista de usuários em memória.
func init() {
	inMemoUsers = make(map[string]User)
}

func (m *MysqlRepositository) InsertUser(user_id string, name string) *User {
	user := User{ID: user_id, Name: name}
	inMemoUsers[user_id] = user
	return &user
}

func (m *MysqlRepositository) GetUserByID(user_id string) *User {
	user, ok := inMemoUsers[user_id]
	if !ok {
		return nil
	}

	return &user
}
