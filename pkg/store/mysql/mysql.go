package mysql

import(
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"github.com/joho/godotenv"
)

type ConfigMysql struct{

	DbName string

	ServerName string

	Port string

	User string

	Password string

}

func init(){

	if err := godotenv.Load(); err != nil {
        fmt.Print("No .env file found")
    }

}

func (configMysql *ConfigMysql) initDefaultConfig(){

	configMysql.DbName = os.Getenv("DB_NAME")
	configMysql.ServerName = os.Getenv("DB_SERVER_NAME")
	configMysql.User = os.Getenv("DB_USER")
	configMysql.Port = os.Getenv("DB_PORT")
	configMysql.Password = os.Getenv("DB_PASSWORD")

}

func (configMysql ConfigMysql) toDsn() string{

	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v", configMysql.User, configMysql.Password, configMysql.ServerName, configMysql.Port, configMysql.DbName)
	fmt.Println(dsn)
	return dsn

}

func ConnectDb(configMysql ConfigMysql) (*sql.DB, error){

	if ((ConfigMysql{}) == configMysql){
		(&configMysql).initDefaultConfig()
	}

	fmt.Println(configMysql)

	dsn := configMysql.toDsn()

	database, err := sql.Open("mysql", dsn)

	if (err != nil){
		return nil, err
	}

	return database, nil

}