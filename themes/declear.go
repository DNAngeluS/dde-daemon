/**
 * Copyright (c) 2011 ~ 2014 Deepin, Inc.
 *               2013 ~ 2014 jouyouyun
 *
 * Author:      jouyouyun <jouyouwen717@gmail.com>
 * Maintainer:  jouyouyun <jouyouwen717@gmail.com>
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

package themes

import (
	"path"
)

const (
	MANAGER_DEST = "com.deepin.daemon.Themes"
	MANAGER_PATH = "/com/deepin/daemon/Themes"
	MANAGER_IFC  = "com.deepin.daemon.Themes"

	THEME_PATH = "/com/deepin/daemon/"
	THEME_IFC  = "com.deepin.daemon.Theme"
)

const (
	THEME_SYS_PATH   = "/usr/share/themes"
	THEME_LOCAL_PATH = ".themes"
	ICON_SYS_PATH    = "/usr/share/icons"
	ICON_LOCAL_PATH  = ".icons"
	SOUND_THEME_PATH = "/usr/share/sounds"

	PERSON_SYS_BASE_PATH   = "/usr/share/personalization"
	PERSON_LOCAL_BASE_PATH = ".local/share/personalization"

	THEME_BG_NAME      = "wallpappers"
	DEFAULT_SYS_BG_DIR = "/usr/share/backgrounds"

	THEME_CONFIG_NAME = "theme.ini"
)

var (
	PERSON_SYS_THEME_PATH   = path.Join(PERSON_SYS_BASE_PATH, "themes")
	PERSON_LOCAL_THEME_PATH = path.Join(PERSON_LOCAL_BASE_PATH, "themes")
)

const (
	THEME_GROUP_THEME     = "Theme"
	THEME_GROUP_COMPONENT = "Component"
	THEME_KEY_NAME        = "Name"
	THEME_KEY_GTK         = "GtkTheme"
	THEME_KEY_ICON        = "IconTheme"
	THEME_KEY_SOUND       = "SoundTheme"
	THEME_KEY_CURSOR      = "CursorTheme"
	THEME_KEY_BACKGROUND  = "BackgroundFile"
	THEME_KEY_FONT_SIZE   = "FontSize"
)

const (
	GS_KEY_CURRENT_THEME = "current-theme"
	GS_KEY_CURRENT_BG    = "current-picture"
	GS_KEY_CURRENT_SOUND = "current-sound-theme"
)