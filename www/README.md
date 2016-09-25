# Go In 5 Minutes

This folder has code to build the https://www.goin5minutes.com website. If you're looking for
the Go In 5 Minutes example code and outlines, please see the [root directory](https://github.com/arschles/go-in-5-minutes).

# Create a New Screencast Page

run `hugo new screencast/episode_x_some_description.md`

if you don't have [Hugo](http://gohugo.io) installed, follow the [install instructions](http://gohugo.io/) (click the "install" button) or run it in a [Docker](https://www.docker.com/) container:

```console
docker run --rm -v $PWD:/pwd -w /pwd quay.io/arschles/hugo:latest hugo new screencast/episode_x_some_description.md
```
