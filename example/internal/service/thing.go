package service

import (
	"context"
	"strings"
	"time"

	"github.com/pkg/errors"

	"github.com/theflyingcodr/things"
)

type thing struct {
	store things.ThingReaderWriter
}

// NewThing will setup and return a new Thing service.
//
// Note: the constructor will ALWAYS take an interface (code to interfaces),
// this means you can pass anything that implements the readerwriter including mocks, facades etc.
// In this example we will inject the data.thingFacade to this under cmd/rest-server.
func NewThing(store things.ThingReaderWriter) *thing {
	return &thing{store: store}
}

// Thing will return a single thing matching args.
func (t *thing) Thing(ctx context.Context, args things.ThingArgs) (*things.Thing, error) {
	if err := args.Validate().Err(); err != nil{
		return nil, err
	}
	// could raise an event here logging that a thing has been read
	resp, err := t.store.Thing(ctx, args)
	if err != nil {
		return nil, errors.WithMessagef(err, "failed to get thing with id %d", args.ThingID)
	}
	return resp, nil
}

// Things will return all things.
func (t *thing) Things(ctx context.Context) ([]*things.Thing, error) {
	// could raise an event here logging that things have been read
	resp, err := t.store.Things(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get things")
	}
	return resp, nil
}

// Create will attempt to add a new Thing to the datastore.
func (t *thing) Create(ctx context.Context, req things.ThingCreate) (*things.Thing, error) {
	if err := req.Validate().Err(); err != nil{
		return nil, err
	}
	// some logic here to illustrate
	if strings.HasPrefix(req.Name, "test"){
		req.Name = strings.Title(req.Name)
	} else{
		req.Name = strings.ToLower(req.Name)
	}
	req.CreatedAt = time.Now().UTC()
	th, err := t.store.Create(ctx, req)
	if err != nil{
		return nil, errors.WithMessage(err, "failed to add thing to data store.")
	}
	// again dummy logic here to show that this service layer handles
	// your business logic, here if the dates don't match we fail, this
	// would be due to some business reason.
	if th.ModifiedAt != th.CreatedAt {
		return nil, errors.New("dates don't match")
	}
	// could raise a business event here.
	return th, nil
}

// Update will update an existing thing if found. Modified date should not be before the created date after update.
func (t *thing) Update(ctx context.Context, args things.ThingArgs, req things.ThingUpdate) (*things.Thing, error) {
	if err := args.Validate().Err(); err != nil{
		return nil, err
	}
	if err := req.Validate().Err(); err != nil{
		return nil, err
	}
	th, err := t.store.Update(ctx, args, req)
	if err != nil{
		return nil, errors.WithMessagef(err, "failed to update thing with ID %d", args.ThingID)
	}
	if th.ModifiedAt.Before(th.CreatedAt){
		return nil, errors.New("modified date isn't right")
	}
	return th, nil
}

// Delete will remove a thing.
func (t *thing) Delete(ctx context.Context, args things.ThingArgs) error {
	if err := args.Validate().Err(); err != nil{
		return err
	}
	return errors.WithMessagef(t.store.Delete(ctx, args), "failed to remove thing with ID %d", args.ThingID)
}
