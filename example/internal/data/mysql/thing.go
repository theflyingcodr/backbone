package mysql

import (
	"context"
	"database/sql"

	"github.com/pkg/errors"

	"github.com/theflyingcodr/things"
)

const (
	sqlThing = `
	SELECT thingID, thing
	FROM things
	WHERE thingID = :thingID
	`
)

type thing struct {
	db *sql.DB
}

// NewThing takes a database object which will be setup in main and injected to the store.
func NewThing(db *sql.DB) *thing {
	return &thing{db: db}
}

// Thing will return a single thing matching the args supplied.
func (t *thing) Thing(ctx context.Context, args things.ThingArgs) (*things.Thing, error) {
	// TODO: get from db
	row := t.db.QueryRow(sqlThing, args)
	var resp *things.Thing
	if err := row.Scan(&resp); err != nil {
		return nil, errors.Wrap(err, "failed to scan row")
	}
	return resp, nil
}

// Things will return all things, you could also pass some args to support filtering.
func (t *thing) Things(ctx context.Context) ([]*things.Thing, error) {
	// TODO: get from db and return
	return nil, nil
}

// Create will add a new thing to the database.
func (t *thing) Create(ctx context.Context, req things.ThingCreate) (*things.Thing, error) {
	// TODO: add to db and return
	return nil, nil
}

// Update will update a single thing in the database.
func (t *thing) Update(ctx context.Context, args things.ThingArgs, req things.ThingUpdate) (*things.Thing, error) {
	// TODO: update in db and return
	return nil, nil
}

// Delete will remove a thing from the database.
func (t *thing) Delete(ctx context.Context, args things.ThingArgs) error {
	// TODO: remove from db
	return nil
}
