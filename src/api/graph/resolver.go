package graph

import "github.com/mnogokotin/sbubnom-backend/graph/model"

type Resolver struct {
	questions []*model.Question
	choices   []*model.Choice
}
