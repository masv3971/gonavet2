package navettypes

var (
	ErrSKVInvalidCall             = &Errors{Message: "Felaktigt anrop"}
	ErrSKVNotLogedIn              = &Errors{Message: "Ej inloggad"}
	ErrSKVAuthorizationNotAllowed = &Errors{Message: "Åtkomst nekad"}
	ErrSKVNoValidContentType      = &Errors{Message: "Felaktigt anrop"}
	ErrSKVTechnicalFail           = &Errors{Message: "Tekniskt fel"}
	ErrSKVNotSpecified            = &Errors{Message: "Not specified"}
	ErrSKVTooManyRequests         = &Errors{Message: "Too many requests"}
	ErrSKVTimeout                 = &Errors{Message: "Timeout"}
)

// Errors is the bespoke error struct
type Errors struct {
	SkvClientCorrelationID string `json:"skv_client_correlation_id"`
	Message                string `json:"message"`
}

func (e *Errors) addID(id string) {
	e.SkvClientCorrelationID = id
}

func (e *Errors) Error() string {
	//	if e.Ladok != nil && len(e.Internal) > 0 {
	//		return fmt.Sprintf("internal error: %v, ladok error: %v", e.Internal, e.Ladok)
	//	} else if len(e.Internal) > 0 {
	//		return fmt.Sprintf("internal error: %v", e.Internal)
	//	} else if e.Ladok != nil {
	//		return fmt.Sprintf("ladok error: %v", e.Ladok)
	//	}
	return ""
}

// Error interface
type Error interface {
	Error() string
}

func oneError(m, t, f, e string) *Errors {
	return nil
}
