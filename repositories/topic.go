package repositories

// TopicCode represents de code for a Topic
type TopicCode string

// Topic contains information about one topic
type Topic struct {
	Code TopicCode `json:"code"`
	Name string    `json:"name"`
}

// TopicRepository interface
type TopicRepository interface {
	List() []Topic
	Get(code TopicCode) (Topic, error)
	Create(code TopicCode, name string) (Topic, error)
}
