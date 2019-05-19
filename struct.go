package main

type Character struct {
	name string

	// Tank
	Paladin    int
	Warrior    int
	DarkKnight int
	// Healer
	WhiteMage   int
	Scholar     int
	Astrologian int

	// DPS
	Monk      int
	Dragoon   int
	Ninja     int
	Samurai   int
	Bard      int
	Machinist int
	BlackMage int
	Summoner  int
	RedMage   int
	BlueMage  int

	// Crafter
	Carpenter     int
	Blacksmith    int
	Armorer       int
	Goldsmith     int
	Leatherworker int
	Weaver        int
	Alchemist     int
	Culinarian    int

	// Gatherer
	Miner    int
	Botanist int
	Fisher   int
}
