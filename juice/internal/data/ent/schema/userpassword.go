package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
)

// UserPassword holds the schema definition for the UserPassword entity.
type UserPassword struct {
	ent.Schema
}

// Fields of the UserPassword.
func (UserPassword) Fields() []ent.Field {

	return []ent.Field{

		field.Int64("user_id").SchemaType(map[string]string{
			dialect.MySQL: "BIGINTUNSIGNED", // Override MySQL.
		}).NonNegative().Optional().Unique(),

		field.String("salt").SchemaType(map[string]string{
			dialect.MySQL: "CHAR(64)", // Override MySQL.
		}),

		field.String("pwd").SchemaType(map[string]string{
			dialect.MySQL: "CHAR(64)", // Override MySQL.
		}),
	}

}

// Edges of the UserPassword.
func (UserPassword) Edges() []ent.Edge {
	return nil
}
