package GoDbChecker

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

const (
	Q_TABLE_LIST = `SELECT column1,column2,column3,column4 FROM table`
)

type db_col struct {
	col_name      string
	period        string
	dat_type      string
	date_col_name string
	str_date      string
	str_time      string
	//defualt_type string
}

func Clean_DB_Hist() {
	//mysql connect
	db, err := sql.Open("mysql", "test:pwd@111.1.1.1:3306/testdb")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	m_tb_list := map[string]db_col{}

	//, COLUMN_DEFAULT AS 'Default'

	rows, err := db.Query(Q_TABLE_LIST)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var tb_name, retention_period, datatype, db_col_name, string_date, string_time string
		err := rows.Scan(&tb_name, &retention_period, &datatype, &db_col_name, &string_date, &string_time)
		if err != nil {
			log.Fatal(err)
		}
		m_tb_list[tb_name] = db_col{tb_name, retention_period, datatype, db_col_name, string_date, string_time}
	}

	for k, v := range m_tb_list {
		fmt.Println("key : ", k, "value : ", v)
	}

	for _, v := range m_tb_list {
		if v.dat_type == "timestamp" {
			query := fmt.Sprintf(Q_TABLE_LIST, v.col_name, v.str_date, v.str_time, v.period)
			rows, err := db.Query(query)
			if err != nil {
				log.Fatal(err)
			}
			defer rows.Close()
		} else if v.dat_type == "string" {
			query := fmt.Sprintf(Q_TABLE_LIST, v.col_name, v.date_col_name, v.period)
			rows, err := db.Query(query)
			if err != nil {
				log.Fatal(err)
			}
			defer rows.Close()
		}
	}
}
