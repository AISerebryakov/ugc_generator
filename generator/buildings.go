package generator

import (
	"context"
	"fmt"
	"math/rand"
	"strconv"

	"github.com/pretcat/ugc_test_task/common/random"
	"github.com/pretcat/ugc_test_task/models"
	buildrepos "github.com/pretcat/ugc_test_task/repositories/buildings"
)

func (gen *Generator) initBuildingRepository() (err error) {
	gen.buildRepos, err = buildrepos.New(buildrepos.NewConfig(gen.conf.GetPgConfig()))
	if err != nil {
		return err
	}
	return nil
}

func (gen Generator) GenerateBuildings(num int) error {
	for i := 0; i <= num; i++ {
		building := models.NewBuilding()
		building.CreateAt = building.CreateAt - rand.Int63n(1000)
		building.Address = "Address_" + random.String(10) + "," + strconv.Itoa(i+1)
		building.Location = genLocation()
		fmt.Printf("%d/%d\n", i+1, num)
		if err := gen.buildRepos.Insert(context.Background(), building); err != nil {
			return err
		}
	}
	return nil
}

func genLocation() (loc models.Location) {
	loc.Latitude = randFloat(-90, 90)
	loc.Longitude = randFloat(-180, 180)
	return loc
}

func randFloat(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}
