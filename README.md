Setup instructions:

PostgreSQL:
1. Start your PostgreSQL server
2. Create a database
3. Ensure all permissions from the database created in 2. are granted to your user.
   Ideally the user is a superuser.

Backend:
Environment setup:
1. Create .env in the top level of api/
2. Create the following environment variables:
```
ADDR=:PORT_NUMBER
DB_URL=postgres://POSTGRES_USER:PASSWORD@localhost:DATABASE_PORT/DATABASE_NAME?sslmode=disable 
//DATABASE_PORT is usually 5432
DB_VERSION=1
DOMAIN=FRONTEND_DOMAIN //The URL of the frontend
ACCESS_SECRET=SOME_SECRET_1 //can be anything really
REFRESH_SECRET=SOME_SECRET_2 //same as above
```
3. In the same folder, run
```
make migrate-up
```

Frontend:
1. In go-gossip/,  call 
```
npm install
```
2. In the same location, create a .env file with the following variables:
```
NEXT_PUBLIC_API_URL=BACKEND_DOMAIN //URL of backend
```

Start instructions:
1. Start your PostgreSQL server
2. Start the backend in api/ by calling
```
make dev
```
3. Start the frontend in go-gossip/ by calling 
```
npm run dev
``` 

AI USE:
- Gemini for searching up syntax of libraries, design patterns, best practices, meanings of error messages 