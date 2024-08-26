package metalcloud

import (
	"encoding/json"
	"testing"

	. "github.com/onsi/gomega"
)

func TestInstanceArrayUnmarshalTest(t *testing.T) {
	RegisterTestingT(t)

	var i Instance
	err := json.Unmarshal([]byte(_instanceFixture1), &i)
	Expect(err).To(BeNil())
	Expect(i).NotTo(BeNil())

	Expect(i.InstanceID).To(Equal(20639))
	Expect(i.InstanceCredentials).NotTo(BeNil())
	Expect(i.InstanceCredentials.SSH).NotTo(BeNil())
	Expect(i.InstanceCredentials.SSH.InitialPassword).To(Equal("asdasd"))
	Expect(i.InstanceInterfaces[0].InstanceInterfaceLabel).To(Equal("if0"))

	found := false
	for _, i := range i.InstanceInterfaces {
		for _, ip := range i.InstanceInterfaceIPs {
			if ip.IPHumanReadable == "172.17.106.22" {
				found = true
			}
		}
	}

	Expect(found).To(BeTrue())

}

func TestInstanceArrayUnmarshalTestWithSharedDrivesAttached(t *testing.T) {
	RegisterTestingT(t)

	var i Instance
	err := json.Unmarshal([]byte(_instanceFixture3), &i)
	Expect(err).To(BeNil())
	Expect(i).NotTo(BeNil())

	Expect(i.InstanceCredentials.ISCSI.Password).To(Equal("sssdsd"))
	Expect(i.InstanceCredentials.SharedDrives["csivolumename3"].TargetIQN).To(Equal("iqn.2013-01.io.metalsoft:storage.91nqwgd.4lys6bq.iel1v3k"))

}

func TestInstanceArrayInstances(t *testing.T) {
	RegisterTestingT(t)

	responseBody = `{"result": {"test":` + _instanceFixture1 + `},"jsonrpc": "2.0","id": 0}`

	mc, err := GetMetalcloudClient("userEmail", "APIKey", httpServer.URL, false, "", "", "")
	Expect(err).To(BeNil())

	ret, err := mc.instanceArrayInstances("test")
	Expect(err).To(BeNil())
	Expect(ret).NotTo(BeNil())

	i := (*ret)["test"]

	Expect(i.InstanceID).To(Equal(20639))
	Expect(i.InstanceCredentials.SSH.InitialPassword).To(Equal("asdasd"))

	body := (<-requestChan).body

	var m map[string]interface{}
	err2 := json.Unmarshal([]byte(body), &m)
	Expect(err2).To(BeNil())
	Expect(m).NotTo(BeNil())

	params := (m["params"].([]interface{}))

	Expect(params[0].(string)).To(Equal("test"))

}

func TestInstancePowerGet(t *testing.T) {
	RegisterTestingT(t)

	responseBody = `{"result": "on","jsonrpc": "2.0","id": 0}`

	mc, err := GetMetalcloudClient("userEmail", "APIKey", httpServer.URL, false, "", "", "")
	Expect(err).To(BeNil())

	ret, err := mc.InstanceServerPowerGet(1)
	Expect(err).To(BeNil())
	Expect(*ret).To(Equal("on"))

	(<-requestChan)

}

const _instanceFixture1 = "{\"instance_id\":20639,\"instance_array_id\":23739,\"server_id\":1,\"server_type_id\":1,\"instance_change_id\":48471,\"instance_service_status\":\"active\",\"drive_id_bootable\":12146,\"instance_label\":\"instance-20639\",\"instance_subdomain\":\"instance-20639.instance-array-23739.vanilla.demo.2.poc.metalcloud.io\",\"instance_subdomain_permanent\":\"instance-20639.poc.metalcloud.io\",\"instance_updated_timestamp\":\"2019-11-27T14:38:26Z\",\"instance_created_timestamp\":\"2019-11-27T14:38:25Z\",\"template_id_origin\":null,\"instance_operation\":{\"instance_change_id\":48471,\"instance_id\":20639,\"instance_array_id\":23739,\"server_id\":1,\"server_type_id\":1,\"instance_deploy_type\":\"create\",\"instance_deploy_status\":\"finished\",\"instance_label\":\"instance-20639\",\"instance_subdomain\":\"instance-20639.instance-array-23739.vanilla.demo.2.poc.metalcloud.io\",\"instance_updated_timestamp\":\"2019-11-27T14:38:26Z\",\"drive_id_bootable\":12146,\"template_id_origin\":null,\"type\":\"InstanceOperation\"},\"type\":\"Instance\",\"instance_interfaces\":[{\"instance_interface_id\":82706,\"network_id\":9013,\"instance_interface_index\":0,\"instance_id\":20639,\"instance_interface_change_id\":203393,\"instance_interface_service_status\":\"active\",\"instance_interface_capacity_mbps\":10000,\"instance_interface_label\":\"if0\",\"instance_interface_subdomain\":\"if0.instance-20639.instance-array-23739.vanilla.demo.2.poc.metalcloud.io\",\"instance_interface_operation\":{\"instance_interface_change_id\":203393,\"instance_interface_id\":82706,\"network_id\":9013,\"instance_interface_index\":0,\"instance_id\":20639,\"instance_interface_deploy_type\":\"create\",\"instance_interface_deploy_status\":\"finished\",\"instance_interface_capacity_mbps\":10000,\"instance_interface_subdomain\":\"if0.instance-20639.instance-array-23739.vanilla.demo.2.poc.metalcloud.io\",\"instance_interface_label\":\"if0\",\"type\":\"InstanceInterfaceOperation\",\"instance_interface_lagg_indexes\":[]},\"type\":\"InstanceInterface\",\"instance_interface_lagg_indexes\":[],\"server_interface\":{\"server_interface_mac_address\":\"24:6e:96:6a:37:9a\",\"type\":\"ServerInterface\"},\"instance_interface_ips\":[]},{\"instance_interface_id\":82707,\"network_id\":9012,\"instance_interface_index\":1,\"instance_id\":20639,\"instance_interface_change_id\":203392,\"instance_interface_service_status\":\"active\",\"instance_interface_capacity_mbps\":10000,\"instance_interface_label\":\"if1\",\"instance_interface_subdomain\":\"if1.instance-20639.instance-array-23739.vanilla.demo.2.poc.metalcloud.io\",\"instance_interface_operation\":{\"instance_interface_change_id\":203392,\"instance_interface_id\":82707,\"network_id\":9012,\"instance_interface_index\":1,\"instance_id\":20639,\"instance_interface_deploy_type\":\"create\",\"instance_interface_deploy_status\":\"finished\",\"instance_interface_capacity_mbps\":10000,\"instance_interface_subdomain\":\"if1.instance-20639.instance-array-23739.vanilla.demo.2.poc.metalcloud.io\",\"instance_interface_label\":\"if1\",\"type\":\"InstanceInterfaceOperation\",\"instance_interface_lagg_indexes\":[]},\"type\":\"InstanceInterface\",\"instance_interface_lagg_indexes\":[],\"server_interface\":{\"server_interface_mac_address\":\"a0:36:9f:ee:b2:f0\",\"type\":\"ServerInterface\"},\"instance_interface_ips\":[{\"ip_id\":59395,\"ip_hex\":\"fd1f8bbb56b308020000000000000002\",\"ip_human_readable\":\"fd1f:8bbb:56b3:0802:0000:0000:0000:0002\",\"ip_type\":\"ipv6\",\"instance_interface_id\":82707,\"subnet_id\":20504,\"ip_lease_expires\":\"0000-00-00T00:00:00Z\",\"ip_change_id\":91240,\"subnet_netmask_human_readable\":\"ffff:ffff:ffff:ffff:0000:0000:0000:0000\",\"subnet_gateway_human_readable\":\"fd1f:8bbb:56b3:0802:0000:0000:0000:0001\",\"subnet_destination\":\"wan\",\"ip_operation\":{\"ip_change_id\":91240,\"ip_id\":59395,\"ip_hex\":\"fd1f8bbb56b308020000000000000002\",\"ip_human_readable\":\"fd1f:8bbb:56b3:0802:0000:0000:0000:0002\",\"ip_type\":\"ipv6\",\"instance_interface_id\":82707,\"subnet_id\":20504,\"ip_lease_expires\":\"0000-00-00T00:00:00Z\",\"ip_updated_timestamp\":\"2019-11-27T14:38:25Z\",\"ip_deploy_type\":\"create\",\"ip_deploy_status\":\"finished\",\"ip_label\":\"ip-59395\",\"ip_subdomain\":\"ip-59395.subnet-20504.wan.demo.2.poc.metalcloud.io\",\"type\":\"IPOperation\"},\"type\":\"IP\"},{\"ip_id\":59396,\"ip_hex\":\"ac116a16\",\"ip_human_readable\":\"172.17.106.22\",\"ip_type\":\"ipv4\",\"instance_interface_id\":82707,\"subnet_id\":20505,\"ip_lease_expires\":\"0000-00-00T00:00:00Z\",\"ip_change_id\":91241,\"subnet_netmask_human_readable\":\"255.255.255.252\",\"subnet_gateway_human_readable\":\"172.17.106.21\",\"subnet_destination\":\"wan\",\"ip_operation\":{\"ip_change_id\":91241,\"ip_id\":59396,\"ip_hex\":\"ac116a16\",\"ip_human_readable\":\"172.17.106.22\",\"ip_type\":\"ipv4\",\"instance_interface_id\":82707,\"subnet_id\":20505,\"ip_lease_expires\":\"0000-00-00T00:00:00Z\",\"ip_updated_timestamp\":\"2019-11-27T14:38:25Z\",\"ip_deploy_type\":\"create\",\"ip_deploy_status\":\"finished\",\"ip_label\":\"ip-59396\",\"ip_subdomain\":\"ip-59396.subnet-20505.wan.demo.2.poc.metalcloud.io\",\"type\":\"IPOperation\"},\"type\":\"IP\"}]},{\"instance_interface_id\":82708,\"network_id\":null,\"instance_interface_index\":2,\"instance_id\":20639,\"instance_interface_change_id\":203390,\"instance_interface_service_status\":\"active\",\"instance_interface_capacity_mbps\":10000,\"instance_interface_label\":\"if2\",\"instance_interface_subdomain\":\"if2.instance-20639.instance-array-23739.vanilla.demo.2.poc.metalcloud.io\",\"instance_interface_operation\":{\"instance_interface_change_id\":203390,\"instance_interface_id\":82708,\"network_id\":null,\"instance_interface_index\":2,\"instance_id\":20639,\"instance_interface_deploy_type\":\"create\",\"instance_interface_deploy_status\":\"finished\",\"instance_interface_capacity_mbps\":10000,\"instance_interface_subdomain\":\"if2.instance-20639.instance-array-23739.vanilla.demo.2.poc.metalcloud.io\",\"instance_interface_label\":\"if2\",\"type\":\"InstanceInterfaceOperation\",\"instance_interface_lagg_indexes\":[]},\"type\":\"InstanceInterface\",\"instance_interface_lagg_indexes\":[],\"server_interface\":{\"server_interface_mac_address\":\"a0:36:9f:f0:ee:3a\",\"type\":\"ServerInterface\"},\"instance_interface_ips\":[]},{\"instance_interface_id\":82709,\"network_id\":null,\"instance_interface_index\":3,\"instance_id\":20639,\"instance_interface_change_id\":203391,\"instance_interface_service_status\":\"active\",\"instance_interface_capacity_mbps\":10000,\"instance_interface_label\":\"if3\",\"instance_interface_subdomain\":\"if3.instance-20639.instance-array-23739.vanilla.demo.2.poc.metalcloud.io\",\"instance_interface_operation\":{\"instance_interface_change_id\":203391,\"instance_interface_id\":82709,\"network_id\":null,\"instance_interface_index\":3,\"instance_id\":20639,\"instance_interface_deploy_type\":\"create\",\"instance_interface_deploy_status\":\"finished\",\"instance_interface_capacity_mbps\":10000,\"instance_interface_subdomain\":\"if3.instance-20639.instance-array-23739.vanilla.demo.2.poc.metalcloud.io\",\"instance_interface_label\":\"if3\",\"type\":\"InstanceInterfaceOperation\",\"instance_interface_lagg_indexes\":[]},\"type\":\"InstanceInterface\",\"instance_interface_lagg_indexes\":[],\"server_interface\":{\"server_interface_mac_address\":\"a0:36:9f:f0:ee:6e\",\"type\":\"ServerInterface\"},\"instance_interface_ips\":[]}],\"instance_credentials\":{\"ipmi\":{\"username\":null,\"initial_password\":null,\"ip_address\":null,\"version\":null,\"type\":\"IPMI\"},\"ilo\":{\"control_panel_url\":null,\"username\":null,\"initial_password\":null,\"type\":\"iLO\"},\"idrac\":{\"control_panel_url\":null,\"username\":null,\"initial_password\":null,\"type\":\"iDRAC\"},\"rdp\":{\"port\":null,\"username\":null,\"initial_password\":null,\"type\":\"RDP\"},\"ssh\":{\"port\":22,\"username\":\"root\",\"initial_password\":\"asdasd\",\"initial_ssh_keys\":{\"marius.boeru@metalsoft.io\":[{\"user_ssh_key_id\":38,\"user_id\":1,\"user_ssh_key\":\"ssh-rsa AAAasdasdasd\",\"user_ssh_key_created_timestamp\":\"2019-10-17T14:14:52Z\",\"user_ssh_key_status\":\"active\",\"type\":\"SSHKey\"}]},\"type\":\"SSH\"},\"remote_console\":{\"remote_protocol\":\"ssh\",\"remote_control_panel_url\":\"?product=instance&id=20639\",\"tunnel_path_url\":\"https://us-santaclara-api.poc.metalcloud.io/remote-console/instance-tunnel\",\"type\":\"RemoteConsole\"},\"telnet\":null,\"iscsi\":{\"username\":\"asdasd\",\"password\":\"asdad\",\"initiator_iqn\":\"iqn.2019-11.com.metalsoft.storage:instance-20639\",\"gateway\":\"100.64.0.1\",\"netmask\":\"255.255.255.248\",\"initiator_ip_address\":\"100.64.0.6\",\"type\":\"iSCSIInitiator\"},\"shared_drives\":[],\"ip_addresses_public\":[{\"ip_id\":59395,\"ip_hex\":\"fd1f8bbb56b308020000000000000002\",\"ip_human_readable\":\"fd1f:8bbb:56b3:0802:0000:0000:0000:0002\",\"ip_type\":\"ipv6\",\"instance_interface_id\":82707,\"subnet_id\":20504,\"ip_lease_expires\":\"0000-00-00T00:00:00Z\",\"ip_change_id\":91240,\"subnet_netmask_human_readable\":\"ffff:ffff:ffff:ffff:0000:0000:0000:0000\",\"subnet_gateway_human_readable\":\"fd1f:8bbb:56b3:0802:0000:0000:0000:0001\",\"subnet_destination\":\"wan\",\"ip_operation\":{\"ip_change_id\":91240,\"ip_id\":59395,\"ip_hex\":\"fd1f8bbb56b308020000000000000002\",\"ip_human_readable\":\"fd1f:8bbb:56b3:0802:0000:0000:0000:0002\",\"ip_type\":\"ipv6\",\"instance_interface_id\":82707,\"subnet_id\":20504,\"ip_lease_expires\":\"0000-00-00T00:00:00Z\",\"ip_updated_timestamp\":\"2019-11-27T14:38:25Z\",\"ip_deploy_type\":\"create\",\"ip_deploy_status\":\"finished\",\"ip_label\":\"ip-59395\",\"ip_subdomain\":\"ip-59395.subnet-20504.wan.demo.2.poc.metalcloud.io\",\"type\":\"IPOperation\"},\"type\":\"IP\"},{\"ip_id\":59396,\"ip_hex\":\"ac116a16\",\"ip_human_readable\":\"172.17.106.22\",\"ip_type\":\"ipv4\",\"instance_interface_id\":82707,\"subnet_id\":20505,\"ip_lease_expires\":\"0000-00-00T00:00:00Z\",\"ip_change_id\":91241,\"subnet_netmask_human_readable\":\"255.255.255.252\",\"subnet_gateway_human_readable\":\"172.17.106.21\",\"subnet_destination\":\"wan\",\"ip_operation\":{\"ip_change_id\":91241,\"ip_id\":59396,\"ip_hex\":\"ac116a16\",\"ip_human_readable\":\"172.17.106.22\",\"ip_type\":\"ipv4\",\"instance_interface_id\":82707,\"subnet_id\":20505,\"ip_lease_expires\":\"0000-00-00T00:00:00Z\",\"ip_updated_timestamp\":\"2019-11-27T14:38:25Z\",\"ip_deploy_type\":\"create\",\"ip_deploy_status\":\"finished\",\"ip_label\":\"ip-59396\",\"ip_subdomain\":\"ip-59396.subnet-20505.wan.demo.2.poc.metalcloud.io\",\"type\":\"IPOperation\"},\"type\":\"IP\"}],\"ip_addresses_private\":[],\"type\":\"InstanceCredentials\"}}"

const _instanceFixture3 = "{\"instance_id\":59422,\"instance_array_id\":37824,\"server_id\":104,\"server_type_id\":23,\"instance_change_id\":291972,\"instance_service_status\":\"active\",\"drive_id_bootable\":75242,\"instance_label\":\"instance-59422\",\"instance_subdomain\":\"instance-59422.instance-array-37824.vanilla.test-kube-csi.7.metalsoft.io\",\"instance_subdomain_permanent\":\"instance-59422.metalsoft.io\",\"instance_updated_timestamp\":\"2020-07-24T19:31:31Z\",\"instance_created_timestamp\":\"2020-07-24T19:29:59Z\",\"template_id_origin\":null,\"instance_operation\":{\"instance_change_id\":291972,\"instance_id\":59422,\"instance_array_id\":37824,\"server_id\":104,\"server_type_id\":23,\"instance_deploy_type\":\"create\",\"instance_deploy_status\":\"finished\",\"instance_label\":\"instance-59422\",\"instance_subdomain\":\"instance-59422.instance-array-37824.vanilla.test-kube-csi.7.metalsoft.io\",\"instance_updated_timestamp\":\"2020-07-24T19:30:04Z\",\"drive_id_bootable\":75242,\"template_id_origin\":null,\"type\":\"InstanceOperation\"},\"type\":\"Instance\",\"instance_tags\":[],\"instance_interfaces\":[{\"instance_interface_id\":237712,\"network_id\":61184,\"instance_interface_index\":0,\"instance_id\":59422,\"instance_interface_change_id\":845304,\"instance_interface_service_status\":\"active\",\"instance_interface_capacity_mbps\":1000,\"instance_interface_label\":\"if0\",\"instance_interface_subdomain\":\"if0.instance-59422.instance-array-37824.vanilla.test-kube-csi.7.metalsoft.io\",\"instance_interface_operation\":{\"instance_interface_change_id\":845304,\"instance_interface_id\":237712,\"network_id\":61184,\"instance_interface_index\":0,\"instance_id\":59422,\"instance_interface_deploy_type\":\"create\",\"instance_interface_deploy_status\":\"finished\",\"instance_interface_capacity_mbps\":1000,\"instance_interface_subdomain\":\"if0.instance-59422.instance-array-37824.vanilla.test-kube-csi.7.metalsoft.io\",\"instance_interface_label\":\"if0\",\"type\":\"InstanceInterfaceOperation\",\"instance_interface_lagg_indexes\":[]},\"type\":\"InstanceInterface\",\"instance_interface_lagg_indexes\":[],\"server_interface\":{\"server_interface_mac_address\":\"34:64:a9:95:0b:00\",\"type\":\"ServerInterface\"},\"instance_interface_ips\":[]},{\"instance_interface_id\":237713,\"network_id\":61183,\"instance_interface_index\":1,\"instance_id\":59422,\"instance_interface_change_id\":845301,\"instance_interface_service_status\":\"active\",\"instance_interface_capacity_mbps\":1000,\"instance_interface_label\":\"if1\",\"instance_interface_subdomain\":\"if1.instance-59422.instance-array-37824.vanilla.test-kube-csi.7.metalsoft.io\",\"instance_interface_operation\":{\"instance_interface_change_id\":845301,\"instance_interface_id\":237713,\"network_id\":61183,\"instance_interface_index\":1,\"instance_id\":59422,\"instance_interface_deploy_type\":\"create\",\"instance_interface_deploy_status\":\"finished\",\"instance_interface_capacity_mbps\":1000,\"instance_interface_subdomain\":\"if1.instance-59422.instance-array-37824.vanilla.test-kube-csi.7.metalsoft.io\",\"instance_interface_label\":\"if1\",\"type\":\"InstanceInterfaceOperation\",\"instance_interface_lagg_indexes\":[]},\"type\":\"InstanceInterface\",\"instance_interface_lagg_indexes\":[],\"server_interface\":{\"server_interface_mac_address\":\"34:64:a9:95:0b:01\",\"type\":\"ServerInterface\"},\"instance_interface_ips\":[{\"ip_id\":144213,\"ip_hex\":\"2a020cb8000028020000000000000002\",\"ip_human_readable\":\"2a02:0cb8:0000:2802:0000:0000:0000:0002\",\"ip_type\":\"ipv6\",\"instance_interface_id\":237713,\"subnet_id\":50944,\"ip_lease_expires\":\"0000-00-00T00:00:00Z\",\"ip_change_id\":224145,\"subnet_netmask_human_readable\":\"ffff:ffff:ffff:ffff:0000:0000:0000:0000\",\"subnet_gateway_human_readable\":\"2a02:0cb8:0000:2802:0000:0000:0000:0001\",\"subnet_destination\":\"wan\",\"ip_operation\":{\"ip_change_id\":224145,\"ip_id\":144213,\"ip_hex\":\"2a020cb8000028020000000000000002\",\"ip_human_readable\":\"2a02:0cb8:0000:2802:0000:0000:0000:0002\",\"ip_type\":\"ipv6\",\"instance_interface_id\":237713,\"subnet_id\":50944,\"ip_lease_expires\":\"0000-00-00T00:00:00Z\",\"ip_updated_timestamp\":\"2020-07-24T19:30:00Z\",\"ip_deploy_type\":\"create\",\"ip_deploy_status\":\"finished\",\"ip_label\":\"ip-144213\",\"ip_subdomain\":\"ip-144213.subnet-50944.wan.test-kube-csi.7.metalsoft.io\",\"type\":\"IPOperation\"},\"type\":\"IP\"},{\"ip_id\":144216,\"ip_hex\":\"54283d82\",\"ip_human_readable\":\"84.40.61.130\",\"ip_type\":\"ipv4\",\"instance_interface_id\":237713,\"subnet_id\":50945,\"ip_lease_expires\":\"0000-00-00T00:00:00Z\",\"ip_change_id\":224148,\"subnet_netmask_human_readable\":\"255.255.255.248\",\"subnet_gateway_human_readable\":\"84.40.61.129\",\"subnet_destination\":\"wan\",\"ip_operation\":{\"ip_change_id\":224148,\"ip_id\":144216,\"ip_hex\":\"54283d82\",\"ip_human_readable\":\"84.40.61.130\",\"ip_type\":\"ipv4\",\"instance_interface_id\":237713,\"subnet_id\":50945,\"ip_lease_expires\":\"0000-00-00T00:00:00Z\",\"ip_updated_timestamp\":\"2020-07-24T19:30:01Z\",\"ip_deploy_type\":\"create\",\"ip_deploy_status\":\"finished\",\"ip_label\":\"ip-144216\",\"ip_subdomain\":\"ip-144216.subnet-50945.wan.test-kube-csi.7.metalsoft.io\",\"type\":\"IPOperation\"},\"type\":\"IP\"}]},{\"instance_interface_id\":237714,\"network_id\":61185,\"instance_interface_index\":2,\"instance_id\":59422,\"instance_interface_change_id\":845307,\"instance_interface_service_status\":\"active\",\"instance_interface_capacity_mbps\":1000,\"instance_interface_label\":\"if2\",\"instance_interface_subdomain\":\"if2.instance-59422.instance-array-37824.vanilla.test-kube-csi.7.metalsoft.io\",\"instance_interface_operation\":{\"instance_interface_change_id\":845307,\"instance_interface_id\":237714,\"network_id\":61185,\"instance_interface_index\":2,\"instance_id\":59422,\"instance_interface_deploy_type\":\"create\",\"instance_interface_deploy_status\":\"finished\",\"instance_interface_capacity_mbps\":1000,\"instance_interface_subdomain\":\"if2.instance-59422.instance-array-37824.vanilla.test-kube-csi.7.metalsoft.io\",\"instance_interface_label\":\"if2\",\"type\":\"InstanceInterfaceOperation\",\"instance_interface_lagg_indexes\":[]},\"type\":\"InstanceInterface\",\"instance_interface_lagg_indexes\":[],\"server_interface\":{\"server_interface_mac_address\":\"0c:c4:7a:14:37:a8\",\"type\":\"ServerInterface\"},\"instance_interface_ips\":[]},{\"instance_interface_id\":237715,\"network_id\":null,\"instance_interface_index\":3,\"instance_id\":59422,\"instance_interface_change_id\":845292,\"instance_interface_service_status\":\"active\",\"instance_interface_capacity_mbps\":1000,\"instance_interface_label\":\"if3\",\"instance_interface_subdomain\":\"if3.instance-59422.instance-array-37824.vanilla.test-kube-csi.7.metalsoft.io\",\"instance_interface_operation\":{\"instance_interface_change_id\":845292,\"instance_interface_id\":237715,\"network_id\":null,\"instance_interface_index\":3,\"instance_id\":59422,\"instance_interface_deploy_type\":\"create\",\"instance_interface_deploy_status\":\"finished\",\"instance_interface_capacity_mbps\":1000,\"instance_interface_subdomain\":\"if3.instance-59422.instance-array-37824.vanilla.test-kube-csi.7.metalsoft.io\",\"instance_interface_label\":\"if3\",\"type\":\"InstanceInterfaceOperation\",\"instance_interface_lagg_indexes\":[]},\"type\":\"InstanceInterface\",\"instance_interface_lagg_indexes\":[],\"server_interface\":{\"server_interface_mac_address\":\"0c:c4:7a:14:37:a9\",\"type\":\"ServerInterface\"},\"instance_interface_ips\":[]},{\"instance_interface_id\":237724,\"network_id\":null,\"instance_interface_index\":4,\"instance_id\":59422,\"instance_interface_change_id\":845310,\"instance_interface_service_status\":\"active\",\"instance_interface_capacity_mbps\":0,\"instance_interface_label\":\"if4\",\"instance_interface_subdomain\":\"if4.instance-59422.instance-array-37824.vanilla.test-kube-csi.7.metalsoft.io\",\"instance_interface_operation\":{\"instance_interface_change_id\":845310,\"instance_interface_id\":237724,\"network_id\":null,\"instance_interface_index\":4,\"instance_id\":59422,\"instance_interface_deploy_type\":\"create\",\"instance_interface_deploy_status\":\"finished\",\"instance_interface_capacity_mbps\":0,\"instance_interface_subdomain\":\"if4.instance-59422.instance-array-37824.vanilla.test-kube-csi.7.metalsoft.io\",\"instance_interface_label\":\"if4\",\"type\":\"InstanceInterfaceOperation\",\"instance_interface_lagg_indexes\":[]},\"type\":\"InstanceInterface\",\"instance_interface_lagg_indexes\":[],\"server_interface\":{\"server_interface_mac_address\":null,\"type\":\"ServerInterface\"},\"instance_interface_ips\":[]}],\"instance_credentials\":{\"ipmi\":{\"username\":null,\"initial_password\":null,\"ip_address\":null,\"version\":null,\"type\":\"IPMI\"},\"ilo\":{\"control_panel_url\":null,\"username\":null,\"initial_password\":null,\"type\":\"iLO\"},\"idrac\":{\"control_panel_url\":null,\"username\":null,\"initial_password\":null,\"type\":\"iDRAC\"},\"rdp\":{\"port\":null,\"username\":null,\"initial_password\":null,\"type\":\"RDP\"},\"ssh\":{\"port\":22,\"username\":\"root\",\"initial_password\":\"R[NDAKpxg6WksWmn\",\"initial_ssh_keys\":{\"alex.bordei@metalsoft.io\":[{\"user_ssh_key_id\":44,\"user_id\":7,\"user_ssh_key\":\"ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQCphfy8RbUCMuwoLMox9ZDOO3YG4zUQ5G0/Bw5u5fcj8HLRtMIEY1vgKTHSgMP20Bbqye5wfNt5YiyZduDxOqnO1uEW3fPUwSgD2f/M1yiSoX6Hj/ayKPEGYtRkS77l0h64UCSFPEopwHzuBkmgdUET3G2wBUxaaUKcZY6UX7KgvpGFws6H4mx3ZC/qX+j6BkT5Qnx7jrUkdfpgXANar8pjkcVik7A5qpyhrwikAVggoaOT2pTSCWrV+z8i3+BbAQ/STLMWneFZqrS33nLOCs+KsdOiu0r8efSKUHOsCAmIWU/MupZkl3mpbLvX2722uOo2XYUDHVtasD8q+g2bgcqx+/zeo/7M2XX/G3V9fzhu46fS6HfWlCs2ariS+2HsG1Hn99PdpdxyTCglRQfUNA+ysbktHUKDYd/dzQN3EptF6tucB8TD76mi9X9K/BbtCTlSNwDGP3sW4hb0s8SVZCxJRvgieDnC47XsOYmrEymsejDyNgfK88+RHr3SVvKhNwi+WTFrChfO3Vx3nAjA34PH1bfDCmWrXknq3nsijjpsG6l4lBufR6531xevcxmC74RSj0bCGcpy+kqi7dtop3dj8ve/2b1fFoZwJ0jBXBgUgwWdY5lNIKtdpwiG41jtS2VIGeqkXSGIKihF8Gu7BLUUJ1GzIh4oqQFGIWuZSUeACQ== alexandrubordei@Alexs-Mac\",\"user_ssh_key_created_timestamp\":\"2015-06-02T20:27:57Z\",\"user_ssh_key_status\":\"active\",\"type\":\"SSHKey\"}],\"cristina.grosu@metalsoft.io\":[{\"user_ssh_key_id\":141,\"user_id\":779,\"user_ssh_key\":\"ssh-rsa AAAAB3NzaC1yc2EAAAABIwAAAQEA2rUCRELfeBdLIh0UDm5OFNIyaCdakqTxylCZ0+yybBLgSTL3MoOyKxjY3SLMq74XPALaZ9YVzxnKvW3q9dJNHkxrjHeDbx3QzjLi/s+3LXz7PMQmMB/59NUfdPB5Pc7Zq1HZ9vqpVh8Em4wHw5Pq/38mDoVXNGZ3KiI/nW+N6fUD/WohEd8PP4YyD4VfNgpB0uD+KJmkp/ZoHn9BW5j6Jq6PIh0aU13Fl1E66lwzOx6TKEZcVtohFrpdHGXEVI84oESwOb8wFUyPXzl+s8OLdyzUu6SFE9XtKZegLXtezMa4aTzt5tbFbYkFvskSXdWpAEQnaHC1KaBFULoh5u3hhw==\",\"user_ssh_key_created_timestamp\":\"2016-04-13T16:54:22Z\",\"user_ssh_key_status\":\"active\",\"type\":\"SSHKey\"},{\"user_ssh_key_id\":223,\"user_id\":779,\"user_ssh_key\":\"ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQDHLOfA2mr1PnjAT3A+ZnO2Wudt3882ePiaKGgfN7O3GS0yuWSVOrW2g0FbvxuW2WCSDYsx5duNleWD5mkW0SIZaIM2ed1f0id1fnW/r/Rlcz5Noq3VLBdXwcNoqDcn/3YwedZOkz3BGaIm6gpERKxLWPzNutGJ8SmjNOX02zleuFEm8tCrxmI8OgC1CAj7/nhaEPzwSbfJDKQYn/c3r0MlPqv8R/1ZGpLlyD7Meky35HzrtjqcnfoxfuTx6eBkptKC0PxSrhO79qT0bL+4XZDAMVk/eAltFmnJzu/9KsPqJ7OVduFu60N8YlwNJre1b0rUwiD3eu2YUvjFxgA7QeP+BFjjVhcIKJB4zGswZQRIrBkalZzyCYfA3Z72ymfAN1CZEU1xG0coJyLw2oL/Xpa2ek20I1YcCWKH9w/pGqD8I/aLrwxODUoRvBW1jwSQMZ/WMOxALQBorBvdwPEwCLa4/KAnOaibnWpkFoAiI3nWJd83tGgiAitjHN2nBVEzzC+KRTHfC5jD9Gx/VA1gEu+KTGEbLPwT8g8fIKaDIgyYtiH3OLBIgBoN90jgKOiQh49YbRPFjWC+rSJJYgSyttVfiRlW+KXQjyCNuIvyAQBpjN8g9QUjTSYkX0bBTLbA8NYabW2S1w6wyybfpx9JjTr4nlkCpLq/uQWBT9XhAiZu4Q==\",\"user_ssh_key_created_timestamp\":\"2017-09-28T11:15:09Z\",\"user_ssh_key_status\":\"active\",\"type\":\"SSHKey\"}]},\"type\":\"SSH\"},\"remote_console\":{\"remote_protocol\":\"ssh\",\"remote_control_panel_url\":\"?product=instance&id=59422\",\"tunnel_path_url\":\"https://uk-reading-api.metalsoft.io/remote-console/instance-tunnel\",\"type\":\"RemoteConsole\"},\"telnet\":null,\"iscsi\":{\"username\":\"sss\",\"password\":\"sssdsd\",\"initiator_iqn\":\"iqn.2020-07.com.metalsoft.storage:instance-59422\",\"gateway\":\"100.64.26.1\",\"netmask\":\"255.255.255.248\",\"initiator_ip_address\":\"100.64.26.6\",\"type\":\"iSCSIInitiator\"},\"shared_drives\":{\"csivolumename3\":{\"storage_ip_address\":\"100.96.0.192\",\"storage_port\":3260,\"target_iqn\":\"iqn.2013-01.io.metalsoft:storage.91nqwgd.4lys6bq.iel1v3k\",\"lun_id\":16,\"type\":\"iSCSI\"}},\"ip_addresses_public\":[{\"ip_id\":144213,\"ip_hex\":\"2a020cb8000028020000000000000002\",\"ip_human_readable\":\"2a02:0cb8:0000:2802:0000:0000:0000:0002\",\"ip_type\":\"ipv6\",\"instance_interface_id\":237713,\"subnet_id\":50944,\"ip_lease_expires\":\"0000-00-00T00:00:00Z\",\"ip_change_id\":224145,\"subnet_netmask_human_readable\":\"ffff:ffff:ffff:ffff:0000:0000:0000:0000\",\"subnet_gateway_human_readable\":\"2a02:0cb8:0000:2802:0000:0000:0000:0001\",\"subnet_destination\":\"wan\",\"ip_operation\":{\"ip_change_id\":224145,\"ip_id\":144213,\"ip_hex\":\"2a020cb8000028020000000000000002\",\"ip_human_readable\":\"2a02:0cb8:0000:2802:0000:0000:0000:0002\",\"ip_type\":\"ipv6\",\"instance_interface_id\":237713,\"subnet_id\":50944,\"ip_lease_expires\":\"0000-00-00T00:00:00Z\",\"ip_updated_timestamp\":\"2020-07-24T19:30:00Z\",\"ip_deploy_type\":\"create\",\"ip_deploy_status\":\"finished\",\"ip_label\":\"ip-144213\",\"ip_subdomain\":\"ip-144213.subnet-50944.wan.test-kube-csi.7.metalsoft.io\",\"type\":\"IPOperation\"},\"type\":\"IP\"},{\"ip_id\":144216,\"ip_hex\":\"54283d82\",\"ip_human_readable\":\"84.40.61.130\",\"ip_type\":\"ipv4\",\"instance_interface_id\":237713,\"subnet_id\":50945,\"ip_lease_expires\":\"0000-00-00T00:00:00Z\",\"ip_change_id\":224148,\"subnet_netmask_human_readable\":\"255.255.255.248\",\"subnet_gateway_human_readable\":\"84.40.61.129\",\"subnet_destination\":\"wan\",\"ip_operation\":{\"ip_change_id\":224148,\"ip_id\":144216,\"ip_hex\":\"54283d82\",\"ip_human_readable\":\"84.40.61.130\",\"ip_type\":\"ipv4\",\"instance_interface_id\":237713,\"subnet_id\":50945,\"ip_lease_expires\":\"0000-00-00T00:00:00Z\",\"ip_updated_timestamp\":\"2020-07-24T19:30:01Z\",\"ip_deploy_type\":\"create\",\"ip_deploy_status\":\"finished\",\"ip_label\":\"ip-144216\",\"ip_subdomain\":\"ip-144216.subnet-50945.wan.test-kube-csi.7.metalsoft.io\",\"type\":\"IPOperation\"},\"type\":\"IP\"}],\"ip_addresses_private\":[],\"type\":\"InstanceCredentials\"}}"
