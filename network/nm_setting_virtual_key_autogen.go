// This file is automatically generated, please don't edit manully.
package main

// All virtual keys data
var virtualKeys = []virtualKey{
	virtualKey{NM_SETTING_VK_802_1X_ENABLE, ktypeBoolean, NM_SETTING_802_1X_SETTING_NAME, NM_SETTING_VK_NONE_RELATED_KEY, false, true, false},
	virtualKey{NM_SETTING_VK_802_1X_EAP, ktypeString, NM_SETTING_802_1X_SETTING_NAME, NM_SETTING_802_1X_EAP, false, true, false},
	virtualKey{NM_SETTING_VK_802_1X_PAC_FILE, ktypeString, NM_SETTING_802_1X_SETTING_NAME, NM_SETTING_802_1X_PAC_FILE, false, true, false},
	virtualKey{NM_SETTING_VK_802_1X_CA_CERT, ktypeString, NM_SETTING_802_1X_SETTING_NAME, NM_SETTING_802_1X_CA_CERT, false, true, false},
	virtualKey{NM_SETTING_VK_802_1X_CLIENT_CERT, ktypeString, NM_SETTING_802_1X_SETTING_NAME, NM_SETTING_802_1X_CLIENT_CERT, false, true, false},
	virtualKey{NM_SETTING_VK_802_1X_PRIVATE_KEY, ktypeString, NM_SETTING_802_1X_SETTING_NAME, NM_SETTING_802_1X_PRIVATE_KEY, false, true, false},
	virtualKey{NM_SETTING_VK_CONNECTION_NO_PERMISSION, ktypeBoolean, NM_SETTING_CONNECTION_SETTING_NAME, NM_SETTING_CONNECTION_PERMISSIONS, false, true, false},
	virtualKey{NM_SETTING_VK_IP4_CONFIG_ADDRESSES_ADDRESS, ktypeString, NM_SETTING_IP4_CONFIG_SETTING_NAME, NM_SETTING_IP4_CONFIG_ADDRESSES, false, true, false},
	virtualKey{NM_SETTING_VK_IP4_CONFIG_ADDRESSES_MASK, ktypeString, NM_SETTING_IP4_CONFIG_SETTING_NAME, NM_SETTING_IP4_CONFIG_ADDRESSES, false, true, false},
	virtualKey{NM_SETTING_VK_IP4_CONFIG_ADDRESSES_GATEWAY, ktypeString, NM_SETTING_IP4_CONFIG_SETTING_NAME, NM_SETTING_IP4_CONFIG_ADDRESSES, false, true, true},
	virtualKey{NM_SETTING_VK_IP4_CONFIG_DNS, ktypeString, NM_SETTING_IP4_CONFIG_SETTING_NAME, NM_SETTING_IP4_CONFIG_DNS, false, true, false},
	virtualKey{NM_SETTING_VK_IP4_CONFIG_ROUTES_ADDRESS, ktypeString, NM_SETTING_IP4_CONFIG_SETTING_NAME, NM_SETTING_IP4_CONFIG_ROUTES, false, false, false},
	virtualKey{NM_SETTING_VK_IP4_CONFIG_ROUTES_MASK, ktypeString, NM_SETTING_IP4_CONFIG_SETTING_NAME, NM_SETTING_IP4_CONFIG_ROUTES, false, false, false},
	virtualKey{NM_SETTING_VK_IP4_CONFIG_ROUTES_NEXTHOP, ktypeString, NM_SETTING_IP4_CONFIG_SETTING_NAME, NM_SETTING_IP4_CONFIG_ROUTES, false, false, false},
	virtualKey{NM_SETTING_VK_IP4_CONFIG_ROUTES_METRIC, ktypeString, NM_SETTING_IP4_CONFIG_SETTING_NAME, NM_SETTING_IP4_CONFIG_ROUTES, false, false, false},
	virtualKey{NM_SETTING_VK_IP6_CONFIG_ADDRESSES_ADDRESS, ktypeString, NM_SETTING_IP6_CONFIG_SETTING_NAME, NM_SETTING_IP6_CONFIG_ADDRESSES, false, true, false},
	virtualKey{NM_SETTING_VK_IP6_CONFIG_ADDRESSES_PREFIX, ktypeUint32, NM_SETTING_IP6_CONFIG_SETTING_NAME, NM_SETTING_IP6_CONFIG_ADDRESSES, false, true, false},
	virtualKey{NM_SETTING_VK_IP6_CONFIG_ADDRESSES_GATEWAY, ktypeString, NM_SETTING_IP6_CONFIG_SETTING_NAME, NM_SETTING_IP6_CONFIG_ADDRESSES, false, true, true},
	virtualKey{NM_SETTING_VK_IP6_CONFIG_DNS, ktypeString, NM_SETTING_IP6_CONFIG_SETTING_NAME, NM_SETTING_IP6_CONFIG_DNS, false, true, false},
	virtualKey{NM_SETTING_VK_IP6_CONFIG_ROUTES_ADDRESS, ktypeString, NM_SETTING_IP6_CONFIG_SETTING_NAME, NM_SETTING_IP6_CONFIG_ROUTES, false, false, false},
	virtualKey{NM_SETTING_VK_IP6_CONFIG_ROUTES_PREFIX, ktypeUint32, NM_SETTING_IP6_CONFIG_SETTING_NAME, NM_SETTING_IP6_CONFIG_ROUTES, false, false, false},
	virtualKey{NM_SETTING_VK_IP6_CONFIG_ROUTES_NEXTHOP, ktypeString, NM_SETTING_IP6_CONFIG_SETTING_NAME, NM_SETTING_IP6_CONFIG_ROUTES, false, false, false},
	virtualKey{NM_SETTING_VK_IP6_CONFIG_ROUTES_METRIC, ktypeUint32, NM_SETTING_IP6_CONFIG_SETTING_NAME, NM_SETTING_IP6_CONFIG_ROUTES, false, false, false},
	virtualKey{NM_SETTING_VK_PPP_ENABLE_LCP_ECHO, ktypeBoolean, NM_SETTING_PPP_SETTING_NAME, NM_SETTING_PPP_LCP_ECHO_FAILURE, false, true, false},
	virtualKey{NM_SETTING_VK_VPN_L2TP_ENABLE_LCP_ECHO, ktypeBoolean, NM_SETTING_VF_VPN_L2TP_PPP_SETTING_NAME, NM_SETTING_VPN_L2TP_KEY_LCP_ECHO_FAILURE, false, true, false},
	virtualKey{NM_SETTING_VK_VPN_OPENVPN_KEY_ENABLE_PORT, ktypeBoolean, NM_SETTING_VF_VPN_OPENVPN_ADVANCED_SETTING_NAME, NM_SETTING_VPN_OPENVPN_KEY_PORT, true, true, false},
	virtualKey{NM_SETTING_VK_VPN_OPENVPN_KEY_ENABLE_RENEG_SECONDS, ktypeBoolean, NM_SETTING_VF_VPN_OPENVPN_ADVANCED_SETTING_NAME, NM_SETTING_VPN_OPENVPN_KEY_RENEG_SECONDS, true, true, false},
	virtualKey{NM_SETTING_VK_VPN_OPENVPN_KEY_ENABLE_TUNNEL_MTU, ktypeBoolean, NM_SETTING_VF_VPN_OPENVPN_ADVANCED_SETTING_NAME, NM_SETTING_VPN_OPENVPN_KEY_TUNNEL_MTU, true, true, false},
	virtualKey{NM_SETTING_VK_VPN_OPENVPN_KEY_ENABLE_FRAGMENT_SIZE, ktypeBoolean, NM_SETTING_VF_VPN_OPENVPN_ADVANCED_SETTING_NAME, NM_SETTING_VPN_OPENVPN_KEY_FRAGMENT_SIZE, true, true, false},
	virtualKey{NM_SETTING_VK_VPN_OPENVPN_KEY_ENABLE_STATIC_KEY_DIRECTION, ktypeBoolean, NM_SETTING_VF_VPN_OPENVPN_SETTING_NAME, NM_SETTING_VPN_OPENVPN_KEY_STATIC_KEY_DIRECTION, true, true, false},
	virtualKey{NM_SETTING_VK_VPN_OPENVPN_KEY_ENABLE_TA_DIR, ktypeBoolean, NM_SETTING_VF_VPN_OPENVPN_TLSAUTH_SETTING_NAME, NM_SETTING_VPN_OPENVPN_KEY_TA_DIR, true, true, false},
	virtualKey{NM_SETTING_VK_VPN_PPTP_ENABLE_LCP_ECHO, ktypeBoolean, NM_SETTING_VF_VPN_PPTP_PPP_SETTING_NAME, NM_SETTING_VPN_PPTP_KEY_LCP_ECHO_FAILURE, false, true, false},
	virtualKey{NM_SETTING_VK_VPN_VPNC_KEY_HYBRID_AUTHMODE, ktypeBoolean, NM_SETTING_VF_VPN_VPNC_SETTING_NAME, NM_SETTING_VPN_VPNC_KEY_AUTHMODE, false, true, false},
	virtualKey{NM_SETTING_VK_VPN_VPNC_KEY_ENCRYPTION_METHOD, ktypeString, NM_SETTING_VF_VPN_VPNC_ADVANCED_SETTING_NAME, NM_SETTING_VPN_VPNC_KEY_SINGLE_DES, false, true, false},
	virtualKey{NM_SETTING_VK_VPN_VPNC_KEY_DISABLE_DPD, ktypeBoolean, NM_SETTING_VF_VPN_VPNC_ADVANCED_SETTING_NAME, NM_SETTING_VPN_VPNC_KEY_DPD_IDLE_TIMEOUT, false, true, false},
	virtualKey{NM_SETTING_VK_WIRED_MTU, ktypeBoolean, NM_SETTING_WIRED_SETTING_NAME, NM_SETTING_WIRED_MTU, true, true, false},
	virtualKey{NM_SETTING_VK_WIRELESS_SECURITY_KEY_MGMT, ktypeString, NM_SETTING_WIRELESS_SECURITY_SETTING_NAME, NM_SETTING_WIRELESS_SECURITY_KEY_MGMT, false, true, false},
}

// Get JSON value generally
func generalGetVirtualKeyJSON(data connectionData, field, key string) (valueJSON string) {
	switch field {
	case NM_SETTING_802_1X_SETTING_NAME:
		switch key {
		case NM_SETTING_VK_802_1X_ENABLE:
			return getSettingVk8021xEnableJSON(data)
		case NM_SETTING_VK_802_1X_EAP:
			return getSettingVk8021xEapJSON(data)
		case NM_SETTING_VK_802_1X_PAC_FILE:
			return getSettingVk8021xPacFileJSON(data)
		case NM_SETTING_VK_802_1X_CA_CERT:
			return getSettingVk8021xCaCertJSON(data)
		case NM_SETTING_VK_802_1X_CLIENT_CERT:
			return getSettingVk8021xClientCertJSON(data)
		case NM_SETTING_VK_802_1X_PRIVATE_KEY:
			return getSettingVk8021xPrivateKeyJSON(data)
		}
	case NM_SETTING_CONNECTION_SETTING_NAME:
		switch key {
		case NM_SETTING_VK_CONNECTION_NO_PERMISSION:
			return getSettingVkConnectionNoPermissionJSON(data)
		}
	case NM_SETTING_IP4_CONFIG_SETTING_NAME:
		switch key {
		case NM_SETTING_VK_IP4_CONFIG_ADDRESSES_ADDRESS:
			return getSettingVkIp4ConfigAddressesAddressJSON(data)
		case NM_SETTING_VK_IP4_CONFIG_ADDRESSES_MASK:
			return getSettingVkIp4ConfigAddressesMaskJSON(data)
		case NM_SETTING_VK_IP4_CONFIG_ADDRESSES_GATEWAY:
			return getSettingVkIp4ConfigAddressesGatewayJSON(data)
		case NM_SETTING_VK_IP4_CONFIG_DNS:
			return getSettingVkIp4ConfigDnsJSON(data)
		case NM_SETTING_VK_IP4_CONFIG_ROUTES_ADDRESS:
			return getSettingVkIp4ConfigRoutesAddressJSON(data)
		case NM_SETTING_VK_IP4_CONFIG_ROUTES_MASK:
			return getSettingVkIp4ConfigRoutesMaskJSON(data)
		case NM_SETTING_VK_IP4_CONFIG_ROUTES_NEXTHOP:
			return getSettingVkIp4ConfigRoutesNexthopJSON(data)
		case NM_SETTING_VK_IP4_CONFIG_ROUTES_METRIC:
			return getSettingVkIp4ConfigRoutesMetricJSON(data)
		}
	case NM_SETTING_IP6_CONFIG_SETTING_NAME:
		switch key {
		case NM_SETTING_VK_IP6_CONFIG_ADDRESSES_ADDRESS:
			return getSettingVkIp6ConfigAddressesAddressJSON(data)
		case NM_SETTING_VK_IP6_CONFIG_ADDRESSES_PREFIX:
			return getSettingVkIp6ConfigAddressesPrefixJSON(data)
		case NM_SETTING_VK_IP6_CONFIG_ADDRESSES_GATEWAY:
			return getSettingVkIp6ConfigAddressesGatewayJSON(data)
		case NM_SETTING_VK_IP6_CONFIG_DNS:
			return getSettingVkIp6ConfigDnsJSON(data)
		case NM_SETTING_VK_IP6_CONFIG_ROUTES_ADDRESS:
			return getSettingVkIp6ConfigRoutesAddressJSON(data)
		case NM_SETTING_VK_IP6_CONFIG_ROUTES_PREFIX:
			return getSettingVkIp6ConfigRoutesPrefixJSON(data)
		case NM_SETTING_VK_IP6_CONFIG_ROUTES_NEXTHOP:
			return getSettingVkIp6ConfigRoutesNexthopJSON(data)
		case NM_SETTING_VK_IP6_CONFIG_ROUTES_METRIC:
			return getSettingVkIp6ConfigRoutesMetricJSON(data)
		}
	case NM_SETTING_PPP_SETTING_NAME:
		switch key {
		case NM_SETTING_VK_PPP_ENABLE_LCP_ECHO:
			return getSettingVkPppEnableLcpEchoJSON(data)
		}
	case NM_SETTING_VF_VPN_L2TP_PPP_SETTING_NAME:
		switch key {
		case NM_SETTING_VK_VPN_L2TP_ENABLE_LCP_ECHO:
			return getSettingVkVpnL2tpEnableLcpEchoJSON(data)
		}
	case NM_SETTING_VF_VPN_OPENVPN_ADVANCED_SETTING_NAME:
		switch key {
		case NM_SETTING_VK_VPN_OPENVPN_KEY_ENABLE_PORT:
			return getSettingVkVpnOpenvpnKeyEnablePortJSON(data)
		case NM_SETTING_VK_VPN_OPENVPN_KEY_ENABLE_RENEG_SECONDS:
			return getSettingVkVpnOpenvpnKeyEnableRenegSecondsJSON(data)
		case NM_SETTING_VK_VPN_OPENVPN_KEY_ENABLE_TUNNEL_MTU:
			return getSettingVkVpnOpenvpnKeyEnableTunnelMtuJSON(data)
		case NM_SETTING_VK_VPN_OPENVPN_KEY_ENABLE_FRAGMENT_SIZE:
			return getSettingVkVpnOpenvpnKeyEnableFragmentSizeJSON(data)
		}
	case NM_SETTING_VF_VPN_OPENVPN_SETTING_NAME:
		switch key {
		case NM_SETTING_VK_VPN_OPENVPN_KEY_ENABLE_STATIC_KEY_DIRECTION:
			return getSettingVkVpnOpenvpnKeyEnableStaticKeyDirectionJSON(data)
		}
	case NM_SETTING_VF_VPN_OPENVPN_TLSAUTH_SETTING_NAME:
		switch key {
		case NM_SETTING_VK_VPN_OPENVPN_KEY_ENABLE_TA_DIR:
			return getSettingVkVpnOpenvpnKeyEnableTaDirJSON(data)
		}
	case NM_SETTING_VF_VPN_PPTP_PPP_SETTING_NAME:
		switch key {
		case NM_SETTING_VK_VPN_PPTP_ENABLE_LCP_ECHO:
			return getSettingVkVpnPptpEnableLcpEchoJSON(data)
		}
	case NM_SETTING_VF_VPN_VPNC_SETTING_NAME:
		switch key {
		case NM_SETTING_VK_VPN_VPNC_KEY_HYBRID_AUTHMODE:
			return getSettingVkVpnVpncKeyHybridAuthmodeJSON(data)
		}
	case NM_SETTING_VF_VPN_VPNC_ADVANCED_SETTING_NAME:
		switch key {
		case NM_SETTING_VK_VPN_VPNC_KEY_ENCRYPTION_METHOD:
			return getSettingVkVpnVpncKeyEncryptionMethodJSON(data)
		case NM_SETTING_VK_VPN_VPNC_KEY_DISABLE_DPD:
			return getSettingVkVpnVpncKeyDisableDpdJSON(data)
		}
	case NM_SETTING_WIRED_SETTING_NAME:
		switch key {
		case NM_SETTING_VK_WIRED_MTU:
			return getSettingVkWiredMtuJSON(data)
		}
	case NM_SETTING_WIRELESS_SECURITY_SETTING_NAME:
		switch key {
		case NM_SETTING_VK_WIRELESS_SECURITY_KEY_MGMT:
			return getSettingVkWirelessSecurityKeyMgmtJSON(data)
		}
	}
	logger.Error("invalid virtual key:", field, key)
	return
}

// Set JSON value generally
func generalSetVirtualKeyJSON(data connectionData, field, key string, valueJSON string) (err error) {
	switch field {
	case NM_SETTING_802_1X_SETTING_NAME:
		switch key {
		case NM_SETTING_VK_802_1X_ENABLE:
			err = logicSetSettingVk8021xEnableJSON(data, valueJSON)
			return
		case NM_SETTING_VK_802_1X_EAP:
			err = logicSetSettingVk8021xEapJSON(data, valueJSON)
			return
		case NM_SETTING_VK_802_1X_PAC_FILE:
			err = logicSetSettingVk8021xPacFileJSON(data, valueJSON)
			return
		case NM_SETTING_VK_802_1X_CA_CERT:
			err = logicSetSettingVk8021xCaCertJSON(data, valueJSON)
			return
		case NM_SETTING_VK_802_1X_CLIENT_CERT:
			err = logicSetSettingVk8021xClientCertJSON(data, valueJSON)
			return
		case NM_SETTING_VK_802_1X_PRIVATE_KEY:
			err = logicSetSettingVk8021xPrivateKeyJSON(data, valueJSON)
			return
		}
	case NM_SETTING_CONNECTION_SETTING_NAME:
		switch key {
		case NM_SETTING_VK_CONNECTION_NO_PERMISSION:
			err = logicSetSettingVkConnectionNoPermissionJSON(data, valueJSON)
			return
		}
	case NM_SETTING_IP4_CONFIG_SETTING_NAME:
		switch key {
		case NM_SETTING_VK_IP4_CONFIG_ADDRESSES_ADDRESS:
			err = logicSetSettingVkIp4ConfigAddressesAddressJSON(data, valueJSON)
			return
		case NM_SETTING_VK_IP4_CONFIG_ADDRESSES_MASK:
			err = logicSetSettingVkIp4ConfigAddressesMaskJSON(data, valueJSON)
			return
		case NM_SETTING_VK_IP4_CONFIG_ADDRESSES_GATEWAY:
			err = logicSetSettingVkIp4ConfigAddressesGatewayJSON(data, valueJSON)
			return
		case NM_SETTING_VK_IP4_CONFIG_DNS:
			err = logicSetSettingVkIp4ConfigDnsJSON(data, valueJSON)
			return
		case NM_SETTING_VK_IP4_CONFIG_ROUTES_ADDRESS:
			err = logicSetSettingVkIp4ConfigRoutesAddressJSON(data, valueJSON)
			return
		case NM_SETTING_VK_IP4_CONFIG_ROUTES_MASK:
			err = logicSetSettingVkIp4ConfigRoutesMaskJSON(data, valueJSON)
			return
		case NM_SETTING_VK_IP4_CONFIG_ROUTES_NEXTHOP:
			err = logicSetSettingVkIp4ConfigRoutesNexthopJSON(data, valueJSON)
			return
		case NM_SETTING_VK_IP4_CONFIG_ROUTES_METRIC:
			err = logicSetSettingVkIp4ConfigRoutesMetricJSON(data, valueJSON)
			return
		}
	case NM_SETTING_IP6_CONFIG_SETTING_NAME:
		switch key {
		case NM_SETTING_VK_IP6_CONFIG_ADDRESSES_ADDRESS:
			err = logicSetSettingVkIp6ConfigAddressesAddressJSON(data, valueJSON)
			return
		case NM_SETTING_VK_IP6_CONFIG_ADDRESSES_PREFIX:
			err = logicSetSettingVkIp6ConfigAddressesPrefixJSON(data, valueJSON)
			return
		case NM_SETTING_VK_IP6_CONFIG_ADDRESSES_GATEWAY:
			err = logicSetSettingVkIp6ConfigAddressesGatewayJSON(data, valueJSON)
			return
		case NM_SETTING_VK_IP6_CONFIG_DNS:
			err = logicSetSettingVkIp6ConfigDnsJSON(data, valueJSON)
			return
		case NM_SETTING_VK_IP6_CONFIG_ROUTES_ADDRESS:
			err = logicSetSettingVkIp6ConfigRoutesAddressJSON(data, valueJSON)
			return
		case NM_SETTING_VK_IP6_CONFIG_ROUTES_PREFIX:
			err = logicSetSettingVkIp6ConfigRoutesPrefixJSON(data, valueJSON)
			return
		case NM_SETTING_VK_IP6_CONFIG_ROUTES_NEXTHOP:
			err = logicSetSettingVkIp6ConfigRoutesNexthopJSON(data, valueJSON)
			return
		case NM_SETTING_VK_IP6_CONFIG_ROUTES_METRIC:
			err = logicSetSettingVkIp6ConfigRoutesMetricJSON(data, valueJSON)
			return
		}
	case NM_SETTING_PPP_SETTING_NAME:
		switch key {
		case NM_SETTING_VK_PPP_ENABLE_LCP_ECHO:
			err = logicSetSettingVkPppEnableLcpEchoJSON(data, valueJSON)
			return
		}
	case NM_SETTING_VF_VPN_L2TP_PPP_SETTING_NAME:
		switch key {
		case NM_SETTING_VK_VPN_L2TP_ENABLE_LCP_ECHO:
			err = logicSetSettingVkVpnL2tpEnableLcpEchoJSON(data, valueJSON)
			return
		}
	case NM_SETTING_VF_VPN_OPENVPN_ADVANCED_SETTING_NAME:
		switch key {
		case NM_SETTING_VK_VPN_OPENVPN_KEY_ENABLE_PORT:
			err = logicSetSettingVkVpnOpenvpnKeyEnablePortJSON(data, valueJSON)
			return
		case NM_SETTING_VK_VPN_OPENVPN_KEY_ENABLE_RENEG_SECONDS:
			err = logicSetSettingVkVpnOpenvpnKeyEnableRenegSecondsJSON(data, valueJSON)
			return
		case NM_SETTING_VK_VPN_OPENVPN_KEY_ENABLE_TUNNEL_MTU:
			err = logicSetSettingVkVpnOpenvpnKeyEnableTunnelMtuJSON(data, valueJSON)
			return
		case NM_SETTING_VK_VPN_OPENVPN_KEY_ENABLE_FRAGMENT_SIZE:
			err = logicSetSettingVkVpnOpenvpnKeyEnableFragmentSizeJSON(data, valueJSON)
			return
		}
	case NM_SETTING_VF_VPN_OPENVPN_SETTING_NAME:
		switch key {
		case NM_SETTING_VK_VPN_OPENVPN_KEY_ENABLE_STATIC_KEY_DIRECTION:
			err = logicSetSettingVkVpnOpenvpnKeyEnableStaticKeyDirectionJSON(data, valueJSON)
			return
		}
	case NM_SETTING_VF_VPN_OPENVPN_TLSAUTH_SETTING_NAME:
		switch key {
		case NM_SETTING_VK_VPN_OPENVPN_KEY_ENABLE_TA_DIR:
			err = logicSetSettingVkVpnOpenvpnKeyEnableTaDirJSON(data, valueJSON)
			return
		}
	case NM_SETTING_VF_VPN_PPTP_PPP_SETTING_NAME:
		switch key {
		case NM_SETTING_VK_VPN_PPTP_ENABLE_LCP_ECHO:
			err = logicSetSettingVkVpnPptpEnableLcpEchoJSON(data, valueJSON)
			return
		}
	case NM_SETTING_VF_VPN_VPNC_SETTING_NAME:
		switch key {
		case NM_SETTING_VK_VPN_VPNC_KEY_HYBRID_AUTHMODE:
			err = logicSetSettingVkVpnVpncKeyHybridAuthmodeJSON(data, valueJSON)
			return
		}
	case NM_SETTING_VF_VPN_VPNC_ADVANCED_SETTING_NAME:
		switch key {
		case NM_SETTING_VK_VPN_VPNC_KEY_ENCRYPTION_METHOD:
			err = logicSetSettingVkVpnVpncKeyEncryptionMethodJSON(data, valueJSON)
			return
		case NM_SETTING_VK_VPN_VPNC_KEY_DISABLE_DPD:
			err = logicSetSettingVkVpnVpncKeyDisableDpdJSON(data, valueJSON)
			return
		}
	case NM_SETTING_WIRED_SETTING_NAME:
		switch key {
		case NM_SETTING_VK_WIRED_MTU:
			err = logicSetSettingVkWiredMtuJSON(data, valueJSON)
			return
		}
	case NM_SETTING_WIRELESS_SECURITY_SETTING_NAME:
		switch key {
		case NM_SETTING_VK_WIRELESS_SECURITY_KEY_MGMT:
			err = logicSetSettingVkWirelessSecurityKeyMgmtJSON(data, valueJSON)
			return
		}
	}
	logger.Error("invalid virtual key:", field, key)
	return
}

// JSON getter
func getSettingVk8021xEnableJSON(data connectionData) (valueJSON string) {
	valueJSON, _ = marshalJSON(getSettingVk8021xEnable(data))
	return
}
func getSettingVk8021xEapJSON(data connectionData) (valueJSON string) {
	valueJSON, _ = marshalJSON(getSettingVk8021xEap(data))
	return
}
func getSettingVk8021xPacFileJSON(data connectionData) (valueJSON string) {
	valueJSON, _ = marshalJSON(getSettingVk8021xPacFile(data))
	return
}
func getSettingVk8021xCaCertJSON(data connectionData) (valueJSON string) {
	valueJSON, _ = marshalJSON(getSettingVk8021xCaCert(data))
	return
}
func getSettingVk8021xClientCertJSON(data connectionData) (valueJSON string) {
	valueJSON, _ = marshalJSON(getSettingVk8021xClientCert(data))
	return
}
func getSettingVk8021xPrivateKeyJSON(data connectionData) (valueJSON string) {
	valueJSON, _ = marshalJSON(getSettingVk8021xPrivateKey(data))
	return
}
func getSettingVkConnectionNoPermissionJSON(data connectionData) (valueJSON string) {
	valueJSON, _ = marshalJSON(getSettingVkConnectionNoPermission(data))
	return
}
func getSettingVkIp4ConfigAddressesAddressJSON(data connectionData) (valueJSON string) {
	valueJSON, _ = marshalJSON(getSettingVkIp4ConfigAddressesAddress(data))
	return
}
func getSettingVkIp4ConfigAddressesMaskJSON(data connectionData) (valueJSON string) {
	valueJSON, _ = marshalJSON(getSettingVkIp4ConfigAddressesMask(data))
	return
}
func getSettingVkIp4ConfigAddressesGatewayJSON(data connectionData) (valueJSON string) {
	valueJSON, _ = marshalJSON(getSettingVkIp4ConfigAddressesGateway(data))
	return
}
func getSettingVkIp4ConfigDnsJSON(data connectionData) (valueJSON string) {
	valueJSON, _ = marshalJSON(getSettingVkIp4ConfigDns(data))
	return
}
func getSettingVkIp4ConfigRoutesAddressJSON(data connectionData) (valueJSON string) {
	valueJSON, _ = marshalJSON(getSettingVkIp4ConfigRoutesAddress(data))
	return
}
func getSettingVkIp4ConfigRoutesMaskJSON(data connectionData) (valueJSON string) {
	valueJSON, _ = marshalJSON(getSettingVkIp4ConfigRoutesMask(data))
	return
}
func getSettingVkIp4ConfigRoutesNexthopJSON(data connectionData) (valueJSON string) {
	valueJSON, _ = marshalJSON(getSettingVkIp4ConfigRoutesNexthop(data))
	return
}
func getSettingVkIp4ConfigRoutesMetricJSON(data connectionData) (valueJSON string) {
	valueJSON, _ = marshalJSON(getSettingVkIp4ConfigRoutesMetric(data))
	return
}
func getSettingVkIp6ConfigAddressesAddressJSON(data connectionData) (valueJSON string) {
	valueJSON, _ = marshalJSON(getSettingVkIp6ConfigAddressesAddress(data))
	return
}
func getSettingVkIp6ConfigAddressesPrefixJSON(data connectionData) (valueJSON string) {
	valueJSON, _ = marshalJSON(getSettingVkIp6ConfigAddressesPrefix(data))
	return
}
func getSettingVkIp6ConfigAddressesGatewayJSON(data connectionData) (valueJSON string) {
	valueJSON, _ = marshalJSON(getSettingVkIp6ConfigAddressesGateway(data))
	return
}
func getSettingVkIp6ConfigDnsJSON(data connectionData) (valueJSON string) {
	valueJSON, _ = marshalJSON(getSettingVkIp6ConfigDns(data))
	return
}
func getSettingVkIp6ConfigRoutesAddressJSON(data connectionData) (valueJSON string) {
	valueJSON, _ = marshalJSON(getSettingVkIp6ConfigRoutesAddress(data))
	return
}
func getSettingVkIp6ConfigRoutesPrefixJSON(data connectionData) (valueJSON string) {
	valueJSON, _ = marshalJSON(getSettingVkIp6ConfigRoutesPrefix(data))
	return
}
func getSettingVkIp6ConfigRoutesNexthopJSON(data connectionData) (valueJSON string) {
	valueJSON, _ = marshalJSON(getSettingVkIp6ConfigRoutesNexthop(data))
	return
}
func getSettingVkIp6ConfigRoutesMetricJSON(data connectionData) (valueJSON string) {
	valueJSON, _ = marshalJSON(getSettingVkIp6ConfigRoutesMetric(data))
	return
}
func getSettingVkPppEnableLcpEchoJSON(data connectionData) (valueJSON string) {
	valueJSON, _ = marshalJSON(getSettingVkPppEnableLcpEcho(data))
	return
}
func getSettingVkVpnL2tpEnableLcpEchoJSON(data connectionData) (valueJSON string) {
	valueJSON, _ = marshalJSON(getSettingVkVpnL2tpEnableLcpEcho(data))
	return
}
func getSettingVkVpnOpenvpnKeyEnablePortJSON(data connectionData) (valueJSON string) {
	valueJSON, _ = marshalJSON(getSettingVkVpnOpenvpnKeyEnablePort(data))
	return
}
func getSettingVkVpnOpenvpnKeyEnableRenegSecondsJSON(data connectionData) (valueJSON string) {
	valueJSON, _ = marshalJSON(getSettingVkVpnOpenvpnKeyEnableRenegSeconds(data))
	return
}
func getSettingVkVpnOpenvpnKeyEnableTunnelMtuJSON(data connectionData) (valueJSON string) {
	valueJSON, _ = marshalJSON(getSettingVkVpnOpenvpnKeyEnableTunnelMtu(data))
	return
}
func getSettingVkVpnOpenvpnKeyEnableFragmentSizeJSON(data connectionData) (valueJSON string) {
	valueJSON, _ = marshalJSON(getSettingVkVpnOpenvpnKeyEnableFragmentSize(data))
	return
}
func getSettingVkVpnOpenvpnKeyEnableStaticKeyDirectionJSON(data connectionData) (valueJSON string) {
	valueJSON, _ = marshalJSON(getSettingVkVpnOpenvpnKeyEnableStaticKeyDirection(data))
	return
}
func getSettingVkVpnOpenvpnKeyEnableTaDirJSON(data connectionData) (valueJSON string) {
	valueJSON, _ = marshalJSON(getSettingVkVpnOpenvpnKeyEnableTaDir(data))
	return
}
func getSettingVkVpnPptpEnableLcpEchoJSON(data connectionData) (valueJSON string) {
	valueJSON, _ = marshalJSON(getSettingVkVpnPptpEnableLcpEcho(data))
	return
}
func getSettingVkVpnVpncKeyHybridAuthmodeJSON(data connectionData) (valueJSON string) {
	valueJSON, _ = marshalJSON(getSettingVkVpnVpncKeyHybridAuthmode(data))
	return
}
func getSettingVkVpnVpncKeyEncryptionMethodJSON(data connectionData) (valueJSON string) {
	valueJSON, _ = marshalJSON(getSettingVkVpnVpncKeyEncryptionMethod(data))
	return
}
func getSettingVkVpnVpncKeyDisableDpdJSON(data connectionData) (valueJSON string) {
	valueJSON, _ = marshalJSON(getSettingVkVpnVpncKeyDisableDpd(data))
	return
}
func getSettingVkWiredMtuJSON(data connectionData) (valueJSON string) {
	valueJSON, _ = marshalJSON(getSettingVkWiredMtu(data))
	return
}
func getSettingVkWirelessSecurityKeyMgmtJSON(data connectionData) (valueJSON string) {
	valueJSON, _ = marshalJSON(getSettingVkWirelessSecurityKeyMgmt(data))
	return
}

// Logic JSON setter
func logicSetSettingVk8021xEnableJSON(data connectionData, valueJSON string) (err error) {
	value, _ := jsonToKeyValueBoolean(valueJSON)
	return logicSetSettingVk8021xEnable(data, value)
}
func logicSetSettingVk8021xEapJSON(data connectionData, valueJSON string) (err error) {
	value, _ := jsonToKeyValueString(valueJSON)
	return logicSetSettingVk8021xEap(data, value)
}
func logicSetSettingVk8021xPacFileJSON(data connectionData, valueJSON string) (err error) {
	value, _ := jsonToKeyValueString(valueJSON)
	return logicSetSettingVk8021xPacFile(data, value)
}
func logicSetSettingVk8021xCaCertJSON(data connectionData, valueJSON string) (err error) {
	value, _ := jsonToKeyValueString(valueJSON)
	return logicSetSettingVk8021xCaCert(data, value)
}
func logicSetSettingVk8021xClientCertJSON(data connectionData, valueJSON string) (err error) {
	value, _ := jsonToKeyValueString(valueJSON)
	return logicSetSettingVk8021xClientCert(data, value)
}
func logicSetSettingVk8021xPrivateKeyJSON(data connectionData, valueJSON string) (err error) {
	value, _ := jsonToKeyValueString(valueJSON)
	return logicSetSettingVk8021xPrivateKey(data, value)
}
func logicSetSettingVkConnectionNoPermissionJSON(data connectionData, valueJSON string) (err error) {
	value, _ := jsonToKeyValueBoolean(valueJSON)
	return logicSetSettingVkConnectionNoPermission(data, value)
}
func logicSetSettingVkIp4ConfigAddressesAddressJSON(data connectionData, valueJSON string) (err error) {
	value, _ := jsonToKeyValueString(valueJSON)
	return logicSetSettingVkIp4ConfigAddressesAddress(data, value)
}
func logicSetSettingVkIp4ConfigAddressesMaskJSON(data connectionData, valueJSON string) (err error) {
	value, _ := jsonToKeyValueString(valueJSON)
	return logicSetSettingVkIp4ConfigAddressesMask(data, value)
}
func logicSetSettingVkIp4ConfigAddressesGatewayJSON(data connectionData, valueJSON string) (err error) {
	value, _ := jsonToKeyValueString(valueJSON)
	return logicSetSettingVkIp4ConfigAddressesGateway(data, value)
}
func logicSetSettingVkIp4ConfigDnsJSON(data connectionData, valueJSON string) (err error) {
	value, _ := jsonToKeyValueString(valueJSON)
	return logicSetSettingVkIp4ConfigDns(data, value)
}
func logicSetSettingVkIp4ConfigRoutesAddressJSON(data connectionData, valueJSON string) (err error) {
	value, _ := jsonToKeyValueString(valueJSON)
	return logicSetSettingVkIp4ConfigRoutesAddress(data, value)
}
func logicSetSettingVkIp4ConfigRoutesMaskJSON(data connectionData, valueJSON string) (err error) {
	value, _ := jsonToKeyValueString(valueJSON)
	return logicSetSettingVkIp4ConfigRoutesMask(data, value)
}
func logicSetSettingVkIp4ConfigRoutesNexthopJSON(data connectionData, valueJSON string) (err error) {
	value, _ := jsonToKeyValueString(valueJSON)
	return logicSetSettingVkIp4ConfigRoutesNexthop(data, value)
}
func logicSetSettingVkIp4ConfigRoutesMetricJSON(data connectionData, valueJSON string) (err error) {
	value, _ := jsonToKeyValueString(valueJSON)
	return logicSetSettingVkIp4ConfigRoutesMetric(data, value)
}
func logicSetSettingVkIp6ConfigAddressesAddressJSON(data connectionData, valueJSON string) (err error) {
	value, _ := jsonToKeyValueString(valueJSON)
	return logicSetSettingVkIp6ConfigAddressesAddress(data, value)
}
func logicSetSettingVkIp6ConfigAddressesPrefixJSON(data connectionData, valueJSON string) (err error) {
	value, _ := jsonToKeyValueUint32(valueJSON)
	return logicSetSettingVkIp6ConfigAddressesPrefix(data, value)
}
func logicSetSettingVkIp6ConfigAddressesGatewayJSON(data connectionData, valueJSON string) (err error) {
	value, _ := jsonToKeyValueString(valueJSON)
	return logicSetSettingVkIp6ConfigAddressesGateway(data, value)
}
func logicSetSettingVkIp6ConfigDnsJSON(data connectionData, valueJSON string) (err error) {
	value, _ := jsonToKeyValueString(valueJSON)
	return logicSetSettingVkIp6ConfigDns(data, value)
}
func logicSetSettingVkIp6ConfigRoutesAddressJSON(data connectionData, valueJSON string) (err error) {
	value, _ := jsonToKeyValueString(valueJSON)
	return logicSetSettingVkIp6ConfigRoutesAddress(data, value)
}
func logicSetSettingVkIp6ConfigRoutesPrefixJSON(data connectionData, valueJSON string) (err error) {
	value, _ := jsonToKeyValueUint32(valueJSON)
	return logicSetSettingVkIp6ConfigRoutesPrefix(data, value)
}
func logicSetSettingVkIp6ConfigRoutesNexthopJSON(data connectionData, valueJSON string) (err error) {
	value, _ := jsonToKeyValueString(valueJSON)
	return logicSetSettingVkIp6ConfigRoutesNexthop(data, value)
}
func logicSetSettingVkIp6ConfigRoutesMetricJSON(data connectionData, valueJSON string) (err error) {
	value, _ := jsonToKeyValueUint32(valueJSON)
	return logicSetSettingVkIp6ConfigRoutesMetric(data, value)
}
func logicSetSettingVkPppEnableLcpEchoJSON(data connectionData, valueJSON string) (err error) {
	value, _ := jsonToKeyValueBoolean(valueJSON)
	return logicSetSettingVkPppEnableLcpEcho(data, value)
}
func logicSetSettingVkVpnL2tpEnableLcpEchoJSON(data connectionData, valueJSON string) (err error) {
	value, _ := jsonToKeyValueBoolean(valueJSON)
	return logicSetSettingVkVpnL2tpEnableLcpEcho(data, value)
}
func logicSetSettingVkVpnOpenvpnKeyEnablePortJSON(data connectionData, valueJSON string) (err error) {
	value, _ := jsonToKeyValueBoolean(valueJSON)
	return logicSetSettingVkVpnOpenvpnKeyEnablePort(data, value)
}
func logicSetSettingVkVpnOpenvpnKeyEnableRenegSecondsJSON(data connectionData, valueJSON string) (err error) {
	value, _ := jsonToKeyValueBoolean(valueJSON)
	return logicSetSettingVkVpnOpenvpnKeyEnableRenegSeconds(data, value)
}
func logicSetSettingVkVpnOpenvpnKeyEnableTunnelMtuJSON(data connectionData, valueJSON string) (err error) {
	value, _ := jsonToKeyValueBoolean(valueJSON)
	return logicSetSettingVkVpnOpenvpnKeyEnableTunnelMtu(data, value)
}
func logicSetSettingVkVpnOpenvpnKeyEnableFragmentSizeJSON(data connectionData, valueJSON string) (err error) {
	value, _ := jsonToKeyValueBoolean(valueJSON)
	return logicSetSettingVkVpnOpenvpnKeyEnableFragmentSize(data, value)
}
func logicSetSettingVkVpnOpenvpnKeyEnableStaticKeyDirectionJSON(data connectionData, valueJSON string) (err error) {
	value, _ := jsonToKeyValueBoolean(valueJSON)
	return logicSetSettingVkVpnOpenvpnKeyEnableStaticKeyDirection(data, value)
}
func logicSetSettingVkVpnOpenvpnKeyEnableTaDirJSON(data connectionData, valueJSON string) (err error) {
	value, _ := jsonToKeyValueBoolean(valueJSON)
	return logicSetSettingVkVpnOpenvpnKeyEnableTaDir(data, value)
}
func logicSetSettingVkVpnPptpEnableLcpEchoJSON(data connectionData, valueJSON string) (err error) {
	value, _ := jsonToKeyValueBoolean(valueJSON)
	return logicSetSettingVkVpnPptpEnableLcpEcho(data, value)
}
func logicSetSettingVkVpnVpncKeyHybridAuthmodeJSON(data connectionData, valueJSON string) (err error) {
	value, _ := jsonToKeyValueBoolean(valueJSON)
	return logicSetSettingVkVpnVpncKeyHybridAuthmode(data, value)
}
func logicSetSettingVkVpnVpncKeyEncryptionMethodJSON(data connectionData, valueJSON string) (err error) {
	value, _ := jsonToKeyValueString(valueJSON)
	return logicSetSettingVkVpnVpncKeyEncryptionMethod(data, value)
}
func logicSetSettingVkVpnVpncKeyDisableDpdJSON(data connectionData, valueJSON string) (err error) {
	value, _ := jsonToKeyValueBoolean(valueJSON)
	return logicSetSettingVkVpnVpncKeyDisableDpd(data, value)
}
func logicSetSettingVkWiredMtuJSON(data connectionData, valueJSON string) (err error) {
	value, _ := jsonToKeyValueBoolean(valueJSON)
	return logicSetSettingVkWiredMtu(data, value)
}
func logicSetSettingVkWirelessSecurityKeyMgmtJSON(data connectionData, valueJSON string) (err error) {
	value, _ := jsonToKeyValueString(valueJSON)
	return logicSetSettingVkWirelessSecurityKeyMgmt(data, value)
}

// Getter for enable key wrapper
func getSettingVkVpnOpenvpnKeyEnablePort(data connectionData) (value bool) {
	if isSettingVpnOpenvpnKeyPortExists(data) {
		return true
	}
	return false
}
func getSettingVkVpnOpenvpnKeyEnableRenegSeconds(data connectionData) (value bool) {
	if isSettingVpnOpenvpnKeyRenegSecondsExists(data) {
		return true
	}
	return false
}
func getSettingVkVpnOpenvpnKeyEnableTunnelMtu(data connectionData) (value bool) {
	if isSettingVpnOpenvpnKeyTunnelMtuExists(data) {
		return true
	}
	return false
}
func getSettingVkVpnOpenvpnKeyEnableFragmentSize(data connectionData) (value bool) {
	if isSettingVpnOpenvpnKeyFragmentSizeExists(data) {
		return true
	}
	return false
}
func getSettingVkVpnOpenvpnKeyEnableStaticKeyDirection(data connectionData) (value bool) {
	if isSettingVpnOpenvpnKeyStaticKeyDirectionExists(data) {
		return true
	}
	return false
}
func getSettingVkVpnOpenvpnKeyEnableTaDir(data connectionData) (value bool) {
	if isSettingVpnOpenvpnKeyTaDirExists(data) {
		return true
	}
	return false
}
func getSettingVkWiredMtu(data connectionData) (value bool) {
	if isSettingWiredMtuExists(data) {
		return true
	}
	return false
}

// Setter for enable key wrapper
func logicSetSettingVkVpnOpenvpnKeyEnablePort(data connectionData, value bool) (err error) {
	if value {
		setSettingVpnOpenvpnKeyPort(data, 1194)
	} else {
		removeSettingVpnOpenvpnKeyPort(data)
	}
	return
}
func logicSetSettingVkVpnOpenvpnKeyEnableRenegSeconds(data connectionData, value bool) (err error) {
	if value {
		setSettingVpnOpenvpnKeyRenegSeconds(data, 0)
	} else {
		removeSettingVpnOpenvpnKeyRenegSeconds(data)
	}
	return
}
func logicSetSettingVkVpnOpenvpnKeyEnableTunnelMtu(data connectionData, value bool) (err error) {
	if value {
		setSettingVpnOpenvpnKeyTunnelMtu(data, 1500)
	} else {
		removeSettingVpnOpenvpnKeyTunnelMtu(data)
	}
	return
}
func logicSetSettingVkVpnOpenvpnKeyEnableFragmentSize(data connectionData, value bool) (err error) {
	if value {
		setSettingVpnOpenvpnKeyFragmentSize(data, 1300)
	} else {
		removeSettingVpnOpenvpnKeyFragmentSize(data)
	}
	return
}
func logicSetSettingVkVpnOpenvpnKeyEnableStaticKeyDirection(data connectionData, value bool) (err error) {
	if value {
		setSettingVpnOpenvpnKeyStaticKeyDirection(data, 0)
	} else {
		removeSettingVpnOpenvpnKeyStaticKeyDirection(data)
	}
	return
}
func logicSetSettingVkVpnOpenvpnKeyEnableTaDir(data connectionData, value bool) (err error) {
	if value {
		setSettingVpnOpenvpnKeyTaDir(data, 0)
	} else {
		removeSettingVpnOpenvpnKeyTaDir(data)
	}
	return
}
func logicSetSettingVkWiredMtu(data connectionData, value bool) (err error) {
	if value {
		setSettingWiredMtu(data, 0)
	} else {
		removeSettingWiredMtu(data)
	}
	return
}
