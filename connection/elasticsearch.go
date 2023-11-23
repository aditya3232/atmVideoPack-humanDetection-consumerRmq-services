package connection

import (
	"fmt"
	"log"

	"github.com/aditya3232/atmVideoPack-humanDetection-consumerRmq-services.git/config"
	esv7 "github.com/elastic/go-elasticsearch/v7"
)

func ConnectElastic() (*esv7.Client, error) {
	// Create a new Elasticsearch client
	esClient, err := esv7.NewClient(
		esv7.Config{
			Addresses: []string{
				"http://" + config.CONFIG.ES_HOST + ":" + config.CONFIG.ES_PORT,
			},
			Username: config.CONFIG.ES_USER, // Replace with your Elasticsearch username
			Password: config.CONFIG.ES_PASS, // Replace with your Elasticsearch password
		},
	)
	if err != nil {
		fmt.Println(err)
		// log.Fatalln(err)
		return nil, err
	}

	// Ping the Elasticsearch server to check if it's reachable
	res, err := esClient.Ping()
	if err != nil {
		// log.Fatalln(err)
		fmt.Println(err)
		return nil, err
	}
	defer res.Body.Close()

	// fmt.Println(res)
	log.Print("ElasticSearch is connected")

	return esClient, nil
}

func ElasticSearch() *esv7.Client {
	return connection.es
}
