# Mercury

## How to run tests
In order to run tests in Mercury, first spin-up the docker-compose environment:
```
docker-compose up --build && docker-compose rm
```
Then run tests:
```
    go test -v ./...
```