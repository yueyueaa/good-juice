package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"

	"time"
)

// UserBaseInfo holds the schema definition for the UserBaseInfo entity.
type UserBaseInfo struct {
	ent.Schema
}

// Fields of the UserBaseInfo.
func (UserBaseInfo) Fields() []ent.Field {

	return []ent.Field{

		field.Int64("user_id").SchemaType(map[string]string{
			dialect.MySQL: "BIGINTUNSIGNED", // Override MySQL.
		}).NonNegative().Unique(),

		field.String("username").SchemaType(map[string]string{
			dialect.MySQL: "CHAR(30)", // Override MySQL.
		}),

		field.Int8("sex").SchemaType(map[string]string{
			dialect.MySQL: "TINYINT", // Override MySQL.
		}),

		field.Time("birth").SchemaType(map[string]string{
			dialect.MySQL: "DATE", // Override MySQL.
		}).Default(time.Now),

		field.Int32("area").SchemaType(map[string]string{
			dialect.MySQL: "INT", // Override MySQL.
		}),

		field.String("user_profile").SchemaType(map[string]string{
			dialect.MySQL: "CHAR(255)", // Override MySQL.
		}).Optional(),

		field.String("user_profile_photo_url").SchemaType(map[string]string{
			dialect.MySQL: "CHAR(255)", // Override MySQL.
		}).Optional(),

		field.Int32("follow_count").SchemaType(map[string]string{
			dialect.MySQL: "INT", // Override MySQL.
		}).Default(0),

		field.Int32("fan_count").SchemaType(map[string]string{
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

// Edges of the UserBaseInfo.
func (UserBaseInfo) Edges() []ent.Edge {
	return nil
}
