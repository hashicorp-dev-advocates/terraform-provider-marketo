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

type resourceEmailType struct{}

func (r resourceEmailType) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
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
			"from_email": {
				Type:     types.StringType,
				Required: true,
			},
			"from_name": {
				Type:     types.StringType,
				Required: true,
			},
			"reply_to": {
				Type:     types.StringType,
				Required: true,
			},
			"operational": {
				Type:     types.BoolType,
				Optional: true,
			},
			"text_only": {
				Type:     types.BoolType,
				Optional: true,
			},
			"subject": {
				Type:     types.StringType,
				Required: true,
			},
			"template": {
				Type:     types.StringType,
				Required: true,
			},
			"content": {
				Optional: true,
				Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{
					"section": {
						Type:     types.StringType,
						Required: true,
					},
					"text": {
						Type:     types.StringType,
						Optional: true,
					},
					"dynamic_content": {
						Type:     types.StringType,
						Optional: true,
					},
					"snippet": {
						Type:     types.StringType,
						Optional: true,
					},
				}, tfsdk.ListNestedAttributesOptions{}),
			},
		},
	}, nil
}

func (r resourceEmailType) NewResource(_ context.Context, p tfsdk.Provider) (tfsdk.Resource, diag.Diagnostics) {
	return resourceEmail{
		p: *(p.(*provider)),
	}, nil
}

type resourceEmail struct {
	p provider
}

func (r resourceEmail) Create(ctx context.Context, req tfsdk.CreateResourceRequest, resp *tfsdk.CreateResourceResponse) {
	if !r.p.configured {
		resp.Diagnostics.AddError(
			"Provider not configured",
			"The provider hasn't been configured before apply, likely because it depends on an unknown value from another resource. This leads to weird stuff happening, so we'd prefer if you didn't do that. Thanks!",
		)
		return
	}

	var plan Email
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	email := marketo.Email{
		Name: plan.Name.Value,
	}

	result, err := r.p.client.CreateEmail(email)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating email",
			"Could not create email, unexpected error: "+err.Error(),
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

func (r resourceEmail) Read(ctx context.Context, req tfsdk.ReadResourceRequest, resp *tfsdk.ReadResourceResponse) {
	var state Email
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	emailID := state.ID.Value
	email, err := r.p.client.GetEmail(emailID)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error reading email",
			"Could not read email with ID "+emailID+": "+err.Error(),
		)
		return
	}

	state.ID = types.String{Value: email.ID}
	state.Name = types.String{Value: email.Name}

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r resourceEmail) Update(ctx context.Context, req tfsdk.UpdateResourceRequest, resp *tfsdk.UpdateResourceResponse) {
	var plan Email
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state Email
	diags = req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	email := marketo.Email{
		Name: plan.Name.Value,
	}

	emailID := state.ID.Value
	result, err := r.p.client.UpdateEmail(emailID, email)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error update order",
			"Could not update orderID "+emailID+": "+err.Error(),
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

func (r resourceEmail) Delete(ctx context.Context, req tfsdk.DeleteResourceRequest, resp *tfsdk.DeleteResourceResponse) {
	var state Email
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	emailID := state.ID.Value
	err := r.p.client.DeleteEmail(emailID)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error deleting email",
			"Could not delete email with ID "+emailID+": "+err.Error(),
		)
		return
	}

	resp.State.RemoveResource(ctx)
}

func (r resourceEmail) ImportState(ctx context.Context, req tfsdk.ImportResourceStateRequest, resp *tfsdk.ImportResourceStateResponse) {
	tfsdk.ResourceImportStatePassthroughID(ctx, tftypes.NewAttributePath().WithAttributeName("id"), req, resp)
}
