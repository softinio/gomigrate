package conf

import (
	"encoding/json"
	"fmt"
	"os"
)

const (
	DB_CONFIG_FILE = "conf/db.json"
)

type Connection struct {
	Dbname   string
	Host     string
	Port     int
	User     string
	Password string
}

func (dbc *Connection) DbConnectionString() string {
	return fmt.Sprintf("user=%s dbname=%s host=%s port=%d password=%s sslmode=disabled", dbc.User, dbc.Dbname, dbc.Host, dbc.Port, dbc.Password)
}

var DbxConfig = new(Connection)

func init() {
	// Open db config file
	dbfile, err := os.Open(DB_CONFIG_FILE)

	if err != nil {
		if len(DB_CONFIG_FILE) > 1 {
			fmt.Printf("Error: Could not read db config file: %s \n %s \n", DB_CONFIG_FILE, err)
		}
	}

	dbdecoder := json.NewDecoder(dbfile)
	err = dbdecoder.Decode(&DbxConfig)
	if err != nil {
		fmt.Printf("Error decoding: %s \n %s \n", dbfile, err)
	}
}
