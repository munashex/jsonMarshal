package main 

import (
	"fmt" 
	 "encoding/json"
)


type StudentJson struct {
	Name string `json:"name"`
	Age int      `json:"age"`
} 

func main() {
	singleStudent:= StudentJson{"munashe", 21}
	result , err:= json.Marshal(singleStudent) 
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(result))
}