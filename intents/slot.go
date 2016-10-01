package intent

type SlotType string

const (
	AzDate      SlotType = "AMAZON.DATE"
	AzDuration  SlotType = "AMAZON.DURATION"
	AzFourDigit SlotType = "AMAZON.FOUR_DIGIT_NUMBER"
	AzNumber    SlotType = "AMAZON.NUMBER"
	AzTime      SlotType = "AMAZON.TIME"
	AzUsCity    SlotType = "AMAZON.US_CITY"
	AzFirstName SlotType = "AMAZON.US_FIRST_NAME"
	AzState     SlotType = "AMAZON.US_STATE"
)

type Slot struct {
	Name string   `json:"name,"`
	Type SlotType `json:"type,"`
}
