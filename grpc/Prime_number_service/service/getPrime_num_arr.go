package service

import (
	"encoding/json"
	"fmt"
)

//type Data []struct {
//	Num int `json:"num"`
//}

func GetPrime(val string) (string, error) {
	fmt.Println(val)
	var numArr []int
	err := json.Unmarshal([]byte(val), &numArr)
	if err != nil {
		fmt.Println("json trans err :", err)
		return "", err
	}
	m := make([]int, 0)
	for _, v := range numArr {
		flag := false
		for i := 2; i < v; i++ {
			if v%i == 0 {
				flag = true
				break
			}
		}
		fmt.Println(v, flag)
		if !flag && v != 1 {
			m = append(m, v)
		}
	}
	marshal, err := json.Marshal(m)
	if err != nil {
		fmt.Println("mao to json err :", marshal)
		return "", err
	}
	return string(marshal), nil
}
