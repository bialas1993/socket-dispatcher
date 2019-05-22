package model

type Socket struct {
	Id   int    `orm:"auto"`
	Hash string `orm:"size(16)"`
	Port int
}

type ProjectPorts struct {
	Id        int `orm:"auto"`
	PortStart int
	PortEnd   int
}
