package test

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/weaviate/weaviate-go-client/v4/weaviate"
	"github.com/weaviate/weaviate/entities/models"
)

func GetSchema() {
	cfg := weaviate.Config{
		Host:   "localhost:8080",
		Scheme: "http",
	}
	client, err := weaviate.NewClient(cfg)
	if err != nil {
		panic(err)
	}

	ctx := context.TODO()
	class := &models.Class{
		Class: "JeopardyQuestion1",
		Properties: []*models.Property{
			{Name: "textProp", DataType: []string{"text"}},
		},
		MultiTenancyConfig: &models.MultiTenancyConfig{
			Enabled:            true,
			AutoTenantCreation: true,
		},
	}
	err = client.Schema().ClassCreator().WithClass(class).Do(ctx)
	if err != nil {
		fmt.Println(err)
	}

	tenants, err := client.Schema().TenantsGetter().
		WithClassName("JeopardyQuestion1").
		Do(ctx)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%v\n", tenants)

	schema, err := client.Schema().Getter().Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", schema)
}

func TestGet(t *testing.T) {
	GetSchema()
}

func TestCreate(t *testing.T) {
	cfg := weaviate.Config{
		Host:   "localhost:8080",
		Scheme: "http",
	}
	client, err := weaviate.NewClient(cfg)
	if err != nil {
		t.Fatal(err)
	}
	w, err := client.Data().Creator().
		WithClassName("JeopardyQuestion1").
		WithProperties(map[string]interface{}{
			"question":    "This vector DB is OSS and supports automatic property type inference on import",
			"answer":      "Weaviate", // schema properties can be omitted
			"newProperty": 123,        // will be automatically added as a number property
		}).
		WithTenant("test1").
		Do(context.TODO())
	if err != nil {
		t.Fatal(err)
	}

	// the returned value is a wrapped object
	b, err := json.MarshalIndent(w.Object, "", "  ")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(b))
}
