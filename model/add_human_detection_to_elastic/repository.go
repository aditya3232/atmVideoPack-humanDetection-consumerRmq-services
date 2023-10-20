package add_human_detection_to_elastic

import (
	"strconv"
	"strings"

	"github.com/elastic/go-elasticsearch"
)

type Repository interface {
	CreateElasticHumanDetection(elasticHumanDetection ElasticHumanDetection) (ElasticHumanDetection, error)
}

type repository struct {
	elasticsearch *elasticsearch.Client
}

func NewRepository(elasticsearch *elasticsearch.Client) *repository {
	return &repository{elasticsearch}
}

func (r *repository) CreateElasticHumanDetection(elasticHumanDetection ElasticHumanDetection) (ElasticHumanDetection, error) {

	// Menggunakan library "github.com/elastic/go-elasticsearch" untuk melakukan operasi penyimpanan
	// Gantilah `indexName` dengan nama index Elasticsearch yang sesuai
	indexName := "human_detection_index"

	// Anda dapat membuat body dokumen yang akan disimpan di Elasticsearch
	// Misalnya, jika Anda ingin menyimpan data deteksi manusia yang diberikan sebagai JSON:
	body := []byte(`{
		"id": "` + elasticHumanDetection.ID + `",
		"tid_id": "` + strconv.Itoa(*elasticHumanDetection.TidID) + `",
		"date_time": "` + elasticHumanDetection.DateTime + `",
		"person": "` + elasticHumanDetection.Person + `",
		"file_name_capture_human_detection": "` + elasticHumanDetection.FileNameCaptureHumanDetection + `"
	}`)

	// Mengirimkan data ke Elasticsearch untuk disimpan
	_, err := r.elasticsearch.Index(indexName, strings.NewReader(string(body)))
	if err != nil {
		return elasticHumanDetection, err
	}

	// Jika operasi berhasil, Anda dapat mengembalikan data yang sama yang Anda terima sebagai argumen.
	return elasticHumanDetection, nil

}
