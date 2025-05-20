package main

type retangulo struct {
	comprimento, altura int
}

func (qualquerCoisa retangulo) area() int {
	return qualquerCoisa.comprimento + qualquerCoisa.altura
}

func (qualquerCoisa retangulo) perimetro() int {
	return 2 * (qualquerCoisa.comprimento + qualquerCoisa.altura)
}

func main() {
	retang := retangulo{comprimento: 50, altura: 25}
	println("Área do retângulo:", retang.area())
	println("Comprimento do retângulo:", retang.perimetro())
}