package hw03frequencyanalysis

import (
	"regexp"
	"sort"
	"strings"
)

type KeyValue struct {
	Key   string
	Value int
}

const Length = 10

var r = regexp.MustCompile(`[\p{L}\d_-]+`)

func wordNormalizer(el string) string {
	result := strings.ToLower(r.FindString(el))
	return result
}

func textToMap(str string) map[string]int {
	result := make(map[string]int)
	for _, el := range strings.Fields(str) {
		result[wordNormalizer(el)]++
	}
	return result
}

func sortedSlice(resultMap map[string]int) []KeyValue {
	sortedResult := make([]KeyValue, 0, len(resultMap))
	for key, value := range resultMap {
		if key != "" && key != "-" {
			sortedResult = append(sortedResult, KeyValue{key, value})
		}
	}
	sort.Slice(sortedResult, func(i, j int) bool {
		if sortedResult[i].Value == sortedResult[j].Value {
			return sortedResult[i].Key < sortedResult[j].Key
		}
		return sortedResult[i].Value > sortedResult[j].Value
	})
	if len(sortedResult) > 10 {
		return sortedResult[:10]
	}
	return sortedResult
}

func Top10(str string) []string {
	resultMap := textToMap(str)
	sortedResult := sortedSlice(resultMap)
	var result []string
	for i := 0; i < len(sortedResult); i++ {
		result = append(result, sortedResult[i].Key)
	}
	return result
}
