package main

import "fmt"

type USB interface {
	Name() string
	Connector
}

type Connector interface {
	Connect()
}

type PhoneConnector struct {
	name string
}

func (pc PhoneConnector) Name() string{
	return pc.name
}

func (pc PhoneConnector) Connect() {
	fmt.Println("Connected:",pc.name)
}

func main(){
	a := PhoneConnector{}
	a.name = "PhoneConnector"
	a.Connect()
	Disconnect(a)

	pc := PhoneConnector{name: "PhoneConnector"}
	var b Connector
	b = Connector(pc)
	b.Connect()
}

func Disconnect(usb interface{}){
	/*//类型判断
	if pc,ok := usb.(PhoneConnector);ok {
		fmt.Println("Disconnected:",pc.name)
		return
	} else {
		fmt.Println("Unknown device")
	}*/

	//类型判断
	switch v := usb.(type) {
	case PhoneConnector:
		fmt.Println("Disconnected:",v.name)
	default:
		fmt.Println("Unknown device")
	}
}