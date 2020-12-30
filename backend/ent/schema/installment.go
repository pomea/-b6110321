package schema

import "github.com/facebookincubator/ent"

// Installment holds the schema definition for the Installment entity.
type Installment struct {
	ent.Schema
}

// Fields of the Installment.
func (Installment) Fields() []ent.Field {
	return nil
}

// Edges of the Installment.
func (Installment) Edges() []ent.Edge {
	return nil
}
