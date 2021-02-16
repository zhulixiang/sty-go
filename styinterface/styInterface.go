package styinterface
import "fmt"

type Usb interface {
	Start()
	Stop()
}

type Phone struct {

}
func (p Phone)Start(){
	fmt.Println("手机开始工作")
}
func (p Phone)Stop(){
	fmt.Println("手机停止工作")
}

type Camera struct {

}

func (c Camera)Start(){
	fmt.Println("相机开始工作")
}
func (c Camera)Stop(){
	fmt.Println("相机停止工作")
}

type Computer struct {

}

func (computer Computer)Working(usb Usb){
	usb.Start()
	usb.Stop()
}

func UseInterface(){
	defer func() {
		if err := recover();err != nil {
			fmt.Println(err)
		}
	}()

	var computer = Computer{}
	var phone = Phone{}
	var camera = Camera{}

	computer.Working(phone)
	computer.Working(camera)

	var t interface{} = computer
	fmt.Println(t)
	var num1 float64 = 8.8
	t = num1
	fmt.Println(t)
}




















