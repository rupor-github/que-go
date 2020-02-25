# que-go

----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------

Original project has not been updated for many years. During this period "ruby-que"" graduated to a different model (schema v4) and "pgx" driver matured to v4.
This fork (at commit b198c7caf054cfc10e9da6b34f33364d80464425) for now leaves ruby schema v3 compatibility but incorporates changes necessary to use latest pgx branch. It also
includes original PR#20 (Stop worker when Shutdown() is called, even when jobs are available).

----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------

[![GoDoc](https://godoc.org/github.com/bgentry/que-go?status.svg)][godoc]

Que-go is a fully interoperable Golang port of [Chris Hanks][chanks]' [Ruby Que
queuing library][ruby-que] for PostgreSQL. Que uses PostgreSQL's advisory locks
for speed and reliability.

Because que-go is an interoperable port of Que, you can enqueue jobs in Ruby
(i.e. from a Rails app) and write your workers in Go. Or if you have a limited
set of jobs that you want to write in Go, you can leave most of your workers in
Ruby and just add a few Go workers on a different queue name. Or you can just
write everything in Go :)

## pgx PostgreSQL driver

This package uses the [pgx][pgx] Go PostgreSQL driver rather than the more
popular [pq][pq]. Because Que uses session-level advisory locks, we have to hold
the same connection throughout the process of getting a job, working it,
deleting it, and removing the lock.

Pq and the built-in database/sql interfaces do not offer this functionality, so
we'd have to implement our own connection pool. Fortunately, pgx already has a
perfectly usable one built for us. Even better, it offers better performance
than pq due largely to its use of binary encoding.

Please see the [godocs][godoc] for more info and examples.

[godoc]: https://godoc.org/github.com/bgentry/que-go
[chanks]: https://github.com/chanks
[ruby-que]: https://github.com/chanks/que
[pgx]: https://github.com/jackc/pgx
[pq]: https://github.com/lib/pq
