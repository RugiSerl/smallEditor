package input

import (
	"github.com/RugiSerl/smallEditor/app/graphic"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Key int32

const (
	KeyNull Key = 0

	// Keyboard Function Keys
	KeySpace        Key = 32
	KeyEscape       Key = 256
	KeyEnter        Key = 257
	KeyTab          Key = 258
	KeyBackspace    Key = 259
	KeyInsert       Key = 260
	KeyDelete       Key = 261
	KeyRight        Key = 262
	KeyLeft         Key = 263
	KeyDown         Key = 264
	KeyUp           Key = 265
	KeyPageUp       Key = 266
	KeyPageDown     Key = 267
	KeyHome         Key = 268
	KeyEnd          Key = 269
	KeyCapsLock     Key = 280
	KeyScrollLock   Key = 281
	KeyNumLock      Key = 282
	KeyPrintScreen  Key = 283
	KeyPause        Key = 284
	KeyF1           Key = 290
	KeyF2           Key = 291
	KeyF3           Key = 292
	KeyF4           Key = 293
	KeyF5           Key = 294
	KeyF6           Key = 295
	KeyF7           Key = 296
	KeyF8           Key = 297
	KeyF9           Key = 298
	KeyF10          Key = 299
	KeyF11          Key = 300
	KeyF12          Key = 301
	KeyLeftShift    Key = 340
	KeyLeftControl  Key = 341
	KeyLeftAlt      Key = 342
	KeyLeftSuper    Key = 343
	KeyRightShift   Key = 344
	KeyRightControl Key = 345
	KeyRightAlt     Key = 346
	KeyRightSuper   Key = 347
	KeyKbMenu       Key = 348
	KeyLeftBracket  Key = 91
	KeyBackSlash    Key = 92
	KeyRightBracket Key = 93
	KeyGrave        Key = 96

	// Keyboard Number Pad Keys
	KeyKp0        Key = 320
	KeyKp1        Key = 321
	KeyKp2        Key = 322
	KeyKp3        Key = 323
	KeyKp4        Key = 324
	KeyKp5        Key = 325
	KeyKp6        Key = 326
	KeyKp7        Key = 327
	KeyKp8        Key = 328
	KeyKp9        Key = 329
	KeyKpDecimal  Key = 330
	KeyKpDivide   Key = 331
	KeyKpMultiply Key = 332
	KeyKpSubtract Key = 333
	KeyKpAdd      Key = 334
	KeyKpEnter    Key = 335
	KeyKpEqual    Key = 336

	// Keyboard Alpha Numeric Keys
	KeyApostrophe Key = 39
	KeyComma      Key = 44
	KeyMinus      Key = 45
	KeyPeriod     Key = 46
	KeySlash      Key = 47
	KeyZero       Key = 48
	KeyOne        Key = 49
	KeyTwo        Key = 50
	KeyThree      Key = 51
	KeyFour       Key = 52
	KeyFive       Key = 53
	KeySix        Key = 54
	KeySeven      Key = 55
	KeyEight      Key = 56
	KeyNine       Key = 57
	KeySemicolon  Key = 59
	KeyEqual      Key = 61
	KeyA          Key = 65
	KeyB          Key = 66
	KeyC          Key = 67
	KeyD          Key = 68
	KeyE          Key = 69
	KeyF          Key = 70
	KeyG          Key = 71
	KeyH          Key = 72
	KeyI          Key = 73
	KeyJ          Key = 74
	KeyK          Key = 75
	KeyL          Key = 76
	KeyM          Key = 77
	KeyN          Key = 78
	KeyO          Key = 79
	KeyP          Key = 80
	KeyQ          Key = 81
	KeyR          Key = 82
	KeyS          Key = 83
	KeyT          Key = 84
	KeyU          Key = 85
	KeyV          Key = 86
	KeyW          Key = 87
	KeyX          Key = 88
	KeyY          Key = 89
	KeyZ          Key = 90

	// Android keys
	KeyBack       Key = 4
	KeyMenu       Key = 82
	KeyVolumeUp   Key = 24
	KeyVolumeDown Key = 25
)

const (
	FULL_AUTO_TRIGGER     = 0.5  // Amount of time before the mode switches to "full auto"
	KEY_COOLDOWN_DURATION = 0.03 // Amount of time to wait when a key is down between each action it triggers in full auto
)

var (
	keyCoolDowns map[Key]float64 = make(map[Key]float64) // Used to regulate the amount of key per second
)

func IsKeyPressed(key Key) bool {
	return rl.IsKeyPressed(int32(key))
}

func IsKeyDown(key Key) bool {
	return rl.IsKeyDown(int32(key))
}

// Must be called ONCE every frame
func UpdateKeyCoolDown() {
	for key := range keyCoolDowns {
		keyCoolDowns[key] += graphic.GetDeltaTime()
	}
}

func IsKeyDownUsingCoolDown(key Key) bool {
	if IsKeyDown(key) {
		if cooldown, ok := keyCoolDowns[key]; ok {
			if cooldown > KEY_COOLDOWN_DURATION || IsKeyPressed(key) {
				resetCoolDown(key)
				return true
			}

		} else { // Register the new key
			resetCoolDown(key)
			return true
		}
	}
	return false
}
func resetCoolDown(key Key) {
	if IsKeyPressed(key) { // Special case -> does not go full auto if the key was just pressed
		keyCoolDowns[key] = -FULL_AUTO_TRIGGER
	} else {
		keyCoolDowns[key] = 0
	}

}

// Get the key pressed in the last frame as a string
func GetKeysPressed() string {
	key := rl.GetCharPressed()
	var word []rune
	for key > 0 {
		word = append(word, key)
		key = rl.GetCharPressed()
	}

	return string(word)
}
