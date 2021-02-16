package styinterface

import (
	"fmt"
	"math/rand"
	"sort"
	"strconv"
)

type Hero struct {
	Id int
	Name string
}

//2.申明一个切片类型
type HeroSlice []Hero


func (heroslice HeroSlice)Len()int{
	return len(heroslice)
}

func (heroslice HeroSlice) Less(i, j int) bool{
	return heroslice[i].Id < heroslice[j].Id
}

func (heroslice HeroSlice) Swap(i, j int){
	/*tmp := *heroslice[i]
	heroslice[i] = heroslice[j]
	heroslice[j] = tmp*/
	heroslice[i],heroslice[j] = heroslice[j],heroslice[i]
}






func UseHero()  {
	var  hs HeroSlice
	for i:=0;i < 10;i++ {
		hero := Hero{Id :rand.Intn(100),Name :"名字" + strconv.Itoa(i)}
		hs = append(hs,hero)
	}
	fmt.Println("排序前的数组：",hs)
	sort.Sort(hs)
	fmt.Println("排序后的数组：",hs)
}





