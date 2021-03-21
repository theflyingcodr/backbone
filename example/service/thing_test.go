package service

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/matryer/is"

	"github.com/theflyingcodr/things"
	"github.com/theflyingcodr/things/mocks"
)

// TestThing_Thing uses table driven tests with a mock of the ThingReaderWriter.Thing method.
// https://dave.cheney.net/2019/05/07/prefer-table-driven-tests
//
// I've only added a test for the THing method, you'd add separate test methods for each method in
// the target service you are testing.
//
// The mock object used is auto generated and can be found in the mocks/ package.
func TestThing_Thing(t *testing.T) {
	t.Parallel()
	is := is.New(t)
	tests := map[string]struct{
		args things.ThingArgs
		thingMockFn func(ctx context.Context, args things.ThingArgs) (*things.Thing, error)
		exp *things.Thing
		err error
	}{
		"valid request should return a thing from the data store": {
			args:        things.ThingArgs{
				ThingID: 1,
			},
			thingMockFn: func(ctx context.Context, args things.ThingArgs) (*things.Thing, error) {
				return &things.Thing{
					ThingID:    1,
					Name:       "test",
					CreatedAt:  time.Time{},
					ModifiedAt: time.Time{},
				}, nil
			},
			exp: &things.Thing{
				ThingID:    1,
				Name:       "test",
				CreatedAt:  time.Time{},
				ModifiedAt: time.Time{},
			},
			err:         nil,
		},"args with thingId of 0 should error": {
			args:        things.ThingArgs{
				ThingID: 0,
			},
			err:         errors.New("[thingID: value 0 is smaller than minimum 1]"),
		},"data store error should be returned with empty thing": {
			args:        things.ThingArgs{
				ThingID: 1,
			},
			thingMockFn: func(ctx context.Context, args things.ThingArgs) (*things.Thing, error) {
				return nil, errors.New("test failed")
			},
			err:         errors.New("failed to get thing with id 1: test failed"),
		},
	}
	for name, test := range tests{
		t.Run(name, func(t *testing.T) {
			is = is.NewRelaxed(t)
			mock := &mocks.ThingReaderWriterMock{
				ThingFunc: test.thingMockFn,
			}
			svc := NewThing(mock)
			resp, err := svc.Thing(context.Background(), test.args)
			if test.err != nil{
				is.True(resp == nil)
				is.True(err != nil)
				is.Equal(test.err.Error(), err.Error())
				return
			}
			is.True(resp != nil)
			is.True(err == nil)
			is.Equal(test.exp, resp)
		})
	}
}
