package helper

import (
	"encoding/json"
	"example/go-crud/forms"
	"fmt"
	"strconv"

	"go.mongodb.org/mongo-driver/bson"
)

func AppendCondition(filter bson.M, key string, value interface{}) bson.M {
	if filter == nil {
		filter = bson.M{}
	}
	filter[key] = value
	return filter
}

func ParseFilterParams(queryParams map[string][]string) forms.Filter {
	filter := forms.Filter{
		Limit:  10, // Default limit
		Offset: 0,  // Default offset
	}

	if limitStr, ok := queryParams["limit"]; ok {
		limit, _ := strconv.ParseInt(limitStr[0], 10, 64)
		filter.Limit = limit
	}

	if offsetStr, ok := queryParams["offset"]; ok {
		offset, _ := strconv.ParseInt(offsetStr[0], 10, 64)
		filter.Offset = offset
	}

	if filterDetailsStr, ok := queryParams["filterDetails"]; ok {
		var filterDetails []forms.FilterDetail
		err := json.Unmarshal([]byte(filterDetailsStr[0]), &filterDetails)
		if err != nil {
			fmt.Println("Error:", err)
			return filter
		}
		filter.FilterDetails = filterDetails
	}
	return filter
}
