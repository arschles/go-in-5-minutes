# The [Buffalo](https://gobuffalo.io) Series #4 - Templates

Go in 5 Minutes, episode 25.

We've been looking at the [Buffalo framework](https://gobuffalo.io) lately, and
we're gonna keep on going today with [templating](https://gobuffalo.io/en/docs/templating).

If you're building a webapp, templating is crucial for you to be able to serve dynamic content to your users. For example, if you need to fetch the user's name and information from your database, you'll need to somehow put that data into HTML that you can serve to their browser.

Templates are how you do it, and Buffalo has a great, easy to understand template engine to help you out!

Here's what we'll learn today:

- Introduction to templating
- The [plush](https://github.com/gobuffalo/plush) template engine
- How to use plush from your Buffalo app

And like always I'll be showing live code.

Check out the screencast for more!

# Show Notes

- Buffalo templates: https://gobuffalo.io/en/docs/templating
- Go standard library [`html/template`](https://godoc.org/html/template) and [`text/template`](https://godoc.org/text/template)
- [CSRF](https://en.wikipedia.org/wiki/Cross-site_request_forgery)
    - [Prevention cheat sheet](https://www.owasp.org/index.php/Cross-Site_Request_Forgery_(CSRF)_Prevention_Cheat_Sheet))
    - See [here](https://stackoverflow.com/questions/941594/understanding-the-rails-authenticity-token) for a great explanation on how `authenticity_token` fits in 

---

If you enjoy these screencasts, please consider 
[becoming a backer](https://www.patreon.com/goin5minutes)
and supporting this project. Cool stickers and more await you if you do!
