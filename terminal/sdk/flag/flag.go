package flag

type Flag struct {
	Long  string
	Short string

	Value *string
}

func Register(flags []Flag) *Set {
	return &Set{
		Flags: flags,
	}
}
