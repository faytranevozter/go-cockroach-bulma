# DOCKERIZING
### Building docker image (once)
#### Creating directory and go into it
`mkdir asudahlah && cd asudahlah`
#### Cloning web app from github repository
`git clone https://github.com/faytranevozter/go-cockroach-bulma.git .` 
#### Build the image
`docker build -t faytranevozter/gocockroachbulma:v1.0-alpine .`  

### Cockroach-ing
#### Pull Cockroach DB (optional)
`docker pull cockroachdb/cockroach`
#### Start cockroachDB
`docker run -d --name=roach -p 26257:26257 -p 8080:8080 cockroachdb/cockroach start --insecure`
#### Save cockroachDB IP Address into variable
This step used later to set `ROACH_HOST` environment  
`ROACH_IP_ADDRESS=$(docker inspect -f '{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' roach)`  

### Create and Import Database
#### Create Database
`docker exec roach ./cockroach sql --insecure --execute="CREATE DATABASE IF NOT EXISTS tugas_cockroach;"`
#### Download SQL
Download using `wget`  
`wget https://raw.githubusercontent.com/faytranevozter/go-cockroach-bulma/master/db.sql`  
or using `curl` instead  
`curl https://raw.githubusercontent.com/faytranevozter/go-cockroach-bulma/master/db.sql -O`
#### Import SQL
`docker exec -i roach ./cockroach sql --database=tugas_cockroach < db.sql --insecure`  

### Running web app
`docker run -d --name=webapp -p 9000:8888 --env ROACH_HOST=$ROACH_IP_ADDRESS faytranevozter/gocockroachbulma:v1.0-alpine`
