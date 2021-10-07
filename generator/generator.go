package generator

import (
	"context"

	"github.com/pretcat/ugc_test_task/pg"
	buildrepos "github.com/pretcat/ugc_test_task/repositories/buildings"
	categrepos "github.com/pretcat/ugc_test_task/repositories/categories"
	companrepos "github.com/pretcat/ugc_test_task/repositories/companies"
)

type Generator struct {
	conf          Config
	buildingRepos buildrepos.Repository
	categoryRepos categrepos.Repository
	companyRepos  companrepos.Repository
	pgClient      pg.Client
}

func New(conf Config) (gen Generator, _ error) {
	gen.conf = conf
	if err := gen.initPgClient(); err != nil {
		return Generator{}, err
	}
	if err := gen.initBuildingRepository(); err != nil {
		return Generator{}, err
	}
	if err := gen.initCategoryRepository(); err != nil {
		return Generator{}, err
	}
	if err := gen.initCompanyRepository(); err != nil {
		return Generator{}, err
	}
	return gen, nil
}

func (gen *Generator) initPgClient() (err error) {
	gen.pgClient, err = pg.Connect(context.Background(), gen.conf.GetPgConfig())
	if err != nil {
		return err
	}
	return nil
}
