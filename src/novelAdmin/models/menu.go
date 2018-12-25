package models

type Menu struct {
	Id        int
	Name      string
	Pid       int
	Url       string
	TopShow   int
	AddUrl    string
	HasAdd    int
	Listorder int
	Status    int
}

func (this *Menu) getAllMenu() {

}
