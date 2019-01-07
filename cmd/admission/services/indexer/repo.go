package indexer

import (
	"context"
	"log"

	"github.com/Shikanime/unicampus/pkg/admission"
	"github.com/olivere/elastic"
)

func NewClient() Repo {
	conn, err := elastic.NewClient(
		elastic.SetURL("http://localhost:9200"),
		elastic.SetSniff(false),
	)
	if err != nil {
		log.Fatalf("failed to connect indexer database: %v", err)
	}

	// bulk := conn.Bulk().
	// 	Index("schools").
	// 	Type("_doc")
	// bulk.Add(elastic.NewBulkIndexRequest().Id("1").Doc(&School{ID: "1", Name: "ETNA", Description: "Desc"}))

	// _, err = bulk.Do(context.Background())
	// if err != nil {
	// 	log.Fatalf("failed to connect indexer database: %v", err)
	// }

	return Repo{
		conn: conn,
	}
}

type Repo struct {
	conn *elastic.Client
}

func (r *Repo) SearchSchools(school *admission.School) ([]*admission.School, error) {
	return nil, nil
}

func (r *Repo) SearchSchoolsByQuery(query string) ([]*admission.School, error) {
	queryBuilder := elastic.NewMultiMatchQuery(query, "name", "description").
		Fuzziness("2").
		MinimumShouldMatch("2")

	result, err := r.conn.Search().
		Index("schools").
		Query(queryBuilder).
		Do(context.Background())
	if err != nil {
		return nil, err
	}

	return newSchoolsIndexerToDomain(result.Hits.Hits), nil
}

func (r *Repo) Close() {
}