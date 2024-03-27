package utils

import "go.mongodb.org/mongo-driver/bson/primitive"

func SomeElementInSlice(slice1, slice2 []string) bool {
	for _, s1 := range slice1 {
		for _, s2 := range slice2 {
			if s1 == s2 {
				return true
			}
		}
	}
	return false
}

func ConvertPrimitiveToSlice(arr primitive.A) []string {
	var result []string
	for _, value := range arr {
		result = append(result, value.(string))
	}
	return result
}

func ConvertInterfacesToSlice(arr []interface{}) []string {
	var result []string
	for _, value := range arr {
		result = append(result, value.(string))
	}
	return result
}
