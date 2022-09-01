package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"github.com/google/uuid"
)

// Cat holds the schema definition for the Cat entity.
type Cat struct {
	ent.Schema
}

// Fields of the Cat.
func (Cat) Fields() []ent.Field {
	return []ent.Field{
		field.
			String("id").
			DefaultFunc(uuid.NewString).
			Unique(),
		field.String("name").Default(""),
		field.String("description").Default(""),
		field.String("image_id").Default(""),
		field.String("owner_id"),
		field.Strings("tags").Default([]string{}),
	}
}

// Edges of the Cat.
func (Cat) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("owner", User.Type).
			Field("owner_id").
			Required().
			Unique(),
	}
}

// Mixin of the Cat.
func (Cat) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}
