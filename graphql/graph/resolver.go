package graph

import "github.com/douglasandradeee/go-expert-course/graphql/internal/database"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.
// Injetando as dependecias do meu projeto para poder usar dentro do schema.Resolvers.
type Resolver struct {
	CategoryDB *database.Category
	CourseDB   *database.Course
}
