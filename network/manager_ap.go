/**
 * Copyright (c) 2014 Deepin, Inc.
 *               2014 Xu FaSheng
 *
 * Author:      Xu FaSheng <fasheng.xu@gmail.com>
 * Maintainer:  Xu FaSheng <fasheng.xu@gmail.com>
 *
 * This program is free software; you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation; either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program; if not, see <http://www.gnu.org/licenses/>.
 **/

package network

import (
	nm "dbus/org/freedesktop/networkmanager"
	"pkg.linuxdeepin.com/lib/dbus"
	"pkg.linuxdeepin.com/lib/utils"
)

type apSecType uint32

const (
	apSecNone apSecType = iota
	apSecWep
	apSecPsk
	apSecEap
)

// TODO refactor code, use accesspoint's secret types defined in network-manager
const (
	NMU_SEC_INVALID apSecType = 0
	NMU_SEC_NONE
	NMU_SEC_STATIC_WEP
	NMU_SEC_LEAP
	NMU_SEC_DYNAMIC_WEP
	NMU_SEC_WPA_PSK
	NMU_SEC_WPA_ENTERPRISE
	NMU_SEC_WPA2_PSK
	NMU_SEC_WPA2_ENTERPRIS
)

type accessPoint struct {
	nmAp *nm.AccessPoint

	Ssid         string
	Secured      bool
	SecuredInEap bool
	Strength     uint8
	Path         dbus.ObjectPath
}

func newAccessPoint(apPath dbus.ObjectPath) (ap accessPoint, err error) {
	nmAp, err := nmNewAccessPoint(apPath)
	if err != nil {
		return
	}

	ap = accessPoint{
		nmAp: nmAp,
		Path: apPath,
	}
	ap.updateProperties()
	return
}

func (a *accessPoint) updateProperties() {
	a.Ssid = string(a.nmAp.Ssid.Get())
	a.Secured = getApSecType(a.nmAp) != apSecNone
	a.SecuredInEap = getApSecType(a.nmAp) == apSecEap
	a.Strength = a.nmAp.Strength.Get()
}

// TODO: remove
func calcApStrength(s uint8) uint8 {
	switch {
	case s <= 10:
		return 0
	case s <= 25:
		return 25
	case s <= 50:
		return 50
	case s <= 75:
		return 75
	case s <= 100:
		return 100
	}
	return 0
}

func getApSecType(ap *nm.AccessPoint) apSecType {
	return doParseApSecType(ap.Flags.Get(), ap.WpaFlags.Get(), ap.RsnFlags.Get())
}

func doParseApSecType(flags, wpaFlags, rsnFlags uint32) apSecType {
	r := apSecNone

	if (flags&NM_802_11_AP_FLAGS_PRIVACY != 0) && (wpaFlags == NM_802_11_AP_SEC_NONE) && (rsnFlags == NM_802_11_AP_SEC_NONE) {
		r = apSecWep
	}
	if wpaFlags != NM_802_11_AP_SEC_NONE {
		r = apSecPsk
	}
	if rsnFlags != NM_802_11_AP_SEC_NONE {
		r = apSecPsk
	}
	if (wpaFlags&NM_802_11_AP_SEC_KEY_MGMT_802_1X != 0) || (rsnFlags&NM_802_11_AP_SEC_KEY_MGMT_802_1X != 0) {
		r = apSecEap
	}
	return r
}

func (m *Manager) clearAccessPoints() {
	for devPath, aps := range m.accessPoints {
		for _, ap := range aps {
			m.removeAccessPoint(devPath, ap.Path)
		}
	}
	m.accessPoints = nil
}
func (m *Manager) addAccessPoint(devPath, apPath dbus.ObjectPath) {
	m.accessPointsLocker.Lock()
	defer m.accessPointsLocker.Unlock()

	if m.isAccessPointExists(devPath, apPath) {
		return
	}
	ap, err := newAccessPoint(apPath)
	if err != nil {
		return
	}
	if len(ap.Ssid) == 0 {
		// ignore hidden access point
		return
	}

	// connect properties changed
	ap.nmAp.ConnectPropertiesChanged(func(properties map[string]dbus.Variant) {
		m.accessPointsLocker.Lock()
		defer m.accessPointsLocker.Unlock()

		ap.updateProperties()
		apJSON, _ := marshalJSON(ap)
		dbus.Emit(m, "AccessPointPropertiesChanged", string(devPath), apJSON)
	})

	apJSON, _ := marshalJSON(ap)
	dbus.Emit(m, "AccessPointAdded", string(devPath), apJSON)

	m.accessPoints[devPath] = append(m.accessPoints[devPath], &ap)
}
func (m *Manager) removeAccessPoint(devPath, apPath dbus.ObjectPath) {
	m.accessPointsLocker.Lock()
	defer m.accessPointsLocker.Unlock()

	// get access point information
	var apJSON string
	if ap := m.getAccessPoint(devPath, apPath); ap != nil {
		apJSON, _ = marshalJSON(ap)
	} else {
		apJSON, _ = marshalJSON(accessPoint{Path: apPath})
	}
	dbus.Emit(m, "AccessPointRemoved", string(devPath), apJSON)
	m.doRemoveAccessPoint(devPath, apPath)
}
func (m *Manager) doRemoveAccessPoint(devPath, apPath dbus.ObjectPath) {
	i := m.getAccessPointIndex(devPath, apPath)
	if i < 0 {
		return
	}

	// destroy object to reset all property connects
	aps := m.accessPoints[devPath]
	ap := aps[i]
	nmDestroyAccessPoint(ap.nmAp)

	copy(aps[i:], aps[i+1:])
	aps[len(aps)-1] = nil
	aps = aps[:len(aps)-1]
	m.accessPoints[devPath] = aps
}
func (m *Manager) getAccessPoint(devPath, apPath dbus.ObjectPath) (ap *accessPoint) {
	i := m.getAccessPointIndex(devPath, apPath)
	if i < 0 {
		logger.Warning("could not found access point:", devPath, apPath)
		return
	}
	ap = m.accessPoints[devPath][i]
	return
}
func (m *Manager) isAccessPointExists(devPath, apPath dbus.ObjectPath) bool {
	if m.getAccessPointIndex(devPath, apPath) >= 0 {
		return true
	}
	return false
}
func (m *Manager) getAccessPointIndex(devPath, apPath dbus.ObjectPath) int {
	for i, ap := range m.accessPoints[devPath] {
		if ap.Path == apPath {
			return i
		}
	}
	return -1
}
func (m *Manager) isSsidExists(devPath dbus.ObjectPath, ssid string) bool {
	for _, ap := range m.accessPoints[devPath] {
		if ap.Ssid == ssid {
			return true
		}
	}
	return false
}

// GetAccessPoints return all access points object which marshaled by json.
func (m *Manager) GetAccessPoints(path dbus.ObjectPath) (apsJSON string, err error) {
	m.accessPointsLocker.Lock()
	defer m.accessPointsLocker.Unlock()

	apsJSON, err = marshalJSON(m.accessPoints[path])
	return
}
func (m *Manager) doGetAccessPoints(devPath dbus.ObjectPath) (aps []accessPoint, err error) {
	apPaths := nmGetAccessPoints(devPath)
	for _, path := range apPaths {
		ap, err := newAccessPoint(path)
		if err != nil {
			continue
		}
		if len(ap.Ssid) == 0 {
			// ignore hidden access point
			continue
		}
		aps = append(aps, ap)
	}
	return
}

// ActivateAccessPoint add and activate connection for access point.
func (m *Manager) ActivateAccessPoint(uuid string, apPath, devPath dbus.ObjectPath) (cpath dbus.ObjectPath, err error) {
	logger.Debugf("ActivateAccessPoint: uuid=%s, apPath=%s, devPath=%s", uuid, apPath, devPath)
	defer logger.Debugf("ActivateAccessPoint end") // TODO test

	if len(uuid) > 0 {
		cpath, err = m.ActivateConnection(uuid, devPath)
	} else {
		// if there is no connection for current access point, create one
		var nmAp *nm.AccessPoint
		nmAp, err = nmNewAccessPoint(apPath)
		if err != nil {
			return
		}
		defer nmDestroyAccessPoint(nmAp)

		uuid = utils.GenUuid()
		data := newWirelessConnectionData(string(nmAp.Ssid.Get()), uuid, []byte(nmAp.Ssid.Get()), getApSecType(nmAp))
		cpath, _, err = nmAddAndActivateConnection(data, devPath)
	}
	return
}

func (m *Manager) CreateConnectionForAccessPoint(apPath, devPath dbus.ObjectPath) (session *ConnectionSession, err error) {
	session, err = newConnectionSessionByCreate(connectionWireless, devPath)
	if err != nil {
		logger.Error(err)
		return
	}

	// setup access point data
	nmAp, err := nmNewAccessPoint(apPath)
	if err != nil {
		return
	}
	defer nmDestroyAccessPoint(nmAp)

	setSettingConnectionId(session.data, string(nmAp.Ssid.Get()))
	setSettingWirelessSsid(session.data, []byte(nmAp.Ssid.Get()))
	secType := getApSecType(nmAp)
	switch secType {
	case apSecNone:
		logicSetSettingVkWirelessSecurityKeyMgmt(session.data, "none")
	case apSecWep:
		logicSetSettingVkWirelessSecurityKeyMgmt(session.data, "wep")
	case apSecPsk:
		logicSetSettingVkWirelessSecurityKeyMgmt(session.data, "wpa-psk")
	case apSecEap:
		logicSetSettingVkWirelessSecurityKeyMgmt(session.data, "wpa-eap")
	}
	session.setProps()

	// install dbus session
	m.addConnectionSession(session)
	return
}
