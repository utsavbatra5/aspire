package database

import (
	"aspire/config"
	"aspire/constants"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" //Required for connecting to PostgreSQL

	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/go-sql-driver/mysql"
)

var dbConnections = make(map[string]*sql.DB)

type DBAttributes struct {
	Server   string
	UserID   string
	Password string
	Port     int
	Database string
}

func GetConnection(dbconnection string) *sql.DB {
	return dbConnections[dbconnection]
}

// InitDB Connections
func InitConnections() {

	dbConfig := config.Conf.Db

	if dbConfig.Driver == constants.PostgresDriver {
		dbConnectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s",
			dbConfig.Host, dbConfig.Port, dbConfig.Username, dbConfig.Password, dbConfig.Database)
		db, err := sql.Open(constants.PostgresDriver, dbConnectionString)
		if err != nil {
			fmt.Errorf("unable to open postgres connection", err.Error())
		}
		dbConnections[dbConfig.Name] = db
	} else if dbConfig.Driver == constants.MySQLDriver {
		connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", dbConfig.Username, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.Database)
		db, err := sql.Open(constants.MySQLDriver, connString)
		if err != nil {
			fmt.Errorf("unable to open mysql connection", err.Error())
		}
		dbConnections[dbConfig.Name] = db
	} else if dbConfig.Driver == constants.MSSQLDriver {
		connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;", dbConfig.Host, dbConfig.Username, dbConfig.Password, dbConfig.Port, dbConfig.Database)
		db, err := sql.Open(constants.MSSQLDriver, connString)
		if err != nil {
			fmt.Errorf("unable to open mssql connection", err.Error())
		}
		dbConnections[dbConfig.Name] = db
	}

}
