package enemy

type Enemy01 struct{}

func (Enemy01) GetSpritesheet() (string, int, int) {
	return "textures/enemy_01.png", 120, 160
}

func (Enemy01) GetShieldLevelMax() int {
	return 0
}
