package things

import (
	"errors"
	"testing"
	"time"

	"github.com/matryer/is"
)

func Test_ThingArgs_Validate(t *testing.T){
	t.Parallel()
	is := is.New(t)
	tests := map[string]struct{
		args ThingArgs
		err error
	}{
		"valid args should not return error": {
			args: ThingArgs{
				ThingID: 10,
			},
			err:  nil,
		},"args with thingID of 0 should fail": {
			args: ThingArgs{
				ThingID: 0,
			},
			err:  errors.New("[thingID: value 0 is smaller than minimum 1]"),
		},
	}
	for name, test := range tests{
		t.Run(name, func(t *testing.T) {
			is = is.NewRelaxed(t)
			res := test.args.Validate().Err()
			if test.err == nil{
				is.NoErr(res)
				return

			}
			is.Equal(test.err.Error(),res.Error() )
		})
	}
}

func Test_ThingCreate_Validate(t *testing.T){
	t.Parallel()
	is := is.New(t)
	tests := map[string]struct{
		req ThingCreate
		err error
	}{
		"valid thing should not return error": {
			req: ThingCreate{
				Name:      "Test",
				CreatedAt: time.Time{},
			},
			err:  nil,
		},"name too short should return error": {
			req: ThingCreate{
				Name:      "T",
				CreatedAt: time.Time{},
			},
			err:  errors.New("[name: value must be between 2 and 25 characters]"),
		},"name too long should return error": {
			req: ThingCreate{
				Name:      "I am a name that is far far far too long",
				CreatedAt: time.Time{},
			},
			err:  errors.New("[name: value must be between 2 and 25 characters]"),
		},"name empty should return error": {
			req: ThingCreate{
				Name:      "",
				CreatedAt: time.Time{},
			},
			err:  errors.New("[name: value must be between 2 and 25 characters]"),
		},
	}
	for name, test := range tests{
		t.Run(name, func(t *testing.T) {
			is = is.NewRelaxed(t)
			res := test.req.Validate().Err()
			if test.err == nil{
				is.NoErr(res)
				return

			}
			is.Equal(test.err.Error(),res.Error() )
		})
	}
}

func Test_ThingUpdate_Validate(t *testing.T){
	t.Parallel()
	is := is.New(t)
	tests := map[string]struct{
		req ThingUpdate
		err error
	}{
		"valid thing should not return error": {
			req: ThingUpdate{
				Name:      "Test",
				ModifiedAt: time.Time{},
			},
			err:  nil,
		},"name too short should return error": {
			req: ThingUpdate{
				Name:      "T",
				ModifiedAt: time.Time{},
			},
			err:  errors.New("[name: value must be between 2 and 25 characters]"),
		},"name too long should return error": {
			req: ThingUpdate{
				Name:      "I am a name that is far far far too long",
				ModifiedAt: time.Time{},
			},
			err:  errors.New("[name: value must be between 2 and 25 characters]"),
		},"name empty should return error": {
			req: ThingUpdate{
				Name:      "",
				ModifiedAt: time.Time{},
			},
			err:  errors.New("[name: value must be between 2 and 25 characters]"),
		},
	}
	for name, test := range tests{
		t.Run(name, func(t *testing.T) {
			is = is.NewRelaxed(t)
			res := test.req.Validate().Err()
			if test.err == nil{
				is.NoErr(res)
				return

			}
			is.Equal(test.err.Error(),res.Error() )
		})
	}
}

