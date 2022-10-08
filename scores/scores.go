package scores

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

	_ "github.com/lib/pq"
)

func readCsvFile(filePath string) [][]string {
	f, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		panic(err)
	}

	return records
}

func openDB(host, user, password, dbname string, port int) (*sql.DB, error) {
	postgresqlDbInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", postgresqlDbInfo)
	return db, err
}

func writeRecords(records [][]string, db *sql.DB) error {
	_, err := db.Exec(
		"DROP TABLE IF EXISTS scores;",
	)
	if err != nil {
		return err
	}

	_, err = db.Exec(
		`CREATE TABLE scores (
			user_id serial PRIMARY KEY, 
			name VARCHAR (50) NOT NULL,
			course VARCHAR (50) NOT NULL,
			grade INT NOT NULL
			)`,
	)
	if err != nil {
		return err
	}

	for _, rec := range records {
		grade, err := strconv.Atoi(rec[2])
		if err != nil {
			panic(err)
		}
		_, err = db.Exec(
			"INSERT INTO scores (name, course, grade) VALUES ($1, $2, $3)",
			rec[0],
			rec[1],
			grade,
		)
		if err != nil {
			return err
		}
	}
	return nil
}

func getTopStudents(db *sql.DB) ([]string, error) {
	topStudents := map[string]struct{}{}
	rows, err := db.Query("SELECT name FROM scores WHERE grade = $1", 20)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var student string
		if err = rows.Scan(&student); err != nil {
			panic(err)
		}
		topStudents[student] = struct{}{}
	}
	var result []string
	for name, _ := range topStudents {
		result = append(result, name)
	}
	return result, nil
}

func runQueries() {

	// records := readCsvFile("scores.csv")

	const (
		host     = "localhost"
		port     = 5432
		user     = "postgres"
		password = "password"
		dbname   = "golangtest"
	)

	db, err := openDB(host, user, password, dbname, port)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// err = writeRecords(records, db)
	// if err != nil {
	// 	panic(err)
	// }

	result, err := getTopStudents(db)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}
