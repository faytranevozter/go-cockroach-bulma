# go-cockroach-bulma
Nothing here!  
This app following by [this](https://gowebexamples.com/) website.

## Requirement
- golang
- cockroachdb (not yet implemented)

## Installation
### Golang [Official Instruction](https://golang.org/doc/install)
[Download](https://golang.org/dl/)
#### Windows
[Download for windows](https://dl.google.com/go/go1.12.5.windows-amd64.msi) and just install
#### Linux
[Download for Linux](https://dl.google.com/go/go1.12.5.linux-amd64.tar.gz)  
Extract file to `/usr/local`

`tar -C /usr/local -xzf go$VERSION.$OS-$ARCH.tar.gz`  

and then  

`export PATH=$PATH:/usr/local/go/bin`

### Install Dependency package
- Gorilla/mux `go get -u github.com/gorilla/mux`

## Running App
1. Clone this repo  
2. Run `go run web.go`
3. Done. See (http://localhost:8888) to see the app.