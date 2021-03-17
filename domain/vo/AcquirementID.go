package vo

type AcquirementID string

type AcquirementIDRequest struct {
	AcquirementID string
}

func NewAcquirementID(req AcquirementIDRequest) (AcquirementID, error) {
	obj := AcquirementID(req.AcquirementID)
	return obj, nil
}
