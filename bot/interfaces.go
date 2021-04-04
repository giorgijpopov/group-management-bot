package bot

type Manager interface {
	Start()
	SetupHandles()
	HandleError(err error) bool
}
