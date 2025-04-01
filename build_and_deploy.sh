#!/usr/bin/env bash
cd /services/webserver/Monkey-Factory
git pull
cd /services/webserver/SpacePortfolio
git pull
cd /services/webserver/WaveCollapseEditor
git pull
cd /services/webserver/vertical_profile
git pull

# VERSION=$(git rev-parse HEAD)

# echo "$(date --utc +%FT%TZ): Building and deploying ..."
# docker compose up

# echo "$(date --utc +%FT%TZ): Building new image..."
# docker compose rm -f --remove-orphans
# docker compose build

# PREV_CONTAINER=$(docker ps -aqf "name=server")
# echo "$(date --utc +%FT%TZ): Building Container 2..."
# docker compose up -d --no-deps --scale server=2 --no-recreate server 

# sleep 30

# echo "$(date --utc +%FT%TZ): Scaling down server 1 ..."
# docker compose rm -f $PREV_CONTAINER --remove-orphans
# docker compose up -d --no-deps --scale server=1 --no-recreate server

# echo "$(date --utc +%FT%TZ): Restart Caddy..."
# CADDY_CS=$(docker ps -aqf "name=caddy")
# docker exec $CADDY_CS caddy reload -c /etc/caddy/Caddyfile


cd /services/DebianWebServer
git pull

docker compose down
docker compose build --no-cache
docker compose up -d 


