package internal

import (
	esv7 "github.com/elastic/go-elasticsearch/v7"

	"github.com/ryanpfalz/go-openapi/internal"
	"github.com/ryanpfalz/go-openapi/internal/envvar"
)

// NewElasticSearch instantiates the ElasticSearch client using configuration defined in environment variables.
func NewElasticSearch(conf *envvar.Configuration) (es *esv7.Client, err error) {
	es, err = esv7.NewDefaultClient()
	if err != nil {
		return nil, internal.WrapErrorf(err, internal.ErrorCodeUnknown, "elasticsearch.Open")
	}

	res, err := es.Info()
	if err != nil {
		return nil, internal.WrapErrorf(err, internal.ErrorCodeUnknown, "es.Info")
	}

	defer func() {
		err = res.Body.Close()
	}()

	return es, nil
}