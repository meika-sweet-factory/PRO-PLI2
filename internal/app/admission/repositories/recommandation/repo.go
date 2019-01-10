package recommandation

import (
	"github.com/Shikanime/unicampus/internal/pkg/services"
	golangNeo4jBoltDriver "github.com/johnnadratowski/golang-neo4j-bolt-driver"
)

func NewRepository(service *services.Neo4jService) *Repo {
	return &Repo{
		conn: service.Driver(),
	}
}

type Repo struct {
	conn golangNeo4jBoltDriver.Conn
}
