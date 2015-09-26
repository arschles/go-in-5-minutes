# Go In Five Minutes, Episode 0 - Mocking in Go

This screencast focuses on mocking functionality in Go.

## Outline

1. When we write unit tests, we want to execute production code in isolation
  - That means it shouldn't talk to external databases, APIs, etc...
2. We want to achieve this isolation by replacing dependencies with local implementations
  - For example, replace a [Redis client](http://godoc.org/github.com/hoisie/redis) with a local, in-memory implementation
3. Other languages use reflection, monkey patching and other methods generally accepted as "magic"
  - We're gonna follow the Go way and keep things simple
4. Make an `interface` for your Redis client, have your production code accept the `interface`
  - And implement and use a local, in-memory version for tests
