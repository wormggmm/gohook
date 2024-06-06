package hook

import "runtime"

var (
	raw2key = map[uint16]string{ // https://github.com/wesbos/keycodes
		0:     "error",
		3:     "break",
		8:     "backspace",
		9:     "tab",
		12:    "clear",
		13:    "enter",
		16:    "shift",
		17:    "ctrl",
		18:    "alt",
		19:    "pause/break",
		20:    "caps lock",
		21:    "hangul",
		25:    "hanja",
		27:    "escape",
		28:    "conversion",
		29:    "non-conversion",
		32:    "spacebar",
		33:    "page up",
		34:    "page down",
		35:    "end",
		36:    "home",
		37:    "left arrow",
		38:    "up arrow",
		39:    "right arrow",
		40:    "down arrow",
		41:    "select",
		42:    "print",
		43:    "execute",
		44:    "Print Screen",
		45:    "insert",
		46:    "delete",
		47:    "help",
		48:    "0",
		49:    "1",
		50:    "2",
		51:    "3",
		52:    "4",
		53:    "5",
		54:    "6",
		55:    "7",
		56:    "8",
		57:    "9",
		58:    ":",
		59:    ";",
		60:    "<",
		61:    "=",
		63:    "ß",
		64:    "@",
		65:    "a",
		66:    "b",
		67:    "c",
		68:    "d",
		69:    "e",
		70:    "f",
		71:    "g",
		72:    "h",
		73:    "i",
		74:    "j",
		75:    "k",
		76:    "l",
		77:    "m",
		78:    "n",
		79:    "o",
		80:    "p",
		81:    "q",
		82:    "r",
		83:    "s",
		84:    "t",
		85:    "u",
		86:    "v",
		87:    "w",
		88:    "x",
		89:    "y",
		90:    "z",
		91:    "l-super",
		92:    "r-super",
		93:    "apps",
		95:    "sleep",
		96:    "numpad 0",
		97:    "numpad 1",
		98:    "numpad 2",
		99:    "numpad 3",
		100:   "numpad 4",
		101:   "numpad 5",
		102:   "numpad 6",
		103:   "numpad 7",
		104:   "numpad 8",
		105:   "numpad 9",
		106:   "multiply",
		107:   "add",
		108:   "numpad period",
		109:   "subtract",
		110:   "decimal point",
		111:   "divide",
		112:   "f1",
		113:   "f2",
		114:   "f3",
		115:   "f4",
		116:   "f5",
		117:   "f6",
		118:   "f7",
		119:   "f8",
		120:   "f9",
		121:   "f10",
		122:   "f11",
		123:   "f12",
		124:   "f13",
		125:   "f14",
		126:   "f15",
		127:   "f16",
		128:   "f17",
		129:   "f18",
		130:   "f19",
		131:   "f20",
		132:   "f21",
		133:   "f22",
		134:   "f23",
		135:   "f24",
		144:   "num lock",
		145:   "scroll lock",
		160:   "^",
		161:   "!",
		162:   "؛",
		163:   "#",
		164:   "$",
		165:   "ù",
		166:   "page backward",
		167:   "page forward",
		168:   "refresh",
		169:   "closing paren (AZERTY)",
		170:   "*",
		171:   "~ + * key",
		172:   "home key",
		173:   "minus (firefox), mute/unmute",
		174:   "decrease volume level",
		175:   "increase volume level",
		176:   "next",
		177:   "previous",
		178:   "stop",
		179:   "play/pause",
		180:   "e-mail",
		181:   "mute/unmute (firefox)",
		182:   "decrease volume level (firefox)",
		183:   "increase volume level (firefox)",
		186:   "semi-colon / ñ",
		187:   "equal sign",
		188:   "comma",
		189:   "dash",
		190:   "period",
		191:   "forward slash / ç",
		192:   "grave accent / ñ / æ / ö",
		193:   "?, / or °",
		194:   "numpad period (chrome)",
		219:   "open bracket",
		220:   "back slash",
		221:   "close bracket / å",
		222:   "single quote / ø / ä",
		223:   "`",
		224:   "left or right ⌘ key (firefox)",
		225:   "altgr",
		226:   "< /git >, left back slash",
		230:   "GNOME Compose Key",
		231:   "ç",
		233:   "XF86Forward",
		234:   "XF86Back",
		235:   "non-conversion",
		240:   "alphanumeric",
		242:   "hiragana/katakana",
		243:   "half-width/full-width",
		244:   "kanji",
		251:   "unlock trackpad (Chrome/Edge)",
		255:   "toggle touchpad",
		65517: "hyper",
	}

	keytoraw = map[string]uint16{
		"error":                           0,
		"break":                           3,
		"backspace":                       8,
		"tab":                             9,
		"clear":                           12,
		"enter":                           13,
		"shift":                           16,
		"ctrl":                            17,
		"alt":                             18,
		"pause/break":                     19,
		"caps lock":                       20,
		"hangul":                          21,
		"hanja":                           25,
		"escape":                          27,
		"conversion":                      28,
		"non-conversion":                  29,
		"spacebar":                        32,
		"page up":                         33,
		"page down":                       34,
		"end":                             35,
		"home":                            36,
		"left arrow":                      37,
		"up arrow":                        38,
		"right arrow":                     39,
		"down arrow":                      40,
		"select":                          41,
		"print":                           42,
		"execute":                         43,
		"Print Screen":                    44,
		"insert":                          45,
		"delete":                          46,
		"help":                            47,
		"0":                               48,
		"1":                               49,
		"2":                               50,
		"3":                               51,
		"4":                               52,
		"5":                               53,
		"6":                               54,
		"7":                               55,
		"8":                               56,
		"9":                               57,
		":":                               58,
		";":                               59,
		"<":                               60,
		"=":                               61,
		"ß":                               63,
		"@":                               64,
		"a":                               65,
		"b":                               66,
		"c":                               67,
		"d":                               68,
		"e":                               69,
		"f":                               70,
		"g":                               71,
		"h":                               72,
		"i":                               73,
		"j":                               74,
		"k":                               75,
		"l":                               76,
		"m":                               77,
		"n":                               78,
		"o":                               79,
		"p":                               80,
		"q":                               81,
		"r":                               82,
		"s":                               83,
		"t":                               84,
		"u":                               85,
		"v":                               86,
		"w":                               87,
		"x":                               88,
		"y":                               89,
		"z":                               90,
		"l-super":                         91,
		"r-super":                         92,
		"apps":                            93,
		"sleep":                           95,
		"numpad 0":                        96,
		"numpad 1":                        97,
		"numpad 2":                        98,
		"numpad 3":                        99,
		"numpad 4":                        100,
		"numpad 5":                        101,
		"numpad 6":                        102,
		"numpad 7":                        103,
		"numpad 8":                        104,
		"numpad 9":                        105,
		"multiply":                        106,
		"add":                             107,
		"numpad period":                   108,
		"subtract":                        109,
		"decimal point":                   110,
		"divide":                          111,
		"f1":                              112,
		"f2":                              113,
		"f3":                              114,
		"f4":                              115,
		"f5":                              116,
		"f6":                              117,
		"f7":                              118,
		"f8":                              119,
		"f9":                              120,
		"f10":                             121,
		"f11":                             122,
		"f12":                             123,
		"f13":                             124,
		"f14":                             125,
		"f15":                             126,
		"f16":                             127,
		"f17":                             128,
		"f18":                             129,
		"f19":                             130,
		"f20":                             131,
		"f21":                             132,
		"f22":                             133,
		"f23":                             134,
		"f24":                             135,
		"num lock":                        144,
		"scroll lock":                     145,
		"^":                               160,
		"!":                               161,
		"؛":                               162,
		"#":                               163,
		"$":                               164,
		"ù":                               165,
		"page backward":                   166,
		"page forward":                    167,
		"refresh":                         168,
		"closing paren (AZERTY)":          169,
		"*":                               170,
		"~ + * key":                       171,
		"home key":                        172,
		"minus (firefox), mute/unmute":    173,
		"decrease volume level":           174,
		"increase volume level":           175,
		"next":                            176,
		"previous":                        177,
		"stop":                            178,
		"play/pause":                      179,
		"e-mail":                          180,
		"mute/unmute (firefox)":           181,
		"decrease volume level (firefox)": 182,
		"increase volume level (firefox)": 183,
		"semi-colon / ñ":                  186,
		"equal sign":                      187,
		"comma":                           188,
		"dash":                            189,
		"period":                          190,
		"forward slash / ç":               191,
		"grave accent / ñ / æ / ö":        192,
		"?, / or °":                       193,
		"numpad period (chrome)":          194,
		"open bracket":                    219,
		"back slash":                      220,
		"close bracket / å":               221,
		"single quote / ø / ä":            222,
		"`":                               223,
		"left or right ⌘ key (firefox)":   224,
		"altgr":                           225,
		"< /git >, left back slash":       226,
		"GNOME Compose Key":               230,
		"ç":                               231,
		"XF86Forward":                     233,
		"XF86Back":                        234,
		"alphanumeric":                    240,
		"hiragana/katakana":               242,
		"half-width/full-width":           243,
		"kanji":                           244,
		"unlock trackpad (Chrome/Edge)":   251,
		"toggle touchpad":                 255,
		"hyper":                           65517,
	}
	WinRaw2Code = map[uint16]uint16{
		160: 16, // Shift
		161: 16, // Shift Right
		162: 17, // Ctrl
		163: 17, // Ctrl Right
		164: 18, // Alt
		165: 18, // Alt Right
	}
	LinuxRaw2Code = map[uint16]uint16{
		96:    192, // Backquote
		126:   192, // Backquote / grave accent
		41:    48,  // Digit 0
		33:    49,  // Digit 1
		64:    50,  // Digit 2
		35:    51,  // Digit 3
		36:    52,  // Digit 4
		37:    53,  // Digit 5
		94:    54,  // Digit 6
		38:    55,  // Digit 7
		42:    57,  // Digit 8
		40:    57,  // Digit 9
		45:    189, // Dash / Minus
		95:    189, // Dash // Minus
		43:    61,  // Equal
		65289: 9,   // Tab
		65056: 9,   // Tab
		97:    65,  // Key A
		98:    66,  // Key B
		99:    67,  // Key C
		100:   68,  // Key D
		101:   69,  // Key E
		102:   70,  // Key F
		103:   71,  // Key G
		104:   72,  // Key H
		105:   73,  // Key I
		106:   74,  // Key J
		107:   75,  // Key K
		108:   76,  // Key L
		109:   77,  // Key M
		110:   78,  // Key N
		111:   79,  // Key O
		112:   80,  // Key P
		113:   81,  // Key Q
		114:   82,  // Key R
		115:   83,  // Key S
		116:   84,  // Key T
		117:   85,  // Key U
		118:   86,  // Key V
		119:   87,  // Key W
		120:   88,  // Key X
		121:   89,  // Key Y
		122:   90,  // Key Z
		91:    219, // Left bracket
		93:    221, // Right bracket
		123:   219, // Left bracket
		125:   221, // Right bracket
		39:    222, // Quote / Double quote
		34:    222, // Quote / Double quote
		92:    220, // Backslash / Pipe
		124:   220, // Backslash / Pipe
		44:    188, // Comma
		60:    188, // Comma / Arrow left
		46:    190, // Period
		62:    190, // Period / arrow right
		47:    191, // Forward slash
		63:    191, // Forward Slash
		65288: 8,   // backspace
		65509: 20,  // Caps Lock
		65505: 16,  // Shift
		65506: 16,  // Shift Right
		65507: 17,  // Ctrl
		65508: 17,  // Ctrl Left
		65515: 91,  // Meta / Win Key (Left)
		65513: 18,  // Alt left
		65514: 18,  // Alt right
		65293: 13,  // Enter
		65421: 13,  // Enter (Numpad)
		65307: 27,  // Escape
		65470: 112, // F1
		65471: 113, // F2
		65472: 114, // F3
		65473: 115, // F4
		65474: 116, // F5
		65475: 117, // F6
		65476: 118, // F7
		65477: 119, // F8
		65478: 120, // F9
		65479: 121, // F10
		65480: 122, // F11
		65481: 123, // F12
		65409: 112, // Fn + F1
		65297: 113, // Fn + F2
		65299: 114, // Fn + F3
		65298: 115, // Fn + F4
		65301: 116, // Fn + F5
		65302: 117, // Fn + F6
		65300: 118, // Fn + F7
		65303: 119, // Fn + F8
		65305: 120, // Fn + F9
		65304: 121, // Fn + F10
		65309: 123, // Fn + F12
		65377: 44,  // Print screen
		// 65300: 135, // Scroll Lock // Duplicate with Fn + F7
		// 65299: 19, // Pause // Duplicate with Fn + F6
		65379: 45,  // Insert
		65535: 46,  // Delete
		65360: 36,  // Home
		65367: 35,  // End
		65365: 33,  // Page Up
		65366: 34,  // Page down
		65362: 38,  // Arrow up
		65364: 40,  // Arrow Down
		65361: 37,  // Arrow left
		65363: 39,  // Arrow Right
		65407: 144, // Num Lock
		65455: 111, // Devide (Numpad)
		65450: 106, // Multiply (Numpad)
		65453: 109, // Substract (Numpad)
		65451: 107, // Add (Numpad)
		65454: 110, // Decimal point / period (Numpad)
		65456: 96,  // Numpad 0
		65457: 97,  // Numpad 1
		65458: 98,  // Numpad 2
		65459: 99,  // Numpad 3
		65460: 100, // Numpad 4
		65461: 101, // Numpad 5
		65462: 102, // Numpad 6
		65463: 103, // Numpad 7
		65464: 104, // Numpad 8
		65465: 105, // Numpad 9
		65429: 36,  // Shift + Numpad 7 = (Home)
		65430: 37,  // Shift + Numpad 4 = (Arrow left)
		65431: 38,  // Shift + Numpad 8 = (Arrow up)
		65432: 39,  // Shift + Numpad 6 = (Arrow right)
		65433: 40,  // Shift + Numpad 2 = (Arrow down)
		65434: 33,  // Shift + Numpad 9 = (Page Up)
		65435: 34,  // Shift + Numpad 3 = (Page Down)
		65436: 35,  // Shift + Numpad 4 = (End)
		65437: 12,  // Shift + Numpad 5 = (clear)
		65438: 45,  // Shift + Numpad 0 = (Insert)
	}
	DarwinRaw2Code = map[uint16]uint16{
		122: 112,
		120: 113,
		99:  114,
		118: 115,
		96:  116,
		97:  117,
		98:  118,
		100: 119,
		101: 120,
		109: 121,
		103: 122,
		111: 123,
		105: 124,
		107: 125,
		113: 126,
		106: 127,
		64:  128,
		79:  129,
		80:  130,
		90:  131,
		82:  96,
		83:  97,
		84:  98,
		85:  99,
		86:  100,
		87:  101,
		88:  102,
		89:  103,
		91:  104,
		92:  105,
		65:  110,
		67:  106,
		69:  107,
		71:  0,
		75:  111,
		76:  13,
		78:  109,
		29:  48,
		18:  49,
		19:  50,
		20:  51,
		21:  52,
		23:  53,
		22:  54,
		26:  55,
		28:  56,
		25:  57,
		48:  9,
		49:  32,
		36:  13,
		51:  8,
		117: 8,
		53:  27,
		55:  0,
		56:  16,
		57:  20,
		58:  0,
		59:  0,
		60:  16,
		61:  0,
		62:  17,
		63:  0,
		0:   65,
		11:  66,
		8:   67,
		2:   68,
		14:  69,
		3:   70,
		5:   71,
		4:   72,
		34:  73,
		38:  74,
		40:  75,
		37:  76,
		46:  77,
		45:  78,
		31:  79,
		35:  80,
		12:  81,
		15:  82,
		1:   83,
		17:  84,
		32:  85,
		9:   86,
		13:  87,
		7:   88,
		16:  89,
		6:   90,
		50:  0,
		27:  0,
		24:  0,
		33:  0,
		30:  0,
		42:  0,
		41:  0,
		39:  0,
		43:  0,
		47:  0,
		44:  0,
		114: 0,
		81:  0,
		123: 37,
		124: 39,
		125: 40,
		126: 38,
		115: 36,
		116: 33,
		119: 35,
		121: 34,
	}
)

func rawToKey(rawCode uint16) string {
	switch runtime.GOOS {
	case "darwin":
		return raw2key[DarwinRaw2Code[rawCode]]
	case "linux":
		return raw2key[LinuxRaw2Code[rawCode]]
	case "windows":
		return raw2key[WinRaw2Code[rawCode]]
	default:
		return raw2key[uint16(rawCode)]
	}
}
