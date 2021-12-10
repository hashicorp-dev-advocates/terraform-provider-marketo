package provider

import (
	"context"
	"os"

	"github.com/eveld/terraform-provider-marketo/marketo"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func New() tfsdk.Provider {
	return &provider{}
}

type provider struct {
	configured bool
	client     *marketo.Client
}

func (p *provider) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Attributes: map[string]tfsdk.Attribute{
			"id": {
				Type:     types.StringType,
				Required: true,
			},
			"secret": {
				Type:     types.StringType,
				Required: true,
			},
			"endpoint": {
				Type:     types.StringType,
				Required: true,
			},
		},
	}, nil
}

type providerData struct {
	Endpoint types.String `tfsdk:"endpoint"`
	ID       types.String `tfsdk:"id"`
	Secret   types.String `tfsdk:"secret"`
}

func (p *provider) Configure(ctx context.Context, req tfsdk.ConfigureProviderRequest, resp *tfsdk.ConfigureProviderResponse) {
	var config providerData

	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	if config.Endpoint.Unknown {
		resp.Diagnostics.AddWarning(
			"Unable to create client",
			"Cannot use unknown value as endpoint",
		)
		return
	}

	var endpoint string
	if config.Endpoint.Null {
		endpoint = os.Getenv("MARKETO_ENDPOINT")
	} else {
		endpoint = config.Endpoint.Value
	}

	if endpoint == "" {
		resp.Diagnostics.AddError(
			"Unable to create client",
			"Endpoint cannot be an empty string",
		)
		return
	}

	if config.ID.Unknown {
		resp.Diagnostics.AddWarning(
			"Unable to create client",
			"Cannot use unknown value as id",
		)
		return
	}

	var id string
	if config.ID.Null {
		id = os.Getenv("MARKETO_ID")
	} else {
		id = config.ID.Value
	}

	if id == "" {
		resp.Diagnostics.AddError(
			"Unable to create client",
			"ID cannot be an empty string",
		)
		return
	}

	if config.Secret.Unknown {
		resp.Diagnostics.AddWarning(
			"Unable to create client",
			"Cannot use unknown value as secret",
		)
		return
	}

	var secret string
	if config.Secret.Null {
		secret = os.Getenv("MARKETO_SECRET")
	} else {
		secret = config.Secret.Value
	}

	if secret == "" {
		resp.Diagnostics.AddError(
			"Unable to create client",
			"Secret cannot be an empty string",
		)
		return
	}

	client, err := marketo.NewClient(endpoint, id, secret)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to create client",
			"Unable to create marketo client:\n\n"+err.Error(),
		)
		return
	}

	p.client = client
	p.configured = true
}

func (p *provider) GetResources(_ context.Context) (map[string]tfsdk.ResourceType, diag.Diagnostics) {
	return map[string]tfsdk.ResourceType{
		"marketo_program":        resourceProgramType{},
		"marketo_folder":         resourceFolderType{},
		"marketo_email":          resourceEmailType{},
		"marketo_email_template": resourceEmailTemplateType{},
		"marketo_smart_campaign": resourceSmartCampaignType{},
		"marketo_smart_list":     resourceSmartListType{},
	}, nil
}

func (p *provider) GetDataSources(_ context.Context) (map[string]tfsdk.DataSourceType, diag.Diagnostics) {
	return map[string]tfsdk.DataSourceType{
		"marketo_channel":    dataSourceChannelType{},
		"marketo_smart_list": dataSourceSmartListType{},
	}, nil
}
