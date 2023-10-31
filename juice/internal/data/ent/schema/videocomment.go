package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"

	"time"
)

// VideoComment holds the schema definition for the VideoComment entity.
type VideoComment struct {
	ent.Schema
}

// Fields of the VideoComment.
func (VideoComment) Fields() []ent.Field {

	return []ent.Field{

		field.Int64("comment_id").SchemaType(map[string]string{
			dialect.MySQL: "BIGINT", // Override MySQL.
		}).Unique(),

		field.Int64("pcomment_id").SchemaType(map[string]string{
			dialect.MySQL: "BIGINT", // Override MySQL.
		}),

		field.Int64("video_id").SchemaType(map[string]string{
			dialect.MySQL: "BIGINT", // Override MySQL.
		}),

		field.Int64("user_id").SchemaType(map[string]string{
			dialect.MySQL: "BIGINT", // Override MySQL.
		}),

		field.Text("comment_text").SchemaType(map[string]string{
			dialect.MySQL: "TEXT", // Override MySQL.
		}).Optional(),

		field.Time("create_time").SchemaType(map[string]string{
			dialect.MySQL: "DATETIME", // Override MySQL.
		}).Default(time.Now),

		field.Time("update_time").SchemaType(map[string]string{
			dialect.MySQL: "DATETIME", // Override MySQL.
		}).Default(time.Now).UpdateDefault(time.Now),
	}

}

// Edges of the VideoComment.
func (VideoComment) Edges() []ent.Edge {
	return nil
}
