# Example

This shows a server implementation of the backbone template.

It exposes a single rest server which uses [Echo](https://echo.labstack.com/) and has two data stores available, an inmemory and mysql data store (mysql is mostly TODO).

This Readme will go from top to bottom in terms of the data flow coming into the app:

Transport -> Service -> Data 

## Thing

There is a single domain in this example named Thing. Usually a server would have more than one domain type, in this pattern each would be added as a separate file to the root of the domain.

The Thing domain supports full CRUD ops so you can see examples of reads, updates, creates and deletes through the 3 layers of the application.

As with the template, the main objects of the domain are created at the root of the repo. This file is named `thing.go` and contains all models, validators and interfaces needed to implement Thing behaviours.

## Transports

There is a single transport type in this example, rest, it uses Echo handler funcs but you could use anything here.

It takes a service injected to the constructor and you will notice there is literally 0 business logic here, it's only job is to parse requests and form responses.

## Service

The service layer implements the full CRUD operations as defined in the ThingService interface.

This layer enforces the validation logic defined in the domain and also asserts and applies some other business logic.

It receives an injected ThingReaderWriter interface and is responsible for deciding when to call the data store and how to handle the errors.

Not implemented but illustrated by comments is where you could raise business events such as via rabbitMQ or Kafka.

### Injected stores

In this example I'm injecting two different ThingReaderWriters in different places, in the service/thing_test I inject a mocked ThingReaderWriter.

In the main.go I inject a thing_facade (described below) which orchestrates between a cache and database.

## Data Facade

This example shows how we could handle caching with this layout or general orchestrating between different data stores.

It adds a thing_facade to the data package and orchestrates between inmemory and mysql data stores, keeping this logic out of the service layer.

A Service should never have to change due to a data store concern.

This is a neat pattern to follow and allows for unit testing of the data store switching logic.

## Data

There are two data stores, an inmemory store used for caching and a mysql store.

The mysql store is purely illustrative and isn't yet implemented.

## Config

I've yet to publish my configuration package but you would add your Configuration structs and setup to the config package.

12 Factor Apps recommend Environment Variable based configuration so I'd recommend getting values into your config via environment variables and writing them to well formed structs so you can easier reference them and inject into your services etc that rely on config.

Try not to have configlib.Get(key) calls throughout your code, it's annoying to test and hard to replace.

## Mocks

Personally I like Mat Ryers [Moq library](https://github.com/matryer/moq) so that's what I use here, I find it very simple to setup (auto generated) and use (it's just functions).

## Validation

You will notice the Validate() calls in the Thing domain root, this uses my own [Validator](https://github.com/theflyingcodr/govalidator) package.


