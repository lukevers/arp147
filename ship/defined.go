package ship

type Defined interface {
	GetSpritesheet() (string, int, int)
	GetShieldLevelMax() int
}
