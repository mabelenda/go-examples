//Package mystr sirve para concatenar strings
package mystr

import "strings"

//Cat sirve para concantenar de manera manual strings
func Cat(xs []string) string {
	s := xs[0]
	for _, v := range xs[1:] {
		s += " "
		s += v
	}
	return s
}

//Join sirve para concantenar de manera manual strings
func Join(xs []string) string {

	return strings.Join(xs, " ")
}
