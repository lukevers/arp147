package enemy

type Enemy02 struct{}

func (Enemy02) GetSpritesheet() (string, int, int) {
	return "textures/enemy_02.png", 172, 226
}

func (Enemy02) GetShieldLevelMax() int {
	return 0
}
