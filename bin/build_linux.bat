set GOOS=linux
set CGO_ENABLED=0
cd ../
del bin/fbrobot
go build -o ./bin/fbrobot ./services/fbrobot/main.go
REM pause