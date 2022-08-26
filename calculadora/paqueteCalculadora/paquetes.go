package paqueteCalculadora

import (
	"fmt"
	"strconv"
	"strings"
)

//Paquete de utilidades de la calculadora

//Clase o Struct Calculadora
type Calculadora struct {
	Identificador     int
	ResSuma           float64
	ResResta          float64
	ResDivision       float64
	ResMultiplicacion float64
	Resultado         float64
}

//Constructor de Calculadora
func NuevaCalculadora(t int) *Calculadora {
	//retornamos la copia de Calculadora
	return &Calculadora{
		Identificador:     t,
		ResSuma:           0,
		ResResta:          0,
		ResDivision:       0,
		ResMultiplicacion: 0,
		Resultado:         0,
	}
}

//funcion Sumar de Calculadora
func (c *Calculadora) Sumar(a, b float64) float64 {
	//convertimos
	result := (a + b)
	c.ResSuma = result
	//Resultado general
	c.Resultado = result
	return result

}

//Funcion Restar de Calculadora
func (c *Calculadora) Restar(a, b float64) float64 {
	//convertimos
	result := (a - b)
	c.ResResta = result
	//Resultado general
	c.Resultado = result
	return result

}

//Funcion Dividir de Calculadora
func (c *Calculadora) Dividir(a, b float64) float64 {
	//convertimos
	result := (a / b)
	c.ResDivision = result
	//Resultado general
	c.Resultado = result
	return result

}

//Funcion Multiplicar de Calculadora
func (c *Calculadora) Multiplicar(a, b float64) float64 {
	//convertimos
	result := (a * b)
	c.ResMultiplicacion = result
	//Resultado general
	c.Resultado = result
	return result

}

func (c *Calculadora) Convertir(op string) (a, b float64, operacion string) {
	//Valida que tipo de operacion es
	sum := strings.Contains(op, "+")
	rest := strings.Contains(op, "-")
	div := strings.Contains(op, "/")
	mult := strings.Contains(op, "*")
	//Tomamos los casos
	if sum == true {
		//Dividimos String op y lo convertimos en un slice
		values := strings.Split(op, "+")
		// Cast valores from text to number
		operador1, _ := strconv.ParseFloat(values[0], 64)
		operador2, _ := strconv.ParseFloat(values[1], 64)
		return operador1, operador2, "+"
	}
	if rest == true {
		//Dividimos String op y lo convertimos en un slice
		values := strings.Split(op, "-")
		// Cast valores from text to number
		operador1, _ := strconv.ParseFloat(values[0], 64)
		operador2, _ := strconv.ParseFloat(values[1], 64)
		return operador1, operador2, "-"
	}
	if div == true {
		//Dividimos String op y lo convertimos en un slice
		values := strings.Split(op, "/")
		// Cast valores from text to number
		operador1, _ := strconv.ParseFloat(values[0], 64)
		operador2, _ := strconv.ParseFloat(values[1], 64)
		return operador1, operador2, "/"
	}
	if mult == true {
		//Dividimos String op y lo convertimos en un slice
		values := strings.Split(op, "*")
		// Cast valores from text to number
		operador1, _ := strconv.ParseFloat(values[0], 64)
		operador2, _ := strconv.ParseFloat(values[1], 64)
		return operador1, operador2, "*"
	}

	return

}

//Se agrega funcion que trabaja el output del struct o clase Calculadora
func (c Calculadora) String() string {
	//se agrega logica para trabajar la salida de la funcion
	//recordar %d para int, %s string, %f para floats y %v time
	//return fmt.Sprintf("Identificador de Calculadora: %d Suma %f , Resta %f Division %f , Multiplicacion %f, Ejcucion : %v", c.Identificador, c.ResSuma, c.ResResta, c.ResDivision, c.ResMultiplicacion, time.Now())
	return fmt.Sprintf("Ejecucion Exitosa! \n Identificador de Calculadora: %d Resultado de la operacion: %f", c.Identificador, c.Resultado)

}
