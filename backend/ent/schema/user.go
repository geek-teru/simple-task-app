package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique().Comment("ID"),
		field.String("name").NotEmpty().Comment("name"),
		field.String("email").Unique().NotEmpty().Comment("email"),
		field.String("password").NotEmpty().Comment("password"),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("tasks", Task.Type).
			Comment("ユーザーが所有するタスク").
			StructTag(`json:"tasks,omitempty"`),
	}
}
