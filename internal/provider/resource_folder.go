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

type resourceFolderType struct{}

func (r resourceFolderType) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
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
		},
	}, nil
}

func (r resourceFolderType) NewResource(_ context.Context, p tfsdk.Provider) (tfsdk.Resource, diag.Diagnostics) {
	return resourceFolder{
		p: *(p.(*provider)),
	}, nil
}

type resourceFolder struct {
	p provider
}

func (r resourceFolder) Create(ctx context.Context, req tfsdk.CreateResourceRequest, resp *tfsdk.CreateResourceResponse) {
	if !r.p.configured {
		resp.Diagnostics.AddError(
			"Provider not configured",
			"The provider hasn't been configured before apply, likely because it depends on an unknown value from another resource. This leads to weird stuff happening, so we'd prefer if you didn't do that. Thanks!",
		)
		return
	}

	var plan Folder
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	folder := marketo.Folder{
		Name: plan.Name.Value,
	}

	result, err := r.p.client.CreateFolder(folder)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating folder",
			"Could not create folder, unexpected error: "+err.Error(),
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

func (r resourceFolder) Read(ctx context.Context, req tfsdk.ReadResourceRequest, resp *tfsdk.ReadResourceResponse) {
	var state Folder
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	folderID := state.ID.Value
	folder, err := r.p.client.GetFolder(folderID)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error reading folder",
			"Could not read folder with ID "+folderID+": "+err.Error(),
		)
		return
	}

	state.ID = types.String{Value: folder.ID}
	state.Name = types.String{Value: folder.Name}

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r resourceFolder) Update(ctx context.Context, req tfsdk.UpdateResourceRequest, resp *tfsdk.UpdateResourceResponse) {
	var plan Folder
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state Folder
	diags = req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	folder := marketo.Folder{
		Name: plan.Name.Value,
	}

	folderID := state.ID.Value
	result, err := r.p.client.UpdateFolder(folderID, folder)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error update order",
			"Could not update orderID "+folderID+": "+err.Error(),
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

func (r resourceFolder) Delete(ctx context.Context, req tfsdk.DeleteResourceRequest, resp *tfsdk.DeleteResourceResponse) {
	var state Folder
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	folderID := state.ID.Value
	err := r.p.client.DeleteFolder(folderID)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error deleting folder",
			"Could not delete folder with ID "+folderID+": "+err.Error(),
		)
		return
	}

	resp.State.RemoveResource(ctx)
}

func (r resourceFolder) ImportState(ctx context.Context, req tfsdk.ImportResourceStateRequest, resp *tfsdk.ImportResourceStateResponse) {
	tfsdk.ResourceImportStatePassthroughID(ctx, tftypes.NewAttributePath().WithAttributeName("id"), req, resp)
}
