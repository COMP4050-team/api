package graph

import (
	"github.com/COMP4050/square-team-5/api/internal/pkg/db"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	DB db.Database
}
