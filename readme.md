# Initialize the api

First we must instantiate the db, we will use postgres.

we build the image:

```bash
cd databases/ && docker build . -t post-db
```

create the container

```bash
docker run -p 54321:5432 post-db
```
