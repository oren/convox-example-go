# go-mysql

```
docker-compose up
curl localhost:3000/users
```
http: panic serving 172.22.0.1:50354: default addr for network 'db' unknown

Access MySQL from the Go web service:
```
docker exec -it gomysql_users-service_1 sh
apk update
apk add --no-cache mysql-client
mysql -h db -u root -p123
```
