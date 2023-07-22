package terraform

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-framework-timeouts/resource/timeouts"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stretchr/testify/suite"

	"github.com/wernerstrydom/conversion"
)

type NestedTag struct {
	Color   string     `json:"color,omitempty"`
	Display string     `json:"display,omitempty"`
	ID      int64      `json:"id,omitempty"`
	Name    *string    `json:"name"`
	Slug    *string    `json:"slug"`
	URL     strfmt.URI `json:"url,omitempty"`
}

type Manufacturer struct {
	Created            *strfmt.DateTime `json:"created,omitempty"`
	CustomFields       interface{}      `json:"custom_fields,omitempty"`
	Description        string           `json:"description,omitempty"`
	DevicetypeCount    int64            `json:"devicetype_count,omitempty"`
	Display            string           `json:"display,omitempty"`
	ID                 int64            `json:"id,omitempty"`
	InventoryitemCount int64            `json:"inventoryitem_count,omitempty"`
	LastUpdated        *strfmt.DateTime `json:"last_updated,omitempty"`
	Name               *string          `json:"name"`
	PlatformCount      int64            `json:"platform_count,omitempty"`
	Slug               *string          `json:"slug"`
	Tags               []*NestedTag     `json:"tags,omitempty"`
	URL                strfmt.URI       `json:"url,omitempty"`
}

type DcimManufacturersCreateParams struct {

	// Data.
	Data *Manufacturer

	Context    context.Context
	HTTPClient *http.Client
}

type manufacturerResourceModel struct {
	ID          types.String   `tfsdk:"id"`
	Name        types.String   `tfsdk:"name"`
	Slug        types.String   `tfsdk:"slug"`
	Description types.String   `tfsdk:"description"`
	Timeouts    timeouts.Value `tfsdk:"timeouts"`
}

type NetBoxTests struct {
	suite.Suite
}

func (suite *NetBoxTests) SetupTest() {
	fmt.Printf("SetupTest\n")
}

func (suite *NetBoxTests) TearDownTest() {
	fmt.Printf("TearDownTest\n")
}

func TestNetBox(t *testing.T) {
	suite.Run(t, new(NetBoxTests))
}

func (suite *NetBoxTests) TestMapFromModelToResource() {
	var err error

	model := manufacturerResourceModel{
		ID:          types.StringValue("1"),
		Name:        types.StringValue("name1"),
		Slug:        types.StringValue("slug1"),
		Description: types.StringValue("description1"),
	}

	params := DcimManufacturersCreateParams{
		Data:    nil,
		Context: context.Background(),
	}

	cfg := conversion.NewConfiguration()
	cfg.RegisterTypeConverters(Converters())
	mapper := cfg.BuildMapper()

	err = mapper.Map(model, &params.Data)
	suite.Nil(err)

	suite.Equal("1", model.ID.ValueString())
	suite.Equal("name1", model.Name.ValueString())
	suite.Equal("slug1", model.Slug.ValueString())
	suite.Equal("description1", model.Description.ValueString())
}
