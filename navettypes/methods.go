package navettypes

func (s SimpleAPIError) error() error {
	switch s.Message {
	case "Too many requests":
		return ErrSKVTooManyRequests
	case "Timeout":
		return ErrSKVTimeout
	}
	return nil
}

func (f Felmeddelande) error() error {
	switch f.Orsakskodsbeskrivning {
	case "Felaktigt anrop: indata ej korrekt ifyllt":
		return ErrSKVInvalidCall
	case "Ej inloggad":
		return ErrSKVNotLogedIn
	case "Åtkomst nekad":
		return ErrSKVAuthorizationNotAllowed
	case "Felaktigt anrop: saknar stöd för Content-Type 'application/xml'":
		return ErrSKVNoValidContentType
	case "Tekniskt fel":
		return ErrSKVTechnicalFail
	default:
		return ErrSKVNotSpecified
	}
}
