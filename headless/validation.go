package headless

type ValidationRule[K comparable] func(value K) (bool, string)

var (
	RuleRequired = func(errMsg string) ValidationRule[string] {
		return func(value string) (bool, string) {
			return value != "", errMsg
		}
	}

	RuleMinLength = func(min int, errMsg string) ValidationRule[string] {
		return func(value string) (bool, string) {
			return len(value) >= min, errMsg
		}
	}

	RuleMaxLength = func(max int, errMsg string) ValidationRule[string] {
		return func(value string) (bool, string) {
			return len(value) <= max, errMsg
		}
	}

	RuleBetween = func(min, max int, errMsg string) ValidationRule[string] {
		return func(value string) (bool, string) {
			return len(value) >= min && len(value) <= max, errMsg
		}
	}

	RuleEmail = func(errMsg string) ValidationRule[string] {
		return func(value string) (bool, string) {
			return true, errMsg
		}
	}
)
