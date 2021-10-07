package generator

import (
	"context"
	"fmt"
	"math/rand"
	"strings"

	"github.com/pretcat/ugc_test_task/random"

	"github.com/pretcat/ugc_test_task/models"
	companrepos "github.com/pretcat/ugc_test_task/repositories/companies"
)

func (gen *Generator) initCompanyRepository() (err error) {
	gen.companyRepos, err = companrepos.New(gen.pgClient, gen.categoryRepos)
	if err != nil {
		return err
	}
	return nil
}

func (gen Generator) GenerateCompanies(n int) error {
	countBuildings, err := gen.buildingRepos.Select(context.Background()).Count()
	if err != nil {
		return err
	}
	countCategories, err := gen.categoryRepos.Select(context.Background()).Count()
	if err != nil {
		return err
	}
	for i := 0; i < n; i++ {
		company := models.NewCompany()
		company.Name = genCompanyName()
		company.PhoneNumbers = genPhoneNumbers(rand.Intn(2) + 1)
		building, err := gen.getRandomBuilding(countBuildings)
		if err != nil {
			return err
		}
		company.BuildingId = building.Id
		categoryIds := make([]string, 0)
		err = gen.getRandomCategories(countCategories, rand.Intn(4)+1, func(category models.Category) {
			categoryIds = append(categoryIds, category.Id)
		})
		if err != nil {
			return err
		}
		fmt.Printf("Gen companies: %d/%d\n", i+1, n)
		if _, err := gen.companyRepos.Insert(context.Background(), company, categoryIds); err != nil {
			fmt.Println("Category ids: ", categoryIds)
			return err
		}
	}
	return nil
}

func genPhoneNumbers(n int) []string {
	phoneNumbers := make([]string, 0, n)
	for i := 0; i <= n; i++ {
		phoneNumber := "+7" + random.Numbers(10)
		phoneNumbers = append(phoneNumbers, phoneNumber)
	}
	return phoneNumbers
}

func genCompanyName() string {
	name := random.Letters(rand.Intn(15) + 4)
	name = strings.ToLower(name)
	name = strings.Title(name)
	return name
}
