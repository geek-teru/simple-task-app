package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Task holds the schema definition for the Task entity.
type Task struct {
	ent.Schema
}

// Fields of the Task.
func (Task) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").
			Unique().
			Immutable().
			Comment("タスクID"),
		field.String("title").
			NotEmpty().
			Comment("タイトル"),
		field.String("description").
			Optional().
			Comment("詳細"),
		field.Time("due_date").
			Optional().
			Nillable().
			Comment("期限日"),
		field.Enum("status").
			Values("TODO", "IN_PROGRESS", "DONE").
			Default("TODO").
			Comment("ステータス"),
		field.Int("user_id").
			Comment("ユーザーID"),
	}
}

// Edges of the Task.
func (Task) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("tasks").
			Field("user_id").
			Unique().
			Required(),
	}
}
