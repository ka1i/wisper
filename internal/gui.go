package internal

import (
	"log"
	"time"

	"github.com/progrium/macdriver/cocoa"
	"github.com/progrium/macdriver/core"
	"github.com/progrium/macdriver/objc"
	"github.com/progrium/macdriver/webkit"
	"github.com/progrium/watcher"
)

func Wisper(localserve string, fw *watcher.Watcher) {
	ticker := time.NewTicker(time.Second * 30)

	cocoa.TerminateAfterWindowsClose = false

	config := webkit.WKWebViewConfiguration_New()
	config.Preferences().SetValueForKey(core.True, core.String("developerExtrasEnabled"))

	url := core.URL(localserve)
	req := core.NSURLRequest_Init(url)

	app := cocoa.NSApp_WithDidLaunch(func(_ objc.Object) {
		wv := webkit.WKWebView_Init(cocoa.NSScreen_Main().Frame(), config)
		wv.Retain()
		wv.SetOpaque(false)
		wv.SetBackgroundColor(cocoa.NSColor_Clear())
		wv.SetValueForKey(core.False, core.String("drawsBackground"))
		wv.LoadRequest(req)
		win := cocoa.NSWindow_Init(cocoa.NSScreen_Main().Frame(),
			cocoa.NSClosableWindowMask|cocoa.NSBorderlessWindowMask,
			cocoa.NSBackingStoreBuffered, false)
		win.SetContentView(wv)
		win.SetBackgroundColor(cocoa.NSColor_Clear())
		win.SetOpaque(false)
		win.SetTitleVisibility(cocoa.NSWindowTitleHidden)
		win.SetTitlebarAppearsTransparent(true)
		win.SetIgnoresMouseEvents(true)
		win.SetLevel(cocoa.NSMainMenuWindowLevel - 25)
		win.MakeKeyAndOrderFront(win)
		win.SetCollectionBehavior(cocoa.NSWindowCollectionBehaviorCanJoinAllSpaces)
		win.Send("setHasShadow:", false)

		statusBar := cocoa.NSStatusBar_System().StatusItemWithLength(cocoa.NSVariableStatusItemLength)
		statusBar.Retain()
		statusBar.Button().SetTitle("wisper")

		menuEnabled := cocoa.NSMenuItem_New()
		menuEnabled.Retain()
		menuEnabled.SetTitle("Enabled")
		menuEnabled.SetState(1)
		menuEnabled.SetAction(objc.Sel("enabled:"))
		cocoa.DefaultDelegateClass.AddMethod("enabled:", func(_ objc.Object) {
			if win.IsVisible() {
				win.Send("orderOut:", win)
				menuEnabled.SetState(0)
			} else {
				win.Send("orderFront:", win)
				menuEnabled.SetState(1)
			}
		})

		menuQuit := cocoa.NSMenuItem_New()
		menuQuit.SetTitle("Quit")
		menuQuit.SetAction(objc.Sel("terminate:"))

		menu := cocoa.NSMenu_New()
		menu.SetAutoenablesItems(false)
		menu.AddItem(menuEnabled)
		menu.AddItem(cocoa.NSMenuItem_Separator())
		menu.AddItem(menuQuit)

		statusBar.SetMenu(menu)

		go func() {
			for {
				select {
				case event := <-fw.Event:
					if event.IsDir() {
						continue
					}
					wv.Reload(nil)
					log.Printf("event - hotload successfully: %v", event.FileInfo.Name())
				case <-fw.Closed:
					return
				case <-ticker.C:
					wv.Reload(nil)
				}
			}
		}()
	})

	app.ActivateIgnoringOtherApps(true)
	app.Run()
}
