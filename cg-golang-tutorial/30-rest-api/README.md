## Resource
1. https://youtu.be/OGhQhFKvMiM?list=PLXQpH_kZIxTWUe-Ee-DZEX5gfeoo4tHV6

## Aim
1. json api 
2. crud applicaiton
3. use thunder clinet to test api endpoints

## Process
1. go mod init github.com/radialhuman/cg-golang-tutorial/30_rest_api
2. cmd/students-api/main.go
    -  go  run cmd/students-api/main.go
    -  cmd is convention
3. config file : config/local.yaml, must be prod.yaml in future
4. use sqlite
5. When the applciaityon starts, it needs to first read the config files and use them as variable 
6. internal/pkg : is also a convention
7. using cleanenv for reading and desearlizing config files
    - go get -u github.com/ilyakaznacheev/cleanenv