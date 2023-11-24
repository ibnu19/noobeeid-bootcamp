docker rm -f postgresql web;
docker build -t ibnu19/web:1.0 .;

docker run -d \
--name=postgresql \
-e POSTGRES_PASSWORD=admin \
-v ${PWD}/data:/var/lib/postgresql/data \
-p 5432:5432 \
postgres:16.1-alpine3.18;

docker run -d \
--name=web \
-e APP_PORT=:8080 \
-p 5000:8080 \
ibnu19/web:1.0