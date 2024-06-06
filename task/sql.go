package task

import (
	"fmt"
	"strings"

	"github.com/illenko/digoflow/container"
)

type SQL struct{}

func (t *SQL) Execute(c *container.Container, input Input) (Output, error) {

	query := input["query"].(string)

	if query == "" {
		return nil, fmt.Errorf("query is required")
	}

	if strings.HasPrefix(query, "SELECT") {
		rows, _ := c.Database.Query(query)
		cols, _ := rows.Columns()
		colsTypes, _ := rows.ColumnTypes()

		if rows.Next() {
			columns := make([]interface{}, len(cols))
			columnPointers := make([]interface{}, len(cols))
			for i := range columns {
				columnPointers[i] = &columns[i]
			}

			if err := rows.Scan(columnPointers...); err != nil {
				return Output{}, err
			}

			m := make(Output)
			for i, colName := range cols {

				colType := colsTypes[i]

				val := columnPointers[i].(*interface{})

				switch colType.DatabaseTypeName() {
				case "JSONB", "UUID":
					m[colName] = string((*val).([]byte))
				default:
					m[colName] = *val
				}
			}

			return m, nil

		}
	} else {

		res, err := c.Database.Exec(input["query"].(string))

		fmt.Println(res)

		if err != nil {
			return nil, err
		}

		return Output{}, nil
	}

	return Output{}, nil

}
