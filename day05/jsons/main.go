package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type response1 struct {
	Page int
	Fruit []string
}

type response2 struct {
	Page int `json:"page"`
	Fruits []string `json:"fruits"`
}


func main(){
	bolB,_ := json.Marshal(true)
	fmt.Println(string(bolB))
	intB,_ := json.Marshal(1)
	fmt.Println(string(intB))
	floatB,_ := json.Marshal(1.01)
	fmt.Println(string(floatB))
	stringB,_ := json.Marshal("what")
	fmt.Println(string(stringB))
	sliceB,_ := json.Marshal([]string{"abc","def"})
	fmt.Println(string(sliceB))
	mapB,_ := json.Marshal(map[string]string{"abc":"1","def":"2"})
	fmt.Println(string(mapB))
	respons1B,_ := json.Marshal(&response1{2,[]string{"peach","orange"}})
	fmt.Println(string(respons1B))
	respons2B,_ := json.Marshal(&response2{2,[]string{"peach","orange"}})
	fmt.Println(string(respons2B))

	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)
	var dat map[string]interface{}
	if err := json.Unmarshal(byt,&dat);err!=nil{
		panic(err)
	}
	fmt.Println(dat)
	num:=dat["num"].(float64)
	fmt.Printf("%T,%v\n",num,num)
	strs := dat["strs"].([]interface{})
	strs1 := strs[0].(string)
	fmt.Printf("%T,%v\n",strs,strs)
	fmt.Printf("%T,%v\n",strs1,strs1)

	str := `{"page": 1, "fruits": ["apple", "peach"]}`
	res := response2{}
	json.Unmarshal([]byte(str),&res)
	fmt.Println(res)
	fmt.Println(res.Fruits[0])

	enc := json.NewEncoder(os.Stdout)
	enc.Encode(map[string]int{"apple": 5, "lettuce": 7})
}
