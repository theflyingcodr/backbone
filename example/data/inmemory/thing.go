package inmemory

import (
	"context"
	"sort"
	"sync"
	"time"
	"github.com/pkg/errors"
	"github.com/theflyingcodr/things"
)

type thing struct{
	sync.RWMutex
	store map[uint64]*things.Thing
}

// NewThing will setup and return a new inmemory thing store.
// This is used for local caching of data. If you wanted a distributed cache you'd
// implement a redis store or equivalent.
func NewThing() *thing{
	return &thing{store: map[uint64]*things.Thing{}}
}

// Thing returns a single thing, if found, or an error if not.
func (t *thing) Thing(ctx context.Context, args things.ThingArgs) (*things.Thing, error){
	t.RLock()
	defer t.RUnlock()
	resp, ok := t.store[args.ThingID]
	if !ok{
		return nil, errors.Errorf("thing with ID %d not found", args.ThingID)
	}
	return resp, nil
}

// Things will return all things currently stored in order of ThingID.
func (t *thing) Things(ctx context.Context) ([]*things.Thing, error){
	t.RLock()
	defer t.RUnlock()
	 tt := make([]*things.Thing,0,len(t.store))
	for _, v := range t.store{
		tt = append(tt, v)
	}
	sort.Slice(tt, func(i, j int) bool {
		return tt[i].ThingID < tt[j].ThingID
	})
	return tt, nil
}

// Cache will add a full Thing to the cache.
func (t *thing) Cache(ctx context.Context, req things.Thing) error{
	t.Lock()
	defer t.Unlock()
	if _, ok := t.store[req.ThingID]; ok{
		return errors.Errorf("thing with id %d already exists in the cache", req.ThingID)
	}
	t.store[req.ThingID] = &req
	return nil
}


// Create will add a new thing to the data store.
func (t *thing) Create(ctx context.Context, req things.ThingCreate) (*things.Thing, error){
	t.Lock()
	defer t.Unlock()
	id := uint64(len(t.store)+1)
	t.store[id] = &things.Thing{
		ThingID:    id,
		Name:       req.Name,
		CreatedAt:  time.Now().UTC(),
		ModifiedAt:  time.Now().UTC(),
	}
	return t.store[id], nil
}


// update will update a thing in the data store.
func (t *thing) Update(ctx context.Context, args things.ThingArgs, req things.ThingUpdate) (*things.Thing, error){
	t.Lock()
	defer t.Unlock()
	th, ok := t.store[args.ThingID]
	if !ok{
		return nil, errors.Errorf("thing with ID %d not found when updating", args.ThingID)
	}
	th.Name = req.Name
	th.ModifiedAt = req.ModifiedAt

	return th, nil
}

// Delete will remove a thing from the data store.
func (t *thing) Delete(ctx context.Context, args things.ThingArgs) error{
	t.Lock()
	defer t.Unlock()
	delete(t.store,args.ThingID)
	return nil
}
