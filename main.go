package main

import "os"
import "fmt"
import "bufio"
import "strings"

var reader *bufio.Reader

func crearUsuario() {
	fmt.Println("crear usuario")
}

func listarUsuarios() {
	fmt.Println("listar usuarios")
}

func actualizarUsuario() {
	fmt.Println("actualizar usuario")
}

func eliminarUsuario() {
	fmt.Println("eliminar usuario")
}

func readLine() string {

	if option, err := reader.ReadString('\n'); err != nil {
		panic("No es posible obtener el valor")
	} else {
		return strings.TrimSuffix(option, "\n")
	}
}

func main() {

	var option string

	reader = bufio.NewReader(os.Stdin)

	for {
		fmt.Println("A) Crear")
		fmt.Println("B) Listar")
		fmt.Println("C) Actualizar")
		fmt.Println("D) Eliminar")

		fmt.Println("Ingresa una opción válida: ")

		option = readLine()

		if option == "salir" || option == "s" {
			break
		}

		switch option {
			case "a", "crear":
				crearUsuario()
			case "b", "listar":
				listarUsuarios()
			case "c", "actualizar":
				actualizarUsuario()
			case "d", "eliminar":
				eliminarUsuario()
			default:
				fmt.Println("Opción no válida")				
		}
	}

	fmt.Println("Adiós 0_0")
}