package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Task holds the schema definition for the Task entity.
type Task struct {
	ent.Schema
}

// Fields of the Task.
func (Task) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").MaxLen(100).Default("").Comment("タイトル"),
		field.String("description").MaxLen(255).Default("").Comment("詳細"),
		field.Time("due_date").Default(func() time.Time { return time.Now() }).Comment("期限日"),
		field.Int("status").Default(0).Comment("ステータス"),
		field.Int("user_id").Comment("ユーザーID"),
		field.Time("created_at").Default(time.Now).Immutable().Comment("作成日時"),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now).Comment("更新日時"),
		field.Time("deleted_at").Optional().Nillable().Comment("削除日時"),
	}
}

// Edges of the Task.
// func (Task) Edges() []ent.Edge {
// 	return []ent.Edge{
// 		edge.From("user", User.Type).
// 			Ref("tasks").
// 			Field("user_id").
// 			Unique().
// 			Required(),
// 	}
// }

// Indexes of the Task.
func (Task) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("user_id", "status", "deleted_at").
			StorageKey("idx_user_id_status_deleted_at"),
	}
}
