package core

type O map[string]interface{}

func (o O) JS() map[string]interface{} {
	return o
}
