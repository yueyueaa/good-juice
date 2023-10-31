package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"

	"time"
)

// VideoLike holds the schema definition for the VideoLike entity.
type VideoLike struct {
	ent.Schema
}

// Fields of the VideoLike.
func (VideoLike) Fields() []ent.Field {

	return []ent.Field{

		field.Int64("user_id").SchemaType(map[string]string{
			dialect.MySQL: "BIGINT", // Override MySQL.
		}),

		field.Int64("video_id").SchemaType(map[string]string{
			dialect.MySQL: "BIGINT", // Override MySQL.
		}),

		field.Int8("status").SchemaType(map[string]string{
			dialect.MySQL: "TINYINT", // Override MySQL.
		}).Default(1),

		field.Time("create_time").SchemaType(map[string]string{
			dialect.MySQL: "DATETIME", // Override MySQL.
		}).Default(time.Now),

		field.Time("update_time").SchemaType(map[string]string{
			dialect.MySQL: "DATETIME", // Override MySQL.
		}).Default(time.Now).UpdateDefault(time.Now),
	}

}

// Edges of the VideoLike.
func (VideoLike) Edges() []ent.Edge {
	return nil
}
