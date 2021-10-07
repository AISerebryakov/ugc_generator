package generator

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"strconv"

	"github.com/pretcat/ugc_test_task/models"
	"github.com/pretcat/ugc_test_task/random"
	buildrepos "github.com/pretcat/ugc_test_task/repositories/buildings"
)

func (gen *Generator) initBuildingRepository() (err error) {
	gen.buildingRepos, err = buildrepos.New(gen.pgClient)
	if err != nil {
		return err
	}
	return nil
}

func (gen Generator) GenerateBuildings(num int) error {
	for i := 0; i < num; i++ {
		building := models.NewBuilding()
		building.CreateAt = building.CreateAt - rand.Int63n(1000)
		building.Address = "Address_" + random.Letters(10) + "," + strconv.Itoa(i+1)
		building.Location = genLocation()
		fmt.Printf("Gen buildings: %d/%d\n", i+1, num)
		if err := gen.buildingRepos.Insert(context.Background(), building); err != nil {
			return err
		}
	}
	return nil
}

func (gen Generator) getRandomBuilding(count int) (models.Building, error) {
	building, found, err := gen.buildingRepos.Select(context.Background()).
		Offset(rand.Intn(count-2) + 1).One()
	if !found {
		return models.Building{}, errors.New("not found")
	}
	if err != nil {
		return models.Building{}, err
	}
	return building, nil
}

func genLocation() (loc models.Location) {
	loc.Latitude = randFloat(-90, 90)
	loc.Longitude = randFloat(-180, 180)
	return loc
}

func randFloat(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}
