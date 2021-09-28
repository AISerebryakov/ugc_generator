package generator

import (
	"context"
	"fmt"
	"math/rand"
	"strings"

	"github.com/pretcat/ugc_test_task/common/random"

	"github.com/pretcat/ugc_test_task/models"
	categrepos "github.com/pretcat/ugc_test_task/repositories/categories"
)

func (gen *Generator) initCategoryRepository() (err error) {
	gen.categoryRepos, err = categrepos.New(categrepos.NewConfig(gen.conf.GetPgConfig()))
	if err != nil {
		return err
	}
	return nil
}

func (gen Generator) GenerateCategories(num int) error {
	for i := 0; i <= num; i++ {
		category := models.NewCategory()
		category.CreateAt = category.CreateAt - rand.Int63n(1000)
		category.Name = genCategoryName(rand.Intn(10) + 5)
		fmt.Printf("%d/%d\n", i+1, num)
		if err := gen.categoryRepos.Insert(context.Background(), category); err != nil {
			return err
		}
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
