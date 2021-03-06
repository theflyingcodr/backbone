// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package mocks

import (
	"context"
	"github.com/theflyingcodr/backbone"
	"sync"
)

var (
	lockThingServiceMockCreate sync.RWMutex
	lockThingServiceMockDelete sync.RWMutex
	lockThingServiceMockThing  sync.RWMutex
	lockThingServiceMockThings sync.RWMutex
	lockThingServiceMockUpdate sync.RWMutex
)

// Ensure, that ThingServiceMock does implement ThingService.
// If this is not the case, regenerate this file with moq.
var _ backbone.ThingService = &ThingServiceMock{}

// ThingServiceMock is a mock implementation of ThingService.
//
//     func TestSomethingThatUsesThingService(t *testing.T) {
//
//         // make and configure a mocked ThingService
//         mockedThingService := &ThingServiceMock{
//             CreateFunc: func(ctx context.Context, req backbone.ThingCreate) (*backbone.Thing, error) {
// 	               panic("mock out the Create method")
//             },
//             DeleteFunc: func(ctx context.Context, args backbone.ThingArgs) error {
// 	               panic("mock out the Delete method")
//             },
//             ThingFunc: func(ctx context.Context, args backbone.ThingArgs) (*backbone.Thing, error) {
// 	               panic("mock out the Thing method")
//             },
//             ThingsFunc: func(ctx context.Context) ([]*backbone.Thing, error) {
// 	               panic("mock out the Things method")
//             },
//             UpdateFunc: func(ctx context.Context, args backbone.ThingArgs, req backbone.ThingUpdate) (*backbone.Thing, error) {
// 	               panic("mock out the Update method")
//             },
//         }
//
//         // use mockedThingService in code that requires ThingService
//         // and then make assertions.
//
//     }
type ThingServiceMock struct {
	// CreateFunc mocks the Create method.
	CreateFunc func(ctx context.Context, req backbone.ThingCreate) (*backbone.Thing, error)

	// DeleteFunc mocks the Delete method.
	DeleteFunc func(ctx context.Context, args backbone.ThingArgs) error

	// ThingFunc mocks the Thing method.
	ThingFunc func(ctx context.Context, args backbone.ThingArgs) (*backbone.Thing, error)

	// ThingsFunc mocks the Things method.
	ThingsFunc func(ctx context.Context) ([]*backbone.Thing, error)

	// UpdateFunc mocks the Update method.
	UpdateFunc func(ctx context.Context, args backbone.ThingArgs, req backbone.ThingUpdate) (*backbone.Thing, error)

	// calls tracks calls to the methods.
	calls struct {
		// Create holds details about calls to the Create method.
		Create []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Req is the req argument value.
			Req backbone.ThingCreate
		}
		// Delete holds details about calls to the Delete method.
		Delete []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Args is the args argument value.
			Args backbone.ThingArgs
		}
		// Thing holds details about calls to the Thing method.
		Thing []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Args is the args argument value.
			Args backbone.ThingArgs
		}
		// Things holds details about calls to the Things method.
		Things []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
		}
		// Update holds details about calls to the Update method.
		Update []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Args is the args argument value.
			Args backbone.ThingArgs
			// Req is the req argument value.
			Req backbone.ThingUpdate
		}
	}
}

// Create calls CreateFunc.
func (mock *ThingServiceMock) Create(ctx context.Context, req backbone.ThingCreate) (*backbone.Thing, error) {
	if mock.CreateFunc == nil {
		panic("ThingServiceMock.CreateFunc: method is nil but ThingService.Create was just called")
	}
	callInfo := struct {
		Ctx context.Context
		Req backbone.ThingCreate
	}{
		Ctx: ctx,
		Req: req,
	}
	lockThingServiceMockCreate.Lock()
	mock.calls.Create = append(mock.calls.Create, callInfo)
	lockThingServiceMockCreate.Unlock()
	return mock.CreateFunc(ctx, req)
}

// CreateCalls gets all the calls that were made to Create.
// Check the length with:
//     len(mockedThingService.CreateCalls())
func (mock *ThingServiceMock) CreateCalls() []struct {
	Ctx context.Context
	Req backbone.ThingCreate
} {
	var calls []struct {
		Ctx context.Context
		Req backbone.ThingCreate
	}
	lockThingServiceMockCreate.RLock()
	calls = mock.calls.Create
	lockThingServiceMockCreate.RUnlock()
	return calls
}

// Delete calls DeleteFunc.
func (mock *ThingServiceMock) Delete(ctx context.Context, args backbone.ThingArgs) error {
	if mock.DeleteFunc == nil {
		panic("ThingServiceMock.DeleteFunc: method is nil but ThingService.Delete was just called")
	}
	callInfo := struct {
		Ctx  context.Context
		Args backbone.ThingArgs
	}{
		Ctx:  ctx,
		Args: args,
	}
	lockThingServiceMockDelete.Lock()
	mock.calls.Delete = append(mock.calls.Delete, callInfo)
	lockThingServiceMockDelete.Unlock()
	return mock.DeleteFunc(ctx, args)
}

// DeleteCalls gets all the calls that were made to Delete.
// Check the length with:
//     len(mockedThingService.DeleteCalls())
func (mock *ThingServiceMock) DeleteCalls() []struct {
	Ctx  context.Context
	Args backbone.ThingArgs
} {
	var calls []struct {
		Ctx  context.Context
		Args backbone.ThingArgs
	}
	lockThingServiceMockDelete.RLock()
	calls = mock.calls.Delete
	lockThingServiceMockDelete.RUnlock()
	return calls
}

// Thing calls ThingFunc.
func (mock *ThingServiceMock) Thing(ctx context.Context, args backbone.ThingArgs) (*backbone.Thing, error) {
	if mock.ThingFunc == nil {
		panic("ThingServiceMock.ThingFunc: method is nil but ThingService.Thing was just called")
	}
	callInfo := struct {
		Ctx  context.Context
		Args backbone.ThingArgs
	}{
		Ctx:  ctx,
		Args: args,
	}
	lockThingServiceMockThing.Lock()
	mock.calls.Thing = append(mock.calls.Thing, callInfo)
	lockThingServiceMockThing.Unlock()
	return mock.ThingFunc(ctx, args)
}

// ThingCalls gets all the calls that were made to Thing.
// Check the length with:
//     len(mockedThingService.ThingCalls())
func (mock *ThingServiceMock) ThingCalls() []struct {
	Ctx  context.Context
	Args backbone.ThingArgs
} {
	var calls []struct {
		Ctx  context.Context
		Args backbone.ThingArgs
	}
	lockThingServiceMockThing.RLock()
	calls = mock.calls.Thing
	lockThingServiceMockThing.RUnlock()
	return calls
}

// Things calls ThingsFunc.
func (mock *ThingServiceMock) Things(ctx context.Context) ([]*backbone.Thing, error) {
	if mock.ThingsFunc == nil {
		panic("ThingServiceMock.ThingsFunc: method is nil but ThingService.Things was just called")
	}
	callInfo := struct {
		Ctx context.Context
	}{
		Ctx: ctx,
	}
	lockThingServiceMockThings.Lock()
	mock.calls.Things = append(mock.calls.Things, callInfo)
	lockThingServiceMockThings.Unlock()
	return mock.ThingsFunc(ctx)
}

// ThingsCalls gets all the calls that were made to Things.
// Check the length with:
//     len(mockedThingService.ThingsCalls())
func (mock *ThingServiceMock) ThingsCalls() []struct {
	Ctx context.Context
} {
	var calls []struct {
		Ctx context.Context
	}
	lockThingServiceMockThings.RLock()
	calls = mock.calls.Things
	lockThingServiceMockThings.RUnlock()
	return calls
}

// Update calls UpdateFunc.
func (mock *ThingServiceMock) Update(ctx context.Context, args backbone.ThingArgs, req backbone.ThingUpdate) (*backbone.Thing, error) {
	if mock.UpdateFunc == nil {
		panic("ThingServiceMock.UpdateFunc: method is nil but ThingService.Update was just called")
	}
	callInfo := struct {
		Ctx  context.Context
		Args backbone.ThingArgs
		Req  backbone.ThingUpdate
	}{
		Ctx:  ctx,
		Args: args,
		Req:  req,
	}
	lockThingServiceMockUpdate.Lock()
	mock.calls.Update = append(mock.calls.Update, callInfo)
	lockThingServiceMockUpdate.Unlock()
	return mock.UpdateFunc(ctx, args, req)
}

// UpdateCalls gets all the calls that were made to Update.
// Check the length with:
//     len(mockedThingService.UpdateCalls())
func (mock *ThingServiceMock) UpdateCalls() []struct {
	Ctx  context.Context
	Args backbone.ThingArgs
	Req  backbone.ThingUpdate
} {
	var calls []struct {
		Ctx  context.Context
		Args backbone.ThingArgs
		Req  backbone.ThingUpdate
	}
	lockThingServiceMockUpdate.RLock()
	calls = mock.calls.Update
	lockThingServiceMockUpdate.RUnlock()
	return calls
}
