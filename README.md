## Deployment architecture

**Solution.**

To make deployment independent of cloud providers, use containerization through docker images. That could be runnable on any server, as result we could switch between providers whenever we want without changing the deployment process.

**File locations.**

All docker files should locate in **${projectname}/deploy** directory.
For each service, at the project, write a separate docker file.

**Naming.**

According to docker files naming convention, it should have name of service before dot (ex.: signer.Dockerfile, projectname.Dockerfile).

If the project has several docker-compose files, these files should also have naming according to docker files naming convention (docker-compose.test.yml, docker-compose.local.yml).

**Deployment.**

For deployment use GitHub actions that trigger commands from Makefile. It will build docker images (with commit hash and latest), and it will be pushed to our docker registry. Images from docker registry will use on deployment server in docker-compose file.

**Rollback to particular version.**

On deployment part, create docker image with name that contains commit hash (docker-registry-address/service-name:commit-hash), as result we could rollback to particular version whenever we want.

**Installing.**

Golang is our backend language.

We are using version 1.17.4. You can download it from the official website [GoLang](https://go.dev/dl/), and install it according to the official [instructions](https://go.dev/doc/install.)

**Database.**

For our project we use a relational database PostgreSQL, version 12.11 which you can download by following the link from the official [website](https://www.postgresql.org/download/) or you can run your database in a Docker container.

**Docker.**

For isolated installation of databases and servers we need a Docker, version 20.10.16 or higher, you can download it at official [website](https://docs.docker.com/engine/install/)

```shell
docker run --name=db -e POSTGRES_PASSWORD='$YOUR_PASSWORD' -p $YOUR_PORTS -d --rm postgres
docker exec -it db createdb -U postgres ultimatedivision_test
```

**Run the main server locally.**

From the root of the project use this commands to create .env file with necessary variables:
```shell
cp deploy/local/.env.test deploy/local/.env
```

Need to feel variables on the path deploy/local/.env .

There is a makefile to run the project, you can run it with the command in root of the project:
```shell
make run_local
```
After this you can open console on localhost:8088 and admin panel on localhost:8087

**Access to logs.**

For access to logs, we use [Dozzle](https://dozzle.dev/).
It's running as a separate service in docker-compose. To create login & password - pass as environment variables to docker-compose and provide credentials to QA and Devs. So that they have easy and fast access to logs.`

**Metrics & graphs.**

To collect standards (like CPU, Memory usage) or custom metrics we use [Prometheus](https://prometheus.io/docs/introduction/overview/).

To make graphs we use [Grafana](https://grafana.com/docs/grafana/latest/introduction/) which uses metrics passed by Prometheus.


**Metric examples.**
>metrics/metrics.go
```go
type Metric struct {
   handler  http.Handler
   newUsers Counter
}
   
// NewUsersInc increment Counter newUsers.
func (metric *Metric) NewUsersInc() {
   metric.newUsers.Inc()
}

// NewMetric is a constructor for a Metric.
func NewMetric() *Metric {
    newUsers := prometheus.NewCounter(prometheus.CounterOpts{
        Name: "number_registrations",
        Help: "The total number of successful registrations.",
    })
    
    // Create a custom registry.
   registry := prometheus.NewRegistry()

   // Register using our custom registry gauge.
   registry.MustRegister(newUsers)
   return &Metric{
        // Expose metrics.
        handler:  promhttp.HandlerFor(registry, promhttp.HandlerOpts{Registry: registry}),
        newUsers: newUsers, 
    }
}
```
>console/consoleserver/controllers/auth.go
```go
// Register creates a new user account.
func (auth *Auth) Register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ctx := r.Context()

	var err error
	var request users.CreateUserFields

	if err = json.NewDecoder(r.Body).Decode(&request); err != nil {
		auth.serveError(w, http.StatusBadRequest, AuthError.Wrap(err))
		return
	}

	if !request.IsValid() {
		auth.serveError(w, http.StatusBadRequest, AuthError.New("did not fill in all the fields"))
		return
	}

	err = auth.userAuth.Register(ctx, request.Email, request.Password, request.NickName, request.FirstName, request.LastName, request.Wallet)
	if err != nil {
		switch {
		case userauth.ErrAddressAlreadyInUse.Has(err):
			auth.serveError(w, http.StatusBadRequest, userauth.ErrAddressAlreadyInUse.Wrap(err))
			return
		default:
			auth.log.Error("Unable to register new user", AuthError.Wrap(err))
			auth.serveError(w, http.StatusInternalServerError, AuthError.Wrap(err))
			return
		}
	}
	
	auth.metric.LoginsInc()
}
```
**Mini servers.**
In general, we use private mini-servers when we work with signing something with a private key to protect personal data from hackers.
List of these servers:
- Currency signer;
- NFT signer;

  Brief information about them:
- we don’t have any api in those servers;
- we have docker files/docker-compose files where these servers are already started as closed;
- we have private keys in config files, each server has its own file;
- each server has an endless cycle with an interval of work and performs its specific logic;
  You will find commands for local startup under the server description.
  Deployment instructions on the remote server according to the docker files in the `deploy` directory.
  
**Currency signer**
  The currency signer runs a infinite cycle with an interval of operation that monitors the records of currency waitlist in which there is no signature and if it finds them then generates a signature and sends a transaction to transfer money.
```shell
go run cmd/currencysigrer/main.go run --config=your_path
```
**NFT signer**
The nft signer runs an infinite cycle with an interval of operation that monitors the records of waitlist in which there is no signature and if it finds them then generates a signature and sends a transaction to mint nft.
```shell
go run cmd/nftsigrer/main.go run --config=your_path
```
>deploy/docker-compose.yml
```yaml
version: "3"

services:
  app: 
    container_name: ultimatedivision_app
    image: ${HOST_FOR_DOCKER_IMAGE}/ultimate_division${ENVIRONMENT}:latest
    ports:
      - "8087:8087" # Forward the exposed port 5000 on the container to port 8088 on the host machine (8088:5000).
      - "8088:8088"
    restart: unless-stopped
    volumes:
      - ${PROJECT_DATA_PATH}/ultimate_division:/app/data
      - ${PROJECT_CONFIGS_PATH}/ultimate_division:/config
      - ${PROJECT_ASSETS_PATH}:/assets
    depends_on:
      - ultimatedivision_db # This service depends on postgres. Start that first.
    networks:
      - fullstack

  nft_signer:
    container_name: ultimatedivision_nft_signer
    image: ${HOST_FOR_DOCKER_IMAGE}/ultimate_division_nft_signer${ENVIRONMENT}:latest
    restart: unless-stopped
    volumes:
      - ${PROJECT_DATA_PATH}/signer:/app/data
      - ${PROJECT_CONFIGS_PATH}/nft_signer:/config
    depends_on:
      - ultimatedivision_db # This service depends on postgres. Start that first.
    networks:
      - fullstack

  currency_signer:
    container_name: ultimatedivision_currency_signer
    image: ${HOST_FOR_DOCKER_IMAGE}/ultimate_division_currency_signer${ENVIRONMENT}:latest
    restart: unless-stopped
    volumes:
      - ${PROJECT_DATA_PATH}/currency_signer:/app/data
      - ${PROJECT_CONFIGS_PATH}/currency_signer:/config
    depends_on:
      - ultimatedivision_db # This service depends on postgres. Start that first.
    networks:
      - fullstack

  dozzle:
    container_name: ultimatedivision_dozzle
    image: amir20/dozzle:latest
    volumes:
      - /var/run/docker.sock:/var/run/docker.sock
    ports:
      - "9999:8080"
    networks:
      - fullstack
    depends_on:
      - app
    environment:
      - DOZZLE_NO_ANALYTICS=true
      - DOZZLE_USERNAME=${DOZZLE_USERNAME}
      - DOZZLE_PASSWORD=${DOZZLE_PASSWORD}
      - DOZZLE_KEY=true

  prometheus:
    image: prom/prometheus
    container_name: ultimatedivision_prometheus
    hostname: prometheus
    restart: always
    volumes:
      - ${PROJECT_CONFIGS_PATH}/prometheus/prometheus.yml:/etc/prometheus/prometheus.yml
    ports:
      - "9090:9090"
    networks:
      - fullstack

  grafana:
    image: grafana/grafana
    container_name: ultimatedivision_grafana
    hostname: grafana
    restart: always
    ports:
      - "3000:3000"
    volumes:
      - ${PROJECT_CONFIGS_PATH}/grafana/provisioning:/etc/grafana/provisioning
      - ${PROJECT_DATA_PATH}/grafana:/var/lib/grafana
    networks:
      - fullstack

  node_exporter:
    image: prom/node-exporter
    container_name: ultimatedivision_node_exporter
    hostname: node-exporter
    restart: always
    ports:
      - "9100:9100"
    networks:
      - fullstack

  ultimatedivision_db:
    restart: always
    image: postgres:latest
    container_name: ultimatedivision_db
    ports:
      - "5635:5432"
    volumes:
      - ${PROJECT_DATA_PATH}/db:/var/lib/postgresql/data
    networks:
      - fullstack
    environment:
      - POSTGRES_DB=${POSTGRES_DB}
      - POSTGRES_USER=${POSTGRES_USER}
      - POSTGRES_PASSWORD=${POSTGRES_PASSWORD}

# Networks to be created to facilitate communication between containers
networks:
  fullstack:
    driver: bridge
```
