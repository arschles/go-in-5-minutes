# Go In 5 Minutes

This folder has all the code to build the [Go in 5 Minutes website](https://www.goin5minutes.com). If you're looking for the Go In 5 Minutes example code and outlines, please see the [root directory](/).

# How the site works

The complete site is composed of:

- Static HTML & CSS generated by [Hugo](https://gohugo.io)
- Javascript generated by [Elm](https://elm-lang.org)
- An API server built with [Buffalo](https://gobuffalo.io)

The static site is hosted by [Netlify](https://netlify.com). When it loads, the Javascript (compiled by the Elm toolchain) begins "hydrating" the page with dynamic content. In other words, it contacts the API server for HTML snippets to insert into the page. Dynamic content includes the screencast summary list (at //screencasts on the site), screencast details (at /screencasts/$SCREENCAST_ID), and so on.

>Netlify proxies requests from the static site's JS to the API server, to avoid the need for CORS requests

## Single page app?

I want to note that the site is not a single page app at the moment. Navigations between pages like the index, the screencasts list, and so on are basic browser reloads (I might use [Turbolinks](https://https://github.com/turbolinks/turbolinks) to speed those up in the future). When any single page loads, the Javascript (reminder: compiled from Elm) takes over and hydrates that page.

# Development

Setting up a local development environment is kinda clunky at the moment. You'll end up needing 4 terminal windows open. Here are the programs of note that you'll need to have installed before you start:

- Hugo (to build the site)
- Caddy (to run the proxy server). You can run `make install-caddy` from inside the `site/` directory to install this
- Buffalo (to run the API server)
- Elm (to compile the Javascript)

When you have those, run the following commands from inside the `site/` directory, each in a separate terminal window:

- `make proxy`
- `make hugo`
- `make api`
- `make elm-compile`

The first three commands will run forever and watch for file changes. You'll need to re-run the last one each time you change any Elm files. I haven't gotten to making that one watch for changes yet.

# Old: create a New Screencast Page

You'll need [Hugo](https://gohugo.io/) installed, then run this:

run `hugo new screencast/episode_x_some_description.md`