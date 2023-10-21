package base

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "127.0.0.1"
	port     = "5432"
	user     = "postgres"
	password = "1212"
	dbname   = "Compiler"
)

// TOMORROW I HAVE TO CREATE TABLE AND SET A ONE TASK TO CHECK
// WTIH NEW COMPILER
func Connecting() sql.DB { //(ct *BaseConnection)
	var connectString string = fmt.Sprintf(
		"host=%s port=%s dbname=%s user=%s password=%s sslmode=disable",
		host, port, dbname, user, password)

	fmt.Println("ConnectString: ", connectString)

	db, err := sql.Open("postgres", connectString)
	defer db.Close()

	if err != nil {
		panic(err)
	} else {

		db.Begin()
		fmt.Println("Back Response:", db.Ping())
		var execString string = `create table if not exists "Tasks"(id serial primary key, task_name_id varchar not null, text varchar(255) not null);`
		result, err := db.Exec(execString)

		fmt.Println("result", result)
		fmt.Println("error", err)

		var CreateFuncs string = `CREATE TABLE IF NOT EXISTS "Functions"(
									id serial primary key, 
									task_id integer REFERENCES "Tasks"(id), 
									func_name varchar not null,
									return_value varchar default 'void'
									);`

		res, secerror := db.Exec(CreateFuncs)
		fmt.Println("result", res, "error: ", secerror)

		var ProperCode string = `CREATE TABLE IF NOT EXISTS "ProperCode"(
								id serial primary key,
								task_id integer REFERENCES "Tasks"(id),
								lang varchar not null,
								code varchar not null
								);`

		res3, _ := db.Exec(ProperCode)
		fmt.Println("result", res3)

		var CreateFuncArgs string = `CREATE TABLE IF NOT EXISTS "FunctionArgs"(
									id serial primary key, 
									task_id integer REFERENCES "Tasks"(id), 
									args_name text[],
									args_types text[]
									);`

		res2, nerr := db.Exec(CreateFuncArgs)
		fmt.Println("FunctionArgs", res2, " error: ", nerr)
		// giving global value to struct
		return *db
	}

}
