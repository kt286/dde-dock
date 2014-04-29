// This file is automatically generated, please don't edit manully.
package main

import (
	"fmt"
)

// Get key type
func getSettingWirelessKeyType(key string) (t ktype) {
	switch key {
	default:
		t = ktypeUnknown
	case NM_SETTING_WIRELESS_SSID:
		t = ktypeWrapperString
	case NM_SETTING_WIRELESS_MODE:
		t = ktypeString
	case NM_SETTING_WIRELESS_BAND:
		t = ktypeString
	case NM_SETTING_WIRELESS_CHANNEL:
		t = ktypeUint32
	case NM_SETTING_WIRELESS_BSSID:
		t = ktypeWrapperString
	case NM_SETTING_WIRELESS_RATE:
		t = ktypeUint32
	case NM_SETTING_WIRELESS_TX_POWER:
		t = ktypeUint32
	case NM_SETTING_WIRELESS_MAC_ADDRESS:
		t = ktypeWrapperMacAddress
	case NM_SETTING_WIRELESS_CLONED_MAC_ADDRESS:
		t = ktypeWrapperMacAddress
	case NM_SETTING_WIRELESS_MAC_ADDRESS_BLACKLIST:
		t = ktypeArrayString
	case NM_SETTING_WIRELESS_MTU:
		t = ktypeUint32
	case NM_SETTING_WIRELESS_SEEN_BSSIDS:
		t = ktypeArrayString
	case NM_SETTING_WIRELESS_SEC:
		t = ktypeString
	case NM_SETTING_WIRELESS_HIDDEN:
		t = ktypeBoolean
	}
	return
}

// Check is key in current setting field
func isKeyInSettingWireless(key string) bool {
	switch key {
	case NM_SETTING_WIRELESS_SSID:
		return true
	case NM_SETTING_WIRELESS_MODE:
		return true
	case NM_SETTING_WIRELESS_BAND:
		return true
	case NM_SETTING_WIRELESS_CHANNEL:
		return true
	case NM_SETTING_WIRELESS_BSSID:
		return true
	case NM_SETTING_WIRELESS_RATE:
		return true
	case NM_SETTING_WIRELESS_TX_POWER:
		return true
	case NM_SETTING_WIRELESS_MAC_ADDRESS:
		return true
	case NM_SETTING_WIRELESS_CLONED_MAC_ADDRESS:
		return true
	case NM_SETTING_WIRELESS_MAC_ADDRESS_BLACKLIST:
		return true
	case NM_SETTING_WIRELESS_MTU:
		return true
	case NM_SETTING_WIRELESS_SEEN_BSSIDS:
		return true
	case NM_SETTING_WIRELESS_SEC:
		return true
	case NM_SETTING_WIRELESS_HIDDEN:
		return true
	}
	return false
}

// Get key's default value
func getSettingWirelessKeyDefaultValueJSON(key string) (valueJSON string) {
	switch key {
	default:
		logger.Error("invalid key:", key)
	case NM_SETTING_WIRELESS_SSID:
		valueJSON = `""`
	case NM_SETTING_WIRELESS_MODE:
		valueJSON = `""`
	case NM_SETTING_WIRELESS_BAND:
		valueJSON = `""`
	case NM_SETTING_WIRELESS_CHANNEL:
		valueJSON = `0`
	case NM_SETTING_WIRELESS_BSSID:
		valueJSON = `""`
	case NM_SETTING_WIRELESS_RATE:
		valueJSON = `0`
	case NM_SETTING_WIRELESS_TX_POWER:
		valueJSON = `0`
	case NM_SETTING_WIRELESS_MAC_ADDRESS:
		valueJSON = `""`
	case NM_SETTING_WIRELESS_CLONED_MAC_ADDRESS:
		valueJSON = `""`
	case NM_SETTING_WIRELESS_MAC_ADDRESS_BLACKLIST:
		valueJSON = `null`
	case NM_SETTING_WIRELESS_MTU:
		valueJSON = `0`
	case NM_SETTING_WIRELESS_SEEN_BSSIDS:
		valueJSON = `null`
	case NM_SETTING_WIRELESS_SEC:
		valueJSON = `""`
	case NM_SETTING_WIRELESS_HIDDEN:
		valueJSON = `false`
	}
	return
}

// Get JSON value generally
func generalGetSettingWirelessKeyJSON(data connectionData, key string) (value string) {
	switch key {
	default:
		logger.Error("generalGetSettingWirelessKeyJSON: invalide key", key)
	case NM_SETTING_WIRELESS_SSID:
		value = getSettingWirelessSsidJSON(data)
	case NM_SETTING_WIRELESS_MODE:
		value = getSettingWirelessModeJSON(data)
	case NM_SETTING_WIRELESS_BAND:
		value = getSettingWirelessBandJSON(data)
	case NM_SETTING_WIRELESS_CHANNEL:
		value = getSettingWirelessChannelJSON(data)
	case NM_SETTING_WIRELESS_BSSID:
		value = getSettingWirelessBssidJSON(data)
	case NM_SETTING_WIRELESS_RATE:
		value = getSettingWirelessRateJSON(data)
	case NM_SETTING_WIRELESS_TX_POWER:
		value = getSettingWirelessTxPowerJSON(data)
	case NM_SETTING_WIRELESS_MAC_ADDRESS:
		value = getSettingWirelessMacAddressJSON(data)
	case NM_SETTING_WIRELESS_CLONED_MAC_ADDRESS:
		value = getSettingWirelessClonedMacAddressJSON(data)
	case NM_SETTING_WIRELESS_MAC_ADDRESS_BLACKLIST:
		value = getSettingWirelessMacAddressBlacklistJSON(data)
	case NM_SETTING_WIRELESS_MTU:
		value = getSettingWirelessMtuJSON(data)
	case NM_SETTING_WIRELESS_SEEN_BSSIDS:
		value = getSettingWirelessSeenBssidsJSON(data)
	case NM_SETTING_WIRELESS_SEC:
		value = getSettingWirelessSecJSON(data)
	case NM_SETTING_WIRELESS_HIDDEN:
		value = getSettingWirelessHiddenJSON(data)
	}
	return
}

// Set JSON value generally
func generalSetSettingWirelessKeyJSON(data connectionData, key, valueJSON string) (ok bool, errMsg string) {
	ok = true
	switch key {
	default:
		logger.Error("generalSetSettingWirelessKeyJSON: invalide key", key)
	case NM_SETTING_WIRELESS_SSID:
		setSettingWirelessSsidJSON(data, valueJSON)
	case NM_SETTING_WIRELESS_MODE:
		setSettingWirelessModeJSON(data, valueJSON)
	case NM_SETTING_WIRELESS_BAND:
		setSettingWirelessBandJSON(data, valueJSON)
	case NM_SETTING_WIRELESS_CHANNEL:
		setSettingWirelessChannelJSON(data, valueJSON)
	case NM_SETTING_WIRELESS_BSSID:
		setSettingWirelessBssidJSON(data, valueJSON)
	case NM_SETTING_WIRELESS_RATE:
		setSettingWirelessRateJSON(data, valueJSON)
	case NM_SETTING_WIRELESS_TX_POWER:
		setSettingWirelessTxPowerJSON(data, valueJSON)
	case NM_SETTING_WIRELESS_MAC_ADDRESS:
		setSettingWirelessMacAddressJSON(data, valueJSON)
	case NM_SETTING_WIRELESS_CLONED_MAC_ADDRESS:
		setSettingWirelessClonedMacAddressJSON(data, valueJSON)
	case NM_SETTING_WIRELESS_MAC_ADDRESS_BLACKLIST:
		setSettingWirelessMacAddressBlacklistJSON(data, valueJSON)
	case NM_SETTING_WIRELESS_MTU:
		setSettingWirelessMtuJSON(data, valueJSON)
	case NM_SETTING_WIRELESS_SEEN_BSSIDS:
		setSettingWirelessSeenBssidsJSON(data, valueJSON)
	case NM_SETTING_WIRELESS_SEC:
		setSettingWirelessSecJSON(data, valueJSON)
	case NM_SETTING_WIRELESS_HIDDEN:
		setSettingWirelessHiddenJSON(data, valueJSON)
	}
	return
}

// Check if key exists
func isSettingWirelessSsidExists(data connectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_SSID)
}
func isSettingWirelessModeExists(data connectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_MODE)
}
func isSettingWirelessBandExists(data connectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_BAND)
}
func isSettingWirelessChannelExists(data connectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_CHANNEL)
}
func isSettingWirelessBssidExists(data connectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_BSSID)
}
func isSettingWirelessRateExists(data connectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_RATE)
}
func isSettingWirelessTxPowerExists(data connectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_TX_POWER)
}
func isSettingWirelessMacAddressExists(data connectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_MAC_ADDRESS)
}
func isSettingWirelessClonedMacAddressExists(data connectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_CLONED_MAC_ADDRESS)
}
func isSettingWirelessMacAddressBlacklistExists(data connectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_MAC_ADDRESS_BLACKLIST)
}
func isSettingWirelessMtuExists(data connectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_MTU)
}
func isSettingWirelessSeenBssidsExists(data connectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_SEEN_BSSIDS)
}
func isSettingWirelessSecExists(data connectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_SEC)
}
func isSettingWirelessHiddenExists(data connectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_HIDDEN)
}

// Ensure field and key exists and not empty
func ensureFieldSettingWirelessExists(data connectionData, errs fieldErrors, relatedKey string) {
	if !isSettingFieldExists(data, NM_SETTING_WIRELESS_SETTING_NAME) {
		rememberError(errs, relatedKey, NM_SETTING_WIRELESS_SETTING_NAME, fmt.Sprintf(NM_KEY_ERROR_MISSING_SECTION, NM_SETTING_WIRELESS_SETTING_NAME))
	}
	fieldData, _ := data[NM_SETTING_WIRELESS_SETTING_NAME]
	if len(fieldData) == 0 {
		rememberError(errs, relatedKey, NM_SETTING_WIRELESS_SETTING_NAME, fmt.Sprintf(NM_KEY_ERROR_EMPTY_SECTION, NM_SETTING_WIRELESS_SETTING_NAME))
	}
}
func ensureSettingWirelessSsidNoEmpty(data connectionData, errs fieldErrors) {
	if !isSettingWirelessSsidExists(data) {
		rememberError(errs, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_SSID, NM_KEY_ERROR_MISSING_VALUE)
	}
	value := getSettingWirelessSsid(data)
	if len(value) == 0 {
		rememberError(errs, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_SSID, NM_KEY_ERROR_EMPTY_VALUE)
	}
}
func ensureSettingWirelessModeNoEmpty(data connectionData, errs fieldErrors) {
	if !isSettingWirelessModeExists(data) {
		rememberError(errs, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_MODE, NM_KEY_ERROR_MISSING_VALUE)
	}
	value := getSettingWirelessMode(data)
	if len(value) == 0 {
		rememberError(errs, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_MODE, NM_KEY_ERROR_EMPTY_VALUE)
	}
}
func ensureSettingWirelessBandNoEmpty(data connectionData, errs fieldErrors) {
	if !isSettingWirelessBandExists(data) {
		rememberError(errs, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_BAND, NM_KEY_ERROR_MISSING_VALUE)
	}
	value := getSettingWirelessBand(data)
	if len(value) == 0 {
		rememberError(errs, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_BAND, NM_KEY_ERROR_EMPTY_VALUE)
	}
}
func ensureSettingWirelessChannelNoEmpty(data connectionData, errs fieldErrors) {
	if !isSettingWirelessChannelExists(data) {
		rememberError(errs, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_CHANNEL, NM_KEY_ERROR_MISSING_VALUE)
	}
}
func ensureSettingWirelessBssidNoEmpty(data connectionData, errs fieldErrors) {
	if !isSettingWirelessBssidExists(data) {
		rememberError(errs, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_BSSID, NM_KEY_ERROR_MISSING_VALUE)
	}
	value := getSettingWirelessBssid(data)
	if len(value) == 0 {
		rememberError(errs, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_BSSID, NM_KEY_ERROR_EMPTY_VALUE)
	}
}
func ensureSettingWirelessRateNoEmpty(data connectionData, errs fieldErrors) {
	if !isSettingWirelessRateExists(data) {
		rememberError(errs, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_RATE, NM_KEY_ERROR_MISSING_VALUE)
	}
}
func ensureSettingWirelessTxPowerNoEmpty(data connectionData, errs fieldErrors) {
	if !isSettingWirelessTxPowerExists(data) {
		rememberError(errs, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_TX_POWER, NM_KEY_ERROR_MISSING_VALUE)
	}
}
func ensureSettingWirelessMacAddressNoEmpty(data connectionData, errs fieldErrors) {
	if !isSettingWirelessMacAddressExists(data) {
		rememberError(errs, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_MAC_ADDRESS, NM_KEY_ERROR_MISSING_VALUE)
	}
	value := getSettingWirelessMacAddress(data)
	if len(value) == 0 {
		rememberError(errs, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_MAC_ADDRESS, NM_KEY_ERROR_EMPTY_VALUE)
	}
}
func ensureSettingWirelessClonedMacAddressNoEmpty(data connectionData, errs fieldErrors) {
	if !isSettingWirelessClonedMacAddressExists(data) {
		rememberError(errs, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_CLONED_MAC_ADDRESS, NM_KEY_ERROR_MISSING_VALUE)
	}
	value := getSettingWirelessClonedMacAddress(data)
	if len(value) == 0 {
		rememberError(errs, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_CLONED_MAC_ADDRESS, NM_KEY_ERROR_EMPTY_VALUE)
	}
}
func ensureSettingWirelessMacAddressBlacklistNoEmpty(data connectionData, errs fieldErrors) {
	if !isSettingWirelessMacAddressBlacklistExists(data) {
		rememberError(errs, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_MAC_ADDRESS_BLACKLIST, NM_KEY_ERROR_MISSING_VALUE)
	}
	value := getSettingWirelessMacAddressBlacklist(data)
	if len(value) == 0 {
		rememberError(errs, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_MAC_ADDRESS_BLACKLIST, NM_KEY_ERROR_EMPTY_VALUE)
	}
}
func ensureSettingWirelessMtuNoEmpty(data connectionData, errs fieldErrors) {
	if !isSettingWirelessMtuExists(data) {
		rememberError(errs, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_MTU, NM_KEY_ERROR_MISSING_VALUE)
	}
}
func ensureSettingWirelessSeenBssidsNoEmpty(data connectionData, errs fieldErrors) {
	if !isSettingWirelessSeenBssidsExists(data) {
		rememberError(errs, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_SEEN_BSSIDS, NM_KEY_ERROR_MISSING_VALUE)
	}
	value := getSettingWirelessSeenBssids(data)
	if len(value) == 0 {
		rememberError(errs, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_SEEN_BSSIDS, NM_KEY_ERROR_EMPTY_VALUE)
	}
}
func ensureSettingWirelessSecNoEmpty(data connectionData, errs fieldErrors) {
	if !isSettingWirelessSecExists(data) {
		rememberError(errs, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_SEC, NM_KEY_ERROR_MISSING_VALUE)
	}
	value := getSettingWirelessSec(data)
	if len(value) == 0 {
		rememberError(errs, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_SEC, NM_KEY_ERROR_EMPTY_VALUE)
	}
}
func ensureSettingWirelessHiddenNoEmpty(data connectionData, errs fieldErrors) {
	if !isSettingWirelessHiddenExists(data) {
		rememberError(errs, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_HIDDEN, NM_KEY_ERROR_MISSING_VALUE)
	}
}

// Getter
func getSettingWirelessSsid(data connectionData) (value []byte) {
	value, _ = getSettingKey(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_SSID).([]byte)
	return
}
func getSettingWirelessMode(data connectionData) (value string) {
	value, _ = getSettingKey(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_MODE).(string)
	return
}
func getSettingWirelessBand(data connectionData) (value string) {
	value, _ = getSettingKey(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_BAND).(string)
	return
}
func getSettingWirelessChannel(data connectionData) (value uint32) {
	value, _ = getSettingKey(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_CHANNEL).(uint32)
	return
}
func getSettingWirelessBssid(data connectionData) (value []byte) {
	value, _ = getSettingKey(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_BSSID).([]byte)
	return
}
func getSettingWirelessRate(data connectionData) (value uint32) {
	value, _ = getSettingKey(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_RATE).(uint32)
	return
}
func getSettingWirelessTxPower(data connectionData) (value uint32) {
	value, _ = getSettingKey(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_TX_POWER).(uint32)
	return
}
func getSettingWirelessMacAddress(data connectionData) (value []byte) {
	value, _ = getSettingKey(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_MAC_ADDRESS).([]byte)
	return
}
func getSettingWirelessClonedMacAddress(data connectionData) (value []byte) {
	value, _ = getSettingKey(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_CLONED_MAC_ADDRESS).([]byte)
	return
}
func getSettingWirelessMacAddressBlacklist(data connectionData) (value []string) {
	value, _ = getSettingKey(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_MAC_ADDRESS_BLACKLIST).([]string)
	return
}
func getSettingWirelessMtu(data connectionData) (value uint32) {
	value, _ = getSettingKey(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_MTU).(uint32)
	return
}
func getSettingWirelessSeenBssids(data connectionData) (value []string) {
	value, _ = getSettingKey(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_SEEN_BSSIDS).([]string)
	return
}
func getSettingWirelessSec(data connectionData) (value string) {
	value, _ = getSettingKey(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_SEC).(string)
	return
}
func getSettingWirelessHidden(data connectionData) (value bool) {
	value, _ = getSettingKey(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_HIDDEN).(bool)
	return
}

// Setter
func setSettingWirelessSsid(data connectionData, value []byte) {
	setSettingKey(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_SSID, value)
}
func setSettingWirelessMode(data connectionData, value string) {
	setSettingKey(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_MODE, value)
}
func setSettingWirelessBand(data connectionData, value string) {
	setSettingKey(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_BAND, value)
}
func setSettingWirelessChannel(data connectionData, value uint32) {
	setSettingKey(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_CHANNEL, value)
}
func setSettingWirelessBssid(data connectionData, value []byte) {
	setSettingKey(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_BSSID, value)
}
func setSettingWirelessRate(data connectionData, value uint32) {
	setSettingKey(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_RATE, value)
}
func setSettingWirelessTxPower(data connectionData, value uint32) {
	setSettingKey(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_TX_POWER, value)
}
func setSettingWirelessMacAddress(data connectionData, value []byte) {
	setSettingKey(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_MAC_ADDRESS, value)
}
func setSettingWirelessClonedMacAddress(data connectionData, value []byte) {
	setSettingKey(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_CLONED_MAC_ADDRESS, value)
}
func setSettingWirelessMacAddressBlacklist(data connectionData, value []string) {
	setSettingKey(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_MAC_ADDRESS_BLACKLIST, value)
}
func setSettingWirelessMtu(data connectionData, value uint32) {
	setSettingKey(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_MTU, value)
}
func setSettingWirelessSeenBssids(data connectionData, value []string) {
	setSettingKey(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_SEEN_BSSIDS, value)
}
func setSettingWirelessSec(data connectionData, value string) {
	setSettingKey(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_SEC, value)
}
func setSettingWirelessHidden(data connectionData, value bool) {
	setSettingKey(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_HIDDEN, value)
}

// JSON Getter
func getSettingWirelessSsidJSON(data connectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_SSID, getSettingWirelessKeyType(NM_SETTING_WIRELESS_SSID))
	return
}
func getSettingWirelessModeJSON(data connectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_MODE, getSettingWirelessKeyType(NM_SETTING_WIRELESS_MODE))
	return
}
func getSettingWirelessBandJSON(data connectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_BAND, getSettingWirelessKeyType(NM_SETTING_WIRELESS_BAND))
	return
}
func getSettingWirelessChannelJSON(data connectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_CHANNEL, getSettingWirelessKeyType(NM_SETTING_WIRELESS_CHANNEL))
	return
}
func getSettingWirelessBssidJSON(data connectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_BSSID, getSettingWirelessKeyType(NM_SETTING_WIRELESS_BSSID))
	return
}
func getSettingWirelessRateJSON(data connectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_RATE, getSettingWirelessKeyType(NM_SETTING_WIRELESS_RATE))
	return
}
func getSettingWirelessTxPowerJSON(data connectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_TX_POWER, getSettingWirelessKeyType(NM_SETTING_WIRELESS_TX_POWER))
	return
}
func getSettingWirelessMacAddressJSON(data connectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_MAC_ADDRESS, getSettingWirelessKeyType(NM_SETTING_WIRELESS_MAC_ADDRESS))
	return
}
func getSettingWirelessClonedMacAddressJSON(data connectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_CLONED_MAC_ADDRESS, getSettingWirelessKeyType(NM_SETTING_WIRELESS_CLONED_MAC_ADDRESS))
	return
}
func getSettingWirelessMacAddressBlacklistJSON(data connectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_MAC_ADDRESS_BLACKLIST, getSettingWirelessKeyType(NM_SETTING_WIRELESS_MAC_ADDRESS_BLACKLIST))
	return
}
func getSettingWirelessMtuJSON(data connectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_MTU, getSettingWirelessKeyType(NM_SETTING_WIRELESS_MTU))
	return
}
func getSettingWirelessSeenBssidsJSON(data connectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_SEEN_BSSIDS, getSettingWirelessKeyType(NM_SETTING_WIRELESS_SEEN_BSSIDS))
	return
}
func getSettingWirelessSecJSON(data connectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_SEC, getSettingWirelessKeyType(NM_SETTING_WIRELESS_SEC))
	return
}
func getSettingWirelessHiddenJSON(data connectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_HIDDEN, getSettingWirelessKeyType(NM_SETTING_WIRELESS_HIDDEN))
	return
}

// JSON Setter
func setSettingWirelessSsidJSON(data connectionData, valueJSON string) {
	setSettingKeyJSON(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_SSID, valueJSON, getSettingWirelessKeyType(NM_SETTING_WIRELESS_SSID))
}
func setSettingWirelessModeJSON(data connectionData, valueJSON string) {
	setSettingKeyJSON(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_MODE, valueJSON, getSettingWirelessKeyType(NM_SETTING_WIRELESS_MODE))
}
func setSettingWirelessBandJSON(data connectionData, valueJSON string) {
	setSettingKeyJSON(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_BAND, valueJSON, getSettingWirelessKeyType(NM_SETTING_WIRELESS_BAND))
}
func setSettingWirelessChannelJSON(data connectionData, valueJSON string) {
	setSettingKeyJSON(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_CHANNEL, valueJSON, getSettingWirelessKeyType(NM_SETTING_WIRELESS_CHANNEL))
}
func setSettingWirelessBssidJSON(data connectionData, valueJSON string) {
	setSettingKeyJSON(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_BSSID, valueJSON, getSettingWirelessKeyType(NM_SETTING_WIRELESS_BSSID))
}
func setSettingWirelessRateJSON(data connectionData, valueJSON string) {
	setSettingKeyJSON(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_RATE, valueJSON, getSettingWirelessKeyType(NM_SETTING_WIRELESS_RATE))
}
func setSettingWirelessTxPowerJSON(data connectionData, valueJSON string) {
	setSettingKeyJSON(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_TX_POWER, valueJSON, getSettingWirelessKeyType(NM_SETTING_WIRELESS_TX_POWER))
}
func setSettingWirelessMacAddressJSON(data connectionData, valueJSON string) {
	setSettingKeyJSON(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_MAC_ADDRESS, valueJSON, getSettingWirelessKeyType(NM_SETTING_WIRELESS_MAC_ADDRESS))
}
func setSettingWirelessClonedMacAddressJSON(data connectionData, valueJSON string) {
	setSettingKeyJSON(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_CLONED_MAC_ADDRESS, valueJSON, getSettingWirelessKeyType(NM_SETTING_WIRELESS_CLONED_MAC_ADDRESS))
}
func setSettingWirelessMacAddressBlacklistJSON(data connectionData, valueJSON string) {
	setSettingKeyJSON(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_MAC_ADDRESS_BLACKLIST, valueJSON, getSettingWirelessKeyType(NM_SETTING_WIRELESS_MAC_ADDRESS_BLACKLIST))
}
func setSettingWirelessMtuJSON(data connectionData, valueJSON string) {
	setSettingKeyJSON(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_MTU, valueJSON, getSettingWirelessKeyType(NM_SETTING_WIRELESS_MTU))
}
func setSettingWirelessSeenBssidsJSON(data connectionData, valueJSON string) {
	setSettingKeyJSON(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_SEEN_BSSIDS, valueJSON, getSettingWirelessKeyType(NM_SETTING_WIRELESS_SEEN_BSSIDS))
}
func setSettingWirelessSecJSON(data connectionData, valueJSON string) {
	setSettingKeyJSON(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_SEC, valueJSON, getSettingWirelessKeyType(NM_SETTING_WIRELESS_SEC))
}
func setSettingWirelessHiddenJSON(data connectionData, valueJSON string) {
	setSettingKeyJSON(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_HIDDEN, valueJSON, getSettingWirelessKeyType(NM_SETTING_WIRELESS_HIDDEN))
}

// Logic JSON Setter

// Remover
func removeSettingWirelessSsid(data connectionData) {
	removeSettingKey(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_SSID)
}
func removeSettingWirelessMode(data connectionData) {
	removeSettingKey(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_MODE)
}
func removeSettingWirelessBand(data connectionData) {
	removeSettingKey(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_BAND)
}
func removeSettingWirelessChannel(data connectionData) {
	removeSettingKey(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_CHANNEL)
}
func removeSettingWirelessBssid(data connectionData) {
	removeSettingKey(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_BSSID)
}
func removeSettingWirelessRate(data connectionData) {
	removeSettingKey(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_RATE)
}
func removeSettingWirelessTxPower(data connectionData) {
	removeSettingKey(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_TX_POWER)
}
func removeSettingWirelessMacAddress(data connectionData) {
	removeSettingKey(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_MAC_ADDRESS)
}
func removeSettingWirelessClonedMacAddress(data connectionData) {
	removeSettingKey(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_CLONED_MAC_ADDRESS)
}
func removeSettingWirelessMacAddressBlacklist(data connectionData) {
	removeSettingKey(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_MAC_ADDRESS_BLACKLIST)
}
func removeSettingWirelessMtu(data connectionData) {
	removeSettingKey(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_MTU)
}
func removeSettingWirelessSeenBssids(data connectionData) {
	removeSettingKey(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_SEEN_BSSIDS)
}
func removeSettingWirelessSec(data connectionData) {
	removeSettingKey(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_SEC)
}
func removeSettingWirelessHidden(data connectionData) {
	removeSettingKey(data, NM_SETTING_WIRELESS_SETTING_NAME, NM_SETTING_WIRELESS_HIDDEN)
}
