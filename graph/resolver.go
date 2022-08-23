package graph

import (
	"context"

	"github.com/COMP4050/square-team-5/api/internal/pkg/config"
	"github.com/COMP4050/square-team-5/api/internal/pkg/db"
	"github.com/COMP4050/square-team-5/api/internal/pkg/db/models"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Config      config.Config
	DB          db.Database
	ExtractUser func(ctx context.Context) *models.User
}
