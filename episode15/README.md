# Concurrency Series: Internet Cafe

[![Watch The Screencast](https://www.goin5minutes.com/img/watch-screencast.svg)](https://www.goin5minutes.com/screencast/episode_15_internet_cafe/)

Go in 5 Minutes, episode 15.

This screencast overviews the "Internet Cafe" problem on [Trivial Concurrency Exercises for the Confused Newbie Gopher](http://whipperstacker.com/2015/10/05/3-trivial-concurrency-exercises-for-the-confused-newbie-gopher/), and provides a solution.

We're finishing up the concurrency series - see https://github.com/arschles/go-in-5-minutes/issues/14 for what we've covered.

Screencast video:
https://www.goin5minutes.com/screencast/episode_15_internet_cafe/

# Other Notes

- Extended screencast: https://gum.co/gifm-x-15
- __Reminder: subscribe to the newsletter at https://www.goin5minutes.com/subscribe/__

# Extended Screencast

The code in this screencast has a bug - it may exit before all of the tourists are done using their computer. The [extended screencast](https://gum.co/gifm-x-15) that accompanies this one shows how to fix it.

Get the extended screencast at https://gum.co/gifm-x-15.

# Outline

1. This problem
  - Broadcasting to many goroutines
  - A new way to use `select`
2. Example code
