package math

//Squares saca el cuadrado de un numero
func Squares(number int) (sum int) {
	sum = 0
	for number != 0 {
		digit := number % 10
		sum += digit * digit
		number /= 10
	}
	return
}

//Cubes saca el cubo de un numero
func Cubes(number int) (sum int) {
	sum = 0
	for number != 0 {
		digit := number % 10
		sum += digit * digit * digit
		number /= 10
	}
	return
}

//SquaresPlusCubes suma el valor del resultado de sacar un cuadrado y un cubo
func SquaresPlusCubes(number int) int {
	return Squares(number) + Cubes(number)
}

//SquaresChan saca el cuadrado de un numero con channels
func SquaresChan(number int, numberChan chan<- int) {
	sum := 0
	for number != 0 {
		digit := number % 10
		sum += digit * digit
		number /= 10
	}
	numberChan <- sum
}

//CubesChan saca el cubo de un numero
func CubesChan(number int, numberChan chan<- int) {
	sum := 0
	for number != 0 {
		digit := number % 10
		sum += digit * digit * digit
		number /= 10
	}
	numberChan <- sum
}

//SquaresPlusCubesChan suma el valor del resultado de sacar un cuadrado y un cubo
func SquaresPlusCubesChan(number int, numberChan chan<- int) {
	var squaresChan = make(chan int)
	var cubesChan = make(chan int)

	go SquaresChan(number, squaresChan)
	go CubesChan(number, cubesChan)

	squares, cubes := <-squaresChan, <-cubesChan
	numberChan <- squares + cubes
}
