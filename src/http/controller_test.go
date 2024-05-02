package http

import (
	"agile-homework/src/app"
	"agile-homework/src/http/routes"
	"agile-homework/src/mocks"
	"agile-homework/src/clients/swapi"
	"agile-homework/src/clients/swapi/models"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func setup(swapiClient app.SwapiClient) (*gin.Engine, *app.ServiceContainer) {
  gin.SetMode(gin.TestMode)
  config := app.GetTestConfig()
  container := app.GetApp(config)
  container.SwapiClient = swapiClient

  gin.SetMode(gin.DebugMode)
  _, e := gin.CreateTestContext(httptest.NewRecorder())

  return e, container
}

func TestCharacterController(t *testing.T) {
  n := "next"
  sc := &mocks.MockSwapiClient{
    ExpectedCharacters: swapi.SwapiResponse[models.Character]{
      Next: &n,
      Results: []models.Character{
        {
          Name: "foo",
          Homeworld: "1/",
          Species: []string{
            "1/",
          },
          Starship: []string{
            "1/",
          },
        },
        {
          Name: "bar",
          Homeworld: "2/",
          Species: []string{
            "1/",
            "2/",
          },
          Starship: []string{},
        },
        {
          Name: "baz",
          Homeworld: "3/",
          Species: []string{},
          Starship: []string{
            "1/",
            "2/",
          },
        },
      },
    },
  }

  router, container := setup(sc)
  routes.Init(router, container)

  req, err := http.NewRequest(http.MethodGet, "/api/characters", nil)
  if err != nil {
      t.Fatalf("error creating request: %v\n", err)
  }
  w := httptest.NewRecorder()
  router.ServeHTTP(w, req)

  if w.Code != http.StatusOK {
    t.Fatalf("asdf: %d\n", w.Code)
  }

}

func TestPlanetController(t *testing.T) {
  sc := &mocks.MockSwapiClient{
    ExpectedPlanets: models.Planet{
      Name: "foo",
      Climate: "chill",
      Population: "1",
    },
  }
  router, container := setup(sc)
  routes.Init(router, container)

  req, err := http.NewRequest(http.MethodGet, "/api/planets/1", nil)
  if err != nil {
      t.Fatalf("error creating request: %v\n", err)
  }
  w := httptest.NewRecorder()
  router.ServeHTTP(w, req)

  if w.Code != http.StatusOK {
    t.Fatalf("asdf: %d\n", w.Code)
  }
}

func TestSpeciesController(t *testing.T) {
  sc := &mocks.MockSwapiClient{
    ExpectedSpecies: models.Species{
      Name: "foo",
      AverageLifespan: "100",
      Language: "golang",
    },
  }
  router, container := setup(sc)
  routes.Init(router, container)

  req, err := http.NewRequest(http.MethodGet, "/api/species/1", nil)
  if err != nil {
      t.Fatalf("error creating request: %v\n", err)
  }
  w := httptest.NewRecorder()
  router.ServeHTTP(w, req)

  if w.Code != http.StatusOK {
    t.Fatalf("asdf: %d\n", w.Code)
  }
}

func TestStarshipsController(t *testing.T) {
  sc := &mocks.MockSwapiClient{
    ExpectedStarships: models.Starship{
      Name: "foo",
      CargoCapacity: "19000mt",
      StarshipClass: "constitution",
    },
  }
  router, container := setup(sc)
  routes.Init(router, container)

  req, err := http.NewRequest(http.MethodGet, "/api/starships/1", nil)
  if err != nil {
      t.Fatalf("error creating request: %v\n", err)
  }
  w := httptest.NewRecorder()
  router.ServeHTTP(w, req)

  if w.Code != http.StatusOK {
    t.Fatalf("asdf: %d\n", w.Code)
  }
}
