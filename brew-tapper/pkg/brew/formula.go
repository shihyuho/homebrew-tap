package brew

type Formula struct {
	Name, Version string
}

func (f *Formula) NotSpecified() bool {
	return len(f.Name) == 0 || len(f.Version) == 0
}
