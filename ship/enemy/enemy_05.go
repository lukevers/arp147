package enemy

type Enemy05 struct{}

func (Enemy05) GetSpritesheet() (string, int, int) {
	return "textures/enemy_05.png", 96, 97
}

func (Enemy05) GetShieldLevelMax() int {
	return 0
}
