package main

func PlanetObfuscator(b *Playfield) (obfed *Playfield) {
	for h, _ := range b.Planets {
		if b.Planets[h].Faceup == false {
			// fmt.Println("OBFUSCATING")
			b.Planets[h] = b.Planets[h].Obfuscate()
			// fmt.Println(b.Planets[h])
		}
	}
	return b

}
