package database

import (
	"FiberAuth/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

/*
Connecting to the mysql container and verify:
docker exec -it mysqldb-fiber bash
mysql -u root -p (check db password in docker-compose.yml)
*/

func Connect() {

	connection, err := gorm.Open(mysql.Open("root:root@/fiberauthdb"), &gorm.Config{})

	if err != nil {
		panic("could not connect to the database")
	}

	connection.AutoMigrate(&models.User{})
}
