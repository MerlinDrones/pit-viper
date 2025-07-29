package module

type Module struct {
	config *ModuleConfig
}

func NewModule() *Module {
	return &Module{
		config: NewModuleConfig(),
	}
}

func (m *Module) String() string {
	return m.config.String()
}
