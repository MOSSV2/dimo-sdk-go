package test

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/MOSSV2/dimo-sdk-go/sdk"
	"github.com/weaviate/weaviate-go-client/v4/weaviate"
	"github.com/weaviate/weaviate/entities/models"
)

const classname = "test"
const tenantname = "tenant6"
const weaviatehost = "localhost:8081"
const hubhost = "http://localhost:8080"

func GetSchema() error {
	cfg := weaviate.Config{
		Host:   weaviatehost,
		Scheme: "http",
	}
	client, err := weaviate.NewClient(cfg)
	if err != nil {
		return err
	}

	ctx := context.TODO()
	class := &models.Class{
		Class: classname,
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
		return err
	}

	tenants, err := client.Schema().TenantsGetter().
		WithClassName(classname).
		Do(ctx)
	if err != nil {
		return err
	}

	fmt.Printf("%v\n", tenants)

	schema, err := client.Schema().Getter().Do(context.Background())
	if err != nil {
		return err
	}
	fmt.Printf("%v\n", schema)
	return nil
}

func TestGet(t *testing.T) {
	err := GetSchema()
	t.Log(err)
}

func TestCreate(t *testing.T) {
	cfg := weaviate.Config{
		Host:   weaviatehost,
		Scheme: "http",
	}
	client, err := weaviate.NewClient(cfg)
	if err != nil {
		t.Fatal(err)
	}
	w, err := client.Data().Creator().
		WithClassName(classname).
		WithProperties(map[string]interface{}{
			"question":    "This vector DB is OSS and supports automatic property type inference on import",
			"answer":      "Weaviate1", // schema properties can be omitted
			"newProperty": 11245,       // will be automatically added as a number property
		}).
		WithTenant(tenantname).
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

	time.Sleep(5 * time.Second)

	resb, err := sdk.DownloadHubData(hubhost, w.Object.ID.String(), w.Object.Tenant)
	if err != nil {
		t.Fatal(err)
	}

	res := new(models.Object)
	err = res.UnmarshalBinary(resb)
	if err != nil {
		t.Fatal(err)
	}

	b, err = json.MarshalIndent(res, "", "  ")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(b))
}
