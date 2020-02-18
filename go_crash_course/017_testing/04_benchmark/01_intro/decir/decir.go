//Package decir nos permite saludar
package decir

import "fmt"

// Saludar saluda a una persona
func Saludar(s string) string {

	return fmt.Sprint("Bienvenido Querido ", s)
}
