package service

// validate проверяет корректность полей name, severity и type.
func validate(name, severity, typeThreat string) bool {
	if name == "" {
		return false
	}

	var validSeverities = map[string]struct{}{
		"low": {}, "medium": {}, "high": {}, "critical": {},
	}
	_, ok := validSeverities[severity]
	if !ok {
		return false
	}

	var validTypes = map[string]struct{}{
		"trojan": {}, "worm": {}, "ransomware": {}, "adware": {}, "spyware": {}, "other": {},
	}
	_, ok = validTypes[typeThreat]
	if !ok {
		return false
	}

	return true

}
