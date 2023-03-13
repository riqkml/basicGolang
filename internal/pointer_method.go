package internal

type Man struct {
	Name string
}

func (m *Man) GetMarried() {
	m.Name = "Mr. " + m.Name
}
