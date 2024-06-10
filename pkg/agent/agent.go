package agent

type Agent struct {
	Docker *Docker
}

func NewAgent() *Agent {
	return &Agent{
		Docker: NewDocker(),
	}
}
func (a *Agent)