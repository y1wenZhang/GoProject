package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type Movie struct {
	Title string `json:"title"`
	Price int `json:"price"`
	Actors []string `json:"actors"`
}

func main() {
	movie := Movie{"喜剧之王", 200, []string{"xingye", "zhangbozhi"}}
	// 结构 --> json string
	jsonStr, err := json.Marshal(movie)
	if err != nil {
		fmt.Println("-----")
	}
	fmt.Println(reflect.TypeOf(jsonStr))
	fmt.Printf("json str %s\n", jsonStr)

	// json string --> struct
	myMovie := Movie{}
	err = json.Unmarshal(jsonStr, &myMovie)
	if err != nil {
		fmt.Println("json unmarshal error\n")
	}
	fmt.Printf("%v\n", myMovie)

}
