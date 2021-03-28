# Architecture

This uses a tried and tested simple 3 tier architecture with an abstraction between each layer provided by interfaces.



As shown in the above diagram, it has 3 layers, Transport, Service & Data.

These 3 layers each have an interface separating them enforcing the clear separation of concerns and enhancing testability.

I will list out their roles and responsibilities below:

## Transport

This layer should be as dumb as possible, no logic or validation should be here, the only thing this layer should do is parse requests or streams, pass to the service layer and then format a response from the service layer.

Examples of transports are:

* RabbitMQ consumer
* Rest API
* GRPC Server
* Json RPC Server
* etc

If it is used to get data INTO your application it should go into this layer.

### Accepts

A transport handler will accept a Service interface in its constructor.

## Service

This is the core of application and works hand in hand with the root objects which form the core of your domain.

The core objects (located in the root of the application) define the properties each object has as well as the Service and Data interfaces.

It is a good idea to add validation in receiver methods to these objects as shown:

```go
// ThingArgs are used to retrieve a single thing.
type ThingArgs struct {
	ThingID int64 `param:"thingID" query:"thingID"`
}

// Validate enforces ThingArgs rules.
func (t ThingArgs) Validate() validator.ErrValidation{
	return validator.New().
		Validate("thingID", validator.MinInt64(t.ThingID,1))
}
```
The Service will then enforce these business validation rules as it's first task:
```go
func (t *thing) Create(ctx context.Context, req backbone.ThingCreate) (*backbone.Thing, error) {
    if err := req.Validate().Err(); err != nil{
        return nil, err
    }
    ...
}
```

The Service key responsibility is enforcing and implementing business rules, but it also has a few other potential tasks:

* Raise application events
* Orchestrates the data calls, if any ie which data methods to call and when.

The service should be unchanged by either the transport or data layer changing ie if you decide to add caching, the service should be agnostic to this, instead the data layer should orchestrate this.

### Accepts
A service constructor will usually take at least a Data Reader/Writer interface, but may also take:

* Configuration struct
* Event Publisher
* Logger

## Data

A general rule for what belongs in the data layer is, if the service would need to perform I/O (database, http, file calls etc) to get the data, it should be added to the data layer.

Each data store type should have a package added under the data folder ie data/mysql data/file data/inmemory etc.

As with the transport layer, this layer should be dumb, no business logic should be here, instead, it's only concerned with adding or retrieving data.

### Accepts

The data constructor will at least accept a data store interface / object, this could be a sql.DB interface, grpc client or an http client for example.



