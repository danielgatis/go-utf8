package utf8

const continuationMask byte = 0x3f

// State represents a state type
type State = byte

// Action represents a action type
type Action = byte

// Performer is an interface for parsing.
type Performer interface {
	// CodePoint is called when a codepoint is emitted.
	CodePoint(r rune)

	// InvalidSequence is called when an invalid sequence is found.
	InvalidSequece()
}

// Parser represents a state machine
type Parser struct {
	codepoint rune
	state     State
	performer Performer
}

// New returns a new parser
func New(performer Performer) *Parser {
	return &Parser{
		codepoint: 0,
		state:     groundState,
		performer: performer,
	}
}

// Advance advances the state machine
func (p *Parser) Advance(b byte) {
	change := stateTable[p.state][b]
	state := change & 0x0f
	action := change >> 4

	p.performAction(action, b)
	p.state = state
}

// Codepoint returns a codepoint
func (p *Parser) Codepoint() rune {
	return p.codepoint
}

// StateName returns the current state name
func (p *Parser) StateName() string {
	return stateNames[p.state]
}

// State returns the current state
func (p *Parser) State() State {
	return p.state
}

func (p *Parser) performAction(action, b byte) {
	switch action {
	case invalidSequenceAction:
		p.codepoint = 0
		p.performer.InvalidSequece()

	case emitByteAction:
		p.performer.CodePoint(rune(b))

	case setByte1Action:
		point := p.codepoint | rune(b&continuationMask)
		p.codepoint = 0
		p.performer.CodePoint(point)

	case setByte2Action:
		p.codepoint = p.codepoint | rune((b&continuationMask))<<6

	case setByte2TopAction:
		p.codepoint = p.codepoint | rune((b&0x1f))<<6

	case setByte3Action:
		p.codepoint = p.codepoint | rune((b&continuationMask))<<12

	case setByte3TopAction:
		p.codepoint = p.codepoint | rune((b&0x0f))<<12

	case setByte4Action:
		p.codepoint = p.codepoint | rune((b&0x07))<<18
	}
}
