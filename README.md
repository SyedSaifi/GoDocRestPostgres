#GoDocRestPostgres
This is a sample go application which shows the usage of restful webservice (using mux framework), postgres DB and Docker for deployment.

### Setup Postgres DB

```
psql -h localhost -p 32300 -d docker -U docker --password

CREATE TABLE recipes (
    RecipeID INT PRIMARY KEY,
    Name TEXT NOT NULL,
    PrepTime TEXT NOT NULL,
    Difficulty INT NOT NULL,
    isVeg BOOLEAN NOT NULL
);
      
CREATE SEQUENCE public.recipes_id_seq NO MINVALUE NO MAXVALUE NO CYCLE;
ALTER TABLE public.recipes ALTER COLUMN RecipeID SET DEFAULT nextval('public.recipes_id_seq');
ALTER SEQUENCE public.recipes_id_seq OWNED BY public.recipes.RecipeID;
```

### Running the project

```
glide install
go run main.go
```

### Docker Compose file
version: '3'
services:

# Application containers
  postgres:
    build: db/.
    restart: unless-stopped
    ports:
        - "5432:5432"
    environment:
        LC_ALL: C.UTF-8
        POSTGRES_USER: docker
        POSTGRES_PASSWORD: docker
        POSTGRES_DB: docker
  go:
    build:
      context: ../
      dockerfile: docker/go/Dockerfile
    command: ["go", "run", "main.go"]
    ports:
        - "6000:6000"
    links:
        - postgres
    environment:
        DEBUG: 'true'
        PORT: '6000'
  # postgres:
  #   build:
  #     context: ../
  #     dockerfile: docker/Postgres/Dockerfile
  #   ports:
  #     - "32300:5432"