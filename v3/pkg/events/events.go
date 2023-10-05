package events

type ApplicationEventType uint
type WindowEventType uint

const (
	FilesDropped WindowEventType = iota
)

var Common = newCommonEvents()

type commonEvents struct {
	ApplicationStarted ApplicationEventType
	WindowMaximise     WindowEventType
	WindowUnMaximise   WindowEventType
	WindowFullscreen   WindowEventType
	WindowUnFullscreen WindowEventType
	WindowRestore      WindowEventType
	WindowMinimise     WindowEventType
	WindowUnMinimise   WindowEventType
	WindowClosing      WindowEventType
	WindowZoom         WindowEventType
	WindowZoomIn       WindowEventType
	WindowZoomOut      WindowEventType
	WindowZoomReset    WindowEventType
	WindowFocus        WindowEventType
	WindowLostFocus    WindowEventType
	WindowShow         WindowEventType
	WindowHide         WindowEventType
	WindowDPIChanged   WindowEventType
	ThemeChanged       ApplicationEventType
}

func newCommonEvents() commonEvents {
	return commonEvents{
		ApplicationStarted: 4096,
		WindowMaximise:     4097,
		WindowUnMaximise:   4098,
		WindowFullscreen:   4099,
		WindowUnFullscreen: 4100,
		WindowRestore:      4101,
		WindowMinimise:     4102,
		WindowUnMinimise:   4103,
		WindowClosing:      4104,
		WindowZoom:         4105,
		WindowZoomIn:       4106,
		WindowZoomOut:      4107,
		WindowZoomReset:    4108,
		WindowFocus:        4109,
		WindowLostFocus:    4110,
		WindowShow:         4111,
		WindowHide:         4112,
		WindowDPIChanged:   4113,
		ThemeChanged:       4114,
	}
}

var Mac = newMacEvents()

type macEvents struct {
	ApplicationDidBecomeActive                              ApplicationEventType
	ApplicationDidChangeBackingProperties                   ApplicationEventType
	ApplicationDidChangeEffectiveAppearance                 ApplicationEventType
	ApplicationDidChangeIcon                                ApplicationEventType
	ApplicationDidChangeOcclusionState                      ApplicationEventType
	ApplicationDidChangeScreenParameters                    ApplicationEventType
	ApplicationDidChangeStatusBarFrame                      ApplicationEventType
	ApplicationDidChangeStatusBarOrientation                ApplicationEventType
	ApplicationDidFinishLaunching                           ApplicationEventType
	ApplicationDidHide                                      ApplicationEventType
	ApplicationDidResignActiveNotification                  ApplicationEventType
	ApplicationDidUnhide                                    ApplicationEventType
	ApplicationDidUpdate                                    ApplicationEventType
	ApplicationWillBecomeActive                             ApplicationEventType
	ApplicationWillFinishLaunching                          ApplicationEventType
	ApplicationWillHide                                     ApplicationEventType
	ApplicationWillResignActive                             ApplicationEventType
	ApplicationWillTerminate                                ApplicationEventType
	ApplicationWillUnhide                                   ApplicationEventType
	ApplicationWillUpdate                                   ApplicationEventType
	ApplicationDidChangeTheme                               ApplicationEventType
	WindowDidBecomeKey                                      WindowEventType
	WindowDidBecomeMain                                     WindowEventType
	WindowDidBeginSheet                                     WindowEventType
	WindowDidChangeAlpha                                    WindowEventType
	WindowDidChangeBackingLocation                          WindowEventType
	WindowDidChangeBackingProperties                        WindowEventType
	WindowDidChangeCollectionBehavior                       WindowEventType
	WindowDidChangeEffectiveAppearance                      WindowEventType
	WindowDidChangeOcclusionState                           WindowEventType
	WindowDidChangeOrderingMode                             WindowEventType
	WindowDidChangeScreen                                   WindowEventType
	WindowDidChangeScreenParameters                         WindowEventType
	WindowDidChangeScreenProfile                            WindowEventType
	WindowDidChangeScreenSpace                              WindowEventType
	WindowDidChangeScreenSpaceProperties                    WindowEventType
	WindowDidChangeSharingType                              WindowEventType
	WindowDidChangeSpace                                    WindowEventType
	WindowDidChangeSpaceOrderingMode                        WindowEventType
	WindowDidChangeTitle                                    WindowEventType
	WindowDidChangeToolbar                                  WindowEventType
	WindowDidChangeVisibility                               WindowEventType
	WindowDidDeminiaturize                                  WindowEventType
	WindowDidEndSheet                                       WindowEventType
	WindowDidEnterFullScreen                                WindowEventType
	WindowDidEnterVersionBrowser                            WindowEventType
	WindowDidExitFullScreen                                 WindowEventType
	WindowDidExitVersionBrowser                             WindowEventType
	WindowDidExpose                                         WindowEventType
	WindowDidFocus                                          WindowEventType
	WindowDidMiniaturize                                    WindowEventType
	WindowDidMove                                           WindowEventType
	WindowDidOrderOffScreen                                 WindowEventType
	WindowDidOrderOnScreen                                  WindowEventType
	WindowDidResignKey                                      WindowEventType
	WindowDidResignMain                                     WindowEventType
	WindowDidResize                                         WindowEventType
	WindowDidUpdate                                         WindowEventType
	WindowDidUpdateAlpha                                    WindowEventType
	WindowDidUpdateCollectionBehavior                       WindowEventType
	WindowDidUpdateCollectionProperties                     WindowEventType
	WindowDidUpdateShadow                                   WindowEventType
	WindowDidUpdateTitle                                    WindowEventType
	WindowDidUpdateToolbar                                  WindowEventType
	WindowDidUpdateVisibility                               WindowEventType
	WindowShouldClose                                       WindowEventType
	WindowWillBecomeKey                                     WindowEventType
	WindowWillBecomeMain                                    WindowEventType
	WindowWillBeginSheet                                    WindowEventType
	WindowWillChangeOrderingMode                            WindowEventType
	WindowWillClose                                         WindowEventType
	WindowWillDeminiaturize                                 WindowEventType
	WindowWillEnterFullScreen                               WindowEventType
	WindowWillEnterVersionBrowser                           WindowEventType
	WindowWillExitFullScreen                                WindowEventType
	WindowWillExitVersionBrowser                            WindowEventType
	WindowWillFocus                                         WindowEventType
	WindowWillMiniaturize                                   WindowEventType
	WindowWillMove                                          WindowEventType
	WindowWillOrderOffScreen                                WindowEventType
	WindowWillOrderOnScreen                                 WindowEventType
	WindowWillResignMain                                    WindowEventType
	WindowWillResize                                        WindowEventType
	WindowWillUnfocus                                       WindowEventType
	WindowWillUpdate                                        WindowEventType
	WindowWillUpdateAlpha                                   WindowEventType
	WindowWillUpdateCollectionBehavior                      WindowEventType
	WindowWillUpdateCollectionProperties                    WindowEventType
	WindowWillUpdateShadow                                  WindowEventType
	WindowWillUpdateTitle                                   WindowEventType
	WindowWillUpdateToolbar                                 WindowEventType
	WindowWillUpdateVisibility                              WindowEventType
	WindowWillUseStandardFrame                              WindowEventType
	MenuWillOpen                                            ApplicationEventType
	MenuDidOpen                                             ApplicationEventType
	MenuDidClose                                            ApplicationEventType
	MenuWillSendAction                                      ApplicationEventType
	MenuDidSendAction                                       ApplicationEventType
	MenuWillHighlightItem                                   ApplicationEventType
	MenuDidHighlightItem                                    ApplicationEventType
	MenuWillDisplayItem                                     ApplicationEventType
	MenuDidDisplayItem                                      ApplicationEventType
	MenuWillAddItem                                         ApplicationEventType
	MenuDidAddItem                                          ApplicationEventType
	MenuWillRemoveItem                                      ApplicationEventType
	MenuDidRemoveItem                                       ApplicationEventType
	MenuWillBeginTracking                                   ApplicationEventType
	MenuDidBeginTracking                                    ApplicationEventType
	MenuWillEndTracking                                     ApplicationEventType
	MenuDidEndTracking                                      ApplicationEventType
	MenuWillUpdate                                          ApplicationEventType
	MenuDidUpdate                                           ApplicationEventType
	MenuWillPopUp                                           ApplicationEventType
	MenuDidPopUp                                            ApplicationEventType
	MenuWillSendActionToItem                                ApplicationEventType
	MenuDidSendActionToItem                                 ApplicationEventType
	WebViewDidStartProvisionalNavigation                    WindowEventType
	WebViewDidReceiveServerRedirectForProvisionalNavigation WindowEventType
	WebViewDidFinishNavigation                              WindowEventType
	WebViewDidCommitNavigation                              WindowEventType
	WindowFileDraggingEntered                               WindowEventType
	WindowFileDraggingPerformed                             WindowEventType
	WindowFileDraggingExited                                WindowEventType
}

func newMacEvents() macEvents {
	return macEvents{
		ApplicationDidBecomeActive:               0,
		ApplicationDidChangeBackingProperties:    1,
		ApplicationDidChangeEffectiveAppearance:  2,
		ApplicationDidChangeIcon:                 3,
		ApplicationDidChangeOcclusionState:       4,
		ApplicationDidChangeScreenParameters:     5,
		ApplicationDidChangeStatusBarFrame:       6,
		ApplicationDidChangeStatusBarOrientation: 7,
		ApplicationDidFinishLaunching:            8,
		ApplicationDidHide:                       9,
		ApplicationDidResignActiveNotification:   10,
		ApplicationDidUnhide:                     11,
		ApplicationDidUpdate:                     12,
		ApplicationWillBecomeActive:              13,
		ApplicationWillFinishLaunching:           14,
		ApplicationWillHide:                      15,
		ApplicationWillResignActive:              16,
		ApplicationWillTerminate:                 17,
		ApplicationWillUnhide:                    18,
		ApplicationWillUpdate:                    19,
		ApplicationDidChangeTheme:                20,
		WindowDidBecomeKey:                       20,
		WindowDidBecomeMain:                      21,
		WindowDidBeginSheet:                      22,
		WindowDidChangeAlpha:                     23,
		WindowDidChangeBackingLocation:           24,
		WindowDidChangeBackingProperties:         25,
		WindowDidChangeCollectionBehavior:        26,
		WindowDidChangeEffectiveAppearance:       27,
		WindowDidChangeOcclusionState:            28,
		WindowDidChangeOrderingMode:              29,
		WindowDidChangeScreen:                    30,
		WindowDidChangeScreenParameters:          31,
		WindowDidChangeScreenProfile:             32,
		WindowDidChangeScreenSpace:               33,
		WindowDidChangeScreenSpaceProperties:     34,
		WindowDidChangeSharingType:               35,
		WindowDidChangeSpace:                     36,
		WindowDidChangeSpaceOrderingMode:         37,
		WindowDidChangeTitle:                     38,
		WindowDidChangeToolbar:                   39,
		WindowDidChangeVisibility:                40,
		WindowDidDeminiaturize:                   41,
		WindowDidEndSheet:                        42,
		WindowDidEnterFullScreen:                 43,
		WindowDidEnterVersionBrowser:             44,
		WindowDidExitFullScreen:                  45,
		WindowDidExitVersionBrowser:              46,
		WindowDidExpose:                          47,
		WindowDidFocus:                           48,
		WindowDidMiniaturize:                     49,
		WindowDidMove:                            50,
		WindowDidOrderOffScreen:                  51,
		WindowDidOrderOnScreen:                   52,
		WindowDidResignKey:                       53,
		WindowDidResignMain:                      54,
		WindowDidResize:                          55,
		WindowDidUpdate:                          56,
		WindowDidUpdateAlpha:                     57,
		WindowDidUpdateCollectionBehavior:        58,
		WindowDidUpdateCollectionProperties:      59,
		WindowDidUpdateShadow:                    60,
		WindowDidUpdateTitle:                     61,
		WindowDidUpdateToolbar:                   62,
		WindowDidUpdateVisibility:                63,
		WindowShouldClose:                        64,
		WindowWillBecomeKey:                      64,
		WindowWillBecomeMain:                     65,
		WindowWillBeginSheet:                     66,
		WindowWillChangeOrderingMode:             67,
		WindowWillClose:                          68,
		WindowWillDeminiaturize:                  69,
		WindowWillEnterFullScreen:                70,
		WindowWillEnterVersionBrowser:            71,
		WindowWillExitFullScreen:                 72,
		WindowWillExitVersionBrowser:             73,
		WindowWillFocus:                          74,
		WindowWillMiniaturize:                    75,
		WindowWillMove:                           76,
		WindowWillOrderOffScreen:                 77,
		WindowWillOrderOnScreen:                  78,
		WindowWillResignMain:                     79,
		WindowWillResize:                         80,
		WindowWillUnfocus:                        81,
		WindowWillUpdate:                         82,
		WindowWillUpdateAlpha:                    83,
		WindowWillUpdateCollectionBehavior:       84,
		WindowWillUpdateCollectionProperties:     85,
		WindowWillUpdateShadow:                   86,
		WindowWillUpdateTitle:                    87,
		WindowWillUpdateToolbar:                  88,
		WindowWillUpdateVisibility:               89,
		WindowWillUseStandardFrame:               90,
		MenuWillOpen:                             91,
		MenuDidOpen:                              92,
		MenuDidClose:                             93,
		MenuWillSendAction:                       94,
		MenuDidSendAction:                        95,
		MenuWillHighlightItem:                    96,
		MenuDidHighlightItem:                     97,
		MenuWillDisplayItem:                      98,
		MenuDidDisplayItem:                       99,
		MenuWillAddItem:                          100,
		MenuDidAddItem:                           101,
		MenuWillRemoveItem:                       102,
		MenuDidRemoveItem:                        103,
		MenuWillBeginTracking:                    104,
		MenuDidBeginTracking:                     105,
		MenuWillEndTracking:                      106,
		MenuDidEndTracking:                       107,
		MenuWillUpdate:                           108,
		MenuDidUpdate:                            109,
		MenuWillPopUp:                            110,
		MenuDidPopUp:                             111,
		MenuWillSendActionToItem:                 112,
		MenuDidSendActionToItem:                  113,
		WebViewDidStartProvisionalNavigation:     114,
		WebViewDidReceiveServerRedirectForProvisionalNavigation: 115,
		WebViewDidFinishNavigation:                              116,
		WebViewDidCommitNavigation:                              117,
		WindowFileDraggingEntered:                               118,
		WindowFileDraggingPerformed:                             119,
		WindowFileDraggingExited:                                120,
	}
}

var Linux = newLinuxEvents()

type linuxEvents struct {
	SystemThemeChanged ApplicationEventType
}

func newLinuxEvents() linuxEvents {
	return linuxEvents{
		SystemThemeChanged: 0,
	}
}

var Windows = newWindowsEvents()

type windowsEvents struct {
	SystemThemeChanged         ApplicationEventType
	APMPowerStatusChange       ApplicationEventType
	APMSuspend                 ApplicationEventType
	APMResumeAutomatic         ApplicationEventType
	APMResumeSuspend           ApplicationEventType
	APMPowerSettingChange      ApplicationEventType
	WebViewNavigationCompleted WindowEventType
	WindowInactive             WindowEventType
	WindowActive               WindowEventType
	WindowClickActive          WindowEventType
	WindowMaximise             WindowEventType
	WindowUnMaximise           WindowEventType
	WindowFullscreen           WindowEventType
	WindowUnFullscreen         WindowEventType
	WindowRestore              WindowEventType
	WindowMinimise             WindowEventType
	WindowUnMinimise           WindowEventType
	WindowClose                WindowEventType
	WindowSetFocus             WindowEventType
	WindowKillFocus            WindowEventType
}

func newWindowsEvents() windowsEvents {
	return windowsEvents{
		SystemThemeChanged:         0,
		APMPowerStatusChange:       1,
		APMSuspend:                 2,
		APMResumeAutomatic:         3,
		APMResumeSuspend:           4,
		APMPowerSettingChange:      5,
		WebViewNavigationCompleted: 6,
		WindowInactive:             7,
		WindowActive:               8,
		WindowClickActive:          9,
		WindowMaximise:             10,
		WindowUnMaximise:           11,
		WindowFullscreen:           12,
		WindowUnFullscreen:         13,
		WindowRestore:              14,
		WindowMinimise:             15,
		WindowUnMinimise:           16,
		WindowClose:                17,
		WindowSetFocus:             18,
		WindowKillFocus:            19,
	}
}

func JSEvent(platform string, event uint) string {
	return eventToJS[platform][event]
}

var eventToJS = map[string]map[uint]string{
	"windows": {
		0:  "windows:SystemThemeChanged",
		1:  "windows:APMPowerStatusChange",
		2:  "windows:APMSuspend",
		3:  "windows:APMResumeAutomatic",
		4:  "windows:APMResumeSuspend",
		5:  "windows:APMPowerSettingChange",
		6:  "windows:WebViewNavigationCompleted",
		7:  "windows:WindowInactive",
		8:  "windows:WindowActive",
		9:  "windows:WindowClickActive",
		10: "windows:WindowMaximise",
		11: "windows:WindowUnMaximise",
		12: "windows:WindowFullscreen",
		13: "windows:WindowUnFullscreen",
		14: "windows:WindowRestore",
		15: "windows:WindowMinimise",
		16: "windows:WindowUnMinimise",
		17: "windows:WindowClose",
		18: "windows:WindowSetFocus",
		19: "windows:WindowKillFocus",
	},
	"mac": {
		0:   "mac:ApplicationDidBecomeActive",
		1:   "mac:ApplicationDidChangeBackingProperties",
		2:   "mac:ApplicationDidChangeEffectiveAppearance",
		3:   "mac:ApplicationDidChangeIcon",
		4:   "mac:ApplicationDidChangeOcclusionState",
		5:   "mac:ApplicationDidChangeScreenParameters",
		6:   "mac:ApplicationDidChangeStatusBarFrame",
		7:   "mac:ApplicationDidChangeStatusBarOrientation",
		8:   "mac:ApplicationDidFinishLaunching",
		9:   "mac:ApplicationDidHide",
		10:  "mac:ApplicationDidResignActiveNotification",
		11:  "mac:ApplicationDidUnhide",
		12:  "mac:ApplicationDidUpdate",
		13:  "mac:ApplicationWillBecomeActive",
		14:  "mac:ApplicationWillFinishLaunching",
		15:  "mac:ApplicationWillHide",
		16:  "mac:ApplicationWillResignActive",
		17:  "mac:ApplicationWillTerminate",
		18:  "mac:ApplicationWillUnhide",
		19:  "mac:ApplicationWillUpdate",
		20:  "mac:WindowDidBecomeKey",
		21:  "mac:WindowDidBecomeMain",
		22:  "mac:WindowDidBeginSheet",
		23:  "mac:WindowDidChangeAlpha",
		24:  "mac:WindowDidChangeBackingLocation",
		25:  "mac:WindowDidChangeBackingProperties",
		26:  "mac:WindowDidChangeCollectionBehavior",
		27:  "mac:WindowDidChangeEffectiveAppearance",
		28:  "mac:WindowDidChangeOcclusionState",
		29:  "mac:WindowDidChangeOrderingMode",
		30:  "mac:WindowDidChangeScreen",
		31:  "mac:WindowDidChangeScreenParameters",
		32:  "mac:WindowDidChangeScreenProfile",
		33:  "mac:WindowDidChangeScreenSpace",
		34:  "mac:WindowDidChangeScreenSpaceProperties",
		35:  "mac:WindowDidChangeSharingType",
		36:  "mac:WindowDidChangeSpace",
		37:  "mac:WindowDidChangeSpaceOrderingMode",
		38:  "mac:WindowDidChangeTitle",
		39:  "mac:WindowDidChangeToolbar",
		40:  "mac:WindowDidChangeVisibility",
		41:  "mac:WindowDidDeminiaturize",
		42:  "mac:WindowDidEndSheet",
		43:  "mac:WindowDidEnterFullScreen",
		44:  "mac:WindowDidEnterVersionBrowser",
		45:  "mac:WindowDidExitFullScreen",
		46:  "mac:WindowDidExitVersionBrowser",
		47:  "mac:WindowDidExpose",
		48:  "mac:WindowDidFocus",
		49:  "mac:WindowDidMiniaturize",
		50:  "mac:WindowDidMove",
		51:  "mac:WindowDidOrderOffScreen",
		52:  "mac:WindowDidOrderOnScreen",
		53:  "mac:WindowDidResignKey",
		54:  "mac:WindowDidResignMain",
		55:  "mac:WindowDidResize",
		56:  "mac:WindowDidUpdate",
		57:  "mac:WindowDidUpdateAlpha",
		58:  "mac:WindowDidUpdateCollectionBehavior",
		59:  "mac:WindowDidUpdateCollectionProperties",
		60:  "mac:WindowDidUpdateShadow",
		61:  "mac:WindowDidUpdateTitle",
		62:  "mac:WindowDidUpdateToolbar",
		63:  "mac:WindowDidUpdateVisibility",
		64:  "mac:WindowWillBecomeKey",
		65:  "mac:WindowWillBecomeMain",
		66:  "mac:WindowWillBeginSheet",
		67:  "mac:WindowWillChangeOrderingMode",
		68:  "mac:WindowWillClose",
		69:  "mac:WindowWillDeminiaturize",
		70:  "mac:WindowWillEnterFullScreen",
		71:  "mac:WindowWillEnterVersionBrowser",
		72:  "mac:WindowWillExitFullScreen",
		73:  "mac:WindowWillExitVersionBrowser",
		74:  "mac:WindowWillFocus",
		75:  "mac:WindowWillMiniaturize",
		76:  "mac:WindowWillMove",
		77:  "mac:WindowWillOrderOffScreen",
		78:  "mac:WindowWillOrderOnScreen",
		79:  "mac:WindowWillResignMain",
		80:  "mac:WindowWillResize",
		81:  "mac:WindowWillUnfocus",
		82:  "mac:WindowWillUpdate",
		83:  "mac:WindowWillUpdateAlpha",
		84:  "mac:WindowWillUpdateCollectionBehavior",
		85:  "mac:WindowWillUpdateCollectionProperties",
		86:  "mac:WindowWillUpdateShadow",
		87:  "mac:WindowWillUpdateTitle",
		88:  "mac:WindowWillUpdateToolbar",
		89:  "mac:WindowWillUpdateVisibility",
		90:  "mac:WindowWillUseStandardFrame",
		91:  "mac:MenuWillOpen",
		92:  "mac:MenuDidOpen",
		93:  "mac:MenuDidClose",
		94:  "mac:MenuWillSendAction",
		95:  "mac:MenuDidSendAction",
		96:  "mac:MenuWillHighlightItem",
		97:  "mac:MenuDidHighlightItem",
		98:  "mac:MenuWillDisplayItem",
		99:  "mac:MenuDidDisplayItem",
		100: "mac:MenuWillAddItem",
		101: "mac:MenuDidAddItem",
		102: "mac:MenuWillRemoveItem",
		103: "mac:MenuDidRemoveItem",
		104: "mac:MenuWillBeginTracking",
		105: "mac:MenuDidBeginTracking",
		106: "mac:MenuWillEndTracking",
		107: "mac:MenuDidEndTracking",
		108: "mac:MenuWillUpdate",
		109: "mac:MenuDidUpdate",
		110: "mac:MenuWillPopUp",
		111: "mac:MenuDidPopUp",
		112: "mac:MenuWillSendActionToItem",
		113: "mac:MenuDidSendActionToItem",
		114: "mac:WebViewDidStartProvisionalNavigation",
		115: "mac:WebViewDidReceiveServerRedirectForProvisionalNavigation",
		116: "mac:WebViewDidFinishNavigation",
		117: "mac:WebViewDidCommitNavigation",
		118: "mac:WindowFileDraggingEntered",
		119: "mac:WindowFileDraggingPerformed",
		120: "mac:WindowFileDraggingExited",
	},
	"linux": {
		0: "linux:SystemThemeChanged",
	},
	"common": {
		4096: "common:ApplicationStarted",
		4097: "common:WindowMaximise",
		4098: "common:WindowUnMaximise",
		4099: "common:WindowFullscreen",
		4100: "common:WindowUnFullscreen",
		4101: "common:WindowRestore",
		4102: "common:WindowMinimise",
		4103: "common:WindowUnMinimise",
		4104: "common:WindowClosing",
		4105: "common:WindowZoom",
		4106: "common:WindowZoomIn",
		4107: "common:WindowZoomOut",
		4108: "common:WindowZoomReset",
		4109: "common:WindowFocus",
		4110: "common:WindowLostFocus",
		4111: "common:WindowShow",
		4112: "common:WindowHide",
		4113: "common:WindowDPIChanged",
		4114: "common:ThemeChanged",
	},
}
