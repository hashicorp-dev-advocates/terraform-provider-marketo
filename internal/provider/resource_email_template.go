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

type resourceEmailTemplateType struct{}

func (r resourceEmailTemplateType) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
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
			"content": {
				Type:     types.StringType,
				Required: true,
			},
		},
	}, nil
}

func (r resourceEmailTemplateType) NewResource(_ context.Context, p tfsdk.Provider) (tfsdk.Resource, diag.Diagnostics) {
	return resourceEmailTemplate{
		p: *(p.(*provider)),
	}, nil
}

type resourceEmailTemplate struct {
	p provider
}

func (r resourceEmailTemplate) Create(ctx context.Context, req tfsdk.CreateResourceRequest, resp *tfsdk.CreateResourceResponse) {
	if !r.p.configured {
		resp.Diagnostics.AddError(
			"Provider not configured",
			"The provider hasn't been configured before apply, likely because it depends on an unknown value from another resource. This leads to weird stuff happening, so we'd prefer if you didn't do that. Thanks!",
		)
		return
	}

	var plan EmailTemplate
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	emailTemplate := marketo.EmailTemplate{
		Name: plan.Name.Value,
	}

	result, err := r.p.client.CreateEmailTemplate(emailTemplate)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating emailTemplate",
			"Could not create emailTemplate, unexpected error: "+err.Error(),
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

func (r resourceEmailTemplate) Read(ctx context.Context, req tfsdk.ReadResourceRequest, resp *tfsdk.ReadResourceResponse) {
	var state EmailTemplate
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	emailTemplateID := state.ID.Value
	emailTemplate, err := r.p.client.GetEmailTemplate(emailTemplateID)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error reading emailTemplate",
			"Could not read emailTemplate with ID "+emailTemplateID+": "+err.Error(),
		)
		return
	}

	state.ID = types.String{Value: emailTemplate.ID}
	state.Name = types.String{Value: emailTemplate.Name}

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r resourceEmailTemplate) Update(ctx context.Context, req tfsdk.UpdateResourceRequest, resp *tfsdk.UpdateResourceResponse) {
	var plan EmailTemplate
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state EmailTemplate
	diags = req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	emailTemplate := marketo.EmailTemplate{
		Name: plan.Name.Value,
	}

	emailTemplateID := state.ID.Value
	result, err := r.p.client.UpdateEmailTemplate(emailTemplateID, emailTemplate)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error update order",
			"Could not update orderID "+emailTemplateID+": "+err.Error(),
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

func (r resourceEmailTemplate) Delete(ctx context.Context, req tfsdk.DeleteResourceRequest, resp *tfsdk.DeleteResourceResponse) {
	var state EmailTemplate
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	emailTemplateID := state.ID.Value
	err := r.p.client.DeleteEmailTemplate(emailTemplateID)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error deleting emailTemplate",
			"Could not delete emailTemplate with ID "+emailTemplateID+": "+err.Error(),
		)
		return
	}

	resp.State.RemoveResource(ctx)
}

func (r resourceEmailTemplate) ImportState(ctx context.Context, req tfsdk.ImportResourceStateRequest, resp *tfsdk.ImportResourceStateResponse) {
	tfsdk.ResourceImportStatePassthroughID(ctx, tftypes.NewAttributePath().WithAttributeName("id"), req, resp)
}
