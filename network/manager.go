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
	"pkg.linuxdeepin.com/lib/dbus"
	"sync"
	"time"
)

const (
	dbusNetworkDest = "com.deepin.daemon.Network"
	dbusNetworkPath = "/com/deepin/daemon/Network"
	dbusNetworkIfs  = "com.deepin.daemon.Network"
)

const (
	opAdded = iota
	opRemoved
	opUpdated
)

type connectionData map[string]map[string]dbus.Variant

type Manager struct {
	config *config

	// update by manager.go
	State uint32 // global networking state

	NetworkingEnabled bool `access:"readwrite"` // airplane mode for NetworkManager
	VpnEnabled        bool `access:"readwrite"`

	// hidden properties
	wirelessEnabled bool `access:"readwrite"`
	wwanEnabled     bool `access:"readwrite"`
	wiredEnabled    bool `access:"readwrite"`

	// update by manager_devices.go
	devicesLocker sync.Mutex
	devices       map[string][]*device
	Devices       string // array of device objects and marshaled by json

	accessPointsLocker sync.Mutex
	accessPoints       map[dbus.ObjectPath][]*accessPoint

	// update by manager_connections.go
	connectionsLocker sync.Mutex
	connections       map[string][]*connection
	Connections       string // array of connection information and marshaled by json

	connectionSessionsLocker sync.Mutex
	connectionSessions       []*ConnectionSession

	// update by manager_active.go
	activeConnectionsLocker sync.Mutex
	activeConnections       map[dbus.ObjectPath]*activeConnection
	ActiveConnections       string // array of connections that activated and marshaled by json

	// signals
	NeedSecrets                  func(connPath, settingName, connectionId string, autoConnect bool)
	AccessPointAdded             func(devPath, apJSON string)
	AccessPointRemoved           func(devPath, apJSON string)
	AccessPointPropertiesChanged func(devPath, apJSON string)
	DeviceEnabled                func(devPath string, enabled bool)

	agent         *agent
	stateHandler  *stateHandler
	dbusWatcher   *dbusWatcher
	switchHandler *switchHandler
}

func (m *Manager) GetDBusInfo() dbus.DBusInfo {
	return dbus.DBusInfo{
		Dest:       dbusNetworkDest,
		ObjectPath: dbusNetworkPath,
		Interface:  dbusNetworkIfs,
	}
}

// initialize slice code
func initSlices() {
	initProxyGsettings()
	initAvailableValuesSecretFlags()
	initAvailableValuesNmPptpSecretFlags()
	initAvailableValuesNmL2tpSecretFlags()
	initAvailableValuesNmVpncSecretFlags()
	initAvailableValuesNmOpenvpnSecretFlags()
	initAvailableValuesWirelessChannel()
	initAvailableValues8021x()
	initAvailableValuesIp4()
	initAvailableValuesIp6()
	initNmStateReasons()
}

func NewManager() (m *Manager) {
	m = &Manager{}
	return
}
func DestroyManager(m *Manager) {
	m.destroyManager()
}

func (m *Manager) initManager() {
	logger.Info("initialize network")
	initDbusObjects()
	disableNotify()
	defer enableNotify()
	m.config = newConfig()
	m.switchHandler = newSwitchHandler(m.config)
	m.dbusWatcher = newDbusWatcher(true)
	m.stateHandler = newStateHandler()
	m.agent = newAgent()

	// initialize device and connection handlers
	m.initDeviceManage()
	m.initConnectionManage()
	m.initActiveConnectionManage()

	// update property "State"
	nmManager.State.ConnectChanged(func() {
		m.setPropState()
	})
	m.setPropState()

	// connect computer suspend signal
	loginManager.ConnectPrepareForSleep(func(active bool) {
		if active {
			// suspend
			disableNotify()
		} else {
			// restore
			m.switchHandler.init()
			enableNotify()
		}
	})
}

func (m *Manager) destroyManager() {
	logger.Info("destroy network")
	destroyDbusObjects()
	destroyAgent(m.agent)
	destroyStateHandler(m.stateHandler)
	destroyDbusWatcher(m.dbusWatcher)
	destroySwitchHandler(m.switchHandler)
	m.clearDevices()
	m.clearAccessPoints()
	m.clearConnections()
	m.clearConnectionSessions()
	m.clearActiveConnections()

	// reset dbus properties
	m.setPropNetworkingEnabled(false)
	m.setPropState()
}

func watchNetworkManagerRestart(m *Manager) {
	dbusDaemon.ConnectNameOwnerChanged(func(name, oldOwner, newOwner string) {
		if name == "org.freedesktop.NetworkManager" {
			// if a new dbus session was installed, the name and newOwner
			// will be not empty, if a dbus session was uninstalled, the
			// name and oldOwner will be not empty
			if len(newOwner) != 0 {
				// network-manager is starting
				logger.Info("network-manager is starting")
				time.Sleep(1 * time.Second)
				m.initManager()
			} else {
				// network-manager stopped
				logger.Info("network-manager stopped")
				m.destroyManager()
			}
		}
	})
}
