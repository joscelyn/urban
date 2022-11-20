# Joscelyn Gainié's technical backend test result @ UrbanCampus - GRPC News 

This is a client and a server communicating using gRPC. The server is a Hacker News proxy. The client can connect to the server to get the latest news or informations on a user.

## Installation

```
go mod download
go build -o build/urban ./urban
```

## Usage

To run the server just run :
```
./build/urban server
```

To get top 10 hacker news posts :
```
./build/urban client --list
```

And to get user information :
```
./build/urban client --whois fra
```

Note: you can also use docker with the Dockerfile

## Todo

Since I'm taking the train tonight, I didn't have time to start/polish these last elements:
- The bonus: create a redis docker service to cache rn data
- Export some variables (for example the grpc server port)
- Make cleaner functions and a better file structure
- Maybe add more comments and better test coverage

## Tools used

- A client for Hacker News' "API" https://github.com/peterhellberg/hn
- A lib for GRPC client/server boilerplate code https://github.com/lileio/lile/
