package main

type Direction string

const (
	North Direction = "North"
	East  Direction = "East"
	South Direction = "South"
	West  Direction = "West"
)

type Boxe string

const (
	Nothing      Boxe = "⬛️"
	Grass        Boxe = "🟩"
	Water        Boxe = "🟦"
	Fire         Boxe = "🔥"
	Hamburger    Boxe = "🍔"
	PinkFlower   Boxe = "🌸"
	JackOLantern Boxe = "🎃"
)