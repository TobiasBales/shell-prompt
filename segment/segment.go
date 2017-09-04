package segment

// Segment is used as one segment in the shellprompt
type Segment interface {
	Placeholder() string
	Value() chan string
}
