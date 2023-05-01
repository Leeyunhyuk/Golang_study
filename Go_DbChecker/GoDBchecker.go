package GoDbChecker

import (
	"database/sql"
	"fmt"
	"time"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

const (
	COLOUM_PROPERTIES = `SELECT COLUMN_NAME AS 'Field', COLUMN_TYPE AS 'Type', IS_NULLABLE AS 'NULL', COLUMN_DEFAULT AS 'Default'
	FROM information_schema.COLUMNS  
	WHERE TABLE_SCHEMA = 'DB_name' AND TABLE_NAME = 'Table_name';`
)

type dbProp struct {
	col_name string
	dat_type string
	null_tpye string
	defualt_type string
}

func SchemaChecker() {
	//mysql connect
	db, err := sql.Open("mysql","test:pwd@111.1.1.1:3306/testdb")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Set table name and properties to monitor
	tableName := S_RUH
	properties := []dbProp{}
	
	properties[0] = {col_name}

	// Initialize last property values
	lastProperties := make(map[string]string)
	for _, prop := range properties {
		lastProperties[prop] = ""
	}

	for {
		// Query table properties
		rows, err := db.Query(fmt.Sprintf("SHOW COLUMNS FROM %s", tableName))
		if err != nil {
			panic(err)
		}
		defer rows.Close()

		// Check if any properties have changed
		propertiesChanged := false
		for rows.Next() {
			var columnName string
			var dataType string
			err = rows.Scan(&columnName, &dataType)
			if err != nil {
				panic(err)
			}
			if lastProperties["COLUMN_NAME"] != columnName || lastProperties["DATA_TYPE"] != dataType {
				propertiesChanged = true
				break
			}
		}

		// Update last property values if they have changed
		if propertiesChanged {
			for _, prop := range properties {
				var value string
				err = db.QueryRow(fmt.Sprintf("SELECT %s FROM %s LIMIT 1", prop, tableName)).Scan(&value)
				if err != nil {
					panic(err)
				}
				lastProperties[prop] = value
			}
			fmt.Println("Table properties have changed.")
		} else {
			fmt.Println("Table properties have not changed.")
		}

		// Wait for 5 seconds before checking again
		time.Sleep(5 * time.Second)
	}
}
