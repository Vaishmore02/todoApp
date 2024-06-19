# todoApp
todoApp is used to add , list to do tasks

# Docker setup commands to run scylla DB
docker run --name scylla-new -p 9042:9042 -d scylladb/scylla

# To find the docker host IP use following command, refer Ethernet adapter vEthernet , IPv4 Address
PS C:\Users\admin> ipconfig

# DB Commands to connect and create basic DB schema

docker exec -it scylla-new cqlsh

cqlsh> CREATE KEYSPACE dev WITH replication = {'class':'SimpleStrategy', 'replication_factor' : 1};

cqlsh> USE dev;

cqlsh> CREATE TABLE task(
   ID text PRIMARY KEY,
   UserID text,
   Title text,
   Description text,
   Status text,
   CreatedAt timestamp,
   UpdatedAt timestamp,
   );

# TO Run go server 
\todoApp> go run .\main.go

# API curl req
# create task
curl --location --request GET 'http://localhost:8080/todo?=' \
--header 'Content-Type: application/json' \
--data '{
    "title": "Homework",
    "description": "bla bla",
    "status": "TODO",
    "user_id": "1"
}'

# Delete task
curl --location --request DELETE 'http://localhost:8080/todo?Id=4d46b41b-5c42-4755-8ab9-1ebd5d59925a' \
--header 'Content-Type: application/json' \
--data '{
    "title": "Homework",
    "description": "bla bla",
    "status": "TODO",
    "user_id": "1"
}'

# Get details
curl --location --request GET 'http://localhost:8080/todo?Id=29f62b91-0fab-4557-a5c6-f3cfa589069a' \
--header 'Content-Type: application/json' \
--data '{
    "title": "Homework",
    "description": "bla bla",
    "status": "TODO",
    "user_id": "1"
}'

# list task
curl --location --request GET 'http://localhost:8080/todo/list?UserId=2' \
--header 'Content-Type: application/json' \
--data '{
    "title": "Homework",
    "description": "bla bla",
    "status": "TODO",
    "user_id": "1"
}'
