package styfactory

type student struct {
	Name string
	Age int
	score float64
}

func GetInstance(name string,age int) *student{
	return &student{Name: name,Age: age}
}
func (s *student) SetScore(score float64){
	s.score = score
}

func (s *student) GetScore()float64{
	return s.score
}

















