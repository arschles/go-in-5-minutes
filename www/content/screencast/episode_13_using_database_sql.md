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

{{< screencast_bottom youtube_id="nros7z5z-7M" github_episode="epsode13">}}
