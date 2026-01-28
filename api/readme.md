# Run applliation

```sh
docker compose up --detach # start docker
# open in `http://localhost:8080`
```

# Useful commands

```sh
docker exec -it apache_local /bin/bash
docker exec -it mysql_local /bin/bash
docker exec -it mysql_local mysql -u root -p
docker compose down # take down docker
docker compose up --detach --force-recreate
mysql -u root -p'P@ssw0rd' < db/database.sql
mysql -u root -p'P@ssw0rd' crowdilm < quran.table.sql
mysql -u root -p'P@ssw0rd' < init.sql
mysql -u root -p'P@ssw0rd' < init.sql
mysql -u root -p'P@ssw0rd' < init.sql
mysql -u root -p'P@ssw0rd' < init.sql


docker cp db/table/quran.table.sql mysql_local:tmp/quran.table.sql
docker exec -i mysql_local mysql -u root -p[P@ssw0rd] crowdilm < /tmp/quran.table.sql
docker exec -i mysql_local mysql -u root -p[P@ssw0rd] crowdilm < db/quran.table.sql


```
