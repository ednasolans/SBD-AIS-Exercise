package mapred

import (
	"regexp"
	"strings"
	"sync"
)

type MapReduce struct {
}

// todo implement mapreduce
func (mr *MapReduce) wordCountMapper(text string) []KeyValue {
	re := regexp.MustCompile("[a-zA-Z]+")
	words := re.FindAllString(strings.ToLower(text), -1)

	var kvs []KeyValue
	for _, word := range words {
		if word != "" {
			kvs = append(kvs, KeyValue{Key: word, Value: 1})
		}
	}
	return kvs
}

func (mr MapReduce) wordCountReducer(key string, values []int) KeyValue {
	sum := 0
	for _, value := range values {
		sum += value
	}
	return KeyValue{Key: key, Value: sum}
}

func (mr MapReduce) Run(input []string) map[string]int {
	mapChan := make(chan []KeyValue, len(input))
	var wg sync.WaitGroup

	for _, line := range input {
		wg.Add(1)
		go func(text string) {
			defer wg.Done()
			mapChan <- mr.wordCountMapper(text)
		}(line)
	}

	wg.Wait()
	close(mapChan)

	intermediate := make(map[string][]int)
	for kkvList := range mapChan {
		for _, kv := range kkvList {
			intermediate[kv.Key] = append(intermediate[kv.Key], kv.Value)
		}
	}

	result := make(map[string]int)
	for key, values := range intermediate {
		reducedKV := mr.wordCountReducer(key, values)
		result[reducedKV.Key] = reducedKV.Value
	}

	return result
}
