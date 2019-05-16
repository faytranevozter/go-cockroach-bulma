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

### CockroachDB [Official Instruction](https://www.cockroachlabs.com/docs/stable/install-cockroachdb.html)
#### Windows
Follow [this](https://www.cockroachlabs.com/docs/stable/install-cockroachdb-windows.html) instruction
#### Linux
1. Download and extract files  
`wget -qO- https://binaries.cockroachdb.com/cockroach-v19.1.0.linux-amd64.tgz | tar  xvz`
2. Copy into `PATH`  
`cp -i cockroach-v19.1.0.linux-amd64/cockroach /usr/local/bin`

### Install Dependency package
- Gorilla/mux  
`go get -u github.com/gorilla/mux`
- Go pq driver  
`go get -u github.com/lib/pq`  

## Running App
1. Clone this repo  
`git clone https://github.com/faytranevozter/go-cockroach-bulma.git`
2. Start cockroach db  
`cockroach start --insecure`
3. Import sql from cockroach sql. (if you've done this before, then ignore this step)  
`cockroach sql --insecure`  
and then import sql from [this](https://raw.githubusercontent.com/faytranevozter/go-cockroach-bulma/master/db.sql)
4. Run `go run web.go`
5. Done. See (http://localhost:8888) to see the app.
