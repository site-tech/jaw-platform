package schema

import (
	"time"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Tennant holds the schema definition for the Tennant entity.
type Tennant struct {
	ent.Schema
}

// nolint
func (Tennant) isEntity() {}

// Fields of the Tennant.
func (Tennant) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Unique(),
		field.String("externalId"),
		field.String("cloud"),
		field.UUID("account_id", uuid.UUID{}),
		field.Time("created_at").Default(time.Now),
	}
}

// Edges of the Tennant.
func (Tennant) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("account", Account.Type).
			Ref("tennants").
			Field("account_id").
			Unique().
			Required(),
	}
}

func (Tennant) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(),
		entgql.RelayConnection(),
	}
}
