package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

type Channel struct {
	ID          types.String `tfsdk:"id"`
	LastUpdated types.String `tfsdk:"last_updated"`
	Name        types.String `tfsdk:"name"`
}

type Program struct {
	ID          types.String `tfsdk:"id"`
	LastUpdated types.String `tfsdk:"last_updated"`
	Name        types.String `tfsdk:"name"`
}

type Folder struct {
	ID          types.String `tfsdk:"id"`
	LastUpdated types.String `tfsdk:"last_updated"`
	Name        types.String `tfsdk:"name"`
}

type Email struct {
	ID          types.String `tfsdk:"id"`
	LastUpdated types.String `tfsdk:"last_updated"`
	Name        types.String `tfsdk:"name"`
}

type EmailTemplate struct {
	ID          types.String `tfsdk:"id"`
	LastUpdated types.String `tfsdk:"last_updated"`
	Name        types.String `tfsdk:"name"`
}

type SmartCampaign struct {
	ID          types.String `tfsdk:"id"`
	LastUpdated types.String `tfsdk:"last_updated"`
	Name        types.String `tfsdk:"name"`
}

type SmartList struct {
	ID          types.String `tfsdk:"id"`
	LastUpdated types.String `tfsdk:"last_updated"`
	Name        types.String `tfsdk:"name"`
}
