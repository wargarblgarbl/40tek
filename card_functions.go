package main

import ()

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
