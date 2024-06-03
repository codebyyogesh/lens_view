## How to run docker for db

```
docker compose up    # to start the db and adminer services

docker compose exec -it db psql -U bagheera -d lensview   # connect to db

```
