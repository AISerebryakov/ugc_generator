package generator

import (
	buildrepos "github.com/pretcat/ugc_test_task/repositories/buildings"
	categrepos "github.com/pretcat/ugc_test_task/repositories/categories"
)

type Generator struct {
	conf          Config
	buildRepos    buildrepos.Repository
	categoryRepos categrepos.Repository
}

func New(conf Config) (gen Generator, _ error) {
	gen.conf = conf
	//todo: handle error
	if err := gen.initBuildingRepository(); err != nil {
		return Generator{}, err
	}
	if err := gen.initCategoryRepository(); err != nil {
		return Generator{}, err
	}
	return gen, nil
}
