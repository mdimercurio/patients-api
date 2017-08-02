## Patients API

Implementation of the API defined in [swagger.yml](swagger.yml)

### Build and run

```sh
./dockerbuild.sh
docker run -p 8080:8080 patients-api:latest
```


### POST

```sh
curl -H "Content-Type: application/json" -d '{"email":"jim@gmail.com","first_name":"jim","last_name":"jimmy","birthdate":"2000-01-01T00:00:00Z","sex":" Male"}' -X POST http://localhost:8080/v1/patients
```

### GET by ID

```sh
curl -i http:/localhost:8080/v1/patients/0
```

### GET list

```sh
curl -i http:/localhost:8080/v1/patients
```

### Benchmark

Benchmark post requests with:

```sh
go test -bench=. -benchtime=20s
```