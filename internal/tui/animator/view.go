package animator

func (m *Model) View() string {
	return m.Frames[m.CurrentFrame].String()
}
