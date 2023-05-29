package main

import "os"
import "fmt"
import "bufio"
import "strings"
import "strconv"

var reader *bufio.Reader

type User struct {
	id int
	username string
	email string
	age int
}

var id int
var users map[int]User

func crearUsuario() {
	fmt.Println("Ingresa nombre de usuario: ")
	username := readLine()

	fmt.Println("Ingresa email: ")
	email := readLine()

	fmt.Println("Ingresa edad: ")
	age, err := strconv.Atoi(readLine())

	if err != nil {
		panic("No es posible convertir de un string a un entero")
	}

	id++
	user := User { id, username, email, age}
	users[id] = user

	fmt.Println(users)
	fmt.Println("Usuario creado exitosamente!")
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
	users = make(map[int]User)

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