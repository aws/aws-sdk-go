package ini

// skipper is used to skip certain blocks of an ini file.
// Currently skipper is used to skip nested blocks of ini
// files. See example below
//
//	[ foo ]
//	nested = // this section will be skipped
//		a=b
//		c=d
//	bar=baz // this will be included
type skipper struct {
	shouldSkip bool
	prevTok    Token
}

func (s *skipper) ShouldSkip(tok Token) bool {
	if s.shouldSkip && s.prevTok != nil && s.prevTok.Type() == TokenNL && tok.Type() != TokenWS {
		s.Continue()
		return false
	}
	s.prevTok = tok

	return s.shouldSkip
}

func (s *skipper) Skip() {
	s.shouldSkip = true
	s.prevTok = nil
}

func (s *skipper) Continue() {
	s.shouldSkip = false
	s.prevTok = nil
}
