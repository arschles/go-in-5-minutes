# Mocking in Go

Go in 5 Minutes, episode 0.

This screencast focuses on mocking external dependencies so you can write fast,
focused unit tests for your code.

Find this screencast at https://www.youtube.com/watch?v=mk4BCLimksY&feature=youtu.be.

## Outline

1. Execute your code in isolation
2. Swap out dependencies with local implementation. Example: [Redis client](http://godoc.org/github.com/hoisie/redis)
3. Other languages:
  - Reflection
  - Monkey patching
  - Magic
4. The Go Way = simplicity. Use an `interface`
