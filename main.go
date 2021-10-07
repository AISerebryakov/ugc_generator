package main

import (
	"fmt"

	"github.com/pretcat/ugc_generator/generator"
)

func main() {
	conf := generator.Config{
		Host:     "localhost",
		Port:     "5432",
		DbName:   "ugc_test_task",
		User:     "postgres",
		Password: "7823",
	}
	gen, err := generator.New(conf)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	//if err := gen.GenerateBuildings(650000); err != nil {
	//	fmt.Println(err.Error())
	//	return
	//}
	//if err := gen.GenerateCategories(40000); err != nil {
	//	fmt.Println(err.Error())
	//	return
	//}
	if err := gen.GenerateCompanies(200000); err != nil {
		fmt.Println(err.Error())
		return
	}
}
