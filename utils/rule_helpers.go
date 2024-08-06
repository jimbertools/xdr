package utils

type RuleFactory struct {
	rules []string
}

func NewRuleFactory() *RuleFactory {
	const abcRule string = `
	rule test {
		meta: 
			author = "Wannes Vantorre"
		strings:
			$str = "abc"
		condition:
			$str
	}`

	const xyzRule string = `
	rule test {
		meta: 
			author = "Wannes Vantorre"
		strings:
			$str = "xyz"
		condition:
			$str
	}`
	rules := []string{abcRule, xyzRule}
	return &RuleFactory{rules: rules}
}

func (factory *RuleFactory) GetAllRules() []string {
	return factory.rules
}
