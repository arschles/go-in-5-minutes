# Using Athens to Serve Your Go Modules

Go in 5 Minutes, episode 28.

In [episode 27](https://www.goin5minutes.com/screencast/episode_27_intro_to_modules/), we talked about how to use [Go Modules](https://github.com/golang/go/wiki/Modules) to manage dependencies, but there's more to the story!

Go modules come with module servers that you can download your dependencies from too! The servers themselves are pretty cool, but here's why you should consider using one:

- Downloading dependencies will almost always be faster (sometimes up to 4x faster)
  - Think about how much faster CI runs can be ;)
- You can avoid broken builds when someone deletes a commit, tag or repository

I'm going to explain how module servers work, introduce Athens, and show how to use it in action.

Check out the screencast for more!

# Show Notes

- The Athens Project: https://docs.gomods.io
- The Go team's module proxy: https://proxy.golang.org
- The JFrog public module proxy: https://gocenter.io
