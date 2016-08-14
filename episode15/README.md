# Concurrency Series: Internet Cafe

[![Watch The Screencast](http://www.goin5minutes.com/img/watch-screencast.svg)](http://www.goin5minutes.com/screencast/episode_15_internet_cafe/)

_Reminder: subscribe to the newsletter at http://www.goin5minutes.com/subscribe/_

Go in 5 Minutes, episode 15.

This screencast overviews the "Intenet Cafe" problem on [Trivial Concurrency Exercises for the Confused Newbie Gopher](http://whipperstacker.com/2015/10/05/3-trivial-concurrency-exercises-for-the-confused-newbie-gopher/), and provides a solution.

We're continuing the concurrency series - see https://github.com/arschles/go-in-5-minutes/issues/14 for progress.

Screencast video:
http://www.goin5minutes.com/screencast/episode_15_intenet_cafe/

# Outline

1. This problem
  - Broadcasting to many goroutines
  - A new way to use `select`
2. Example code

# Extended Screencast

The code in this screencast has a bug - it may exit before all of the tourists are done using their computer. The [extended screencast](#TODO) that accompanies this one shows how to fix it.

[Get the Extended Screencast](#TODO)
