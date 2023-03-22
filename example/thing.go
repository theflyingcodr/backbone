// Package things is the main domain package and contains all
// models, validation and interface definitions.
package things

import (
	"context"
	"time"

	validator "github.com/theflyingcodr/govalidator"
)

// Thing defines a single thing.
type Thing struct {
	ThingID    uint64    `json:"thingID" db:"thingID"`
	Name       string    `json:"name" db:"name"`
	CreatedAt  time.Time `json:"createdAt" db:"createdAt"`
	ModifiedAt time.Time `json:"modifiedAt" db:"modifiedAt"`
}

// ThingArgs are used to retrieve a single thing.
type ThingArgs struct {
	ThingID uint64 `param:"thingID" query:"thingID"`
}

// Validate will ensure ThingArgs meet expectations.
func (t *ThingArgs) Validate() validator.ErrValidation {
	return validator.New().
		Validate("thingID", validator.MinUInt64(t.ThingID, 1))
}

// ThingCreate is used to create a new thing.
type ThingCreate struct {
	Name string `json:"name" db:"name"`
	// CreatedAt set in the data layer before storing - could also be a db default value of current_date.
	CreatedAt time.Time `json:"-" db:"createdAt"`
}

// Validate will ensure ThingCreate requests meet expectations.
func (t *ThingCreate) Validate() validator.ErrValidation {
	return validator.New().
		Validate("name", validator.Length(t.Name, 2, 25))
}

// ThingUpdate is used to update a thing.
type ThingUpdate struct {
	Name string `json:"name" db:"name"`
	// ModifiedAt used to set the update time in the data layer.
	ModifiedAt time.Time `json:"-" db:"modifiedAt"`
}

// Validate will ensure ThingUpdate requests meet expectations.
func (t *ThingUpdate) Validate() validator.ErrValidation {
	return validator.New().
		Validate("name", validator.Length(t.Name, 2, 25))
}

// ThingService is used to enforce business rules, orchestrate data store calls and raise business events.
//
// NOTE: the naming conventions, we don't prefix with Get or Fetch etc, simply the method name
// is the name of the ting we are retrieving.
// In terms of the Arg and request objects, these follow the convention <Name><Action> ie ThingCreate.
// Doing this means you can group your structs so it's easier to find when you type domain.Th...
type ThingService interface {
	// Thing will return a single Thing.
	Thing(ctx context.Context, args ThingArgs) (*Thing, error)
	// Things will get all things.
	Things(ctx context.Context) ([]*Thing, error)
	// Create will add a new thing.
	Create(ctx context.Context, req ThingCreate) (*Thing, error)
	// Update will update an existing thing.
	Update(ctx context.Context, args ThingArgs, req ThingUpdate) (*Thing, error)
	// Delete will remove a thing.
	Delete(ctx context.Context, args ThingArgs) error
}

// ThingReader is used to read thing or things from a data store.
type ThingReader interface {
	Thing(ctx context.Context, args ThingArgs) (*Thing, error)
	Things(ctx context.Context) ([]*Thing, error)
}

// ThingWriter can be implemented to write and modify things stored in a datastore.
type ThingWriter interface {
	Create(ctx context.Context, req ThingCreate) (*Thing, error)
	Update(ctx context.Context, args ThingArgs, req ThingUpdate) (*Thing, error)
	Delete(ctx context.Context, args ThingArgs) error
}

// ThingReaderWriter combines the reader and writer interface to define a read/write datastore for
// working with things.
type ThingReaderWriter interface {
	ThingReader
	ThingWriter
}

// ThingCacher implements a RW but also adds a Cache method.
type ThingCacher interface {
	ThingReaderWriter
	// Cache can be used to update datastores with a freshly created thing, inc id etc.
	Cache(ctx context.Context, req Thing) error
}
