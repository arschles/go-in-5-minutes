# Building Login with GitHub into your Buffalo App

Go in 5 Minutes, episode 29.

Welcome back to the Buffalo series! The last Buffalo episode we did was [episode 25](https://www.goin5minutes.com/screencast/episode_25_buffalo_templating_with_plush/), where we talked about templating with [templating with Plush](https://gobuffalo.io/en/docs/templating/).

Today, we're going to continue talking about frontend technologies and look at how to let your users log into your webapp with GitHub

We're going to be using the [Goth library](https://gobuffalo.io/en/docs/goth) to do this, so I'll run through what that is, and then we'll look at the code for how to use it, and then we'll see it in action!

Check out the screencast for more!

# How To Set Up Goth with [`buffalo-goth`](https://github.com/gobuffalo/buffalo-goth)

With our Buffalo app created (I ran `buffalo new episode29` to create this very app!), we can use plugins to customize it. In this episode, we're going to use the `buffalo-goth` plugin to write most of the code for us to implement GitHub auth.

>These instructions are adapted from https://gobuffalo.io/en/docs/goth

First, get the `buffalo-goth` plugin for Buffalo. You'll only need to do this once, then you can re-use the plugin for all your Buffalo projects:

```console
$ go get github.com/gobuffalo/buffalo-goth
```

>The `buffalo-goth` executable is going to be put into `$GOPATH/bin`, so make sure that directory is in your `PATH`.

Next, generate the code for GitHub login:

```console
$ buffalo generate goth github
```

>You can pass more auth "providers" to the `buffalo generate goth` command. See [here](https://github.com/markbates/goth#supported-providers) for a complete list of providers.

# Show Notes

- [Buffalo docs on Goth](https://gobuffalo.io/en/docs/goth)
- [`buffalo-goth` code generator on GitHub](https://github.com/gobuffalo/buffalo-goth)
- [The Goth Library on GitHub](https://github.com/markbates/goth)
