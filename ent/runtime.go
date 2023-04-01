// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/site-tech/jaw-platform/ent/account"
	"github.com/site-tech/jaw-platform/ent/schema"
	"github.com/site-tech/jaw-platform/ent/tennant"
	"github.com/google/uuid"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	accountFields := schema.Account{}.Fields()
	_ = accountFields
	// accountDescCreatedAt is the schema descriptor for created_at field.
	accountDescCreatedAt := accountFields[2].Descriptor()
	// account.DefaultCreatedAt holds the default value on creation for the created_at field.
	account.DefaultCreatedAt = accountDescCreatedAt.Default.(func() time.Time)
	// accountDescID is the schema descriptor for id field.
	accountDescID := accountFields[0].Descriptor()
	// account.DefaultID holds the default value on creation for the id field.
	account.DefaultID = accountDescID.Default.(func() uuid.UUID)
	tennantFields := schema.Tennant{}.Fields()
	_ = tennantFields
	// tennantDescCreatedAt is the schema descriptor for created_at field.
	tennantDescCreatedAt := tennantFields[4].Descriptor()
	// tennant.DefaultCreatedAt holds the default value on creation for the created_at field.
	tennant.DefaultCreatedAt = tennantDescCreatedAt.Default.(func() time.Time)
	// tennantDescID is the schema descriptor for id field.
	tennantDescID := tennantFields[0].Descriptor()
	// tennant.DefaultID holds the default value on creation for the id field.
	tennant.DefaultID = tennantDescID.Default.(func() uuid.UUID)
}