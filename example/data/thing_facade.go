package data

import (
	"context"

	"github.com/labstack/gommon/log"
	"github.com/pkg/errors"

	"github.com/theflyingcodr/things"
)

type thingFacade struct{
	inMem things.ThingCacher
	mysql things.ThingReaderWriter
}

// NewThingFacade will setup and return a new thingFacade that contains a caching layer on top of the db.
//
// Not everything will need this, I am adding to illustrate one pattern for handing multiple data stores, a
// common reason would be caching. IE your main expensive db will usually be slow, so we want a faster layer
// to cache and return the data quickly.
//
// This pattern is nice because it removes this logic from the service layer to the data layer. The service
// layer should be entirely agnostic to the data stores, so the fact we are caching is a data layer concern.
// If we decided to change our cache or to remove it, if this was added to the service, the service would need
// to change. The only reason for a service to change is business logic changing.
func NewThingFacade(inMem things.ThingCacher, mysql things.ThingReaderWriter) *thingFacade{
	return &thingFacade{
		inMem: inMem,
		mysql: mysql,
	}
}

// Thing will read a thing from the datastore, attempting the local cache first before falling back
// to the database. If found in the db, we add it to the cache and return the result.
func (t *thingFacade) Thing(ctx context.Context, args things.ThingArgs) (*things.Thing, error){
	// attempt to get from the cache first, if it errors or returns empty, read from db.
	th, err := t.inMem.Thing(ctx, args)
	if err == nil && th != nil{
		return th, nil
	}
	// Not cached, fall back to expensive store.
	th, err = t.mysql.Thing(ctx, args)
	if err != nil{
		return nil, errors.WithStack(err)
	}
	// cache it - could be multiple cache calls if you've multiple caching layers
	// or it could emit a caching event to be handled async via a consumer so you don't block here.
	if err := t.inMem.Cache(ctx, *th); err != nil{
		// just log this error, don't prevent the user getting their data
		// due to a caching problem.
		// In a production system this should alert somewhere.
		log.Error(err)
	}
	return th, nil
}

// Things will attempt to return all things via the cache falling back to the
// sql store if not found there.
func (t *thingFacade) Things(ctx context.Context) ([]*things.Thing, error){
	tt, err := t.inMem.Things(ctx)
	if err == nil && tt != nil{
		return tt, nil
	}
	// could cache them all after getting them depending on your needs.
	return t.mysql.Things(ctx)
}

// Create will add a new thing to the main datastore then cache it.
func (t *thingFacade) Create(ctx context.Context, req things.ThingCreate) (*things.Thing, error){
	// create in main data store first to get IDs, set createdAt etc
	th, err := t.mysql.Create(ctx, req)
	if err != nil{
		return nil, errors.WithStack(err)
	}
	if err := t.inMem.Cache(ctx, *th); err != nil{
		// if we fail to cache just log to be picked up by alerting
		// systems etc.
		// Next time this is fetched it will be cached.
		log.Error(err)
	}
	return th, nil
}

// Update will update a thing in the datastore and attempt to update the caches.
func (t *thingFacade) Update(ctx context.Context, args things.ThingArgs, req things.ThingUpdate) (*things.Thing, error){
	th, err := t.mysql.Update(ctx, args, req)
	if err != nil{
		// failed to update our main db, fail the request.
		return nil, errors.WithStack(err)
	}
	if _, err := t.inMem.Update(ctx, args, req); err != nil{
		// if we fail to cache just log to be picked up by alerting
		// systems etc.
		log.Error(err)
		// remove the item from the cache so it's freshly cached next time.
		if err := t.inMem.Delete(ctx, args); err != nil{
			log.Error(err)
		}
	}
	return th, nil
}

// Delete will remove a thing from the datastore and caches.
func (t *thingFacade) Delete(ctx context.Context, args things.ThingArgs) error{
	if err := t.mysql.Delete(ctx, args); err != nil{
		return errors.WithStack(err)
	}
	if err := t.inMem.Delete(ctx, args); err != nil{
		log.Error(err)
	}
	return nil
}
