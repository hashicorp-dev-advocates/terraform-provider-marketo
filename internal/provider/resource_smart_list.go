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

type resourceSmartListType struct{}

func (r resourceSmartListType) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
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
			"source": {
				Type:     types.StringType,
				Required: true,
			},
		},
	}, nil
}

func (r resourceSmartListType) NewResource(_ context.Context, p tfsdk.Provider) (tfsdk.Resource, diag.Diagnostics) {
	return resourceSmartList{
		p: *(p.(*provider)),
	}, nil
}

type resourceSmartList struct {
	p provider
}

func (r resourceSmartList) Create(ctx context.Context, req tfsdk.CreateResourceRequest, resp *tfsdk.CreateResourceResponse) {
	if !r.p.configured {
		resp.Diagnostics.AddError(
			"Provider not configured",
			"The provider hasn't been configured before apply, likely because it depends on an unknown value from another resource. This leads to weird stuff happening, so we'd prefer if you didn't do that. Thanks!",
		)
		return
	}

	var plan SmartList
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	smartList := marketo.SmartList{
		Name: plan.Name.Value,
	}

	result, err := r.p.client.CreateSmartList(smartList)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating smartList",
			"Could not create smartList, unexpected error: "+err.Error(),
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

func (r resourceSmartList) Read(ctx context.Context, req tfsdk.ReadResourceRequest, resp *tfsdk.ReadResourceResponse) {
	var state SmartList
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	smartListID := state.ID.Value
	smartList, err := r.p.client.GetSmartList(smartListID)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error reading smartList",
			"Could not read smartList with ID "+smartListID+": "+err.Error(),
		)
		return
	}

	state.ID = types.String{Value: smartList.ID}
	state.Name = types.String{Value: smartList.Name}

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r resourceSmartList) Update(ctx context.Context, req tfsdk.UpdateResourceRequest, resp *tfsdk.UpdateResourceResponse) {
	var plan SmartList
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state SmartList
	diags = req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	smartList := marketo.SmartList{
		Name: plan.Name.Value,
	}

	smartListID := state.ID.Value
	result, err := r.p.client.UpdateSmartList(smartListID, smartList)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error update order",
			"Could not update orderID "+smartListID+": "+err.Error(),
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

func (r resourceSmartList) Delete(ctx context.Context, req tfsdk.DeleteResourceRequest, resp *tfsdk.DeleteResourceResponse) {
	var state SmartList
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	smartListID := state.ID.Value
	err := r.p.client.DeleteSmartList(smartListID)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error deleting smartList",
			"Could not delete smartList with ID "+smartListID+": "+err.Error(),
		)
		return
	}

	resp.State.RemoveResource(ctx)
}

func (r resourceSmartList) ImportState(ctx context.Context, req tfsdk.ImportResourceStateRequest, resp *tfsdk.ImportResourceStateResponse) {
	tfsdk.ResourceImportStatePassthroughID(ctx, tftypes.NewAttributePath().WithAttributeName("id"), req, resp)
}
