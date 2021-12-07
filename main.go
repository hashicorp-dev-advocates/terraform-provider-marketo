package main

import (
	"context"

	"github.com/eveld/terraform-provider-marketo/internal/provider"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
)

func main() {
	tfsdk.Serve(context.Background(), provider.New, tfsdk.ServeOpts{
		Name: "marketo",
	})
}
