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

## Notes
* The jokes are short, so sometimes the markov chain doesn't have much affect. :(
* You forgot to add test coverage for some of the code.


## Acknowledgements
CSS/HTML was leveraged from https://github.com/jaridwarren/simpsons-quotes

I think of Homer when I think of dad(s). I liked this minimal simple design and was wanting some sort of starting template. So, in summary thank you Jarid Warren. I don't know you. You don't me, but I like your design. Nice and simple. <3
