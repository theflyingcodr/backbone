# Data

If an application needs to read or write data from anywhere, this is the layer that handles the data store.

A data store could be many things, not limited to but including:

* SQL Database
* NoSql database
* File system
* Grpc endpoint
* Http endpoints

## Responsibilities

A general rule of thumb is that if I/O is required, it should be in this layer and not the service layer.

As with the Transports layer, there should be no business logic in this layer, ONLY data store speicifc code relating to retrieval, modification or storage of data.

This would be SQL get statements, file system reads or http client calls to an external api for example.

The Service layer will orchestrate these calls.

## DTOs

This layer MAY use the domain models to fetch and retrieve data, but, they may need to create their own Data Transfer Objects depending on requirements. 
For example, you may require nested at the domain level, but the data store may not support mapping of this type of data. 

In this instance you'd create a `data/datastore/models` folder, add your object and code to map between the DTO and Domain objects and add your unit tests.

This way you keep your Domain objects separate from database specific implementation details.


