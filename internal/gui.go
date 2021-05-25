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

func Wisper(localserve string, fw *watcher.Watcher, guichan chan Api) {
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

		// browser
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
		// text
		c, tr := BrowserGate(&localserve)
		w := cocoa.NSWindow_Init(core.Rect(0, 0, 0, 0),
			cocoa.NSBorderlessWindowMask, cocoa.NSBackingStoreBuffered, false)
		w.SetContentView(c)
		w.SetTitlebarAppearsTransparent(true)
		w.SetTitleVisibility(cocoa.NSWindowTitleHidden)
		w.SetOpaque(false)
		w.SetBackgroundColor(cocoa.NSColor_Clear())
		w.SetLevel(cocoa.NSMainMenuWindowLevel + 2)
		w.SetFrameDisplay(tr, true)
		w.MakeKeyAndOrderFront(nil)

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
				case d := <-guichan:
					if d.op == "update" {
						//log.Println(d.parms)
						nurl := core.URL(d.parms)
						nreq := core.NSURLRequest_Init(nurl)
						wv.LoadRequest(nreq)
					} else {
						wv.Reload(nil)
					}
				}
			}
		}()
	})

	app.ActivateIgnoringOtherApps(true)
	app.Run()
}

func BrowserGate(localserve *string) (cocoa.NSView, core.NSRect) {
	screen := cocoa.NSScreen_Main().Frame().Size
	text := localserve
	tr, fontSize := func() (rect core.NSRect, size float64) {
		t := cocoa.NSTextView_Init(core.Rect(0, 0, 0, 0))
		t.SetString(*text)
		for s := 12.0; s < 24; s += 6 {
			t.SetFont(cocoa.Font("Monaco", s))
			t.SetBackgroundColor(cocoa.Color(0.6, 0.3, 0.7, 0.5))
			t.LayoutManager().EnsureLayoutForTextContainer(t.TextContainer())
			rect = t.LayoutManager().UsedRectForTextContainer(t.TextContainer())
			size = s
			if rect.Size.Width >= screen.Width*0.8 {
				break
			}
		}
		return rect, size
	}()

	height := tr.Size.Height * 1
	tr.Origin.Y = (height / 2) - (tr.Size.Height / 2)
	t := cocoa.NSTextView_Init(tr)
	t.SetString(*text)
	t.SetFont(cocoa.Font("Monaco", fontSize))
	t.SetTextColor(cocoa.Color(0.6, 0.3, 0.7, 1))
	t.SetBackgroundColor(cocoa.Color(0.6, 0.3, 0.7, 0.8))
	t.SetEditable(false)
	t.SetImportsGraphics(false)
	t.SetDrawsBackground(false)

	// t1 := cocoa.NSTextView_Init(tr)
	// t1.SetString(*text)
	// t1.SetFont(cocoa.Font("Monaco", 48))
	// t1.SetTextColor(cocoa.Color(0.6, 0.3, 0.7, 1))
	// t1.SetBackgroundColor(cocoa.Color(0.6, 0.3, 0.7, 0.8))
	// t1.SetEditable(false)
	// t1.SetImportsGraphics(false)
	// t1.SetDrawsBackground(false)

	c := cocoa.NSView_Init(core.Rect(0, 0, 0, 0))
	//c.SetBackgroundColor(cocoa.Color(0, 0, 0, 0.75))
	c.SetWantsLayer(true)
	//c.Layer().SetCornerRadius(32.0)
	c.AddSubviewPositionedRelativeTo(t, cocoa.NSWindowAbove, nil)
	//c.AddSubviewPositionedRelativeTo(t1, cocoa.NSWindowAbove, nil)

	tr.Size.Height = 200 //height
	//tr.Size.Width = 1500
	tr.Origin.X = 0 //(screen.Width / 2) - (tr.Size.Width / 2)
	tr.Origin.Y = 0 //(screen.Height / 2) - (tr.Size.Height / 2)

	return c, tr
}
