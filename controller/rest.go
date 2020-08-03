package controller

import (
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"gopkg.in/olivere/elastic.v5"
	"net/http"
	"searcher/config"
	"searcher/domain"
)

func GetESRes(c *gin.Context) {
	var esRes domain.EsResult
	keyword := c.Query("keyword")
	config.Logger.Log("keyword is ", keyword)
	var q_product, q_live, q_member elastic.Query

	q_product = elastic.NewFuzzyQuery("name", keyword)

	q_live = elastic.NewFuzzyQuery("live_title", keyword)
	q_member = elastic.NewFuzzyQuery("username", keyword)
	product, err := config.EsClient.Search().Index("msy_product1").Type("pms_product").Query(q_product).Do(context.Background())
	if err != nil {
		panic(err)
	}
	live, _ := config.EsClient.Search().Index("msy_live1").Type("live").Query(q_live).Do(context.Background())
	member, _ := config.EsClient.Search().Index("msy_member1").Type("ums_member").Query(q_member).Do(context.Background())

	if product.Hits.TotalHits > 0 {
		var items []string
		esRes.ProductNum = product.Hits.TotalHits
		var i int64
		var item string
		b, _ := json.Marshal(product.Hits.Hits[i].Source)
		item = string(b)
		config.Logger.Log("product item :", item)
		items = append(items, item)
		esRes.Product = items
	}
	if live.Hits.TotalHits > 0 {
		var items []string
		esRes.LiveNum = live.Hits.TotalHits
		var i int64
		var item string
		b, _ := json.Marshal(live.Hits.Hits[i].Source)
		item = string(b)
		config.Logger.Log("live item :", item)
		items = append(items, item)
		esRes.Live = items

	}
	if member.Hits.TotalHits > 0 {
		var items []string
		esRes.MemberNum = member.Hits.TotalHits
		var i int64
		var item string
		b, _ := json.Marshal(member.Hits.Hits[i].Source)
		item = string(b)
		config.Logger.Log("member item :", item)
		items = append(items, item)
		esRes.Member = items
	}
	da, _ := json.MarshalIndent(esRes, "", " ")
	config.Logger.Log("esRes is ", string(da))
	c.JSON(200, esRes)
}
func PostApi(c *gin.Context) {
	c.String(http.StatusOK, "ok")
}
