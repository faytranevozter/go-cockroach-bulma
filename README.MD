# go-cockroach-bulma
Nothing here!  
This app following by [this](https://gowebexamples.com/) website.

## Requirement
- golang
- cockroachdb

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

#### Docker
1. Pull Cockroach from docker hub  
`docker pull cockroachdb/cockroach:v19.1.0`  

### Install Dependency package
- Gorilla/mux  
`go get -u github.com/gorilla/mux`
- Go pq driver  
`go get -u github.com/lib/pq`  

## Running App
1. Clone this repo  
`git clone https://github.com/faytranevozter/go-cockroach-bulma.git`
2. Start cockroach db  
for Linux  
`cockroach start --insecure`  
for Docker  
`docker run -d -p 26257:26257 -p 8080:8080 cockroachdb/cockroach:v19.1.0 start --insecure`
3. Create Database. (if you've done this before, then ignore this step)  
for Linux  
`cockroach sql --insecure --execute="CREATE DATABASE IF NOT EXISTS tugas_cockroach;"`  
for Docker  
`docker exec roach ./cockroach sql --insecure --execute="CREATE DATABASE IF NOT EXISTS tugas_cockroach;"`
3. Import sql from cockroach sql. (if you've done this before, then ignore this step)  
for Linux  
`cockroach sql --insecure --database=tugas_cockroach < db.sql`  
for Docker  
`docker exec -it container_name ./cockroach sql --insecure --database=tugas_cockroach < db.sql`  
and then import sql from [this](https://raw.githubusercontent.com/faytranevozter/go-cockroach-bulma/master/db.sql)
4. Run `go run *.go`
5. Done. See (http://localhost:8888) to see the app.

## Deploy App
[See here](https://github.com/faytranevozter/go-cockroach-bulma/blob/master/DOCKERIZING.MD)  

# Contributor
[@faytranevozter](https://github.com/faytranevozter/) M. Fahrur Rifai 165410011  
[@indrahad](https://github.com/indrahad/) Indra Purnama H. 165410005  
[@riskme](https://github.com/riskme/) Rismiyati 165410108  
[@lailyr](https://github.com/lailyr/) Laily Rohmawati 165410174  
[@okayeah](https://github.com/okayeah/) Oky Riyanto 165410231  