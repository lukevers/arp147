package enemy

type Enemy04 struct{}

func (Enemy04) GetSpritesheet() (string, int, int) {
	return "textures/enemy_04.png", 96, 132
}

func (Enemy04) GetShieldLevelMax() int {
	return 0
}
