package generator

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"strings"

	"github.com/pretcat/ugc_test_task/random"

	"github.com/pretcat/ugc_test_task/models"
	categrepos "github.com/pretcat/ugc_test_task/repositories/categories"
)

func (gen *Generator) initCategoryRepository() (err error) {
	gen.categoryRepos, err = categrepos.New(gen.pgClient)
	if err != nil {
		return err
	}
	return nil
}

func (gen Generator) GenerateCategories(n int) error {
	for i := 0; i < n; i++ {
		category := models.NewCategory()
		category.CreateAt = category.CreateAt - rand.Int63n(1000)
		category.Name = genCategoryName(rand.Intn(10) + 5)
		fmt.Printf("Gen categories: %d/%d\n", i+1, n)
		if err := gen.categoryRepos.Insert(context.Background(), category); err != nil {
			return err
		}
	}
	return nil
}

func (gen Generator) getRandomCategory(count int) (models.Category, error) {
	category, found, err := gen.categoryRepos.Select(context.Background()).
		Offset(rand.Intn(count-2) + 1).One()
	if !found {
		return models.Category{}, errors.New("not found")
	}
	if err != nil {
		return models.Category{}, err
	}
	return category, nil
}

func (gen Generator) getRandomCategories(count, n int, callback func(models.Category)) error {
	err := gen.categoryRepos.Select(context.Background()).Offset(rand.Intn(count-2) + 1).
		Limit(n).Iter(func(category models.Category) error {
		callback(category)
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func genCategoryName(n int) (name string) {
	for i := 0; i <= n; i++ {
		if len(name) > 0 {
			name = name + "."
		}
		name = name + random.Letters(rand.Intn(10)+3)
		name = strings.ToLower(name)
		name = strings.Title(name)
	}
	return name
}
