package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	es "github.com/elastic/go-elasticsearch/v6"
	"strings"
)

func getClient() *es.Client {
	client, _ := es.NewClient(es.Config{
		Addresses: []string{"http://localhost:9200"},
	})
	return client
}

func createIndex() {
	client := getClient()
	response, err := client.Indices.Create("test", client.Indices.Create.WithBody(strings.NewReader(`
	{
		"settings": {
			"number_of_shards": 3
		},
		"mappings": {
			"default": {
				"properties": {
					"name": {
						"type": "keyword"
					},
					"age": {
						"type": "integer"
					},
					"desc": {
						"type": "text"
					}
				}
			}
		}
	}
	`)))

	println(response.String())
	if err != nil {
		println(err)
	}
}

func getIndex() string {
	client := getClient()
	response, _ := client.Indices.Get([]string{"test"})
	return response.String()
}

// 都定义了json结构，因为ES请求使用的是json格式，在发送ES请求的时候，会自动转换成json格式。
type TestDocument struct {
	Name string		`json:"name"`
	Age int			`json:"age"`
	Desc string		`json:"desc"`
}

func insertDocument() {
	// 创建文档结构
	doc := TestDocument{
		Name: "test",
		Age: 10,
		Desc: "test document",
	}

	// 使用client创建一个新的文档
	client := getClient()
	body := &bytes.Buffer{}
	json.NewEncoder(body).Encode(doc)
	response, _ := client.Create("test", "100", body, client.Create.WithDocumentType("default"))
	println(response.String())
}

func getDocument(id string) map[string]interface{} {
	client := getClient()
	response, _ := client.Get("test", id, client.Get.WithDocumentType("default"))
	println(response.String())
	var doc map[string]interface{}
	if err := json.NewDecoder(response.Body).Decode(&doc); err != nil {
		println(err.Error())
	}
	return doc
}

func searchDocument() {
	client := getClient()
	body := &bytes.Buffer{}
	body.WriteString(`
		{
		  "query": {
			"match_all": {
			}
		  }
		}
	`)
	response, _ := client.Search(client.Search.WithIndex("test"),
		client.Search.WithBody(body),
		client.Search.WithDocumentType("default"))
	//var docs []map[string]interface{}
	fmt.Println(response.String())
}

func main() {
	//createIndex()
	//index := getIndex()
	//println(index)

	//insertDocument()

	doc := getDocument("100")
	fmt.Println(doc)

	searchDocument()
}
