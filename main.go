package main

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/nokia/eda/apps/terraform-provider-topologies/internal/provider"
)

func main() {
	opts := providerserver.ServeOpts{
		Address: "github.com/nokia-eda/topologies-v1alpha1",
	}

	err := providerserver.Serve(context.Background(), provider.New("v1alpha1"), opts)
	if err != nil {
		log.Fatal(err.Error())
	}
}
