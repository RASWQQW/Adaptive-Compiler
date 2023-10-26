package abst

import (
	"database/sql"
	BaseFunc "example/base"
	"fmt"
	"strings"

	"github.com/lib/pq"
)

// Here it gonna be stopped for a while
// to be concerned lately
type BaseConnection struct {
	dbOb sql.DB
}

var db BaseConnection = BaseConnection{BaseFunc.Connecting()}

func GetBase() BaseConnection {
	// var base = &BaseFunc.BaseConnection{dbOb = db}
	return db
}

func (ctod *BaseConnection) Fof() {}

// IT GOES FOR ALL REQUESTS WHICH IS QUITE SIMPLE BY GETTING OBVIOUS VALUES
func (ctx *BaseConnection) Reader(get_all bool, selCols []string, table string, addString string) (ret_vals []map[string]string) {

	if get_all == true {
		var query string = `SELECT * FROM $1 $2`
		var query2 string = `SELECT column_name from information_schema.columns where table_name = '$1'`
		req, _ := ctx.dbOb.Query(query, addString)
		selCols, _ = req.Columns()

		req2, _ := ctx.dbOb.Query(query2, addString)
		var column_names []string = []string{}
		for req2.Next() {
			var colName string
			req2.Scan(&colName)
			column_names = append(column_names, colName)
		}
		selCols = column_names
	}

	var getArrayer string = strings.Join(selCols, "::text, ")
	req, _ := ctx.dbOb.Query(`SELECT array[$1] as values FROM $2 $3`, getArrayer, table, addString)
	for req.Next() {
		var GetArray []string
		req.Scan(pq.Array(&GetArray))
		Mapper := map[string]string{}
		for in, _ := range selCols {
			Mapper[selCols[in]] = GetArray[in]
		}
		ret_vals = append(ret_vals, Mapper)
	}
	return
}

// HERE GOES BASE METHODS AND LOGICS TO CREATE SAVE AND GET TASKS
func (ctx *BaseConnection) GetFunction(task_id int) []string {
	funcVals, _ := ctx.dbOb.Query(`SELECT return_value, func_name FROM "Functions" WHERE task_id = $1`, task_id)
	var returnType string
	var funcName string
	funcVals.Next()
	funcVals.Scan(&returnType, &funcName)
	return []string{returnType, funcName}
}

func (ctx *BaseConnection) GetFuncParams(task_id int) [][]string {
	vals, _ := ctx.dbOb.Query(`SELECT id, args_name, args_types FROM "FunctionArgs" WHERE task_id = $1 LIMIT 1`, task_id)

	var args [][]string
	if vals != nil {
		var args_id int
		var arg_names []string
		var args_types []string

		vals.Next()
		vals.Scan(&args_id, &arg_names, &args_types)

		for iter, val := range arg_names {
			args = append(args, []string{val, args_types[iter]})
		}
	}
	return args
}

func (ctx *BaseConnection) GetProperCode(task_id int, lang string) string {
	req, _ := ctx.dbOb.Query(`SELECT id, code, lang FROM "ProperCode" WHERE task_id = $1 AND lang = $2 LIMIT 1;
	`, task_id, lang)

	var Values [][]string = [][]string{}
	if req != nil {
		var id int
		var code string
		var lang string
		for req.Next() {
			req.Scan(&id, &code, &lang)
			Values = append(Values, []string{string(id), code, lang})
		}
	}
	// fmt.Println("Values: ", Values, "Len: ", len(Values), "Lang: ", lang, task_id)
	return Values[0][1] // mere simple to get one proper code val return
}

func (ctx *BaseConnection) GetTaskByName(task_name_id string) (int, string, string) {
	fmt.Println("checking: ", ctx.dbOb.Ping(), task_name_id)
	res, err := ctx.dbOb.Query(`SELECT "id", "tasktype" FROM "Tasks" WHERE "Tasks".task_name_id = $1 LIMIT 1`, task_name_id)

	if err != nil {
		panic(err)
	} else {
		var retid int
		var taskType int
		for res.Next() {
			res.Scan(&retid, &taskType)
			// fmt.Println("Current id in for scan: ", retid)
		}
		var taskTypeVal = ctx.Reader(false, []string{"type_name"}, "TaskType", "")
		return retid, taskTypeVal[0]["type_name"], task_name_id
	}
}

func (ctx *BaseConnection) GetTaskById(task_id int) string {
	fmt.Println("Ins estate ", ctx, "Db esatate", ctx.dbOb)
	fmt.Println(ctx.dbOb.Ping())

	if ctx.dbOb.Ping() == nil {
		rr, err := ctx.dbOb.Query(fmt.Sprintf(
			`SELECT "task_name_id" FROM "Tasks" WHERE "Tasks".id = '%d' LIMIT 1;`, task_id))

		if err != nil {
			panic(err)
		} else {
			var task_name string
			var values []interface{} = []interface{}{&task_name}
			for rr.Next() {
				rr.Scan(values...)
			}
			// fmt.Println("Values", values, "task_n"+task_name)
			return string(task_name)
		}
	}
	return ""
}

// func (ctx *BaseConnection) ExecuteMany(resqs []string) {
// 	res, _ := ctx.dbOb.Query(fmt.Sprintf(`%s`, string(strings.Join(resqs, "; "))))
// 	for val := true; val; val = res.NextResultSet() {
// 		var columns, _ = res.Columns()
// 		var columnTypes, _ = res.ColumnTypes()
// 		var getVals []interface{}
// 		for col, val := range columns {

// 			// there have to be one thing that creates type calling it and giving type and name
// 			//  columnTypes[col] - type val val name

// 			var vaap interface{}
// 			// var tot []int = vaap.(columnTypes[col])[12:]

// 			// getVals = append(getVals, make(reflect.ValueOf(columnTypes[col]).Type(), val))
// 		}

// 		for res.Next() {
// 			res.Scan(getVals...)
// 		}
// 	}

// }