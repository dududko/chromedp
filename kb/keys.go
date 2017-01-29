package kb

// DOM keys.
const (
	Backspace            = "\b"
	Tab                  = "\t"
	Enter                = "\r"
	Escape               = "\u001b"
	Quote                = "'"
	Backslash            = "\\"
	Delete               = "\u007f"
	Alt                  = "\u0102"
	CapsLock             = "\u0104"
	Control              = "\u0105"
	Fn                   = "\u0106"
	FnLock               = "\u0107"
	Hyper                = "\u0108"
	Meta                 = "\u0109"
	NumLock              = "\u010a"
	ScrollLock           = "\u010c"
	Shift                = "\u010d"
	Super                = "\u010e"
	ArrowDown            = "\u0301"
	ArrowLeft            = "\u0302"
	ArrowRight           = "\u0303"
	ArrowUp              = "\u0304"
	End                  = "\u0305"
	Home                 = "\u0306"
	PageDown             = "\u0307"
	PageUp               = "\u0308"
	Clear                = "\u0401"
	Copy                 = "\u0402"
	Cut                  = "\u0404"
	Insert               = "\u0407"
	Paste                = "\u0408"
	Redo                 = "\u0409"
	Undo                 = "\u040a"
	Again                = "\u0502"
	Cancel               = "\u0504"
	ContextMenu          = "\u0505"
	Find                 = "\u0507"
	Help                 = "\u0508"
	Pause                = "\u0509"
	Props                = "\u050b"
	Select               = "\u050c"
	ZoomIn               = "\u050d"
	ZoomOut              = "\u050e"
	BrightnessDown       = "\u0601"
	BrightnessUp         = "\u0602"
	Eject                = "\u0604"
	LogOff               = "\u0605"
	Power                = "\u0606"
	PrintScreen          = "\u0608"
	WakeUp               = "\u060b"
	Convert              = "\u0705"
	NonConvert           = "\u070d"
	HangulMode           = "\u0711"
	HanjaMode            = "\u0712"
	Hiragana             = "\u0716"
	KanaMode             = "\u0718"
	Katakana             = "\u071a"
	ZenkakuHankaku       = "\u071d"
	F1                   = "\u0801"
	F2                   = "\u0802"
	F3                   = "\u0803"
	F4                   = "\u0804"
	F5                   = "\u0805"
	F6                   = "\u0806"
	F7                   = "\u0807"
	F8                   = "\u0808"
	F9                   = "\u0809"
	F10                  = "\u080a"
	F11                  = "\u080b"
	F12                  = "\u080c"
	F13                  = "\u080d"
	F14                  = "\u080e"
	F15                  = "\u080f"
	F16                  = "\u0810"
	F17                  = "\u0811"
	F18                  = "\u0812"
	F19                  = "\u0813"
	F20                  = "\u0814"
	F21                  = "\u0815"
	F22                  = "\u0816"
	F23                  = "\u0817"
	F24                  = "\u0818"
	Close                = "\u0a01"
	MailForward          = "\u0a02"
	MailReply            = "\u0a03"
	MailSend             = "\u0a04"
	MediaPlayPause       = "\u0a05"
	MediaStop            = "\u0a07"
	MediaTrackNext       = "\u0a08"
	MediaTrackPrevious   = "\u0a09"
	New                  = "\u0a0a"
	Open                 = "\u0a0b"
	Print                = "\u0a0c"
	Save                 = "\u0a0d"
	SpellCheck           = "\u0a0e"
	AudioVolumeDown      = "\u0a0f"
	AudioVolumeUp        = "\u0a10"
	AudioVolumeMute      = "\u0a11"
	LaunchCalculator     = "\u0b01"
	LaunchCalendar       = "\u0b02"
	LaunchMail           = "\u0b03"
	LaunchMediaPlayer    = "\u0b04"
	LaunchMusicPlayer    = "\u0b05"
	LaunchMyComputer     = "\u0b06"
	LaunchScreenSaver    = "\u0b07"
	LaunchSpreadsheet    = "\u0b08"
	LaunchWebBrowser     = "\u0b09"
	LaunchContacts       = "\u0b0c"
	LaunchPhone          = "\u0b0d"
	BrowserBack          = "\u0c01"
	BrowserFavorites     = "\u0c02"
	BrowserForward       = "\u0c03"
	BrowserHome          = "\u0c04"
	BrowserRefresh       = "\u0c05"
	BrowserSearch        = "\u0c06"
	BrowserStop          = "\u0c07"
	ChannelDown          = "\u0d0a"
	ChannelUp            = "\u0d0b"
	ClosedCaptionToggle  = "\u0d12"
	Exit                 = "\u0d15"
	Guide                = "\u0d22"
	Info                 = "\u0d25"
	MediaFastForward     = "\u0d2c"
	MediaLast            = "\u0d2d"
	MediaPlay            = "\u0d2f"
	MediaRecord          = "\u0d30"
	MediaRewind          = "\u0d31"
	Settings             = "\u0d43"
	ZoomToggle           = "\u0d4e"
	AudioBassBoostToggle = "\u0e02"
	SpeechInputToggle    = "\u0f02"
	AppSwitch            = "\u1001"
)

// Keys is the map of unicode characters to their DOM key data.
var Keys = map[rune]*Key{
	'\b':     &Key{"Backspace", "Backspace", "", "", 8, 8, false, false},
	'\t':     &Key{"Tab", "Tab", "", "", 9, 9, false, false},
	'\r':     &Key{"Enter", "Enter", "\r", "\r", 13, 13, false, true},
	'\u001b': &Key{"Escape", "Escape", "", "", 27, 27, false, false},
	' ':      &Key{"Space", " ", " ", " ", 32, 32, false, true},
	'!':      &Key{"Digit1", "!", "!", "1", 49, 49, true, true},
	'"':      &Key{"Quote", "\"", "\"", "'", 222, 222, true, true},
	'#':      &Key{"Digit3", "#", "#", "3", 51, 51, true, true},
	'$':      &Key{"Digit4", "$", "$", "4", 52, 52, true, true},
	'%':      &Key{"Digit5", "%", "%", "5", 53, 53, true, true},
	'&':      &Key{"Digit7", "&", "&", "7", 55, 55, true, true},
	'\'':     &Key{"Quote", "'", "'", "'", 222, 222, false, true},
	'(':      &Key{"Digit9", "(", "(", "9", 57, 57, true, true},
	')':      &Key{"Digit0", ")", ")", "0", 48, 48, true, true},
	'*':      &Key{"Digit8", "*", "*", "8", 56, 56, true, true},
	'+':      &Key{"Equal", "+", "+", "=", 187, 187, true, true},
	',':      &Key{"Comma", ",", ",", ",", 188, 188, false, true},
	'-':      &Key{"Minus", "-", "-", "-", 189, 189, false, true},
	'.':      &Key{"Period", ".", ".", ".", 190, 190, false, true},
	'/':      &Key{"Slash", "/", "/", "/", 191, 191, false, true},
	'0':      &Key{"Digit0", "0", "0", "0", 48, 48, false, true},
	'1':      &Key{"Digit1", "1", "1", "1", 49, 49, false, true},
	'2':      &Key{"Digit2", "2", "2", "2", 50, 50, false, true},
	'3':      &Key{"Digit3", "3", "3", "3", 51, 51, false, true},
	'4':      &Key{"Digit4", "4", "4", "4", 52, 52, false, true},
	'5':      &Key{"Digit5", "5", "5", "5", 53, 53, false, true},
	'6':      &Key{"Digit6", "6", "6", "6", 54, 54, false, true},
	'7':      &Key{"Digit7", "7", "7", "7", 55, 55, false, true},
	'8':      &Key{"Digit8", "8", "8", "8", 56, 56, false, true},
	'9':      &Key{"Digit9", "9", "9", "9", 57, 57, false, true},
	':':      &Key{"Semicolon", ":", ":", ";", 186, 186, true, true},
	';':      &Key{"Semicolon", ";", ";", ";", 186, 186, false, true},
	'<':      &Key{"Comma", "<", "<", ",", 188, 188, true, true},
	'=':      &Key{"Equal", "=", "=", "=", 187, 187, false, true},
	'>':      &Key{"Period", ">", ">", ".", 190, 190, true, true},
	'?':      &Key{"Slash", "?", "?", "/", 191, 191, true, true},
	'@':      &Key{"Digit2", "@", "@", "2", 50, 50, true, true},
	'A':      &Key{"KeyA", "A", "A", "a", 65, 65, true, true},
	'B':      &Key{"KeyB", "B", "B", "b", 66, 66, true, true},
	'C':      &Key{"KeyC", "C", "C", "c", 67, 67, true, true},
	'D':      &Key{"KeyD", "D", "D", "d", 68, 68, true, true},
	'E':      &Key{"KeyE", "E", "E", "e", 69, 69, true, true},
	'F':      &Key{"KeyF", "F", "F", "f", 70, 70, true, true},
	'G':      &Key{"KeyG", "G", "G", "g", 71, 71, true, true},
	'H':      &Key{"KeyH", "H", "H", "h", 72, 72, true, true},
	'I':      &Key{"KeyI", "I", "I", "i", 73, 73, true, true},
	'J':      &Key{"KeyJ", "J", "J", "j", 74, 74, true, true},
	'K':      &Key{"KeyK", "K", "K", "k", 75, 75, true, true},
	'L':      &Key{"KeyL", "L", "L", "l", 76, 76, true, true},
	'M':      &Key{"KeyM", "M", "M", "m", 77, 77, true, true},
	'N':      &Key{"KeyN", "N", "N", "n", 78, 78, true, true},
	'O':      &Key{"KeyO", "O", "O", "o", 79, 79, true, true},
	'P':      &Key{"KeyP", "P", "P", "p", 80, 80, true, true},
	'Q':      &Key{"KeyQ", "Q", "Q", "q", 81, 81, true, true},
	'R':      &Key{"KeyR", "R", "R", "r", 82, 82, true, true},
	'S':      &Key{"KeyS", "S", "S", "s", 83, 83, true, true},
	'T':      &Key{"KeyT", "T", "T", "t", 84, 84, true, true},
	'U':      &Key{"KeyU", "U", "U", "u", 85, 85, true, true},
	'V':      &Key{"KeyV", "V", "V", "v", 86, 86, true, true},
	'W':      &Key{"KeyW", "W", "W", "w", 87, 87, true, true},
	'X':      &Key{"KeyX", "X", "X", "x", 88, 88, true, true},
	'Y':      &Key{"KeyY", "Y", "Y", "y", 89, 89, true, true},
	'Z':      &Key{"KeyZ", "Z", "Z", "z", 90, 90, true, true},
	'[':      &Key{"BracketLeft", "[", "[", "[", 219, 219, false, true},
	'\\':     &Key{"Backslash", "\\", "\\", "\\", 220, 220, false, true},
	']':      &Key{"BracketRight", "]", "]", "]", 221, 221, false, true},
	'^':      &Key{"Digit6", "^", "^", "6", 54, 54, true, true},
	'_':      &Key{"Minus", "_", "_", "-", 189, 189, true, true},
	'`':      &Key{"Backquote", "`", "`", "`", 192, 192, false, true},
	'a':      &Key{"KeyA", "a", "a", "a", 65, 65, false, true},
	'b':      &Key{"KeyB", "b", "b", "b", 66, 66, false, true},
	'c':      &Key{"KeyC", "c", "c", "c", 67, 67, false, true},
	'd':      &Key{"KeyD", "d", "d", "d", 68, 68, false, true},
	'e':      &Key{"KeyE", "e", "e", "e", 69, 69, false, true},
	'f':      &Key{"KeyF", "f", "f", "f", 70, 70, false, true},
	'g':      &Key{"KeyG", "g", "g", "g", 71, 71, false, true},
	'h':      &Key{"KeyH", "h", "h", "h", 72, 72, false, true},
	'i':      &Key{"KeyI", "i", "i", "i", 73, 73, false, true},
	'j':      &Key{"KeyJ", "j", "j", "j", 74, 74, false, true},
	'k':      &Key{"KeyK", "k", "k", "k", 75, 75, false, true},
	'l':      &Key{"KeyL", "l", "l", "l", 76, 76, false, true},
	'm':      &Key{"KeyM", "m", "m", "m", 77, 77, false, true},
	'n':      &Key{"KeyN", "n", "n", "n", 78, 78, false, true},
	'o':      &Key{"KeyO", "o", "o", "o", 79, 79, false, true},
	'p':      &Key{"KeyP", "p", "p", "p", 80, 80, false, true},
	'q':      &Key{"KeyQ", "q", "q", "q", 81, 81, false, true},
	'r':      &Key{"KeyR", "r", "r", "r", 82, 82, false, true},
	's':      &Key{"KeyS", "s", "s", "s", 83, 83, false, true},
	't':      &Key{"KeyT", "t", "t", "t", 84, 84, false, true},
	'u':      &Key{"KeyU", "u", "u", "u", 85, 85, false, true},
	'v':      &Key{"KeyV", "v", "v", "v", 86, 86, false, true},
	'w':      &Key{"KeyW", "w", "w", "w", 87, 87, false, true},
	'x':      &Key{"KeyX", "x", "x", "x", 88, 88, false, true},
	'y':      &Key{"KeyY", "y", "y", "y", 89, 89, false, true},
	'z':      &Key{"KeyZ", "z", "z", "z", 90, 90, false, true},
	'{':      &Key{"BracketLeft", "{", "{", "[", 219, 219, true, true},
	'|':      &Key{"Backslash", "|", "|", "\\", 220, 220, true, true},
	'}':      &Key{"BracketRight", "}", "}", "]", 221, 221, true, true},
	'~':      &Key{"Backquote", "~", "~", "`", 192, 192, true, true},
	'\u007f': &Key{"Delete", "Delete", "", "", 46, 46, false, false},
	'¥':      &Key{"IntlYen", "¥", "¥", "¥", 220, 220, false, true},
	'\u0102': &Key{"AltLeft", "Alt", "", "", 164, 164, false, false},
	'\u0104': &Key{"CapsLock", "CapsLock", "", "", 20, 20, false, false},
	'\u0105': &Key{"ControlLeft", "Control", "", "", 162, 162, false, false},
	'\u0106': &Key{"Fn", "Fn", "", "", 0, 0, false, false},
	'\u0107': &Key{"FnLock", "FnLock", "", "", 0, 0, false, false},
	'\u0108': &Key{"Hyper", "Hyper", "", "", 0, 0, false, false},
	'\u0109': &Key{"MetaLeft", "Meta", "", "", 91, 91, false, false},
	'\u010a': &Key{"NumLock", "NumLock", "", "", 144, 144, false, false},
	'\u010c': &Key{"ScrollLock", "ScrollLock", "", "", 145, 145, false, false},
	'\u010d': &Key{"ShiftLeft", "Shift", "", "", 160, 160, false, false},
	'\u010e': &Key{"Super", "Super", "", "", 0, 0, false, false},
	'\u0301': &Key{"ArrowDown", "ArrowDown", "", "", 40, 40, false, false},
	'\u0302': &Key{"ArrowLeft", "ArrowLeft", "", "", 37, 37, false, false},
	'\u0303': &Key{"ArrowRight", "ArrowRight", "", "", 39, 39, false, false},
	'\u0304': &Key{"ArrowUp", "ArrowUp", "", "", 38, 38, false, false},
	'\u0305': &Key{"End", "End", "", "", 35, 35, false, false},
	'\u0306': &Key{"Home", "Home", "", "", 36, 36, false, false},
	'\u0307': &Key{"PageDown", "PageDown", "", "", 34, 34, false, false},
	'\u0308': &Key{"PageUp", "PageUp", "", "", 33, 33, false, false},
	'\u0401': &Key{"NumpadClear", "Clear", "", "", 12, 12, false, false},
	'\u0402': &Key{"Copy", "Copy", "", "", 0, 0, false, false},
	'\u0404': &Key{"Cut", "Cut", "", "", 0, 0, false, false},
	'\u0407': &Key{"Insert", "Insert", "", "", 45, 45, false, false},
	'\u0408': &Key{"Paste", "Paste", "", "", 0, 0, false, false},
	'\u0409': &Key{"Redo", "Redo", "", "", 0, 0, false, false},
	'\u040a': &Key{"Undo", "Undo", "", "", 0, 0, false, false},
	'\u0502': &Key{"Again", "Again", "", "", 0, 0, false, false},
	'\u0504': &Key{"Abort", "Cancel", "", "", 0, 0, false, false},
	'\u0505': &Key{"ContextMenu", "ContextMenu", "", "", 93, 93, false, false},
	'\u0507': &Key{"Find", "Find", "", "", 0, 0, false, false},
	'\u0508': &Key{"Help", "Help", "", "", 47, 47, false, false},
	'\u0509': &Key{"Pause", "Pause", "", "", 19, 19, false, false},
	'\u050b': &Key{"Props", "Props", "", "", 0, 0, false, false},
	'\u050c': &Key{"Select", "Select", "", "", 41, 41, false, false},
	'\u050d': &Key{"ZoomIn", "ZoomIn", "", "", 0, 0, false, false},
	'\u050e': &Key{"ZoomOut", "ZoomOut", "", "", 0, 0, false, false},
	'\u0601': &Key{"BrightnessDown", "BrightnessDown", "", "", 216, 0, false, false},
	'\u0602': &Key{"BrightnessUp", "BrightnessUp", "", "", 217, 0, false, false},
	'\u0604': &Key{"Eject", "Eject", "", "", 0, 0, false, false},
	'\u0605': &Key{"LogOff", "LogOff", "", "", 0, 0, false, false},
	'\u0606': &Key{"Power", "Power", "", "", 152, 0, false, false},
	'\u0608': &Key{"PrintScreen", "PrintScreen", "", "", 44, 44, false, false},
	'\u060b': &Key{"WakeUp", "WakeUp", "", "", 0, 0, false, false},
	'\u0705': &Key{"Convert", "Convert", "", "", 28, 28, false, false},
	'\u070d': &Key{"NonConvert", "NonConvert", "", "", 29, 29, false, false},
	'\u0711': &Key{"Lang1", "HangulMode", "", "", 21, 21, false, false},
	'\u0712': &Key{"Lang2", "HanjaMode", "", "", 25, 25, false, false},
	'\u0716': &Key{"Lang4", "Hiragana", "", "", 0, 0, false, false},
	'\u0718': &Key{"KanaMode", "KanaMode", "", "", 21, 21, false, false},
	'\u071a': &Key{"Lang3", "Katakana", "", "", 0, 0, false, false},
	'\u071d': &Key{"Lang5", "ZenkakuHankaku", "", "", 0, 0, false, false},
	'\u0801': &Key{"F1", "F1", "", "", 112, 112, false, false},
	'\u0802': &Key{"F2", "F2", "", "", 113, 113, false, false},
	'\u0803': &Key{"F3", "F3", "", "", 114, 114, false, false},
	'\u0804': &Key{"F4", "F4", "", "", 115, 115, false, false},
	'\u0805': &Key{"F5", "F5", "", "", 116, 116, false, false},
	'\u0806': &Key{"F6", "F6", "", "", 117, 117, false, false},
	'\u0807': &Key{"F7", "F7", "", "", 118, 118, false, false},
	'\u0808': &Key{"F8", "F8", "", "", 119, 119, false, false},
	'\u0809': &Key{"F9", "F9", "", "", 120, 120, false, false},
	'\u080a': &Key{"F10", "F10", "", "", 121, 121, false, false},
	'\u080b': &Key{"F11", "F11", "", "", 122, 122, false, false},
	'\u080c': &Key{"F12", "F12", "", "", 123, 123, false, false},
	'\u080d': &Key{"F13", "F13", "", "", 124, 124, false, false},
	'\u080e': &Key{"F14", "F14", "", "", 125, 125, false, false},
	'\u080f': &Key{"F15", "F15", "", "", 126, 126, false, false},
	'\u0810': &Key{"F16", "F16", "", "", 127, 127, false, false},
	'\u0811': &Key{"F17", "F17", "", "", 128, 128, false, false},
	'\u0812': &Key{"F18", "F18", "", "", 129, 129, false, false},
	'\u0813': &Key{"F19", "F19", "", "", 130, 130, false, false},
	'\u0814': &Key{"F20", "F20", "", "", 131, 131, false, false},
	'\u0815': &Key{"F21", "F21", "", "", 132, 132, false, false},
	'\u0816': &Key{"F22", "F22", "", "", 133, 133, false, false},
	'\u0817': &Key{"F23", "F23", "", "", 134, 134, false, false},
	'\u0818': &Key{"F24", "F24", "", "", 135, 135, false, false},
	'\u0a01': &Key{"Close", "Close", "", "", 0, 0, false, false},
	'\u0a02': &Key{"MailForward", "MailForward", "", "", 0, 0, false, false},
	'\u0a03': &Key{"MailReply", "MailReply", "", "", 0, 0, false, false},
	'\u0a04': &Key{"MailSend", "MailSend", "", "", 0, 0, false, false},
	'\u0a05': &Key{"MediaPlayPause", "MediaPlayPause", "", "", 179, 179, false, false},
	'\u0a07': &Key{"MediaStop", "MediaStop", "", "", 178, 178, false, false},
	'\u0a08': &Key{"MediaTrackNext", "MediaTrackNext", "", "", 176, 176, false, false},
	'\u0a09': &Key{"MediaTrackPrevious", "MediaTrackPrevious", "", "", 177, 177, false, false},
	'\u0a0a': &Key{"New", "New", "", "", 0, 0, false, false},
	'\u0a0b': &Key{"Open", "Open", "", "", 43, 43, false, false},
	'\u0a0c': &Key{"Print", "Print", "", "", 0, 0, false, false},
	'\u0a0d': &Key{"Save", "Save", "", "", 0, 0, false, false},
	'\u0a0e': &Key{"SpellCheck", "SpellCheck", "", "", 0, 0, false, false},
	'\u0a0f': &Key{"AudioVolumeDown", "AudioVolumeDown", "", "", 174, 174, false, false},
	'\u0a10': &Key{"AudioVolumeUp", "AudioVolumeUp", "", "", 175, 175, false, false},
	'\u0a11': &Key{"AudioVolumeMute", "AudioVolumeMute", "", "", 173, 173, false, false},
	'\u0b01': &Key{"LaunchApp2", "LaunchCalculator", "", "", 183, 183, false, false},
	'\u0b02': &Key{"LaunchCalendar", "LaunchCalendar", "", "", 0, 0, false, false},
	'\u0b03': &Key{"LaunchMail", "LaunchMail", "", "", 180, 180, false, false},
	'\u0b04': &Key{"MediaSelect", "LaunchMediaPlayer", "", "", 181, 181, false, false},
	'\u0b05': &Key{"LaunchMusicPlayer", "LaunchMusicPlayer", "", "", 0, 0, false, false},
	'\u0b06': &Key{"LaunchApp1", "LaunchMyComputer", "", "", 182, 182, false, false},
	'\u0b07': &Key{"LaunchScreenSaver", "LaunchScreenSaver", "", "", 0, 0, false, false},
	'\u0b08': &Key{"LaunchSpreadsheet", "LaunchSpreadsheet", "", "", 0, 0, false, false},
	'\u0b09': &Key{"LaunchWebBrowser", "LaunchWebBrowser", "", "", 0, 0, false, false},
	'\u0b0c': &Key{"LaunchContacts", "LaunchContacts", "", "", 0, 0, false, false},
	'\u0b0d': &Key{"LaunchPhone", "LaunchPhone", "", "", 0, 0, false, false},
	'\u0c01': &Key{"BrowserBack", "BrowserBack", "", "", 166, 166, false, false},
	'\u0c02': &Key{"BrowserFavorites", "BrowserFavorites", "", "", 171, 171, false, false},
	'\u0c03': &Key{"BrowserForward", "BrowserForward", "", "", 167, 167, false, false},
	'\u0c04': &Key{"BrowserHome", "BrowserHome", "", "", 172, 172, false, false},
	'\u0c05': &Key{"BrowserRefresh", "BrowserRefresh", "", "", 168, 168, false, false},
	'\u0c06': &Key{"BrowserSearch", "BrowserSearch", "", "", 170, 170, false, false},
	'\u0c07': &Key{"BrowserStop", "BrowserStop", "", "", 169, 169, false, false},
	'\u0d0a': &Key{"ChannelDown", "ChannelDown", "", "", 0, 0, false, false},
	'\u0d0b': &Key{"ChannelUp", "ChannelUp", "", "", 0, 0, false, false},
	'\u0d12': &Key{"ClosedCaptionToggle", "ClosedCaptionToggle", "", "", 0, 0, false, false},
	'\u0d15': &Key{"Exit", "Exit", "", "", 0, 0, false, false},
	'\u0d22': &Key{"Guide", "Guide", "", "", 0, 0, false, false},
	'\u0d25': &Key{"Info", "Info", "", "", 0, 0, false, false},
	'\u0d2c': &Key{"MediaFastForward", "MediaFastForward", "", "", 0, 0, false, false},
	'\u0d2d': &Key{"MediaLast", "MediaLast", "", "", 0, 0, false, false},
	'\u0d2f': &Key{"MediaPlay", "MediaPlay", "", "", 0, 0, false, false},
	'\u0d30': &Key{"MediaRecord", "MediaRecord", "", "", 0, 0, false, false},
	'\u0d31': &Key{"MediaRewind", "MediaRewind", "", "", 0, 0, false, false},
	'\u0d43': &Key{"Settings", "Settings", "", "", 0, 0, false, false},
	'\u0d4e': &Key{"ZoomToggle", "ZoomToggle", "", "", 251, 251, false, false},
	'\u0e02': &Key{"AudioBassBoostToggle", "AudioBassBoostToggle", "", "", 0, 0, false, false},
	'\u0f02': &Key{"SpeechInputToggle", "SpeechInputToggle", "", "", 0, 0, false, false},
	'\u1001': &Key{"SelectTask", "AppSwitch", "", "", 0, 0, false, false},
}
