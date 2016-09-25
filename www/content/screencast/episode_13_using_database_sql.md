+++
date = "2016-05-15T14:45:27-07:00"
title = "Using the database/sql package"
type = "screencast"

+++

_Episode 13_

We briefly describe how to use the `database/sql` package and show example code on how to do [CRUD](https://en.wikipedia.org/wiki/Create,_read,_update_and_delete) actions against a SQL databases using the flexible [`database/sql`](https://godoc.org/database/sql) package.

<!--more-->

The screencast includes example code to show best practices in reading, writing and updating data to the database. Note that if the [standard library documentation](https://godoc.org/database/sql) is lacking for you, http://go-database-sql.org/ is another wonderful source of documentation.

_Note_: the examples in this screencast operate against an in-memory [SQLite](https://www.sqlite.org/) database, but we point out the line of code to change to operate against any other supported database.

<iframe
  class="ytplayer"
  type="text/html"
  width="640"
  height="390"
  src="https://www.youtube.com/embed/nros7z5z-7M?autoplay=0&origin=https://www.goin5minutes.com"
  frameborder="0"
></iframe>

Check out the example code [on Github](https://github.com/arschles/go-in-5-minutes/tree/master/episode13).
