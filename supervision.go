package actor

type Directive int
const (
    ResumeDirective Directive = iota
    RestartDirective
    StopDirective
    EscalateDirective
)

type Decider func(child ActorRef, cause interface{}) Directive

type SupervisionStrategy interface {
    Handle(child ActorRef, cause interface{}) Directive
}

type OneForOneStrategy struct {
	maxNrOfRetries              int
	withinTimeRangeMilliseconds int
	decider                     Decider
}

func (strategy *OneForOneStrategy) Handle(child ActorRef, reason interface{}) Directive {
	return strategy.decider(child, reason)
}

func NewOneForOneStrategy(maxNrOfRetries int, withinTimeRangeMilliseconds int, decider Decider) SupervisionStrategy {
	return &OneForOneStrategy{
		maxNrOfRetries:              maxNrOfRetries,
		withinTimeRangeMilliseconds: withinTimeRangeMilliseconds,
		decider:                     decider,
	}
}

func DefaultDecider(child ActorRef, reason interface{}) Directive {
	return RestartDirective
}

var defaultStrategy SupervisionStrategy = NewOneForOneStrategy(10, 30000, DefaultDecider)

func DefaultStrategy() SupervisionStrategy {
	return defaultStrategy
}