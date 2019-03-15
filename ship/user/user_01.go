package user

type User01 struct{}

func (User01) GetSpritesheet() (string, int, int) {
	return "textures/user_01.png", 171, 171
}

func (User01) GetShieldLevelMax() int {
	return 3
}
