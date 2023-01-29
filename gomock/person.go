package gomock

type Person interface {
	Get(id int64) (string, error)
}

type Male struct {
	Name string
}

func (m *Male) Get(id int64) (string, error) {
	return m.Name, nil
}
