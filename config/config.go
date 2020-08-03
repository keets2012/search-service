package config

import (
	"context"
	"fmt"
	"github.com/go-kit/kit/log"
	"gopkg.in/olivere/elastic.v5"
	plog "log"
	"os"
)

var EsClient *elastic.Client

var host = "http://localhost:9200/"
var Logger log.Logger

func init() {
	Logger = log.NewLogfmtLogger(os.Stdin)
	Logger = log.With(Logger, "ts", log.DefaultTimestampUTC)
	Logger = log.With(Logger, "caller", log.DefaultCaller)
	//es 配置
	errorlog := plog.New(os.Stdout, "APP", plog.LstdFlags)
	var err error

	EsClient, err = elastic.NewClient(elastic.SetErrorLog(errorlog), elastic.SetURL(host),
		elastic.SetSniff(true))
	if err != nil {
		panic(err)
	}
	info, code, err := EsClient.Ping(host).Do(context.Background())
	if err != nil {
		panic(err)
	}

	fmt.Printf("Elasticsearch returned with code %d and version %s\n", code, info.Version.Number)

	esversion, err := EsClient.ElasticsearchVersion(host)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Elasticsearch version %s\n", esversion)
}
