package terraform

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/wernerstrydom/conversion"
)

type MyResourceModel struct {
	Name        types.String `tfsdk:"name"`
	Description types.String `tfsdk:"description"`
}

type ApiModel struct {
	Name        *string `json:"name"`
	Description *string `json:"description"`
}

func Example_terraformResource() {
	name := "name1"
	description := "description1"

	var model MyResourceModel
	apiModel := ApiModel{
		Name:        &name,
		Description: &description,
	}

	var converters = Converters()
	cfg := conversion.NewConfiguration()
	cfg.RegisterTypeConverters(converters)
	mapper := cfg.BuildMapper()

	err := mapper.Map(apiModel, &model)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(model.Name.ValueString())
	fmt.Println(model.Description.ValueString())
	// Output:
	// name1
	// description1
}
