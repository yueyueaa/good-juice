package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"

	"time"
)

// UserFollowInfo holds the schema definition for the UserFollowInfo entity.
type UserFollowInfo struct {
	ent.Schema
}

// Fields of the UserFollowInfo.
func (UserFollowInfo) Fields() []ent.Field {

	return []ent.Field{

		field.Int64("user_id").SchemaType(map[string]string{
			dialect.MySQL: "BIGINTUNSIGNED", // Override MySQL.
		}).NonNegative(),

		field.Int64("follow_id").SchemaType(map[string]string{
			dialect.MySQL: "BIGINTUNSIGNED", // Override MySQL.
		}).NonNegative(),

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

// Edges of the UserFollowInfo.
func (UserFollowInfo) Edges() []ent.Edge {
	return nil
}
