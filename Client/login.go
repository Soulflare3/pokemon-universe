/*Pokemon Universe MMORPG
Copyright (C) 2010 the Pokemon Universe Authors

This program is free software; you can redistribute it and/or
modify it under the terms of the GNU General Public License
as published by the Free Software Foundation; either version 2
of the License, or (at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program; if not, write to the Free Software
Foundation, Inc., 51 Franklin Street, Fifth Floor, Boston, MA  02110-1301, USA.*/
package main

import (
	"sdl"
)

type PU_LoginControls struct {
	txtUsername *PU_Textfield
	txtPassword *PU_Textfield
	txtStatus *PU_Textfield
}

var g_loginControls *PU_LoginControls = NewLoginControls()

func NewLoginControls() *PU_LoginControls {
	return &PU_LoginControls{}
}

func (l *PU_LoginControls) Show() {
	if l.txtUsername == nil {
		l.txtUsername = NewTextfield(NewRect(453, 396, 160, 20), FONT_PURITANBOLD_14)
		l.txtUsername.SetColor(57, 92, 196)
		l.txtUsername.SetStyle(true, false, false)
		l.txtUsername.readonly = false
		l.txtUsername.KeyDownCallback = LoginKeydown
	}
	
	if l.txtPassword == nil {
		l.txtPassword = NewTextfield(NewRect(453, 424, 160, 20), FONT_PURITANBOLD_14)
		l.txtPassword.SetColor(57, 92, 196)
		l.txtPassword.SetStyle(true, false, false)
		l.txtPassword.readonly = false
		l.txtPassword.password = true
		l.txtPassword.KeyDownCallback = LoginKeydown
	}
	
	if l.txtStatus == nil {
		l.txtStatus = NewTextfield(NewRect(393, 514, 250, 20), FONT_PURITANBOLD_14)
		l.txtStatus.SetColor(0, 185, 47)
		l.txtStatus.SetStyle(true, false, false)
		l.txtStatus.readonly = true
	}
	
	g_gui.SetFocus(l.txtUsername)
}

func (l *PU_LoginControls) Hide() {
	if l.txtUsername != nil {
		g_gui.RemoveElement(l.txtUsername)
		l.txtUsername = nil
	}
	
	if l.txtPassword != nil {
		g_gui.RemoveElement(l.txtPassword)
		l.txtPassword = nil
	} 
	
	if l.txtStatus != nil {
		g_gui.RemoveElement(l.txtStatus)
		l.txtStatus = nil
	}
}

func LoginKeydown(_keysym int, _scancode int) {
	if _scancode == 13 { // enter/return
		var username string
		if g_loginControls.txtUsername != nil {
			username = g_loginControls.txtUsername.text
		}
		
		var password string
		if g_loginControls.txtPassword != nil {
			password = g_loginControls.txtPassword.text
		}
		
		go StartLogin(username, password)
	}
}

func StartLogin(_username string, _password string) {
	g_loginControls.txtStatus.SetColor(0, 185, 47)
	g_loginControls.txtUsername.readonly = true
	g_loginControls.txtPassword.readonly = true
	
	if DoLogin(_username, _password) {
		g_loginControls.Hide()
	} else {
		g_loginControls.txtStatus.SetColor(202, 0, 0)
		g_loginControls.txtUsername.readonly = false
		g_loginControls.txtPassword.readonly = false
		
		g_conn.Close()
	}
}

func DoLogin(_username string, _password string) bool {
	g_loginControls.txtStatus.text = "Connecting to server..."
	g_conn.loginStatus = LOGINSTATUS_IDLE
	
	if !g_conn.Connect() {
		g_loginControls.txtStatus.text = "Could not connect to server."
		return false
	}
	
	g_loginControls.txtStatus.text = "Verifying username and password..."
	g_conn.protocol.SendLogin(_username, _password)
	
	timeout := uint16(0)
	for g_conn.loginStatus == LOGINSTATUS_IDLE {
		sdl.Delay(500)
		timeout += 500
		
		if timeout >= 10000 { //10 sec
			g_loginControls.txtStatus.text = "Timeout while verifying data. Please retry."
			return false
		}
	}
	
	switch g_conn.loginStatus {
		case LOGINSTATUS_WRONGACCOUNT:
			g_loginControls.txtStatus.text = "Invalid username and/or password."
			return false
			
		case LOGINSTATUS_SERVERERROR:
			g_loginControls.txtStatus.text = "The server has returned an error. Please retry."
			return false
			
		case LOGINSTATUS_DATABASEERROR:
			g_loginControls.txtStatus.text = "The database has returned an error. Please retry."
			return false
			
		case LOGINSTATUS_ALREADYLOGGEDIN:
			g_loginControls.txtStatus.text = "You are already logged in."
			return false
			
		case LOGINSTATUS_CHARBANNED:
			g_loginControls.txtStatus.text = "This account is banned from the game."
			return false
			
		case LOGINSTATUS_SERVERCLOSED:
			g_loginControls.txtStatus.text = "The server is currently closed."
			return false
			
		case LOGINSTATUS_WRONGVERSION:
			g_loginControls.txtStatus.text = "This client version is outdated."
			return false
			
		case LOGINSTATUS_FAILPROFILELOAD:
			g_loginControls.txtStatus.text = "Your profile could not be loaded. Please retry."
			return false
	}
	g_conn.loginStatus = LOGINSTATUS_IDLE
	
	//the following part is only temporary until the loginserver is ready to be used
	g_loginControls.txtStatus.text = "Loading gameworld..."
	g_conn.Game().SendRequestLoginPackets()
	
	timeout = 0
	for g_conn.loginStatus != LOGINSTATUS_READY {
		sdl.Delay(500)
		timeout += 500
		
		if timeout >= 30000 { //30 sec
			g_loginControls.txtStatus.text = "Timeout while loading gameworld. Please retry."
			return false
		}
	}
	
	g_game.panel = NewGamePanel()
	g_game.CreateChat()
	g_game.state = GAMESTATE_WORLD
	
	return true
}
