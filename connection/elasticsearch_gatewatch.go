package connection

import (
	"fmt"

	"github.com/aditya3232/gatewatchApp-services.git/config"
	"github.com/elastic/go-elasticsearch"
)

func ConnectElastic() (*elasticsearch.Client, error) {
	// Create a new Elasticsearch client
	esClient, err := elasticsearch.NewClient(
		elasticsearch.Config{
			Addresses: []string{
				"http://" + config.CONFIG.ES_HOST + ":" + config.CONFIG.ES_PORT,
				// "http://localhost:9200/",

			},
		},
	)
	if err != nil {
		return nil, err
	}

	// Ping the Elasticsearch server to check if it's reachable
	res, err := esClient.Ping()
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	fmt.Println(res)

	return esClient, nil
}

func ElasticSearchGatewatch() *elasticsearch.Client {
	return database.es
}
