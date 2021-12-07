package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

type Program struct {
	ID          types.String `tfsdk:"id"`
	LastUpdated types.String `tfsdk:"last_updated"`
	Name        types.String `tfsdk:"name"`
}
