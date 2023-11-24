docker run -d --name=app \
-e DB_HOST=postgresql \
-e DB_PORT=5432 \
-e DB_USER=postgres \
-e DB_PASSWORD=admin \
-e DB_NAME=noobeeid \
-e APP_PORT=5000 \
-p 3000:5000 \
--network=network-app \
ibnu19/app:1.0