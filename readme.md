# go-mysql

Convox Base image for Go.

This repository uses two containers: Go Webservice and MySQL Database.

Run containers
```
docker-compose build
docker-compose up
```

Hit the HTTP endpoint
```
curl localhost/users
```

Files
```
├── database
│   ├── Dockerfile
│   └── setup.sql
├── docker-compose.yml
└── webservice
    ├── config.json.sample
    ├── Dockerfile
    ├── main.go
    └── webservice
```
