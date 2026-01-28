#! /bin/bash
prep() {
  rm -rf output

  cd ../web
  npx vite build --mode stage --outDir ../devops/output/stage/web
  npx vite build --mode prod --outDir ../devops/output/www/web
  cd ../devops
}

stage() {
  # ssh -p 2222 -i ~/.ssh/id_rsa tzather@192.185.33.142

  # delete existing files
  ssh -p 2222 tzather@zuhid.com 'rm -r ~/crowdilm/stage/web/assets/*'

  # copy the files to remote server
  scp -P 2222 -rp output/stage/web/. tzather@zuhid.com:/home3/tzather/crowdilm/stage/web

  # set the correct permission
  ssh -p 2222 tzather@zuhid.com 'chmod -R 755 ~/crowdilm/stage'
}


prod() {
  # ssh -p 2222 -i ~/.ssh/id_rsa tzather@192.185.33.142

  # delete existing files
  ssh -p 2222 tzather@zuhid.com 'rm -r ~/crowdilm/www/web/assets/*'

  # copy the files to remote server
  scp -P 2222 -rp output/www/web/. tzather@zuhid.com:/home3/tzather/crowdilm/www/web

  # set the correct permission
  ssh -p 2222 tzather@zuhid.com 'chmod -R 755 ~/crowdilm/www'
}

# prep
# stage
# prod
