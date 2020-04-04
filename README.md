# golang-app-framework
Golang Application Framework

#### Based on:
1. github.com/gin-gonic/gin
2. github.com/go-redis/redis
3. github.com/go-sql-driver/mysql
4. github.com/jinzhu/gorm
5. github.com/sirupsen/logrus
6. github.com/spf13/viper
7. gopkg.in/natefinch/lumberjack.v2

#### Usage:
```
go mod download && go build -o framework
./framework -f ./config.yml
```

#### Docker:
```
cp docker/* . && chmod +x *.sh
./build.sh
./test-run.sh
```
