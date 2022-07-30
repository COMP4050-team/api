package graph

import (
	"testing"

	"github.com/99designs/gqlgen/client"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/introspection"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"

	"github.com/COMP4050/square-team-5/api/fixtures/mocks"
	"github.com/COMP4050/square-team-5/api/graph/generated"
	"github.com/COMP4050/square-team-5/api/internal/pkg/db"
)

func TestRootResolver(t *testing.T) {
	ctrl := gomock.NewController(t)
	db := mocks.NewMockDatabase(ctrl)

	c := client.New(handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &Resolver{DB: db}})))

	t.Run("introspection", func(t *testing.T) {
		// Make sure we can run the graphiql introspection query without errors
		var resp interface{}
		c.MustPost(introspection.Query, &resp)
	})
}

func TestUnitResolver(t *testing.T) {
	t.Parallel()

	t.Run("Get Unit", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		mockDB := mocks.NewMockDatabase(ctrl)
		c := client.New(handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &Resolver{DB: mockDB}})))

		mockDB.EXPECT().GetUnitByID("1", false).Return(&db.Unit{Model: gorm.Model{ID: 1}, Name: "COMP1000"}, nil)

		var resp struct {
			Unit struct{ ID, Name string }
		}
		c.MustPost(`{ unit(id:"1") { id name } }`, &resp)

		assert.Equal(t, "1", resp.Unit.ID)
		assert.Equal(t, "COMP1000", resp.Unit.Name)
	})

	t.Run("Get Unit Not Found", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		mockDB := mocks.NewMockDatabase(ctrl)
		c := client.New(handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &Resolver{DB: mockDB}})))

		mockDB.EXPECT().GetUnitByID("1", false).Return(nil, db.ErrRecordNotFound)

		var resp struct {
			Unit struct{ ID, Name string }
		}
		err := c.Post(`{ unit(id:"1") { id name } }`, &resp)

		assert.ErrorContains(t, err, "record not found")
		assert.NotEqual(t, "1", resp.Unit.ID)
	})

	t.Run("Create Unit", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		mockDB := mocks.NewMockDatabase(ctrl)
		c := client.New(handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &Resolver{DB: mockDB}})))

		mockDB.EXPECT().GetUnitByName("COMP1000").Return(nil, db.ErrRecordNotFound)
		mockDB.EXPECT().CreateUnit("COMP1000").Return(&db.Unit{Model: gorm.Model{ID: 1}, Name: "COMP1000"}, nil)

		var resp struct {
			CreateUnit struct{ ID, Name string }
		}
		c.MustPost(`mutation { createUnit(input: {name: "COMP1000"}) { id name } }`, &resp)

		assert.Equal(t, "1", resp.CreateUnit.ID)
		assert.Equal(t, "COMP1000", resp.CreateUnit.Name)
	})
}

func TestClassResolver(t *testing.T) {
	t.Parallel()

	t.Run("Get Class", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		mockDB := mocks.NewMockDatabase(ctrl)
		c := client.New(handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &Resolver{DB: mockDB}})))

		mockDB.EXPECT().GetClass("1").Return(&db.Class{Model: gorm.Model{ID: 1}, Name: "Class 1"}, nil)

		var resp struct {
			Class struct{ ID, Name string }
		}
		c.MustPost(`{ class(id:"1") { id name } }`, &resp)

		assert.Equal(t, "1", resp.Class.ID)
		assert.Equal(t, "Class 1", resp.Class.Name)
	})

	t.Run("Get Class Not Found", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		mockDB := mocks.NewMockDatabase(ctrl)
		c := client.New(handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &Resolver{DB: mockDB}})))

		mockDB.EXPECT().GetClass("1").Return(nil, db.ErrRecordNotFound)

		var resp struct {
			Class struct{ ID, Name string }
		}
		err := c.Post(`{ class(id:"1") { id name } }`, &resp)

		assert.ErrorContains(t, err, "record not found")
		assert.NotEqual(t, "1", resp.Class.ID)
	})

	t.Run("Create Class", func(t *testing.T) {
		t.Parallel()
		ctrl := gomock.NewController(t)

		mockDB := mocks.NewMockDatabase(ctrl)
		c := client.New(handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &Resolver{DB: mockDB}})))

		mockDB.EXPECT().GetUnitByID("1", false).Return(&db.Unit{Model: gorm.Model{ID: 1}, Name: "COMP1000"}, nil)
		mockDB.EXPECT().CreateClass("Class 1", uint(1)).Return(&db.Class{Model: gorm.Model{ID: 1}, Name: "Class 1"}, nil)

		var resp struct {
			CreateClass struct{ ID, Name string }
		}
		c.MustPost(`mutation { createClass(input: {name: "Class 1", unitID: "1"}) { id name } }`, &resp)

		assert.Equal(t, "1", resp.CreateClass.ID)
		assert.Equal(t, "Class 1", resp.CreateClass.Name)
	})
}
