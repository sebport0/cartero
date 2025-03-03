package modes

type Mode interface{}

type (
	Normal struct {
		Name string
	}
	Sprinter     struct{}
	Orchestrator struct{}
)
