package connection

import (
	"log"

	"github.com/aditya3232/gatewatchApp-services.git/config"
	"github.com/elastic/go-elasticsearch"
)

func ConnectElastic() (*elasticsearch.Client, error) {
	// Create a new Elasticsearch client
	esClient, err := elasticsearch.NewClient(
		elasticsearch.Config{
			Addresses: []string{
				"http://" + config.CONFIG.ES_HOST + ":" + config.CONFIG.ES_PORT,
			},
		},
	)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	// Ping the Elasticsearch server to check if it's reachable
	res, err := esClient.Ping()
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}
	defer res.Body.Close()

	// fmt.Println(res)
	log.Print("ElasticSearch is connected")

	return esClient, nil
}

func ElasticSearch() *elasticsearch.Client {
	return connection.es
}
