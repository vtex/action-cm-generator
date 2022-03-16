package gen

// Exporter is used to write config final result somewhere.
type Exporter interface {
	// Export receives a chan of configs and export the config.
	Export(in <-chan Config) error
}
