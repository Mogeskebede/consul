// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

//go:build !consulent
// +build !consulent

package types

import (
	multiclusterv1alpha1 "github.com/hashicorp/consul/proto-public/pbmulticluster/v1alpha1"
)

func validExportedServicesWithPartition() *multiclusterv1alpha1.ExportedServices {
	consumers := []*multiclusterv1alpha1.ExportedServicesConsumer{
		{
			ConsumerTenancy: &multiclusterv1alpha1.ExportedServicesConsumer_Partition{
				Partition: "",
			},
		},
	}
	return &multiclusterv1alpha1.ExportedServices{
		Services:  []string{"api", "frontend", "backend"},
		Consumers: consumers,
	}
}

func validExportedServicesWithSamenessGroup() *multiclusterv1alpha1.ExportedServices {
	consumers := []*multiclusterv1alpha1.ExportedServicesConsumer{
		{
			ConsumerTenancy: &multiclusterv1alpha1.ExportedServicesConsumer_SamenessGroup{
				SamenessGroup: "",
			},
		},
	}
	return &multiclusterv1alpha1.ExportedServices{
		Services:  []string{"api", "frontend", "backend"},
		Consumers: consumers,
	}
}

func inValidExportedServices() *multiclusterv1alpha1.ExportedServices {
	return &multiclusterv1alpha1.ExportedServices{}
}