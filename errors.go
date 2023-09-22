package weddinginvitation

type RecordNotFound struct {
	Model string
}

type InvalidOperation struct {
	Message string
}

type InvalidData struct {
	Message string
}

type DuplicateData struct {
	Message string
}

func (r RecordNotFound) Error() string {
	return r.Model + " not found"
}

func (r InvalidOperation) Error() string {
	return r.Message
}

func (r InvalidData) Error() string {
	return r.Message
}

func (r DuplicateData) Error() string {
	return r.Message
}

type WeddingInvitationError struct {
	Message string
}

func (t WeddingInvitationError) Error() string {
	return t.Message
}

func NewWeddingInvitationError(message string) error {
	return WeddingInvitationError{
		Message: message,
	}
}
