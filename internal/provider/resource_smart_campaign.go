package provider

import (
	"context"
	"time"

	"github.com/eveld/terraform-provider-marketo/marketo"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

type resourceSmartCampaignType struct{}

func (r resourceSmartCampaignType) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Attributes: map[string]tfsdk.Attribute{
			"id": {
				Type:     types.StringType,
				Computed: true,
			},
			"last_updated": {
				Type:     types.StringType,
				Computed: true,
			},
			"name": {
				Type:     types.StringType,
				Required: true,
			},
			"description": {
				Type:     types.StringType,
				Optional: true,
			},
			"folder": {
				Type:     types.StringType,
				Optional: true,
			},
			"program": {
				Type:     types.StringType,
				Optional: true,
			},
			"schedule": {
				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{
					"run_at": {
						Type:     types.StringType,
						Required: true,
					},
					"tokens": {
						Type: types.MapType{
							ElemType: types.StringType,
						},
						Optional: true,
					},
				}),
				Optional: true,
			},
		},
	}, nil
}

func (r resourceSmartCampaignType) NewResource(_ context.Context, p tfsdk.Provider) (tfsdk.Resource, diag.Diagnostics) {
	return resourceSmartCampaign{
		p: *(p.(*provider)),
	}, nil
}

type resourceSmartCampaign struct {
	p provider
}

func (r resourceSmartCampaign) Create(ctx context.Context, req tfsdk.CreateResourceRequest, resp *tfsdk.CreateResourceResponse) {
	if !r.p.configured {
		resp.Diagnostics.AddError(
			"Provider not configured",
			"The provider hasn't been configured before apply, likely because it depends on an unknown value from another resource. This leads to weird stuff happening, so we'd prefer if you didn't do that. Thanks!",
		)
		return
	}

	var plan SmartCampaign
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	smartCampaign := marketo.SmartCampaign{
		Name: plan.Name.Value,
	}

	result, err := r.p.client.CreateSmartCampaign(smartCampaign)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating smartCampaign",
			"Could not create smartCampaign, unexpected error: "+err.Error(),
		)
		return
	}

	// update more fields once they can differ between result and plan.

	plan.ID = types.String{Value: result.ID}
	plan.LastUpdated = types.String{Value: string(time.Now().Format(time.RFC850))}

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r resourceSmartCampaign) Read(ctx context.Context, req tfsdk.ReadResourceRequest, resp *tfsdk.ReadResourceResponse) {
	var state SmartCampaign
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	smartCampaignID := state.ID.Value
	smartCampaign, err := r.p.client.GetSmartCampaign(smartCampaignID)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error reading smartCampaign",
			"Could not read smartCampaign with ID "+smartCampaignID+": "+err.Error(),
		)
		return
	}

	state.ID = types.String{Value: smartCampaign.ID}
	state.Name = types.String{Value: smartCampaign.Name}

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r resourceSmartCampaign) Update(ctx context.Context, req tfsdk.UpdateResourceRequest, resp *tfsdk.UpdateResourceResponse) {
	var plan SmartCampaign
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state SmartCampaign
	diags = req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	smartCampaign := marketo.SmartCampaign{
		Name: plan.Name.Value,
	}

	smartCampaignID := state.ID.Value
	result, err := r.p.client.UpdateSmartCampaign(smartCampaignID, smartCampaign)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error update order",
			"Could not update orderID "+smartCampaignID+": "+err.Error(),
		)
		return
	}

	// update more fields once they can differ between result and plan.

	plan.ID = types.String{Value: result.ID}
	plan.LastUpdated = types.String{Value: string(time.Now().Format(time.RFC850))}

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r resourceSmartCampaign) Delete(ctx context.Context, req tfsdk.DeleteResourceRequest, resp *tfsdk.DeleteResourceResponse) {
	var state SmartCampaign
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	smartCampaignID := state.ID.Value
	err := r.p.client.DeleteSmartCampaign(smartCampaignID)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error deleting smartCampaign",
			"Could not delete smartCampaign with ID "+smartCampaignID+": "+err.Error(),
		)
		return
	}

	resp.State.RemoveResource(ctx)
}

func (r resourceSmartCampaign) ImportState(ctx context.Context, req tfsdk.ImportResourceStateRequest, resp *tfsdk.ImportResourceStateResponse) {
	tfsdk.ResourceImportStatePassthroughID(ctx, tftypes.NewAttributePath().WithAttributeName("id"), req, resp)
}
