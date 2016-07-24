# Mocking in Go

[![Watch The Screencast](./../watch-screencast.svg)](http://www.goin5minutes.com/screencast/episode_0_writing_testable_code_and_fast_unit_tests_using_mocking/)

Go in 5 Minutes, episode 0.

This screencast focuses on mocking external dependencies so you can write fast,
focused unit tests for your code.

Screencast video:
http://www.goin5minutes.com/screencast/episode_0_writing_testable_code_and_fast_unit_tests_using_mocking/

## Outline

1. Execute your code in isolation
2. Swap out dependencies with local implementation. Example: [Redis client](http://godoc.org/github.com/hoisie/redis)
3. Other languages:
  - Reflection
  - Monkey patching
  - Magic
4. The Go Way = simplicity. Use an `interface`
