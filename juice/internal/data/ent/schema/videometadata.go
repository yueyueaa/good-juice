package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"

	"time"
)

// VideoMetadata holds the schema definition for the VideoMetadata entity.
type VideoMetadata struct {
	ent.Schema
}

// Fields of the VideoMetadata.
func (VideoMetadata) Fields() []ent.Field {

	return []ent.Field{

		field.Int64("video_id").SchemaType(map[string]string{
			dialect.MySQL: "BIGINT", // Override MySQL.
		}).Unique(),

		field.Int64("user_id").SchemaType(map[string]string{
			dialect.MySQL: "BIGINT", // Override MySQL.
		}),

		field.String("cover_url").SchemaType(map[string]string{
			dialect.MySQL: "CHAR(255)", // Override MySQL.
		}),

		field.String("video_url").SchemaType(map[string]string{
			dialect.MySQL: "CHAR(255)", // Override MySQL.
		}),

		field.String("video_intro").SchemaType(map[string]string{
			dialect.MySQL: "CHAR(255)", // Override MySQL.
		}).Optional(),

		field.Int64("video_type").SchemaType(map[string]string{
			dialect.MySQL: "BIGINT", // Override MySQL.
		}),

		field.Int32("publish_address").SchemaType(map[string]string{
			dialect.MySQL: "INT", // Override MySQL.
		}).Default(0),

		field.Time("create_time").SchemaType(map[string]string{
			dialect.MySQL: "DATETIME", // Override MySQL.
		}).Default(time.Now),

		field.Time("update_time").SchemaType(map[string]string{
			dialect.MySQL: "DATETIME", // Override MySQL.
		}).Default(time.Now).UpdateDefault(time.Now),
	}

}

// Edges of the VideoMetadata.
func (VideoMetadata) Edges() []ent.Edge {
	return nil
}
