*Do Not Use In Production. This was my first attempt at playing with Go and wasn't intended for real world use.*

GKV
===

A simple REST API driven, persistent key value store written in Go, that leverages leveldb.

## To build GKV:
    git clone https://github.com/shaunhess/GKV.git
    go get github.com/codegangsta/martini
    go get github.com/codegangsta/martini-contrib/auth
	go get github.com/syndtr/goleveldb/leveldb
    go build github.com\shaunhess\GKV
    go install github.com\shaunhess\GKV

## To execute:
    gkv

## Add a Key/Value pair:
Using curl: The -k option is required if you use a self-signed certificate. The -u option specifies the user:password, which in our case is simply token: (empty password). The -i option prints the whole response, including headers.

    curl -i -k -u token: -X POST --data "key=Author&value=Shaun Hess" "https://localhost:8001/gkv"

## Get a Key/Value pair:
Responses can be requested in JSON, XML or plain text. Response format will be determined based on the endpoint’s extension (.json, .xml or .text, defaulting to JSON).

	curl -i -k -u token: "https://localhost:8001/gkv.json?key=Author"
	curl -i -k -u token: "https://localhost:8001/gkv.xml?key=Author"
	curl -i -k -u token: "https://localhost:8001/gkv.text?key=Author"
	curl -i -k -u token: "https://localhost:8001/gkv?key=Author"
	
## Delete content:
    curl -i -k -u token: -X DELETE "https://localhost:8001/gkv?key=Author"
