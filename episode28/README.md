# Using Athens to Serve Your Go Modules

Go in 5 Minutes, episode 28.

In [episode 27](https://www.goin5minutes.com/screencast/episode_27_intro_to_modules/), we talked about how to use [Go Modules](https://github.com/golang/go/wiki/Modules) to manage dependencies, but there's more to the story!

Go modules come with module servers that you can download your dependencies from too! The servers themselves are pretty cool, but here's why you should consider using one:

- Downloading dependencies will almost always be faster (sometimes up to 4x faster)
  - Think about how much faster CI runs can be ;)
- You can avoid broken builds when someone deletes a commit, tag or repository

I'm going to explain how module servers work, introduce Athens, and show how to use it in action.

Check out the screencast for more!

# How To Run Athens

We're going to run an Athens and build a little server with it as our module proxy. We'll actually do _two_ builds here. The first will be to run Athens when connected to the internet, and the second will be to do it when disconnected from the internet - to show how well Athens works in an isolated environment (like inside a firewall).

These instructions are for Linux/Mac OS X systems.

## Build #1: Build With Athens and an Upstream VCS

Athens maintains its own database of modules. When you do a `go get` and request a module from Athens, it checks its database and sends the module back to you if it's there. If not, then Athens does this:

1. Fetch the module from version control
1. Store it in the database
1. Send it back to you

Since we're starting Athens with nothing in its storage, every dependency we request in this build will make it download modules from version control.

>By the way, you can also configure Athens to download from module mirrors like [proxy.golang.org](https://proxy.golang.org) or [gocenter.io](https://gocenter.io), instead of version control hosts!

### Run The Server!

We try hard to make it easy to run your own Athens. See [here](https://docs.gomods.io/install) for instructions for running the server a few different ways. Today, we're going to use [Docker](https://www.docker.com/) to run ours.

First, run this to start Athens up:

```console
$ docker run -p 3000:3000 -e GO_ENV=development -e ATHENS_GO_GET_WORKERS=5 gomods/athens:v0.5.0
```

And then to set your `GOPROXY` environment variable to tell modules to use the local server:

```console
$ export GOPROXY=http://localhost:3000
```

Also, the Go tool keeps a read-only on-disk cache of every module version you've downloaded for any build. To make it read-only, it stores each file in the cache with `-r--r--r--` permissions. Since that's the case, you need to use `sudo` to clear the cache.

```console
$ sudo rm -rf $(go env GOPATH)/pkg/mod
```

And then build and run the server!

```console
$ go run .
```

## Second Way: Use Your Athens While Offline :scream:

Did I mention that Athens stores your dependencies in storage? Well, it actually stores your dependencies in storage _forever_! That means that you can build your code without access to the internet. And it's faster. Let's do a build and see.

First, make sure not to shut down the Athens server from last time - its storage is inside the Docker container!

Next, clear out your cache again:

```console
$ sudo rm -rf $(go env GOPATH)/pkg/mod
```

And then, **shut down your internet connection** :see_no_evil:.

And finally, do the build & run again!

```console
$ go run .
```

And you're done!

# Show Notes

- [The module download protocol](https://docs.gomods.io/intro/protocol/)
- [The Athens Project](https://docs.gomods.io)
- [Joining the Gophers Slack](https://invite.slack.golangbridge.org/)
  - Come say hi in the `#athens` channel!
- [List of all available module proxies](https://github.com/golang/go/wiki/Modules#are-there-always-on-module-repositories-and-enterprise-proxies)
- [The Go team's module proxy](https://proxy.golang.org)
- [The JFrog public module proxy](https://gocenter.io)
