package generate

func Database() {
	createFile(build_database, `#!/bin/bash

################################################## variables ##################################################
DOCKER_CONTAINER='mysql_local'
HOST='localhost'
DATABASE='crowdilm'
USER='root'
PASSWORD='P@ssw0rd'

################################################## dataload ##################################################
docker exec --interactive $DOCKER_CONTAINER mysql --host $HOST --database $DATABASE --user $USER --password=$PASSWORD --default-character-set=utf8 < _dataload.sql
`)
}
