package service

import (
	"context"

	"github.com/pkg/errors"

	"github.com/theflyingcodr/backbone"
)

type thing struct {
	store backbone.ThingReaderWriter
}

func NewThing(store backbone.ThingReaderWriter) *thing {
	return &thing{store: store}
}

func (t *thing) Thing(ctx context.Context, args backbone.ThingArgs) (*backbone.Thing, error) {
	if args.ThingID == 0 {
		return nil, errors.New("thing cannot be 0")
	}
	// could raise an event here logging that a thing has been read
	resp, err := t.store.Thing(ctx, args)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to get thing with id %d", args.ThingID)
	}
	return resp, nil
}

func (t *thing) Things(ctx context.Context) ([]*backbone.Thing, error) {
	// could raise an event here logging that things have been read
	resp, err := t.store.Things(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get things")
	}
	return resp, nil
}

func (t *thing) Create(ctx context.Context, req backbone.ThingCreate) (*backbone.Thing, error) {
	if err := req.Validate().Err(); err != nil{
		return nil, err
	}
	// TODO: implement your logic
	return nil, nil
}

func (t *thing) Update(ctx context.Context, args backbone.ThingArgs, req backbone.ThingUpdate) (*backbone.Thing, error) {
	// TODO: implement your logic
	return nil, nil
}

func (t *thing) Delete(ctx context.Context, args backbone.ThingArgs) error {
	// TODO: implement your logic
	return nil
}
