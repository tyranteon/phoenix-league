@echo off
docker stack rm phoenix-league
ping 127.0.0.1 -n 8 > nul
for /f "delims=" %%i in ('docker volume ls -q') do docker volume rm %%i
ping 127.0.0.1 -n 8 > nul
docker stack deploy -c stack.yml phoenix-league
pushd flyway
ping 127.0.0.1 -n 11 > nul
flyway migrate && popd