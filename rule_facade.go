package gopass

type RuleFacade struct {
}

func NewRuleFacade() *RuleFacade {
	return &RuleFacade{}
}
func (r *RuleFacade) Required() string {
	return "required"
}
