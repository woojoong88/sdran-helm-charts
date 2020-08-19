// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package main

import (
	"github.com/onosproject/helmit/pkg/registry"
	"github.com/onosproject/helmit/pkg/test"
	"github.com/onosproject/sdran-helm-charts/aether-umbrella/tests"
	ric "github.com/onosproject/sdran-helm-charts/onos-ric/tests"
	sdran "github.com/onosproject/sdran-helm-charts/sd-ran/tests"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
)

func main() {
	registry.RegisterTestSuite("onos-ric", &ric.ONOSRICSuite{})
	registry.RegisterTestSuite("sd-ran", &sdran.SDRANSuite{})
	registry.RegisterTestSuite("aether-umbrella", &tests.AetherUmbrellaSuite{})
	test.Main()
}
