package backbone

import (
	"context"
	"time"
)

// Thing defines a single thing.
type Thing struct {
	ThingID    int64     `json:"thingID" db:"thingID"`
	Name       string    `json:"name" db:"name"`
	CreatedAt  time.Time `json:"createdAt" db:"createdAt"`
	ModifiedAt time.Time `json:"modifiedAt" db:"modifiedAt"`
}

// ThingArgs are used to retrieve a single thing.
type ThingArgs struct {
	ThingID int64 `param:"thingID" query:"thingID"`
}

// ThingCreate is used to create a new thing.
type ThingCreate struct {
	Name string `json:"name" db:"name"`
	// CreatedAt set in the data layer before storing - could also be a db default value of current_date.
	CreatedAt time.Time `json:"-" db:"createdAt"`
}

// ThingUpdate is used to update a thing.
type ThingUpdate struct {
	Name string `json:"name" db:"name"`
	// ModifiedAt used to set the update time in the data layer.
	ModifiedAt time.Time `json:"-" db:"modifiedAt"`
}

// ThingService is used to enforce business rules, orchestrate data store calls and raise business events.
type ThingService interface {
	Thing(ctx context.Context, args ThingArgs) (*Thing, error)
	Things(ctx context.Context) ([]*Thing, error)
	Create(ctx context.Context, req ThingCreate) (*Thing, error)
	Update(ctx context.Context, args ThingArgs, req ThingUpdate) (*Thing, error)
	Delete(ctx context.Context, args ThingArgs) error
}

// ThingReader is used to read thing or things from a data store.
type ThingReader interface {
	Thing(ctx context.Context, args ThingArgs) (*Thing, error)
	Things(ctx context.Context) ([]*Thing, error)
}

// ThingWriter can be implemented to write and modify things stored in a datastore.
type ThingWriter interface {
	Create(ctx context.Context, args ThingCreate) (*Thing, error)
	Update(ctx context.Context, args ThingArgs, req ThingUpdate) (*Thing, error)
	Delete(ctx context.Context, args ThingArgs) error
}

// ThingReaderWriter combines the reader and writer interface to define a read/write datastore for
// working with things.
type ThingReaderWriter interface {
	ThingReader
	ThingWriter
}
