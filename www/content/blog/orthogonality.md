+++
date = "2016-08-28T16:01:02-07:00"
title = "On Orthogonality"
type = "blog"
teaser = "Orthogonality is a powerful concept in Go. I'll explain why, and how to achieve it in your programs."
author = "Aaron Schlesinger"
+++

First of all, welcome everyone to the first post in the Go In 5 Minutes blog series.

Today I'm going to talk about __orthogonality__. I wrote a [tweet storm](https://twitter.com/goin5minutes/status/769325705226227713) (shameless plug: [follow Go In 5 Minutes](https://twitter.com/goin5minutes) if you haven't already!!!) on the concept a few days ago. In it, I explained that this is an easy concept to explain in theory, but it's much harder to explain how to put it into practice. I'm going to try to do the latter today.

# Why Care?

First of all, let me explain why orthogonality is worth understanding and achieving in your programs. Look up the term [in Wikipedia](https://en.wikipedia.org/wiki/Orthogonality_(programming)), and you'll find this:

`orthogonality in a programming language means that a relatively small set of primitive constructs can be combined in a relatively small number of ways to buil d the control and data structures of the language`

As Go programmers, we generally split up our programs into discrete `package`s, `type`s and `func`s. I'll group all of these terms together by saying `components` herefter.

When we make design our program using components that are orthogonal to each other, we implicitly end up with fewer components that have strict rules for how they can be used together. And those rules are generally _enforced by the compiler_. Before I continue, I'd encourage you to re-read that last sentence. Any time we can get the compiler to enforce something in our programs, we reduce the testing, documentation, and complexity burden we put on ourselves.

Simply put, making our components orthogonal to each other leads to less code, fewer edge cases, and a more understandable program. And finally, in many cases our programs are _more_ powerful than they would otherwise be.

# "Doing" Orthogonality

A program made up of orthogonal components is itself orthogonal. There are some hard-and-fast rules for getting there, but there's also some art to it. I put some guidelines in my [tweet storm](https://twitter.com/goin5minutes/status/769325705226227713), and I'm going to expand on the first two of them here: interfaces and concurrency.

## Interfaces

If you have interfaces with lots of functions, it'll be harder to use them as parameters in your functions.

Check this interface out:

```go
type HugeThinger interface {
  Foo(int, string) int
  Bar(int, float64) float64
  Baz(string, string) string
  Qux(string, float64) string
  Oof(int, int) *int
  Rab(string, string) *string
}
```

Pretty good! Until we have to pass it into a `func` that doesn't need all of its functionality:

```go
func DoSomething(th HugeThing, i int, s string) int {
  return th.Foo(i, s) + int(th.Foo(i, float64(i)))
}
```

Now, anyone who needs to call `DoSomething` needs to have an entire implementation of `HugeThing` available to them. This pattern isn't very orthogonal because you need a lot of code (something complex) to implement `HugeThing` in order to combine it with anything else. Another warning sign is that any implementation of `HugeThing` (and `HugeThing`) itself needs a lot of documentation.

The solution is to split `HugeThing` up into appropriate "chunks" of functionality:

```go
type FooBarer interface {
  Foo(int, string) int
  Bar(int, float64) float64
}

type BazQuzer interface {
  Baz(string, string) string
  Qux(string, float64) string
}
type OofRaber interface {
  Oof(int, int) *int
  Rab(string, string) *string
}
```

And compose them if you still need `HugeThing`:

```go
type HugeThing interface {
  FooBarer
  BazQuxer
  OofRaber
}
```

If you have code that already implements `HugeThing`, nothing will break. And you can even change the parameter in `DoSomething` to take in a much smaller, simpler type. The implementation stays exactly the same:

```go
func DoSomething(th FooBarer, i int, s string) int {
  return th.Foo(i, s) + int(th.Foo(i, float64(i)))
}
```

## Concurrency

If you have Goroutines that communicate by writing to shared variables, you're probably doing it wrong. Yes, you might have a really specific reason for using shared memory (and locks, etc...), and that's ok; keep doing it. But even you might get something out of this section!

I'll be basing much of the content herein on a [golang.org blog post on a very similar topic](https://blog.golang.org/share-memory-by-communicating).

If two or more Goroutines write to the same shared memory, they are now a single component. You can't just use one without at least considering the other. And they'll likely have to use a lock or some other synchronization mechanism (unless you _really, really, really_ know what you're doing!), so you're gonna have to consider that too. In other words, you're not dealing with a simple construct anymore, and it definitely can't be combined with other simple constructs in a simple way.

Sometimes using shared memory and a `sync.Mutex`, etc... is the right way to design your program, and if you do, abstract it away behind a `struct` that holds non-exported fields. In essence, make it look like a good old class in Java. And document it well.

But I'm here to talk about channels. You should make your goroutines communicate with each other by sending messages over channels (e.g. share memory by communicating). Do this because you'll make your code less complex, and simpler to compose.

Consider a web server that needs to do some asynchronous work:

```go
func handlerToDoAsyncWork(w http.ResponseWriter, r *http.Request) {
  // schedule the work to be done asynchronously
  w.Write([]byte("stuff will happen soon!"))
}
```

You have two main options for scheduling the work:

1. Send it to a shared queue
2. Send it on a channel

Option #1 requires an implementation of a shared queue and a pool of consumers on the other side of the queue. And, the handler will need a reference to the shared queue so it can publish onto it. Option #2 simply requires the handler have access to the channel. The latter option makes for a simpler and easier to compose system.

```go
func handlerToDoAsyncWork(w http.ResponseWriter, r *http.Request) {
  // will unblock after the work has been scheduled
  asyncCh <- r.URL.Path
  w.Write([]byte("stuff will happen soon!"))
}

// run as many of these as you need with 'go processor()'
func processor() {
  for {
    // get the submission from the handler
    urlPath := <-asyncCh
    // do the work asynchronously
    go doSomethingWithURLPath(urlPath)
  }
}
```
