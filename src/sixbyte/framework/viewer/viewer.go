package viewer

type Viewer struct {
	ViewDir string
}

func NewViewer() *Viewer {
	return &Viewer{}
}
