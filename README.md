# backbone - WIP

Go doesn't provide a default opinionated project template like other languages, I think this is one of its strengths, it allows you to get on with it.

However, I have found I do tend to end up with the same kind of structure over and over again, so decided to template it for my own use and others.

Of course, layout is a subjective thing so in no way is this "the perfect go layout" or "how to build go apps", everyone has their own ideas and preferences, this represents how best fits my brains view of a go project.

Backbone is a template for a go server, this is what I use to bootstrap all my projects and adheres as closely as I can get to DDD and Hexagonal architecture.

## WIP

This isn't yet finished, I'll be adding a few things over time:

* Config structs and a default implementation using Viper
* Code generation - for auto generating domain objects, the related transport, service and data layers

## Goals

It's very easy to write a service quickly in go, this is great because you can test out ideas very rapidly and move quickly.

This however falls over when building enterprise grade services in teams. This rapidly developed code usually ends up hard, if not impossible to unit test, usually relies on global variable or abuse of context values and has no clear separation of application concerns.

This template aims to remedy that by:

* Adhering to well established software engineering princples
* Using Effective Go best practises
* Adding clear separation of concerns by splitting the app into 3 logical layers (see [ARCHITECTURE.md](ARCHITECTURE.md) for more info)
* Highly testable as we follow SOLID principles and code to interfaces
* Replacing global variables with Dependency Injection
* Use context purely for timeouts and request specific info like requestID

## Inspiration

Some inspiration I found when coming up with this is listed below:

* [EffectiveGo](https://golang.org/doc/effective_go)
* [Solid Go - video](https://www.youtube.com/watch?v=zzAdEt3xZ1M)
* [Go Best Practises - video](https://www.youtube.com/watch?v=MzTcsI6tn-0)
* [Clean Architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)
