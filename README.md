# Brandless Take Home

```
git clone https://github.com/chriskaukis/brandless
cd brandless
```

Run with Docker (not a Docker guru yet... maybe this isn't the right way, but it's how I did it):
```
docker build --tag=brandless .
docker run -p 8080:8080 brandless
```
(Use any tag name you want.)

or not
```
go build
./brandless
```

```
http://localhost:8080
```

The port can be changed by setting the environment variable `PORT`

E.g.
`PORT=3000 ./brandless`
