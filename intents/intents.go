package intent

type Intent struct {
	Name  string  `json:"intent,"`
	Slots []*Slot `json:"slots,omitempty"`
}

func NewCancelIntent(slots []*Slot) *Intent {
	return &Intent{
		Name:  "AMAZON.CancelIntent",
		Slots: slots,
	}
}

func NewStartOverIntent(slots []*Slot) *Intent {
	return &Intent{
		Name:  "AMAZON.StartOverIntent",
		Slots: slots,
	}
}

func NewHelpIntent(slots []*Slot) *Intent {
	return &Intent{
		Name:  "AMAZON.HelpIntent",
		Slots: slots,
	}
}

func NewNextIntent(slots []*Slot) *Intent {
	return &Intent{
		Name:  "AMAZON.NextIntent",
		Slots: slots,
	}
}

func NewNoIntent(slots []*Slot) *Intent {
	return &Intent{
		Name:  "AMAZON.NoIntent",
		Slots: slots,
	}
}

func NewPauseIntent(slots []*Slot) *Intent {
	return &Intent{
		Name:  "AMAZON.PauseIntent",
		Slots: slots,
	}
}

func NewPreviousIntent(slots []*Slot) *Intent {
	return &Intent{
		Name:  "AMAZON.PreviousIntent",
		Slots: slots,
	}
}

func NewRepeatIntent(slots []*Slot) *Intent {
	return &Intent{
		Name:  "AMAZON.RepeatIntent",
		Slots: slots,
	}
}

func NewResumeIntent(slots []*Slot) *Intent {
	return &Intent{
		Name:  "AMAZON.ResumeIntent",
		Slots: slots,
	}
}

func NewStartOverIntent(slots []*Slot) *Intent {
	return &Intent{
		Name:  "AMAZON.StartOverIntent",
		Slots: slots,
	}
}

func NewStopIntent(slots []*Slot) *Intent {
	return &Intent{
		Name:  "AMAZON.StopIntent",
		Slots: slots,
	}
}
func NewYesIntent(slots []*Slot) *Intent {
	return &Intent{
		Name:  "AMAZON.YesIntent",
		Slots: slots,
	}
}
