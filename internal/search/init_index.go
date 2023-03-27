package search

import (
	"fmt"

	"github.com/quanxiang-cloud/initdb/pkg/toolkits"
)

type SearchConf struct {
	Host string `envconfig:"ESHOST"`
}

func (s *SearchConf) InitSearch(index string) error {
	userIndex := `
	{
        "settings": {
        "index": {
          "number_of_shards": 3,  
          "number_of_replicas": 2 
        }
      }
    }`
	URL := fmt.Sprintf("%s/%s?pretty", s.Host, index)
	return toolkits.DoPUT(URL, userIndex)
}
