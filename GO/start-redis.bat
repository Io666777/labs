@echo off

docker-compose up --build

echo start Redis...
docker stop my-redis 2>nul
docker rm my-redis 2>nul
docker run -d --name my-redis -p 6379:6379 redis redis-server --bind 0.0.0.0
timeout /t 2 /nobreak >nul

 
go run text.go

echo.
 
echo.
docker exec my-redis redis-cli KEYS "*"
echo.

 
for /f "tokens=*" %%i in ('docker exec my-redis redis-cli KEYS "*"') do (
  echo Key: %%i
  docker exec my-redis redis-cli GET "%%i"
  echo.
)

pause