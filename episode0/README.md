# Mocking in Go

Go in 5 Minutes, episode 0.

This screencast focuses on mocking external dependencies so you can write fast,
focused unit tests for your code.

## Outline

1. We want to execute our code ("business logic") in isolation
  - It shouldn't talk to external databases, APIs, etc...
2. Achieve this isolation by swapping out dependencies with local implementations
  - Example: replace a [Redis client](http://godoc.org/github.com/hoisie/redis) with a local implementation
3. Other languages use reflection, monkey patching and other methods. These are magic
  - The Go Way = simplicity
4. Pass an `interface`, swap out the Redis client with your local impl
