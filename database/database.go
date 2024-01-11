package database
import(
	"log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"user-api/models"
)

var Db *gorm.DB
func ConnectDatabase(){
	dsn := "root:admin@tcp(127.0.0.1:3306)/todo?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// AutoMigrate will create the table. Make sure your MySQL server is running.
	Db.AutoMigrate(&models.Todo{})
}