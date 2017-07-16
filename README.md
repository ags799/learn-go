# learn-go

For learning [Go](https://golang.org/).

## Development

Build with

    make

Run with

    make run

Test with

    make test

Check style with

    make style

Integration test by first starting the Docker containers

    make docker-run

then running integration tests

    make integration-test

then stopping the Docker containers

    make docker-stop

We use [gvt](https://github.com/FiloSottile/gvt) for handling dependencies.
Add a dependency with

    gvt fetch path/to/dependency

Ensure that `vendor/` is checked in.
