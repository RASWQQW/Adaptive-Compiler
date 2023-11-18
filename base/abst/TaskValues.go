package abst

import (
	"database/sql"
	"ep/Execs/obj"
	lv "ep/LevelFuncs"
	BaseFunc "ep/base"
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

	if get_all {
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
	for v, _ := range selCols {
		selCols[v] = selCols[v] + "::text"
	}
	var getArrayer string = strings.Join(selCols, ", ")
	fmt.Println("Reader params:", getArrayer, table, addString)
	req, reqerr := ctx.dbOb.Query(fmt.Sprintf(`SELECT array[%s] as values FROM "%s" %s;`, getArrayer, table, addString)) // []interface{}{getArrayer}...
	if reqerr != nil {
		panic(reqerr)
	}
	for req.Next() {
		var GetArray []string
		req.Scan(pq.Array(&GetArray))
		fmt.Println(GetArray)
		Mapper := map[string]string{}
		for in, _ := range selCols {
			Mapper[selCols[in][:len(selCols[in])-6]] = GetArray[in]
		}

		fmt.Println("Add val:", Mapper)
		ret_vals = append(ret_vals, Mapper)
	}
	fmt.Println("Last result:", ret_vals)
	return
}

// HERE GOES BASE METHODS AND LOGICS TO CREATE SAVE AND GET TASKS
func (ctx *BaseConnection) GetFunction(worker *obj.Career) /*[]string*/ {
	var task_id string = lv.ToString(worker.ValFinder("task_id", "out", -1))

	funcVals, _ := ctx.dbOb.Query(`SELECT return_value, func_name FROM "Functions" WHERE task_id = $1`, task_id)
	var returnType string
	var funcName string
	funcVals.Next()
	funcVals.Scan(&returnType, &funcName)

	worker.OUTS["return_type"] = returnType
	worker.INOUTS["func_name"] = funcName
	//return []string{returnType, funcName}
}

func (ctx *BaseConnection) GetFuncParams(valsl *obj.Career) { //values map[string]interface{}, Chans map[string]chan []string, Pointers map[string]interface{}
	// var task_id string = lv.ToString(values["task_id"])

	var task_id string = lv.ToString(valsl.ValFinder("task_id", "out", -1))
	vals, _ := ctx.dbOb.Query(`SELECT id, args_names, args_types FROM "FunctionArgs" WHERE task_id = $1 LIMIT 1`, task_id)

	var args [][]string
	var args_id int
	var args_names []string
	var args_types []string

	vals.Next()
	vals.Scan(&args_id, pq.Array(&args_names), pq.Array(&args_types))

	if len(args_names) == len(args_types) {
		for iter, val := range args_names {
			args = append(args, []string{args_types[iter], val})
		}
	} else {
		panic("Length of arg names and types doesn't match")
	}
	valsl.OUTS["arg_names"] = args_names
	valsl.OUTS["arg_types"] = args_types
	valsl.OUTS["args"] = args

	// if len(Chans) > 0 {
	// 	var keyLists = maps.Keys(Chans)
	// 	if strings.Contains(strings.Join(keyLists, ","), "prNames") {
	// 		Chans["prNames"] <- arg_names
	// 	}
	// 	if strings.Contains(strings.Join(keyLists, ","), "prTypes") {
	// 		Chans["prTypes"] <- args_types
	// 	}
	// } else {
	// 	Pointers["MainParams"] = args
	// }

}

func (ctx *BaseConnection) GetProperCode(worker *obj.Career) /*string*/ {
	var task_id string = lv.ToString(worker.ValFinder("task_id", "out", -1))
	var lang string = lv.ToString(worker.INOUTS["lang"])

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
	worker.OUTS["prop_code"] = Values[0][1]
	//return Values[0][1] // mere simple to get one proper code val return
}

func (ctx *BaseConnection) GetTaskByName(values *obj.Career) { //(int, string, string)
	// task_name_id string
	var task_name_id string = lv.ToString(values.INOUTS["task_name_id"])
	// var task_name_id string

	// var valp []*string = []*string{&task_name_id}

	// There you can use referencing by & to set values
	// var valsd []interface{} = append([]interface{}{}, vals...)
	// if typeN := reflect.ValueOf(valsd[0]); typeN.Kind() == reflect.String {
	// 	var task_name_id string = typeN.String()
	// }

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

		values.OUTS["task_type"] = []string{fmt.Sprintf("%d", taskType)}
		values.OUTS["task_id"] = []string{fmt.Sprintf("%d", retid)}
		// var taskTypeVal = ctx.Reader(false, []string{"type_name"}, "TaskTypes", string(fmt.Sprintf("WHERE id = %d", taskType)))
		// return retid, taskTypeVal[0]["type_name"], task_name_id
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
