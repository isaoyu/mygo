package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

func main() {
	fmt.Println("入口文件")
	PersonDream()
}

//结构体 https://www.liwenzhou.com/posts/Go/10_struct/#autoid-2-4-1
//构体中字段大写开头表示可公开访问，小写表示私有（仅在定义当前结构体的包中可访问）。

/***** 结构体和方法补充知识点 */
type Person struct {
	name   string
	dreams []string
}

// SetDream 设置梦想
func (p *Person) SetDream(dreams []string) {
	//dreams值改变，p.dreams也改变
	//p.dreams = dreams

	//拷贝一下，dreams值改变，p.dreams不变
	p.dreams = make([]string, len(dreams))
	copy(p.dreams, dreams)

	fmt.Printf("%s的梦想是%s\n", p.name, p.dreams)
}
func PersonDream() {
	p := Person{name: "大白"}
	fmt.Println(p)
	data := []string{"吃饭", "睡觉", "打豆豆"}
	fmt.Println(data)
	p.SetDream(data)
	fmt.Println(p)

	data[1] = "不睡觉"
	fmt.Println(data)
	fmt.Println(p.dreams)
}

/***** 结构体标签 *****/

type StudentTag struct {
	ID   int `json:"id"` //通过指定tag实现json序列化该字段时的key
	Name string
}

func StudentTagJson() {
	s := StudentTag{
		ID:   123,
		Name: "小冯",
	}
	fmt.Println(s)
	data, err := json.Marshal(s)
	if err != nil {
		fmt.Println("错误")
		return
	}
	fmt.Printf("%s\n", data)

}

/***** 结构体与json序列化 *************/

type Student struct {
	ID     int
	Gender string
	Name   string
}

type Class struct {
	Title    string
	Students []*Student
}

func StudentJson() {
	c := Class{
		Title:    "一年级二班",
		Students: make([]*Student, 0, 200),
	}
	fmt.Println(c)
	for i := 0; i < 10; i++ {
		stu := &Student{
			ID:   i,
			Name: fmt.Sprintf("stu%02d", i),
		}
		c.Students = append(c.Students, stu)
	}
	fmt.Println(c)
	//json序列化：结构体->json格式化字符串
	data, err := json.Marshal(c)
	if err != nil {
		fmt.Println("转json失败！")
		return
	}
	fmt.Println(data)
	fmt.Printf("json:%s\n", data)

	//json反序列化：json字符串->结构体
	//Class属性必须是：Title Students,Student属性必须是ID Gender Name，对应
	str := `{"Title":"101","Students":[{"ID":0,"Gender":"男","Name":"stu00"},{"ID":1,"Gender":"男","Name":"stu01"},{"ID":2,"Gender":"男","Name":"stu02"},{"ID":3,"Gender":"男","Name":"stu03"},{"ID":4,"Gender":"男","Name":"stu04"},{"ID":5,"Gender":"男","Name":"stu05"},{"ID":6,"Gender":"男","Name":"stu06"},{"ID":7,"Gender":"男","Name":"stu07"},{"ID":8,"Gender":"男","Name":"stu08"},{"ID":9,"Gender":"男","Name":"stu09"}]}`
	strData := []byte(str) //字符串转换成字节切片
	fmt.Println(str)
	fmt.Println(strData)
	c1 := &Class{}
	err = json.Unmarshal(strData, c1)
	fmt.Println(err)
	if err != nil {
		fmt.Println("json unmarshal failed!")
		return
	}
	fmt.Printf("%#v\n", c1)
	fmt.Printf("%v\n", c1)

	//遍历 c1 Students
	for k, v := range c1.Students {
		fmt.Println(k)
		fmt.Println(v)
	}
}

/***** 结构体的继承 ***************************/

type Animal struct {
	name string
}

func (a Animal) move() {
	fmt.Printf("%s会跑", a.name)
}

type Dog struct {
	Feet    int
	*Animal //通过嵌套匿名结构体实现继承
}

func (d Dog) wang() {
	fmt.Printf("%s会汪汪", d.name)
}

func DogWangWang() {
	dog := Dog{
		Feet:   2,
		Animal: &Animal{name: "豆豆"}, //嵌套的结构体指针
	}
	fmt.Println(dog)
	dog.move()
	dog.wang()
}

/***** 嵌套结构体匿名字段 ************/

type UserAddress struct {
	City string
}

type User2 struct {
	Name string
	Age  int
	UserAddress
}

func User2Init() {
	var user User2
	user.Name = "小周"
	user.Age = 24
	//user.UserAddress.City = "北京"
	user.City = "北京"
	fmt.Println(user)
}

/***** 嵌套结构体 *******/

type User struct {
	Name    string
	Gender  string
	Address Address
}

type Address struct {
	Province string
	City     string
}

func UserInit() {
	user := User{
		Name:   "小明",
		Gender: "男",
		Address: Address{
			Province: "山东",
			City:     "青岛",
		},
	}
	fmt.Println(user)
}

/************ 面试题 ***************/
type student struct {
	name string
	age  int
}

func statusStruct() {
	m := make(map[string]*student)
	stus := []student{
		{name: "小王", age: 19},
		{name: "小李", age: 29},
		{name: "小赵", age: 39},
	}

	for _, stu := range stus {
		fmt.Printf("stu %p\n", &stu)
		fmt.Printf("p2=%#v\n", &stu)
		fmt.Printf("%T\n", stu)     //*main.person
		fmt.Printf("p2=%#v\n", stu) //p2=&main.person{name:"", city:"", age:0}
		//指定的是同一个内存地址
		//m[stu.name] = &stu

		//改进
		value := stu
		m[stu.name] = &value
	}

	for k, v := range m {
		fmt.Println(k, "=>", v.name)
	}

	fmt.Printf("%T === %v", m, m)
	fmt.Printf("%T === %v", stus, stus)
}

/***************************/

/*********** person *******
	p := newPerson("小孙", "上海", 28)
	fmt.Println(p)
	p.Dream()
	p.SetAge(21)
	fmt.Println(p)
*********/

// NewInt 类型定义
type NewInt int

// MyInt 类型别名
type MyInt = int

//结构体
type person struct {
	name string
	age  int
	city string
}

// Dream 接收者
//方法与函数的区别是，函数不属于任何类型，方法属于特定的类型。
func (p person) Dream() {
	fmt.Printf("%s的梦想是吃冰糖葫芦麦旋风", p.name)
}

// SetAge 指针类型的接收者
func (p *person) SetAge(newAge int) {
	p.age = newAge
}

func newPerson(name, city string, age int) *person {
	return &person{
		name: name,
		city: city,
		age:  age,
	}
}

func myStruct() {
	var a NewInt
	var b MyInt
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)

	var p person
	p.name = "哒哒"
	p.age = 19
	p.city = "上海"

	var user struct {
		Name string
		Age  int
	}
	user.Name = "小驼驼"
	user.Age = 23

	fmt.Printf("%Tn", p)
	fmt.Println(p)
	fmt.Printf("%#v\n", user)

}

/********* person end ******************/

//strconv字符串操作
func strconvOperation() {
	s1 := "100"
	i1, err := strconv.Atoi(s1)
	if err != nil {
		fmt.Println("转换失败")
		return
	}
	fmt.Printf("type:%T value:%#v\n", i1, i1) //type:int value:100
}

func readFileWithFor() {
	file, err := os.Open("./main.go")
	if err != nil {
		fmt.Println("打开文件失败")
		return
	}
	defer file.Close()
	//循环读取文件
	var content []byte
	var tmp = make([]byte, 128)
	for {
		n, err := file.Read(tmp)
		if err == io.EOF {
			fmt.Println("文件读完了")
			break
		}
		if err != nil {
			fmt.Println("读取文件失败")
			return
		}
		content = append(content, tmp[:n]...)
	}
	fmt.Println(string(content))
}

func openMyFile() {
	file, err := os.Open("./main.go")
	if err != nil {
		fmt.Println("open file failed!,err:", err)
		return
	}
	fmt.Println(file)
	fmt.Printf("%T\n", file)
	//关闭文件
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("close file failed!,err:", err)
		}
	}(file)

	//read方法读取文件
	var tmp = make([]byte, 128)
	n, err := file.Read(tmp)
	if err == io.EOF {
		fmt.Println("文件读完了")
		return
	}
	if err != nil {
		fmt.Println("文件读取失败！")
		return
	}
	fmt.Printf("读取了%d字节数据\n", n)
	fmt.Println(string(tmp[:n]))
}

func runTime() {
	now := time.Now()
	timestamp1 := now.Unix()      //时间戳
	timestamp2 := now.UnixNano()  //纳秒时间戳
	timestamp3 := now.UnixMicro() //纳秒时间戳
	fmt.Println(timestamp1)
	fmt.Println(timestamp2)
	fmt.Println(timestamp3)
	n := 0
	for i := 0; i < 100; i++ {
		n += i
	}
	timestamp4 := now.UnixNano() //纳秒时间戳
	fmt.Println(timestamp4)
	fmt.Println(n)
	fmt.Println(timestamp4 - timestamp2)
}

func timeData() {
	now := time.Now()
	fmt.Println(now)
	fmt.Println(now.Format("2006/01/02 15:04"))

}

/*
你有50枚金币，需要分配给以下几个人：Matthew,Sarah,Augustus,Heidi,Emilie,Peter,Giana,Adriano,Aaron,Elizabeth。
分配规则如下：
a. 名字中每包含1个'e'或'E'分1枚金币
b. 名字中每包含1个'i'或'I'分2枚金币
c. 名字中每包含1个'o'或'O'分3枚金币
d: 名字中每包含1个'u'或'U'分4枚金币
写一个程序，计算每个用户分到多少金币，以及最后剩余多少金币？
程序结构如下，请实现 ‘dispatchCoin’ 函数
*/
func coins() {
	var (
		coins = 50
		users = []string{
			"Matthew", "Sarah", "Augustus", "Heidi", "Emilie", "Peter", "Giana", "Adriano", "Aaron", "Elizabeth",
		}
	)
	fmt.Printf("users type : %T", users)
	left := dispatchCoin(coins, users)
	fmt.Println("剩下：", left)
}

func dispatchCoin(coins int, users []string) int {
	distribution := make(map[string]int, len(users))
	for _, name := range users {
		fmt.Println(name)
		for _, c := range name {
			switch c {
			case 'e', 'E':
				distribution[name]++
				coins--
			case 'i', 'I':
				distribution[name] += 2
				coins -= 2
			case 'o', 'O':
				distribution[name] += 3
				coins -= 3
			case 'u', 'U':
				distribution[name] += 4
				coins -= 4
			}
		}
	}
	fmt.Println(distribution)
	return coins
}

func mapCode() {
	type Map map[string][]int
	m := make(Map)
	s := []int{1, 2}
	s = append(s, 3)
	fmt.Printf("%+v\n", s)
	m["q1mi"] = s
	s = append(s[:1], s[2:]...)
	fmt.Printf("%+v\n", s)
	fmt.Printf("%+v\n", m["q1mi"])
}

func statStrNum() {
	str := "how do you do"

	strArr := strings.Split(str, " ")

	strMap := make(map[string]int, 6)
	for _, val := range strArr {
		strMap[val]++
	}
	fmt.Println(strArr)
	fmt.Println(strMap)
}

func sortArr() {
	var arr = [...]int{3, 7, 8, 9, 1}
	sort.Ints(arr[:])
	fmt.Println(arr)
}

func makeSlice() {
	s := make([]int, 2, 10)
	fmt.Println(s)

	var a = make([]string, 5, 10)
	for i := 0; i < 10; i++ {
		a = append(a, fmt.Sprintf("%v", i))
	}
	fmt.Println(a)
}

func arrSum() {
	arr := [...]int{1, 3, 5, 7, 8}
	total := 0
	for _, item := range arr {
		total += item
	}
	fmt.Println(total)
}

func arrSearchSameSum() {
	arr := [...]int{1, 3, 5, 7, 8}
	for i := 0; i < len(arr); i++ {
		for j := i; j < len(arr); j++ {
			if arr[i]+arr[j] == 8 {
				fmt.Println(i, j)
			}
		}
	}
}
