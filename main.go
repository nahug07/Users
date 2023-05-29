package main

import "os"
import "fmt"
import "bufio"
import "strings"

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

func main() {

	var reader *bufio.Reader
	var option string

	reader = bufio.NewReader(os.Stdin)

	fmt.Println("A) Crear")
	fmt.Println("B) Listar")
	fmt.Println("C) Actualizar")
	fmt.Println("D) Eliminar")

	fmt.Println("Ingresa una opción válida: ")
	option, err := reader.ReadString('\n')

	if err != nil {
		panic("No es posible obtener el valor")
	}

	option = strings.TrimSuffix(option, "\n")

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