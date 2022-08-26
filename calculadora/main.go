package main

//Paquete de Ejecucion Principal
import (
	"bufio"
	paqueteCalcuadora "calculadora/paqueteCalculadora"
	"fmt"
	"os"
	"strings"
)

//Ejecucion Principal
func main() {
	//Mensaje introductorio
	fmt.Println("Ingrese la operacion (Ejemplo de Suma, de la forma numero + numero, ej: 2+2) para Resta con - para Divison con / y para Multiplicacion *: ")
	//definimos nuestro scanner
	scanner := bufio.NewScanner(os.Stdin)
	//Escaneamos
	scanner.Scan()
	op := scanner.Text()
	fmt.Println("\nla Operacion ingresada es:", op)
	//instanciamos
	MyCalculadora := paqueteCalcuadora.NuevaCalculadora(1)
	//Convertimos lo que ingresamos y definimos que operacion es
	v1, v2, oper := MyCalculadora.Convertir(op)
	//Validamos
	switch {
	case strings.Contains(oper, "+"):
		//do sum
		fmt.Println("Tipo de operacion: Suma")
		MyCalculadora.Sumar(v1, v2)

	case strings.Contains(oper, "-"):
		//do rest
		fmt.Println("Tipo de operacion: Resta")
		MyCalculadora.Restar(v1, v2)

	case strings.Contains(oper, "/"):
		//do div
		fmt.Println("Tipo de operacion: Dividir")
		MyCalculadora.Dividir(v1, v2)

	case strings.Contains(oper, "*"):
		//do mult
		fmt.Println("Tipo de operacion: Suma")
		MyCalculadora.Multiplicar(v1, v2)

	default:
		s := fmt.Sprintf("Operacion: %s no valida!", op)
		fmt.Println(s)
	}

	defer fmt.Println(MyCalculadora.Resultado)

}
