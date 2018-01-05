package main

import ()

func (f *Card) SetPos(pos int) {
	f.Position = pos
}

func (f *Card) Tap() {
	f.Tapped = true
}

func (f *Card) Untap() {
	f.Tapped = false
}

func (f *Card) FlipUp() {
	f.Faceup = true
}

func (f *Card) FlipDown() {
	f.Faceup = false
}

func (f *Card) Obfuscate() Card {
	f = &Card{}
	return *f
}

func (f *Card) Hit(dmg int) {
	f.Damage = f.Damage + dmg
}

func (f *Card) Heal(dmg int) {
	if dmg >= f.Damage {
		f.Damage = 0
	} else {
	f.Damage = f.Damage - dmg
}
}

func (f *Card) Pump(pump int){
	f.Buff = pump
}
