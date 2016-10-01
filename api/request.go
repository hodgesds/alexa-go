package api

type Request string

type Reason string

const (
	UserInitiated Reason = "USER_INITIATED"
	Error         Reason = "ERROR"
	MaxReprompts  Reason = "MAX_REPROMPTS"

	Intent     Request = "INTENT"
	Launch     Request = "LAUNCH"
	SessionEnd Request = "SESSION_END"
)

type Application struct {
	ID string `json:"applicationId"`
}

type User struct {
	ID    string `json:"userId,"`
	Token string `json:"accessToken,"`
}

type Session struct {
	Application *Application           `json:"application,"`
	Attributes  map[string]interface{} `json:"attributes,"`
	ID          string                 `json:"sessionId,"`
	New         bool                   `json:"new,"`
	User        *User                  `json:"user,"`
}

type AlexaRequest struct {
	ID        string         `json:"requestId,"`
	Intent    *intent.Intent `json:"intent,omitempty"`
	Reason    Reason         `json:"reason,omitempty"`
	Timestamp string         `json:"timestamp,"`
	Type      Request        `json:"type,"`
}

func (ar *AlexaRequest) Validate() error {
	if err := ar.ValidateTimestamp(); err != nil {
		return err
	}

	return nil
}

func (ar *AlexaRequest) ValidateTimestamp() error {
	ts, err := time.Parse(time.RFC3339, ar.Timestamp)
	if err != nil {
		return err
	}

	now := time.Now()

	if cutoff := now.Add(-150 * time.Second); cutoff.After(ts) {
		return fmt.Errorf("Too old")
	}

	return nil
}

type APIRequest struct {
	AlexaRequest *AlexaRequest `json:"request,"`
	Session      *Session      `json:"session,"`
	Version      string        `json:"version,"`
}
