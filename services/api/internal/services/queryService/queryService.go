package queryService

import (
	// "database/sql"
	"fmt"
	"net/http"
	"net/url"

	// "reflect"
	"strconv"
	"strings"

	"github.com/MeibisuX673/GoCrud/services/api/internal/domain/query"

)

const(

	STRATEGY_EXACT = "eq"
	STRATEGY_PARTIAL = "includes"


)

func GetQueries(r *http.Request) *query.Queries{

	value, err := url.ParseQuery(r.URL.RawQuery)

	if err != nil{
		panic(err.Error())
	}

	queris := GetQueryFromMap(value)

	fmt.Println(queris)

	return queris

}

func GetQueryFromMap(queryMap url.Values) *query.Queries{

	var queris *query.Queries = &query.Queries{}

	page, err := strconv.Atoi(queryMap.Get("page"))

	if err != nil || page <= 0{
		page = 1
	}

	limit, err := strconv.Atoi(queryMap.Get("limit"))

	if err != nil || page <= 0{
		limit = 3
	}

	queris.Page = page
	queris.Limit = limit


	for key, value := range queryMap{

		replaser := strings.NewReplacer("["," ", "]"," ")

		queryData := strings.Fields(replaser.Replace(key))

		filterName := queryData[0]
		property := queryData[1]

		fmt.Println(queryData, value[0])

		filter := ""
		order := ""
		rangeFilter := ""

		switch filterName{
			case "filter":
				value := value[0]
				strategy := queryData[2]
				switch strategy{ 
					case STRATEGY_EXACT:
						filter = fmt.Sprintf("%s = '%s'", property, value)
					case STRATEGY_PARTIAL:
						filter = fmt.Sprintf("%s LIKE '%%%s%%'", property, value)
				}
			case "order":
				value := value[0]
				sort := strings.ToUpper(value)
				if sort == "DESC"{
					order = fmt.Sprintf("ORDER BY %s DESC", property)
				}else if sort == "ASC"{
					order = fmt.Sprintf("ORDER BY %s ASC", property)
				}
			case "range":

				start, err := strconv.Atoi(queryData[2])
				
				if err != nil{
					break
				}

				end, err := strconv.Atoi(queryData[3])

				if  err != nil{
					break
				}

				rangeFilter = fmt.Sprintf("%s BETWEEN %d AND %d", property, start, end)

		}
		
		if len(filter) != 0 {
			queris.Filters = append(queris.Filters, filter)
		}

		if len(order) != 0{
			queris.Order = order
		}

		if len(rangeFilter) != 0{
			queris.RangeFilter = rangeFilter
		}

	}

	// fmt.Println("filters: ", queris.Filters)

	return queris

}


func ConfigurationDbQuery(queryDb string, queris *query.Queries) string{

	for _, filter := range queris.Filters{
		if len(filter) == 0{
			continue
		}
		queryDb = where(queryDb, filter)
	}

	if len(queris.RangeFilter) != 0{
		queryDb = between(queryDb, queris.RangeFilter)
	}

	if len(queris.Order) != 0{
		queryDb = order(queryDb, queris.Order)
	}

	page := queris.Page
	border := queris.Limit

	queryDb = limit(queryDb, page, border)

	fmt.Println(queryDb)

	return queryDb

}

func where(queryDb string, condition string) string{


	if !strings.Contains(queryDb, "WHERE"){

		queryDb += " WHERE "

	}else{

		queryDb += " AND "
	}

	queryDb += condition

	return queryDb

}

func order(queryDb string, condition string) string{

	queryDb += " " + condition + " "

	return queryDb

}

func limit(queryDb string, page, limit int) string{

	pre := page * limit - limit

	queryDb += fmt.Sprintf(" LIMIT %d, %d", pre, limit)

	return queryDb

}

func between(queryDb string, condition string) string{

	if !strings.Contains(queryDb, "WHERE"){

		queryDb += " WHERE "

	}else{

		queryDb += " AND "
	}

	queryDb += condition

	return queryDb

}