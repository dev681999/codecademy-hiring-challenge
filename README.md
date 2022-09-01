# Catinator

Catinator is the best platform to store all ypu CAT pictures.

## Tech Stack -
**Backend** - Golang

**Authentication** - Simple Bearer Token Auth (NOT PRODUCTION READY)

**Database** - Postgres

**UI** - ReactJS

**API** - Swagger OpenAPI 3 Documentation https://localhost:8080/swagger-ui/


# Authentication-
A simple Bearer Token Auth. It checks the **Authorization** header in format **Bearer token**
Only logged/signedup users are allowed to upload their favourite pictures.
Only the user who has uploaded the a cat picture can edit it.


# Image -
Image is stored locally in file system.

# How To Run?
1. Setup base folders
```
make setup
#or
mkdir .local
```
2. Then build the Docker image 
```
make build
#or
docker build -t catinator-backend:latest .
```
3. Then setup database 
```
make db-up
#or
docker compose up -d db adminer
#or
docker-compose up -d db adminer
```
4. Then run the backend
```
make docker-run
#or
docker compose run backend
#or
docker-compose run backend
```
5. Visit Swagger UI - http://localhost:8080/swagger-ui/
6. Register a User
7. Login with user and Authorize with generate token


## Testing - 
A basic e2e integraiton test is written.
If there was enough time more unit tests could have been added to make tetsing more robust.



# More things -
1. I really wanted to seperate business logic from http handlers but didnt have time
2. Testing could be better
3. API Security is not very best
