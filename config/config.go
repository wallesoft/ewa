package config

//Config config interface
type Config interface {
	Get(key string, def ...interface{}) interface{}
	Set(key string, val interface{})
	Containos(key string) bool
}
