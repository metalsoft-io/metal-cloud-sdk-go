package metalcloud

import (
	"encoding/json"
	"testing"

	. "github.com/onsi/gomega"
)

func TestDatacenterConfiglUnmarshalTest(t *testing.T) {

	RegisterTestingT(t)

	var dc DatacenterConfig

	err := json.Unmarshal([]byte(_DCConfigFixture), &dc)
	if err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	Expect(dc).NotTo(BeNil())
	/*
		Expect(dc.SANRoutedSubnet).To(Equal("100.64.0.0/21"))
		Expect(dc.SwitchProvisioner.Type).To(Equal("VPLSProvisioner"))
		Expect(dc.SwitchProvisioner.Provisioner.(VPLSProvisioner).ACLSAN).To(Equal(3999))
	*/

	Expect(dc.SANRoutedSubnet).To(Equal("100.64.0.0/21"))
	Expect(dc.SwitchProvisioner["type"]).To(Equal("VPLSProvisioner"))
	Expect(dc.SwitchProvisioner["ACLSAN"]).To(Equal(3999.0))

}

const _DCFixture = "{\"datacenter_id\":4,\"datacenter_name\":\"slavedatacenter-140\",\"datacenter_name_parent\":null,\"user_id\":null,\"datacenter_is_master\":false,\"datacenter_is_maintenance\":false,\"datacenter_type\":\"metal_cloud\",\"datacenter_display_name\":\"Slave, Datacenter 140\",\"datacenter_hidden\":false,\"datacenter_created_timestamp\":\"2021-04-27T17:16:20Z\",\"datacenter_updated_timestamp\":\"2021-04-27T17:16:20Z\",\"type\":\"Datacenter\",\"datacenter_tags\":[]}"
const _DCConfigFixture = "{\"SANRoutedSubnet\":\"100.64.0.0/21\",\"BSIVRRPListenIPv4\":\"172.16.10.6\",\"BSIMachineListenIPv4List\":[\"172.16.10.6\"],\"BSIMachinesSubnetIPv4CIDR\":\"10.255.226.0/24\",\"BSIExternallyVisibleIPv4\":\"89.36.24.2\",\"repoURLRoot\":\"https://repointegrationpublic.metalsoft.io\",\"repoURLRootQuarantineNetwork\":\"https://repointegrationpublic.metalsoft.io\",\"DNSServers\":[\"84.40.63.27\"],\"NTPServers\":[\"84.40.58.44\",\"84.40.58.45\"],\"KMS\":\"\",\"TFTPServerWANVRRPListenIPv4\":\"172.16.10.6\",\"dataLakeEnabled\":false,\"monitoringGraphitePlainTextSocketHost\":\"\",\"monitoringGraphiteRenderURLHost\":\"\",\"latitude\":0,\"longitude\":0,\"address\":\"\",\"switchProvisioner\":{\"type\":\"VPLSProvisioner\",\"ACLSAN\":3999,\"ACLWAN\":3399,\"SANACLRange\":\"3700-3998\",\"ToRLANVLANRange\":\"400-699\",\"ToRSANVLANRange\":\"700-999\",\"ToRWANVLANRange\":\"100-300\",\"quarantineVLANID\":5,\"NorthWANVLANRange\":\"1001-2000\"},\"childDatacentersConfigDefault\":[],\"provisionUsingTheQuarantineNetwork\":true,\"enableDHCPRelaySecurityForQuarantineNetwork\":true,\"enableDHCPRelaySecurityForClientNetworks\":true,\"defaultCleanupAndRegistrationMechanism\":\"bmc\",\"defaultDeploymentMechanism\":\"virtual_media\",\"NFSServer\":\"\",\"Option82ToIPMapping\":{\"eth1\":\"10.0.0.1\"}}"

func TestDatacenterConfigMarshalTest(t *testing.T) {

	RegisterTestingT(t)

	var dc DatacenterConfig

	err := json.Unmarshal([]byte(_DCConfigFixture), &dc)
	if err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	Expect(dc).NotTo(BeNil())

	b, err := json.Marshal(dc)
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}
	Expect(b).NotTo(BeNil())

	var dc2 DatacenterConfig
	err = json.Unmarshal(b, &dc2)
	if err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}

	Expect(dc2.SANRoutedSubnet).To(Equal("100.64.0.0/21"))
	Expect(dc2.SwitchProvisioner["type"]).To(Equal("VPLSProvisioner"))
	Expect(dc2.SwitchProvisioner["ACLSAN"]).To(Equal(3999.0))
	Expect(dc2.ProvisionUsingTheQuarantineNetwork).To(Equal(true))
	Expect(dc2.EnableDHCPRelaySecurityForQuarantineNetwork).To(Equal(true))
	Expect(dc2.EnableDHCPRelaySecurityForClientNetworks).To(Equal(true))
}

func TestDatacenterCreateOrUpdate(t *testing.T) {

	RegisterTestingT(t)

	responseBody = `{"result": ` + _driveArrayFixture1 + `,"jsonrpc": "2.0","id": 0}`

	mc, err := GetMetalcloudClient("userEmail", "APIKey", httpServer.URL, false, "", "", "")
	Expect(err).To(BeNil())

	obj := Datacenter{
		DatacenterName: "dctest",
		UserID:         1,
		DatacenterConfig: &DatacenterConfig{
			SANRoutedSubnet:                       "100.64.0.0/21",
			BSIVRRPListenIPv4:                     "172.16.10.6",
			BSIMachineListenIPv4List:              []string{"172.16.10.6"},
			BSIMachinesSubnetIPv4CIDR:             "10.255.226.0/24",
			BSIExternallyVisibleIPv4:              "89.36.24.2",
			RepoURLRoot:                           "https://repointegrationpublic.metalsoft.io",
			RepoURLRootQuarantineNetwork:          "https://repointegrationpublic.metalsoft.io",
			DNSServers:                            []string{"84.40.63.27"},
			NTPServers:                            []string{"84.40.58.44", "84.40.58.45"},
			KMS:                                   "",
			TFTPServerWANVRRPListenIPv4:           "172.16.10.6",
			DataLakeEnabled:                       false,
			MonitoringGraphitePlainTextSocketHost: "",
			MonitoringGraphiteRenderURLHost:       "",
			Latitude:                              0,
			Longitude:                             0,
			SwitchProvisioner: map[string]interface{}{
				"type":                          "VPLSProvisioner",
				"ACLSAN":                        3399,
				"SANACLRange":                   "3700-3998",
				"ToRLANVLANRange":               "400-699",
				"ToRSANVLANRange":               "700-999",
				"ToRWANVLANRange":               "100-300",
				"quarantineVLANID":              5,
				"NorthWANVLANRange":             "1001-2000",
				"childDatacentersConfigDefault": []string{},
			},
		},
	}

	err = obj.CreateOrUpdate(mc)
	if err != nil {
		t.Fatalf("CreateOrUpdate failed: %v", err)
	}

	body := (<-requestChan).body
	var m map[string]interface{}
	err2 := json.Unmarshal([]byte(body), &m)
	Expect(err2).To(BeNil())
	Expect(m).NotTo(BeNil())

	Expect(m["method"].(string)).To(Equal("datacenter_get"))

	params := (m["params"].([]interface{}))

	Expect(params[1].(string)).To(Equal("dctest"))

	body = (<-requestChan).body

	err2 = json.Unmarshal([]byte(body), &m)
	Expect(err2).To(BeNil())
	Expect(m).NotTo(BeNil())

	Expect(m["method"].(string)).To(Equal("datacenter_config_update"))

	params = (m["params"].([]interface{}))

	Expect(params[0].(string)).To(Equal("dctest"))

	responseBody = `{"error": {"message": "Datacenter not found.","code": 269}, "jsonrpc": "2.0", "id": 0}`

	err = obj.CreateOrUpdate(mc)
	// if err != nil {
	// 	t.Fatalf("CreateOrUpdate failed: %v", err)
	// }

	body = (<-requestChan).body
	err2 = json.Unmarshal([]byte(body), &m)
	Expect(err2).To(BeNil())
	Expect(m).NotTo(BeNil())

	Expect(m["method"].(string)).To(Equal("datacenter_get"))

	params = (m["params"].([]interface{}))

	Expect(params[1].(string)).To(Equal("dctest"))

	body = (<-requestChan).body

	err2 = json.Unmarshal([]byte(body), &m)
	Expect(err2).To(BeNil())
	Expect(m).NotTo(BeNil())

	Expect(m["method"].(string)).To(Equal("datacenter_create"))

	params = (m["params"].([]interface{}))
	Expect(params[0].(map[string]interface{})["user_id"].(float64)).To(Equal(1.0))
}

func TestDatacenterDeleteForApply(t *testing.T) {

	RegisterTestingT(t)

	responseBody = `{"result": {"DatacenterName": "dctest"}, "jsonrpc": "2.0", "id": 0}`

	//TODO: Test why this is failing
	// mc, err := GetMetalcloudClient("userEmail", "APIKey", httpServer.URL, false, "", "", "")
	// Expect(err).To(BeNil())

	// obj := Datacenter{
	// 	DatacenterName: "dctest",
	// 	UserID:         1,
	// 	DatacenterConfig: &DatacenterConfig{
	// 		SANRoutedSubnet:                       "100.64.0.0/21",
	// 		BSIVRRPListenIPv4:                     "172.16.10.6",
	// 		BSIMachineListenIPv4List:              []string{"172.16.10.6"},
	// 		BSIMachinesSubnetIPv4CIDR:             "10.255.226.0/24",
	// 		BSIExternallyVisibleIPv4:              "89.36.24.2",
	// 		RepoURLRoot:                           "https://repointegrationpublic.metalsoft.io",
	// 		RepoURLRootQuarantineNetwork:          "https://repointegrationpublic.metalsoft.io",
	// 		DNSServers:                            []string{"84.40.63.27"},
	// 		NTPServers:                            []string{"84.40.58.44", "84.40.58.45"},
	// 		KMS:                                   "",
	// 		TFTPServerWANVRRPListenIPv4:           "172.16.10.6",
	// 		DataLakeEnabled:                       false,
	// 		MonitoringGraphitePlainTextSocketHost: "",
	// 		MonitoringGraphiteRenderURLHost:       "",
	// 		Latitude:                              0,
	// 		Longitude:                             0,
	// 		SwitchProvisioner: map[string]interface{}{
	// 			"type":                          "VPLSProvisioner",
	// 			"ACLSAN":                        3399,
	// 			"SANACLRange":                   "3700-3998",
	// 			"ToRLANVLANRange":               "400-699",
	// 			"ToRSANVLANRange":               "700-999",
	// 			"ToRWANVLANRange":               "100-300",
	// 			"quarantineVLANID":              5,
	// 			"NorthWANVLANRange":             "1001-2000",
	// 			"childDatacentersConfigDefault": []string{},
	// 		},
	// 	},
	// }
	// err = obj.Delete(mc)

	// Expect(err).To(BeNil())
}

const _DCFixture2 = "{\"datacenter_id\":4,\"datacenter_name\":\"slavedatacenter-140\",\"datacenter_name_parent\":null,\"user_id\":null,\"datacenter_is_master\":false,\"datacenter_is_maintenance\":false,\"datacenter_type\":\"metal_cloud\",\"datacenter_display_name\":\"Slave, Datacenter 140\",\"datacenter_hidden\":false,\"datacenter_created_timestamp\":\"2021-04-27T17:16:20Z\",\"datacenter_updated_timestamp\":\"2021-04-27T17:16:20Z\",\"type\":\"Datacenter\",\"datacenter_tags\":[]}"

func TestDatacenterUnmarshalTest(t *testing.T) {

	RegisterTestingT(t)

	var dc Datacenter

	err := json.Unmarshal([]byte(_DCFixture), &dc)
	if err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	Expect(dc).NotTo(BeNil())

}
