package stystruct

import (
	"encoding/json"
	"fmt"
)

type Stu struct {
	Name string `json:"name""`
	Age int `json:"age"`
}

func ConvertStruct2JsonStr(){
	var stu  = Stu{"住",123}
	jsonStr,_ := json.Marshal(stu)
	fmt.Println("jsonStr:",string(jsonStr))
}