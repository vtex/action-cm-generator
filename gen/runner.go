package gen

// Runner is used to compose generating pipeline and run it.
type Runner struct {
	Reader    Reader
	Compiler  Compiler
	Parser    Parser
	Validator Validator
	Exporter  Exporter
}

// Run executes the config generate pipeline.
func (s *Runner) Run() error {
	return s.Exporter.Export(
		s.Validator.Validate(
			s.Parser.Parse(
				s.Compiler.Compile(
					s.Reader.Read(),
				),
			),
		),
	)
}
