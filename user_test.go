package metalcloud

import (
	"encoding/json"
	"testing"

	. "github.com/onsi/gomega"
)

func TestUserSearch(t *testing.T) {
	RegisterTestingT(t)

	responseBody = _UserSearchFixture

	mc, err := GetMetalcloudClient("userEmail", "APIKey", httpServer.URL, false, "", "", "")
	Expect(err).To(BeNil())

	ret, err := mc.UserSearch("*")
	Expect(err).To(BeNil())
	Expect(ret).NotTo(BeNil())
	r := *ret

	Expect(r[0].UserID).To(Equal(8))
	Expect(r[0].UserCreatedTimestamp).To(Equal("2021-02-17T10:58:45Z"))

	Expect(r[1].UserID).To(Equal(7))
	Expect(r[1].UserCreatedTimestamp).To(Equal("2021-02-17T10:58:42Z"))

	body := (<-requestChan).body

	var m map[string]interface{}
	err2 := json.Unmarshal([]byte(body), &m)
	Expect(err2).To(BeNil())
	Expect(m).NotTo(BeNil())
}

const _UserSearchFixture = "{\"result\":{\"_users\":{\"duration_milliseconds\":0.013108968734741211,\"rows\":[{\"franchise\":\"uk.metalsoft.io\",\"user_access_level\":\"user\",\"user_auth_failed_attempts_since_last_login\":0,\"user_authenticator_created_timestamp\":\"1970-01-01T00:00:00Z\",\"user_authenticator_is_mandatory\":false,\"user_authenticator_must_change\":false,\"user_blocked\":false,\"user_brand\":\"default\",\"user_created_timestamp\":\"2021-02-17T10:58:45Z\",\"user_custom_prices_json\":null,\"user_display_name\":\"bsi@development\",\"user_email\":\"kerberos_d4f9ce582f5b21c99446429221eb0a31@void.metalsoft.io\",\"user_email_status\":\"verified\",\"user_exclude_from_reports\":false,\"user_experimental_tags_json\":null,\"user_external_ids_json\":null,\"user_gui_settings_json\":null,\"user_id\":8,\"user_infrastructure_id_default\":null,\"user_is_billable\":false,\"user_is_brand_manager\":false,\"user_is_datastore_publisher\":false,\"user_is_suspended\":false,\"user_is_test_account\":false,\"user_is_testing_mode\":false,\"user_kerberos_principal_name\":\"bsi@development\",\"user_language\":\"en\",\"user_last_login_timestamp\":\"0000-00-00T00:00:00Z\",\"user_last_login_type\":\"md5\",\"user_limits_json\":null,\"user_password_change_required\":false,\"user_permissions_json\":null,\"user_plan_type\":\"vanilla\",\"user_promotion_tags_json\":null},{\"franchise\":\"uk.metalsoft.io\",\"user_access_level\":\"user\",\"user_auth_failed_attempts_since_last_login\":0,\"user_authenticator_created_timestamp\":\"1970-01-01T00:00:00Z\",\"user_authenticator_is_mandatory\":false,\"user_authenticator_must_change\":false,\"user_blocked\":false,\"user_brand\":\"default\",\"user_created_timestamp\":\"2021-02-17T10:58:42Z\",\"user_custom_prices_json\":null,\"user_display_name\":\"kiprop/vm@development\",\"user_email\":\"kerberos_169017a009b3800275f22042a791b3ea@void.metalsoft.io\",\"user_email_status\":\"verified\",\"user_exclude_from_reports\":false,\"user_experimental_tags_json\":null,\"user_external_ids_json\":null,\"user_gui_settings_json\":null,\"user_id\":7,\"user_infrastructure_id_default\":null,\"user_is_billable\":false,\"user_is_brand_manager\":false,\"user_is_datastore_publisher\":false,\"user_is_suspended\":false,\"user_is_test_account\":false,\"user_is_testing_mode\":false,\"user_kerberos_principal_name\":\"kiprop/vm@development\",\"user_language\":\"en\",\"user_last_login_timestamp\":\"0000-00-00T00:00:00Z\",\"user_last_login_type\":\"md5\",\"user_limits_json\":null,\"user_password_change_required\":false,\"user_permissions_json\":null,\"user_plan_type\":\"vanilla\",\"user_promotion_tags_json\":null},{\"franchise\":\"uk.metalsoft.io\",\"user_access_level\":\"user\",\"user_auth_failed_attempts_since_last_login\":0,\"user_authenticator_created_timestamp\":\"1970-01-01T00:00:00Z\",\"user_authenticator_is_mandatory\":false,\"user_authenticator_must_change\":false,\"user_blocked\":false,\"user_brand\":\"default\",\"user_created_timestamp\":\"2021-02-17T10:58:40Z\",\"user_custom_prices_json\":null,\"user_display_name\":\"kadmin/changepw@development\",\"user_email\":\"kerberos_f4d670f0f217597904258ded485f4e35@void.metalsoft.io\",\"user_email_status\":\"verified\",\"user_exclude_from_reports\":false,\"user_experimental_tags_json\":null,\"user_external_ids_json\":null,\"user_gui_settings_json\":null,\"user_id\":6,\"user_infrastructure_id_default\":null,\"user_is_billable\":false,\"user_is_brand_manager\":false,\"user_is_datastore_publisher\":false,\"user_is_suspended\":false,\"user_is_test_account\":false,\"user_is_testing_mode\":false,\"user_kerberos_principal_name\":\"kadmin/changepw@development\",\"user_language\":\"en\",\"user_last_login_timestamp\":\"0000-00-00T00:00:00Z\",\"user_last_login_type\":\"md5\",\"user_limits_json\":null,\"user_password_change_required\":false,\"user_permissions_json\":null,\"user_plan_type\":\"vanilla\",\"user_promotion_tags_json\":null},{\"franchise\":\"uk.metalsoft.io\",\"user_access_level\":\"user\",\"user_auth_failed_attempts_since_last_login\":0,\"user_authenticator_created_timestamp\":\"1970-01-01T00:00:00Z\",\"user_authenticator_is_mandatory\":false,\"user_authenticator_must_change\":false,\"user_blocked\":false,\"user_brand\":\"default\",\"user_created_timestamp\":\"2021-02-17T10:58:37Z\",\"user_custom_prices_json\":null,\"user_display_name\":\"kadmin/admin@development\",\"user_email\":\"kerberos_1ff739969d0875d86bffb479f0b98376@void.metalsoft.io\",\"user_email_status\":\"verified\",\"user_exclude_from_reports\":false,\"user_experimental_tags_json\":null,\"user_external_ids_json\":null,\"user_gui_settings_json\":null,\"user_id\":5,\"user_infrastructure_id_default\":null,\"user_is_billable\":false,\"user_is_brand_manager\":false,\"user_is_datastore_publisher\":false,\"user_is_suspended\":false,\"user_is_test_account\":false,\"user_is_testing_mode\":false,\"user_kerberos_principal_name\":\"kadmin/admin@development\",\"user_language\":\"en\",\"user_last_login_timestamp\":\"0000-00-00T00:00:00Z\",\"user_last_login_type\":\"md5\",\"user_limits_json\":null,\"user_password_change_required\":false,\"user_permissions_json\":null,\"user_plan_type\":\"vanilla\",\"user_promotion_tags_json\":null},{\"franchise\":\"uk.metalsoft.io\",\"user_access_level\":\"user\",\"user_auth_failed_attempts_since_last_login\":0,\"user_authenticator_created_timestamp\":\"1970-01-01T00:00:00Z\",\"user_authenticator_is_mandatory\":false,\"user_authenticator_must_change\":false,\"user_blocked\":false,\"user_brand\":\"default\",\"user_created_timestamp\":\"2021-02-17T10:58:34Z\",\"user_custom_prices_json\":null,\"user_display_name\":\"kadmin/vm@development\",\"user_email\":\"kerberos_5cc6b51846be477118f40122ed484cee@void.metalsoft.io\",\"user_email_status\":\"verified\",\"user_exclude_from_reports\":false,\"user_experimental_tags_json\":null,\"user_external_ids_json\":null,\"user_gui_settings_json\":null,\"user_id\":4,\"user_infrastructure_id_default\":null,\"user_is_billable\":false,\"user_is_brand_manager\":false,\"user_is_datastore_publisher\":false,\"user_is_suspended\":false,\"user_is_test_account\":false,\"user_is_testing_mode\":false,\"user_kerberos_principal_name\":\"kadmin/vm@development\",\"user_language\":\"en\",\"user_last_login_timestamp\":\"0000-00-00T00:00:00Z\",\"user_last_login_type\":\"md5\",\"user_limits_json\":null,\"user_password_change_required\":false,\"user_permissions_json\":null,\"user_plan_type\":\"vanilla\",\"user_promotion_tags_json\":null},{\"franchise\":\"uk.metalsoft.io\",\"user_access_level\":\"user\",\"user_auth_failed_attempts_since_last_login\":0,\"user_authenticator_created_timestamp\":\"1970-01-01T00:00:00Z\",\"user_authenticator_is_mandatory\":false,\"user_authenticator_must_change\":false,\"user_blocked\":false,\"user_brand\":\"default\",\"user_created_timestamp\":\"2021-02-17T10:58:31Z\",\"user_custom_prices_json\":null,\"user_display_name\":\"krbtgt/development@development\",\"user_email\":\"kerberos_d55e604c3ccca09b977d3ea514336b3f@void.metalsoft.io\",\"user_email_status\":\"verified\",\"user_exclude_from_reports\":false,\"user_experimental_tags_json\":null,\"user_external_ids_json\":null,\"user_gui_settings_json\":null,\"user_id\":3,\"user_infrastructure_id_default\":null,\"user_is_billable\":false,\"user_is_brand_manager\":false,\"user_is_datastore_publisher\":false,\"user_is_suspended\":false,\"user_is_test_account\":false,\"user_is_testing_mode\":false,\"user_kerberos_principal_name\":\"krbtgt/development@development\",\"user_language\":\"en\",\"user_last_login_timestamp\":\"0000-00-00T00:00:00Z\",\"user_last_login_type\":\"md5\",\"user_limits_json\":null,\"user_password_change_required\":false,\"user_permissions_json\":null,\"user_plan_type\":\"vanilla\",\"user_promotion_tags_json\":null},{\"franchise\":\"uk.metalsoft.io\",\"user_access_level\":\"user\",\"user_auth_failed_attempts_since_last_login\":0,\"user_authenticator_created_timestamp\":\"1970-01-01T00:00:00Z\",\"user_authenticator_is_mandatory\":false,\"user_authenticator_must_change\":false,\"user_blocked\":false,\"user_brand\":\"default\",\"user_created_timestamp\":\"2021-02-17T10:58:29Z\",\"user_custom_prices_json\":null,\"user_display_name\":\"K/M@development\",\"user_email\":\"kerberos_5dd20972af05766fb268833bd25705c3@void.metalsoft.io\",\"user_email_status\":\"verified\",\"user_exclude_from_reports\":false,\"user_experimental_tags_json\":null,\"user_external_ids_json\":null,\"user_gui_settings_json\":null,\"user_id\":2,\"user_infrastructure_id_default\":null,\"user_is_billable\":false,\"user_is_brand_manager\":false,\"user_is_datastore_publisher\":false,\"user_is_suspended\":false,\"user_is_test_account\":false,\"user_is_testing_mode\":false,\"user_kerberos_principal_name\":\"K/M@development\",\"user_language\":\"en\",\"user_last_login_timestamp\":\"0000-00-00T00:00:00Z\",\"user_last_login_type\":\"md5\",\"user_limits_json\":null,\"user_password_change_required\":false,\"user_permissions_json\":null,\"user_plan_type\":\"vanilla\",\"user_promotion_tags_json\":null},{\"franchise\":\"uk.metalsoft.io\",\"user_access_level\":\"root\",\"user_auth_failed_attempts_since_last_login\":0,\"user_authenticator_created_timestamp\":\"1970-01-01T00:00:00Z\",\"user_authenticator_is_mandatory\":false,\"user_authenticator_must_change\":false,\"user_blocked\":false,\"user_brand\":\"default\",\"user_created_timestamp\":\"2021-02-17T10:35:51Z\",\"user_custom_prices_json\":null,\"user_display_name\":\"alexandru.corman@metalsoft.io\",\"user_email\":\"alexandru.corman@metalsoft.io\",\"user_email_status\":\"verified\",\"user_exclude_from_reports\":false,\"user_experimental_tags_json\":null,\"user_external_ids_json\":null,\"user_gui_settings_json\":null,\"user_id\":1,\"user_infrastructure_id_default\":null,\"user_is_billable\":true,\"user_is_brand_manager\":false,\"user_is_datastore_publisher\":false,\"user_is_suspended\":false,\"user_is_test_account\":false,\"user_is_testing_mode\":false,\"user_kerberos_principal_name\":null,\"user_language\":\"en\",\"user_last_login_timestamp\":\"2021-08-23T08:38:05Z\",\"user_last_login_type\":\"md5\",\"user_limits_json\":null,\"user_password_change_required\":false,\"user_permissions_json\":\"{\\\"admin_password_reveal\\\":{\\\"hdfs\\\":true,\\\"drive\\\":true,\\\"server\\\":true,\\\"switch\\\":true,\\\"cluster\\\":true,\\\"instance\\\":true,\\\"workflow\\\":true,\\\"container\\\":true,\\\"datacenter\\\":true,\\\"chassis_rack\\\":true,\\\"storage_pool\\\":true}}\",\"user_plan_type\":\"vanilla\",\"user_promotion_tags_json\":null}],\"rows_total\":8,\"rows_order\":[[\"user_id\",\"DESC\"]],\"xls_for_pivot_tables_download_url\":\"https://bsi.vmware.local/api/url?rqi=br.ztB-JV5Hvuz-byKjYL9PARA6pd7SaNYGLlq1TSBwiwA0lztCz6NXmLgYQCdQWAT_K9QvMnU_AC5DRZ9dmBY5yX0GNS5GSZvTQ-YZz85VCiWQKVgw2y5TyLFD4cJyE360xxzlDo3McKREOaoGNcuAv7guqFthELlODXxbA9PMbInsTKII9z3SD39QD41pjqOANK6JdBFJ317s82ZlvePZfy0bI96WilUz16DeEHP3TDdT6abg0luHOyE2roc1TZKYY7dn-Zi2katJbWB0c1jcPLbJnF9MawnF5yOnNMbu5cqC6dBMq96wvEsOpIhXDC_CPPGU3HTOB1JkMaTOkwWwpbyaOcj4y0mSxhp8l9Sgx-EfSrkTYqh7AqxI0qhFJhpZACJEDeCuAWhU7C0BVhueet0g1Prn-TgY-8TIhMljATxsGrrlhvRnZznE-YmDkC0g06nPzFOkz2edzRrYDpUK0N5gkBcXcZzEGXd1gjIAxK04Y77IWD4TyMJb27zQ1CONlNkD_fuJohU6rFlTCTwzoctmLtQg13taYo_RLEDYr0PfJU9-EzFRUz3W_NeFjDX96b4wm4LnNz1YpbvciBbj4b6RfponbSRzo0eD5ZxLyYYwNnnpnpMLhcBBGUC1d5Z00d0TgnOHGW1uIfJtABVVItw7-Mv--Vs1IpvL7YCmF197tujlb1IE9rPdraSqMeRPruPWbH-hueg-seebyhbjrH_yZeeg54a9Zh-drxR_Oim_3oZirc3c_WT8ZTTgUzxiexiuF-Gl-bA-OzGGvgs1D3HSGSdhba0NZwuCHM_lvU3LRlRteS1lbiHomVlCuD5hlZmMNjVDp1uWn2emT7t2DmFi1K5Jvsqb0FmVOOpS4UX9fDxd9MyXj2ItJDlfaG--WGyiln6Tj8CbwKJ6P1padm7tuI3HuNBc-IbcpbkFe-DPLdqe00IefEVaPi_16ZSM&v=F7eY3P-R0NeQMnDtaeVJoA\"}},\"jsonrpc\":\"2.0\",\"id\":3}"
