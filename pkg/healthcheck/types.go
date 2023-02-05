package healthcheck

//go:generate easyjson types.go

type ApplicationRuntimeDirective uint8

const (
	ApplicationRuntimeDirectiveTerm = iota + 1
	ApplicationRuntimeDirectiveReload
	ApplicationRuntimeDirectiveNothing
)

const (
	UnitNameLiveness = "liveness_checker_unit"
	UnitNameRediness = "rediness_checker_unit"
	UnitNameStartup  = "startup_checker_unit"
)

// easyjson:json
type Status struct {
	IsHealed  bool                        `json:"is_healed"`
	Directive ApplicationRuntimeDirective `json:"directive"`
	Message   string                      `json:"message"`
	Error     error                       `json:"error,omitempty"`
}
