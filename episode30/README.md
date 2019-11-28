# Hooking your Buffalo Webapp to a Database

Go in 5 Minutes, episode 30.

Another episode in the [Buffalo](https://gobuffalo.io) series! [Last episode](https://gifm.dev/screencast/episode_29_buffalo_login_with_github_using_goth/), we talked about building a "login with GitHub" button into your Buffalo app.

Today, we're gonna travel from that part of your app up to the database layer. If you're building a website, you'll almost certainly need a database to store information about your users. Luckily, Buffalo has great support for some of the most popular SQL databases.

>This episode is tangentially related to [#168](https://github.com/arschles/go-in-5-minutes/issues/168)

We're going to use the [pop](https://github.com/gobuffalo/pop) library -- which has great Buffalo integration -- to do some basic database integration in our app.

I'll briefly introduce the library and how to use it, then we'll go right to the code and see it in action!

Check out the screencast for more!

## How To Run The App

This is a standard Buffalo app that requires a PostgresDB server to be running locally (on `127.0.0.1`). Here's what you need to do to run the app.

First, run the database. We've provided a [Docker Compose](https://docs.docker.com/compose/) configuration file, so the easiest way to do this is to download the `docker-compose` binary and ensure you have the [Docker](https://docs.docker.com/install/) server running. After you have those dependencies, run the database:

```console
$ docker-compose -p ep30 up -d dev-env
```

This starts the database in a container in the background. Next, you should be all set to run the app:

```console
$ buffalo dev
```

Now, you can access the app on `http://localhost:3000`.

When you're done, clean up the database with this command:

```console
$ docker-compose -p ep30 down
```

# Show Notes

- [Buffalo docs on database interaction](https://gobuffalo.io/en/docs/db/getting-started/)
- [The pop repository on GitHub](https://github.com/gobuffalo/pop)
