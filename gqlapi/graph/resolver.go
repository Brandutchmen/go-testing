//go:generate go run github.com/99designs/gqlgen generate
package graph

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.
import (
  "brandon/gqlapi/graph/model"
)

type Resolver struct{
  todos []*model.Todo
}
