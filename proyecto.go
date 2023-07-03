package main

import (
	"fmt"

	f "paquetes.com/funciones"
)

func main() {
	fmt.Println("=====================================================")
	fmt.Println("||          Programa de Control Remoto    	   ||")
	fmt.Println("=====================================================")

	for {
		var opcion string
		// --- Mostrar opciones ---
		fmt.Println("\nSeleccione una opción:")
		fmt.Println("1. Servidor de comandos")
		fmt.Println("2. Conectar como cliente y enviar comandos")
		fmt.Println("3. Salir")
		// --- Leer opción ---
		fmt.Scanf("%s", &opcion)
		
		// --- Evaluar opción ---
		switch opcion {
		case "1":
			for {
				var opcionBD string
				//--- Mostrar opciones ---
				fmt.Println("\nSeleccione una opción:")
				fmt.Println("1. Iniciar servidor")
				fmt.Println("2. Crear Archivo de usuarios")
				fmt.Println("3. Adicionar Usuario")
				fmt.Println("4. Leer Archivo de Usuarios")
				fmt.Println("5. Salir")
				//--- Leer opcion ---
				fmt.Scanf("%s",&opcionBD)
				
				//--- Evaluar opcion ---
				switch opcionBD {
				case "1":
					f.IniciarServidor()
				case "2":
					f.CreaArchivo()
				case "3":
					f.AdicionarUsuario()
				case "4":
					f.LeerArchivo()
				case "5":
					fmt.Println("Saliendo...")
					fmt.Println("Programa finalizado")
					return
				default:
					fmt.Println("Opción no válida. Por favor, seleccione una opción válida.")
				}
			}
		case "2":
			f.ConectarYEnviarComandos()
		case "3":
			fmt.Println("Programa finalizado.")
			return
		default:
			fmt.Println("Opción no válida. Por favor, seleccione una opción válida.")
		}
	}
}
