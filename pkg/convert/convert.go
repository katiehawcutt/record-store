package convert

import (
	"log"
	"strconv"
)

func ConvertStringToInt32(s string) (i int32) {
	int64Value, err := strconv.ParseInt(s, 10, 32)

	if err != nil {
		log.Fatal("Error while converting string to int64")
	}

	int32Value := int32(int64Value)

	return int32Value
}
