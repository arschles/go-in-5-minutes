# Using Go Modules and [Athens](https://docs.gomods.io)

Go in 5 Minutes, episode 25.

We're going to take a break from our series on the Buffalo Framework and talk today about [Go Modules](https://github.com/golang/go/wiki/Modules) and the [Athens module repository](https://docs.gomods.io) to manage our dependencies.

Modules are a brand new way to manage our dependencies, and they're officially a part of the `go` toolchain that you're used to, starting with Go version 1.11. Modules let you:

- Define your dependencies in a file that you can check in with your code
- Fetch all your dependencies without learning any other tools
- Delete your `vendor/` directory if you want
- Work outside the `GOPATH`
- ... Did I mention you can work outside the `GOPATH`??? I love this part

We're going to show a little bit of how modules work, talk about how the Athens server works, and show it in action by building [Hugo](https://gohugo.io) with modules and Athens.

Check out the screencast for more!

# Show Notes

- Just For Func _Intro to Go Modules and SemVer_: https://www.youtube.com/watch?v=aeF3l-zmPsY
- Just For Func _Migrating Go Modules to v2+_: https://www.youtube.com/watch?v=H_4eRD8aegk
- The Go modules wiki: https://github.com/golang/go/wiki/Modules
- The Athens docs page: https://docs.gomods.io
- The Athens GitHub repository: https://github.com/gomods/athens

---

If you enjoy these screencasts, please consider 
[becoming a backer](https://www.patreon.com/goin5minutes)
and supporting this project. Cool stickers and more await you if you do!
