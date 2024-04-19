package cmd

import (
	"log"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type MainFlex struct {
	*tview.Flex

	guildsTree   *GuildsTree
	messagesText *MessagesText
	messageInput *MessageInput
	userList     *UserList

	guildsTreeVisible bool
	userListVisible   bool
}

func newMainFlex() *MainFlex {
	mf := &MainFlex{
		Flex: tview.NewFlex(),

		guildsTree:        newGuildsTree(),
		messagesText:      newMessagesText(),
		messageInput:      newMessageInput(),
		userList:          newUserList(),
		userListVisible:   true,
		guildsTreeVisible: true,
	}

	mf.init()
	mf.SetInputCapture(mf.onInputCapture)
	return mf
}

func (mf *MainFlex) init() {
	mf.Clear()

	chat := tview.NewFlex()
	chat.SetDirection(tview.FlexRow)
	chat.AddItem(mf.messagesText, 0, 1, false)
	chat.AddItem(mf.messageInput, 3, 1, false)
	// The guilds tree is always focused first at start-up.
	if mf.guildsTreeVisible {
		mf.AddItem(mf.guildsTree, 0, 1, mf.guildsTreeVisible)
	}
	mf.AddItem(chat, 0, 4, false)
	if mf.userListVisible {
		mf.AddItem(mf.userList, 0, 1, mf.userListVisible)
	}
}

func (mf *MainFlex) onInputCapture(event *tcell.EventKey) *tcell.EventKey {
	log.Println(event.Name())
	switch event.Name() {
	case cfg.Keys.FocusGuildsTree:
		app.SetFocus(mf.guildsTree)
		return nil
	case cfg.Keys.FocusMessagesText:
		app.SetFocus(mf.messagesText)
		return nil
	case cfg.Keys.FocusMessageInput:
		app.SetFocus(mf.messageInput)
		return nil
	case cfg.Keys.FocusUserList:
		app.SetFocus(mf.userList)
		return nil
	case cfg.Keys.ToggleGuildsTree:
		// The guilds tree is visible if the numbers of items is two.
		if mf.guildsTreeVisible {
			mf.RemoveItem(mf.guildsTree)
			if mf.guildsTree.HasFocus() {
				app.SetFocus(mf)
			}
			mf.guildsTreeVisible = false
		} else {
			mf.guildsTreeVisible = true
			mf.init()
			app.SetFocus(mf.guildsTree)
			return nil
		}
	case cfg.Keys.ToggleUserList:
		// The user list is visible if the state boolean says so
		if mf.userListVisible {
			mf.RemoveItem(mf.userList)
			if mf.userList.HasFocus() {
				app.SetFocus(mf)
			}
			mf.userListVisible = false
		} else {
			mf.userListVisible = true
			mf.init()
			app.SetFocus(mf.userList)
			return nil
		}
	}

	return event
}
