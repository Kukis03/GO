package funciones

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"golang.org/x/crypto/ssh/terminal"
)

func CreaArchivo() {
	passBD := ""
	datoArchivar := []byte(passBD)
	nombreArchivo := "users.pw"
	ioutil.WriteFile(nombreArchivo, datoArchivar, 0777)
	fmt.Println("Archivo de Usuarios Creado")
}

func AdicionarUsuario() {
	var usr, psw, passBD string
	fmt.Println("Nombre de usuario: ")
	fmt.Scanf("%s", &usr)
	fmt.Println("Contraseña: ")
	bytePassword, _ := terminal.ReadPassword(int(os.Stdin.Fd()))
	psw = string(bytePassword)
	pswArray := sha256.Sum256([]byte(psw))
	passBD = usr + ":" + hex.EncodeToString(pswArray[:]) + ":"
	archivo, err := os.OpenFile("users.pw", os.O_APPEND|os.O_WRONLY, 0777)
	if err != nil {
		panic(err)
	}
	archivo.WriteString(passBD)
	archivo.Close()
	fmt.Println("Usuario añadido")
}

func LeerArchivo() {
	archivo, _ := ioutil.ReadFile("users.pw")
	sArchivo := string(archivo)
	fmt.Println("Contenido: ", sArchivo)

	credenciales := strings.Split(sArchivo, ":")
	for i := 0; i < len(credenciales); i++ {
		fmt.Println(" credenciales[", i, "]", credenciales[i])
	}
}
