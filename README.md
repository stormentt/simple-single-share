# Overview
simple-single-share is a utility for sharing data exactly once, simply

# Config
Example `.simple-single-share.yml`
```yaml
uploads:
    api_key: 8OM3gRHVu1ogm1k4nE2IUKK4PhHa9gTv
```

# Usage
`simple-single-share server` will listen on port 8080
The client isn't written yet so you will need to upload with curl

```sh
curl -XPOST --header "Content-Type: text/plain" --header "API-Key: 8OM3gRHVu1ogm1k4nE2IUKK4PhHa9gTv" --data "Hello World" http://localhost:8080/share
```

This will return a uuid, which can then be retrieved at `http://localhost:8080/share/:uuid`

One the data is retrieved it is deleted from the server's memory.

# FAQ
## Password protected shares?
No, that would make this complex-single-share instead of simple-single-share

## Expire shares after x days?
No, that would make this complex-single-share instead of simple-single-share

## Built in end to end encryption?
No, that would make this complex-single-share instead of simple-single-share

## Retrieve the data more than once?
No, that would make this simple-many-share instead of simple-single-share

## Can't someone just guess my uuid?
yeah sure

## If too much data is uploaded won't this OOM?
probably

## I want transport encryption
put a reverse proxy in front of it
