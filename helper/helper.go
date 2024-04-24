package helper

import (
	"encoding/json"
	"example/go-crud/models"
	"fmt"
	"net/url"
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

func ParseFilterParams(queryParams map[string][]string) models.Filter {
	filter := models.Filter{
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

	if _, ok := queryParams["filterDetails"]; ok {
		var filterDetails []models.FilterDetail
		filterDetailsStr, err := url.PathUnescape(queryParams["filterDetails"][0])
		if err != nil {
			fmt.Println("Error:", err)
			return filter
		}
		err = json.Unmarshal([]byte(filterDetailsStr), &filterDetails)
		if err != nil {
			fmt.Println("Error:", err)
			return filter
		}
		filter.FilterDetails = filterDetails
	}
	return filter
}
