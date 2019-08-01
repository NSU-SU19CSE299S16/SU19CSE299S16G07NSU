package company_details

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func ReadJson(filename string) map[string]interface{}{
	plan, er := ioutil.ReadFile(filename)
	if er!= nil{
		fmt.Println(er)
	}
	var data map[string]interface{}
	err := json.Unmarshal(plan, &data)
	if err != nil {
		panic(err)
	}
	return data

}

func CompanyDetails() []string {
	var company_names []string
	data := ReadJson("Database_Handler/Company Details.json")
	for key,_ := range(data) {
		company_names = append(company_names,key)
	}
	return company_names
}