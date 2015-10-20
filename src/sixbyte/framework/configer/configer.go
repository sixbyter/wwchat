package configer

type Configer struct {
	Configs map[string]string
}

func NewConfiger() *Configer {
	return &Configer{
		Configs: make(map[string]string),
	}
}
