package enemy

type Enemy03 struct{}

func (Enemy03) GetSpritesheet() (string, int, int) {
	return "textures/enemy_03.png", 120, 128
}

func (Enemy03) GetShieldLevelMax() int {
	return 0
}
