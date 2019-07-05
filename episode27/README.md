# Using Go Modules For Your Dependencies

Go in 5 Minutes, episode 27.

In this screencast, we're going to talk about [Go Modules](https://github.com/golang/go/wiki/Modules) to manage our dependencies.

Modules are a brand new dependency management system in Go, and they're here to stay. The modules technology is officially part of the `go` toolchain starting with version 1.11, and they're really nice. Modules have some features that are familiar to previous dependency management systems, like these:

- Define your dependencies in a file that you can check in with your code
- Keep track of all the dependencies -- including dependencies-of-your-dependencies (AKA transitive dependencies) -- in another file that you can check in with your code
- Keep checksums of all your dependencies (including transitive dependencies), so that things don't change out from under you

But they add some exciting new things too!

- Fetch all your dependencies without learning any other tools
- Delete your `vendor/` directory if you want
- Work outside the `GOPATH`
- Fetch your dependency code from somewhere _other than GitHub, GitLab, and other version control systems!_
  - More on this in a future episode ...

We're going to show a little bit of how modules work and how you can use them in your code.

Check out the screencast for more!

# How To Build This Code

Because Go Modules is built into the `go` toolchain, building is pretty easy - you don't even need a Makefile or a `GOPATH` set up!

First, make sure you either don't have this `episode27` directory inside your `GOPATH`, _or_ that you have `GO111MODULE` set to `on` like this:

```console
$ export GO111MODULE=on
```

And then, you should be ready to go!

```console
$ go build -o heygophers .
$ ./heygophers
```

Once you run `./heygophers`, your web server should start up on port 8080, and you can open up the app at `http://localhost:8080` in your browser.

Enjoy!

# Show Notes

- Just For Func _Intro to Go Modules and SemVer_: https://www.youtube.com/watch?v=aeF3l-zmPsY
- Just For Func _Migrating Go Modules to v2+_: https://www.youtube.com/watch?v=H_4eRD8aegk
- The Go modules wiki: https://github.com/golang/go/wiki/Modules
