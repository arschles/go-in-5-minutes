# Go In 5 Minutes

This folder has part of the code to build the https://www.goin5minutes.com website. If you're looking for the Go In 5 Minutes example code and outlines, please see the [root directory](https://github.com/arschles/go-in-5-minutes).

# How this is organized

Static site = the "skeleton"

API server (in ../api) that serves HTML snippets for dynamic content

[Elm](https://elm-lang.org) to fetch snippets and insert them into page

- Elm files are in ./elm, and the compiled JS gets spit out to static/js/bundle.js
- [Hugo pipes](https://gohugo.io/hugo-pipes/fingerprint/) fingerprints it (Netlify already minifies and fingerprints, so consider just using that instead)
- The JS is just for "hydrating" dynamic content
- This is not a SPA, it still have distinct page loads for each screencast etc...
  - Page loads are managed with [Turbolinks](https://github.com/turbolinks/turbolinks)

# TODO

Add [this script](https://guide.elm-lang.org/optimization/asset_size.html#scripts) for elm builds

# Create a New Screencast Page

You'll need [Hugo](https://gohugo.io/) installed, then run this:

run `hugo new screencast/episode_x_some_description.md`
