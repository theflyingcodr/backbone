# Transports

This is the entry point of your application and where it will accept data from outside the application, parse it and pass onto the services to handle it in some way.

Transports is a deliberately open term, API implies that this is for GRPC or Rest endpoints, which it is, but that connotation is too limiting in my opinion.

This package, as well as being used for the usual API endpoints can also be used for:

* GRPC handlers
* JSON RPC handlers
* GOB handlers
* Messaging consumers
* anything else that takes data to be processed.

## Responsibilities

The only responsibility of this layer of the application is to parse requests and responses.

This could involve reading query strings, listening to a binary stream of data, parsing XML etc.

Once it has parsed a request it will then pass onto the service layer.

The service layer will handle validation, business rules and data store orchestration before passing back to the transport which will then build and serialise the response.

There should be NO business logic of ANY KIND in this layer of the application, keep it dumb.

## Naming

I'd recommend naming each top level folder with the name of the transport type ie:

* Rest
* JSON_Rpc
* GRPC
* etc

There would then be a corresponding /cmd/... package suffixed with server like /cmd/rest_server.
