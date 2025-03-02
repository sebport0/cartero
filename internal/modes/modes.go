package modes

type (
	Mode struct {
		Name string
	}
	Normal       struct{ Mode }
	Sprinter     struct{ Mode }
	Orchestrator struct{ Mode }
)
