package util

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func GetMovieActiveStatusById(num uint8) string {
	//	Cinema : movie_master : movie_is_active = 0-Deleted, 1-Active, 2-Inactive
	movie_is_active := map[int]string{0: "Deleted", 1: "Active", 2: "Inactive"}

	out, status := movie_is_active[int(num)]
	if status {
		return out
	}
	return "Not Available"
}

func GetMovieStatusById(num uint8) string {
	//	Cinema : movie_master : movie_status = 1. Announced / 2. Work In Progress / 3. Upcoming / 4. Released / 5. Stalled / 6. Cancelled / 7. Now Running / 8. Next Change / 9. Pre-production / 10. Post-production / 11. Unreleased
	movie_status := map[int]string{1: "Announced", 2: "Work In Progress", 3: "Upcoming", 4: "Released", 5: "Stalled", 6: "Cancelled", 7: "Now Running", 8: "Next Change", 9: "Pre-production", 10: "Post-production", 11: "Unreleased"}

	out, status := movie_status[int(num)]
	if status {
		return out
	}
	return "Not Available"
}

func GetSplit(stringvalues, strseparate, strjoin string) string { //ids := "546,545,2171"
	idsslice := strings.Split(stringvalues, strseparate)

	//fmt.Printf("idsslice : %v", idsslice)
	idsstr := strings.Join(idsslice, strjoin)
	//fmt.Printf("idsstr : %v", idsstr)
	return idsstr
}

// GetSliceInt64 slice of int64 for given string comma separated
func GetSliceInt64(infovalueidsstr string) []int64 {
	idstringSlice := strings.Split(infovalueidsstr, ",")
	idsslice := []int64{}
	for _, v := range idstringSlice {
		//log.Println(i, v) //log.Printf("Value %d of type %T", v, v)
		n, err := strconv.ParseInt(v, 10, 64)
		if err == nil {
			log.Printf("%d of type %T", n, n)
		}
		idsslice = append(idsslice, n)
	}
	return idsslice
}

func GetBucketNumberFileName(id, uniqueIdName string) (string, error) {
	i, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		log.Printf("Convert failed id : %v, error: %v", id, err)
		return "", err
	}

	bucket := i % 10000
	bucketName := fmt.Sprintf("%v/%v%s", bucket, uniqueIdName, ".json")
	return bucketName, nil
}
