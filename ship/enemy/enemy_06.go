package enemy

type Enemy06 struct{}

func (Enemy06) GetSpritesheet() (string, int, int) {
	return "textures/enemy_06.png", 184, 192
}

func (Enemy06) GetShieldLevelMax() int {
	return 0
}
