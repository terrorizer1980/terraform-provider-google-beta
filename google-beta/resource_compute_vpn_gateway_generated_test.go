// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package google

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccComputeVpnGateway_TargetVpnGatewayBasicExample(t *testing.T) {
	t.Parallel()

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckComputeVpnGatewayDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccComputeVpnGateway_TargetVpnGatewayBasicExample(acctest.RandString(10)),
			},
			{
				ResourceName:      "google_compute_vpn_gateway.target_gateway",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccComputeVpnGateway_TargetVpnGatewayBasicExample(val string) string {
	return fmt.Sprintf(`
resource "google_compute_vpn_gateway" "target_gateway" {
  name    = "vpn1-%s"
  network = "${google_compute_network.network1.self_link}"
}

resource "google_compute_network" "network1" {
  name       = "network1-%s"
}

resource "google_compute_address" "vpn_static_ip" {
  name   = "vpn-static-ip-%s"
}

resource "google_compute_forwarding_rule" "fr_esp" {
  name        = "fr-esp-%s"
  ip_protocol = "ESP"
  ip_address  = "${google_compute_address.vpn_static_ip.address}"
  target      = "${google_compute_vpn_gateway.target_gateway.self_link}"
}

resource "google_compute_forwarding_rule" "fr_udp500" {
  name        = "fr-udp500-%s"
  ip_protocol = "UDP"
  port_range  = "500"
  ip_address  = "${google_compute_address.vpn_static_ip.address}"
  target      = "${google_compute_vpn_gateway.target_gateway.self_link}"
}

resource "google_compute_forwarding_rule" "fr_udp4500" {
  name        = "fr-udp4500-%s"
  ip_protocol = "UDP"
  port_range  = "4500"
  ip_address  = "${google_compute_address.vpn_static_ip.address}"
  target      = "${google_compute_vpn_gateway.target_gateway.self_link}"
}

resource "google_compute_vpn_tunnel" "tunnel1" {
  name          = "tunnel1-%s"
  peer_ip       = "15.0.0.120"
  shared_secret = "a secret message"

  target_vpn_gateway = "${google_compute_vpn_gateway.target_gateway.self_link}"

  depends_on = [
    "google_compute_forwarding_rule.fr_esp",
    "google_compute_forwarding_rule.fr_udp500",
    "google_compute_forwarding_rule.fr_udp4500",
  ]
}

resource "google_compute_route" "route1" {
  name       = "route1-%s"
  network    = "${google_compute_network.network1.name}"
  dest_range = "15.0.0.0/24"
  priority   = 1000

  next_hop_vpn_tunnel = "${google_compute_vpn_tunnel.tunnel1.self_link}"
}
`, val, val, val, val, val, val, val, val,
	)
}

func testAccCheckComputeVpnGatewayDestroy(s *terraform.State) error {
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "google_compute_vpn_gateway" {
			continue
		}

		config := testAccProvider.Meta().(*Config)

		url, err := replaceVarsForTest(rs, "https://www.googleapis.com/compute/beta/projects/{{project}}/regions/{{region}}/targetVpnGateways/{{name}}")
		if err != nil {
			return err
		}

		_, err = sendRequest(config, "GET", url, nil)
		if err == nil {
			return fmt.Errorf("ComputeVpnGateway still exists at %s", url)
		}
	}

	return nil
}
