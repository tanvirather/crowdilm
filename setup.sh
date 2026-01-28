#!/bin/bash

################################################## variables ##################################################

DOCKER_CONTAINER='mysql_local'
HOST='localhost'
DATABASE='crowdilm'
USER='root'
PASSWORD='P@ssw0rd'
OUTPUT_DIRECTORY="tmp" # Output file

################################################## methods ##################################################

create_database(){
  echo -e "################################################## database ##################################################" >> "$OUTPUT_DIRECTORY/database.sql"
  cat "db/database.sql" >> "$OUTPUT_DIRECTORY/database.sql" # database file
  # Loop through all table.sql files and append them to the output file
  for file in db/table/*.table.sql; do
      echo -e "\n################################################## $file ##################################################" >> "$OUTPUT_DIRECTORY/database.sql"
      cat "$file" >> "$OUTPUT_DIRECTORY/database.sql"
  done
  # Loop through all .fk.sql files and append them to the output file
  for file in db/table/*.fk.sql; do
      echo -e "\n################################################## $file ##################################################" >> "$OUTPUT_DIRECTORY/database.sql"
      cat "$file" >> "$OUTPUT_DIRECTORY/database.sql"
  done
  # run the script against the database
  docker exec --interactive $DOCKER_CONTAINER mysql --host $HOST --user $USER --password=$PASSWORD < "$OUTPUT_DIRECTORY/database.sql"
}

download_quran_files(){
  cd etl
  ./download.sh
  cd ..
}

generate_etl(){
  cd etl
  rm -rf output
  mkdir output
  go run .
    cd output
    chmod +x _build_database.sh
    ./_build_database.sh
    cd ..
  cd ..
}

################################################## execution ##################################################

clear
docker compose -f api/docker-compose.yml up --detach # start the docker for api in http://localhost:8080
# rm -rf tmp
# mkdir tmp
# create_database 
# # download_quran_files
# generate_etl
