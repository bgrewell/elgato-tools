package types

type Lights struct {
	NumberOfLights int      `json:"numberOfLights" yaml:"numberOfLights"`
	Lights         []*Light `json:"lights"`
}

type Light struct {
	On          int `json:"on" yaml:"on"`
	Brightness  int `json:"brightness" yaml:"brightness"`
	Temperature int `json:"temperature" yaml:"temperature"`
}
