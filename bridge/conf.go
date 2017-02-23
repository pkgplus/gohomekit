package bridge

type Configure struct {
	Bridge struct {
		Name string `json:"name"`
		PIN  string `json:"pin"`
	} `json:"bridge"`
	PlatForms []*PlatFormConf `json:"platforms"`
}

type PlatFormConf struct {
	PlatForm string   `json:"platform"`
	Sid      []string `json:"sid"`
	Password []string `json:"password"`
}
