+++
date = "2016-08-28T16:01:02-07:00"
title = "On Orthogonality"
type = "blog"
teaser = "Orthogonality is a powerful concept in Go. I'll explain why, and how to achieve it in your programs."
author = "Aaron Schlesinger"
+++

First of all, welcome everyone to the first post in the Go In 5 Minutes blog series.

Today I'm going to talk about __orthogonality__. I wrote a [tweet storm](https://twitter.com/goin5minutes/status/769325705226227713) (go [follow Go In 5 Minutes](https://twitter.com/goin5minutes) if you haven't already) on the concept a few days ago. In it, I explained that this is an easy concept to explain in theory, but it's much harder to explain how to put it into practice. I'm going to try to do the latter today.

# Why Care?

First of all, let me explain why orthogonality is worth understanding and achieving in your programs. Look up the term [in Wikipedia](https://en.wikipedia.org/wiki/Orthogonality_(programming)), and you'll find this:

`orthogonality in a programming language means that a relatively small set of primitive constructs can be combined in a relatively small number of ways to build the control and data structures of the language`

As Go programmers, we generally split up our programs into discrete `package`s, `type`s and `func`s. I'll group all of these terms together by saying `components` herefter.

When we design our programs using components that are orthogonal to each other, we implicitly end up with fewer components that have strict rules for how they can be used together. And those rules are usually _enforced by the compiler_. That last sentence is important; any time the compiler enforces our rules for us, we reduce complexity and our testing and documentation burden.

Simply put, making our components orthogonal to each other leads to less code, fewer edge cases, and a more understandable program. And in many cases our programs are _more_ powerful than they would otherwise be.

# "Doing" Orthogonality

There are some hard-and-fast rules for determining whether a program's components are orthogonal to each other, but there's also some art to it. I put a few guidelines in my [tweet storm](https://twitter.com/goin5minutes/status/769325705226227713) and I'm going to expand on the first two of them here: interfaces and concurrency.

## Interfaces

If you have interfaces with lots of functions, it'll be harder to use them in your programs.

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

Nothing wrong with it on the surface. We can use it inside a function just fine:

```go
func DoSomething(th HugeThing, i int, s string) int {
  return th.Foo(i, s) + int(th.Foo(i, float64(i)))
}
```

But any caller of `DoSomething` needs to have an entire implementation of `HugeThing` available to them.

This pattern isn't very orthogonal because you need a lot of code to implement `HugeThing`, but `DoSomething` is only using 2 of the 6 functions in it. Another warning sign is that any implementation of `HugeThing` (and `HugeThing`) itself needs a lot of documentation.

The solution? Split `HugeThing` up into appropriate "chunks" of functionality:

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

And compose them if you still need a `HugeThing`:

```go
type HugeThing interface {
  FooBarer
  BazQuxer
  OofRaber
}
```

If you have code that already implements `HugeThing`, nothing will break and you can even change the parameter in `DoSomething` to take in the simpler type. The implementation stays exactly the same:

```go
func DoSomething(th FooBarer, i int, s string) int {
  return th.Foo(i, s) + int(th.Foo(i, float64(i)))
}
```

## Concurrency

If you have Goroutines that communicate by writing to shared variables, you're probably doing it wrong. Yes, you might have a really specific reason for using shared memory (and locks, etc...), and that's ok; keep doing it. But even you might get something out of this section!

I'll be basing much of the content herein on a [golang.org blog post on a very similar topic](https://blog.golang.org/share-memory-by-communicating).

If two or more Goroutines write to the same shared memory, they are now a single component. You can't just use one without at least considering the other. And they'll likely have to use a lock or some other synchronization mechanism (unless you _really, really, really_ know what you're doing!), so you're gonna have to consider that too. In other words, you're not dealing with a simple construct anymore, and it definitely can't be combined with other simple constructs in a simple way.

Sometimes using shared memory and a `sync.Mutex` (or similar) is the right thing to do for your use-case. If so, abstract it away behind a `struct`, hide the shared state the the mutex, and document it well. In essence, make it look like a good old class in Java so it can be composed with other components as easily as possible.

Enough about shared memory, though. I'm here to talk about channels!

You should make your goroutines communicate with each other by sending messages over channels (e.g. share memory by communicating). Do this to make your code less complex and simpler to compose.

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

Option #1 requires an implementation of a shared queue and a pool of consumers on the other side of the queue. And, the handler will need a reference to the shared queue, which will almost certainly be a non-trivial implementation. Option #2 simply requires the handler have access to the channel. The latter option makes for a simpler and easier to compose system.

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

As I indicated in the comment above `processor`, adding "consumers" to the channel is as trivial as calling `go processor()` again. And, the consumer of the channel can be anything - an RPC call, a no-op, etc...

Note that there are elements of [DRY](https://en.wikipedia.org/wiki/Don%27t_repeat_yourself) here, because the `chan` is already a synchronized queue, implemented in the Go runtime.

# Wrapping Up

In many cases, you won't make or break your program if you don't design it with orthogonal pieces. And in almost all cases, you can refactor tightly-coupled, complex systems to be more loosely-coupled with orthogonal pieces. But regardless, orthogonal components will produce a simpler, easier to build and less complex program.


*Credit [Eric Raymond's 17 Unix Rules](https://en.wikipedia.org/wiki/Unix_philosophy#Eric_Raymond.E2.80.99s_17_Unix_Rules) to inspire much of this post, in addition to all other resources linked*
