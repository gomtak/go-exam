package model

type result string

func NewModel() (result, error) {
	return "", nil
}

func (m result) Run(s string) string {
	return "run"
}
func (m result) Jump(s string) string {
	return "jump"
}
func (m result) Sleep(s string) string {
	return "sleep"
}
func (m result) Walk(s string) string {
	return "walk"
}
func (m result) Fly(s string) string {
	return "fly"
}
