/*
 * Copyright (C) 2016 ~ 2018 Deepin Technology Co., Ltd.
 *
 * Author:     jouyouyun <jouyouwen717@gmail.com>
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package keybinding

import (
	"github.com/linuxdeepin/go-dbus-factory/com.deepin.daemon.display"
	"github.com/linuxdeepin/go-dbus-factory/com.deepin.daemon.helper.backlight"
	. "pkg.deepin.io/dde/daemon/keybinding/shortcuts"
	"pkg.deepin.io/lib/dbus1"
)

type DisplayController struct {
	display         *display.Display
	backlightHelper *backlight.Backlight
}

func NewDisplayController(backlightHelper *backlight.Backlight, sessionConn *dbus.Conn) *DisplayController {
	return &DisplayController{
		backlightHelper: backlightHelper,
		display:         display.NewDisplay(sessionConn),
	}
}

func (*DisplayController) Name() string {
	return "Display"
}

func (c *DisplayController) ExecCmd(cmd ActionCmd) error {
	switch cmd {
	case DisplayModeSwitch:
		showOSD("SwitchMonitors")
		return nil

	case MonitorBrightnessUp:
		return c.changeBrightness(true)

	case MonitorBrightnessDown:
		return c.changeBrightness(false)

	default:
		return ErrInvalidActionCmd{cmd}
	}
}

func (c *DisplayController) changeBrightness(raised bool) error {
	var osd = "BrightnessUp"
	if !raised {
		osd = "BrightnessDown"
	}

	err := c.display.ChangeBrightness(dbus.FlagNoAutoStart, raised)
	if err != nil {
		return err
	}

	showOSD(osd)
	return nil
}
