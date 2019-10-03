package render

import (
	"fmt"
	"sync"

	"github.com/gopherjs/gopherjs/js"
)

// Listener GoRender listener util
type Listener struct {
	IMMRootSelector        string
	CurListenEventType     []string
	CurListenToChannel     map[string]chan *js.Object
	CurListenToChannelLock sync.RWMutex
}

// NewListener create new listener
func NewListener(rootSelector string) *Listener {
	lis := &Listener{
		IMMRootSelector:    rootSelector,
		CurListenToChannel: make(map[string]chan *js.Object),
	}
	js.Global.Set("GoRender_ListListenEvent", func() {
		lis.listListenEvent()
	})
	return lis
}

func (lis *Listener) listListenEvent() {
	fmt.Println("Current Listen:")
	for key := range lis.CurListenToChannel {
		fmt.Println(key)
	}
	fmt.Println("Current Listen show finish")
}

func (lis *Listener) checkEventAlreadyListenOnRuntime(eventType string) {
	for _, eventTypeItem := range lis.CurListenEventType {
		if eventType == eventTypeItem {
			return
		}
	}
	lis.CurListenEventType = append(lis.CurListenEventType, eventType)
	element := js.Global.Get("document").Call("querySelector", lis.IMMRootSelector)
	element.Call("addEventListener", eventType, func(event *js.Object) {
		lis.handleEvent(event.Get("type").String(), event.Get("target").Call("getAttribute", "grid").String(), event)
	})
}

func (lis *Listener) handleEvent(eventType, grid string, event *js.Object) {
	lis.CurListenToChannelLock.RLock()
	defer lis.CurListenToChannelLock.RUnlock()
	if channel, ok := lis.CurListenToChannel[lis.getChannelKey(eventType, grid)]; ok {
		channel <- event
	}
}

func (lis *Listener) getChannelKey(eventType, grid string) string {
	return eventType + "_" + grid
}

// Bind bind a new event
func (lis *Listener) Bind(eventType, grid string) chan *js.Object {
	lis.checkEventAlreadyListenOnRuntime(eventType)
	newchan := make(chan *js.Object, 1000)
	lis.CurListenToChannelLock.Lock()
	defer lis.CurListenToChannelLock.Unlock()
	lis.CurListenToChannel[lis.getChannelKey(eventType, grid)] = newchan
	return newchan
}

// Unbind unbind a event
func (lis *Listener) Unbind(eventType, grid string) {
	lis.CurListenToChannelLock.RLock()
	defer lis.CurListenToChannelLock.RUnlock()
	if channel, ok := lis.CurListenToChannel[lis.getChannelKey(eventType, grid)]; ok {
		close(channel)
	} else {
		fmt.Println("Warning: event listener ", lis.getChannelKey(eventType, grid), " not found!")
	}
}
