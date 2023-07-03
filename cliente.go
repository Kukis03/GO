package funciones

import (
	"bufio"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"time"

	"golang.org/x/crypto/ssh/terminal"
)

var salir = false

func ConectarYEnviarComandos() {
	fmt.Println("=====================================================")
	fmt.Println("||	   	 Cliente de Comandos               ||")
	fmt.Println("=====================================================")

	var ip, puerto string
	var reportes int
	fmt.Println("Cliente# Ingrese la IP del servidor al que se desea conectar: ")
	fmt.Scanf("%s", &ip)
	fmt.Println("Cliente# Puerto: ")
	fmt.Scanf("%s", &puerto)

	direccion := fmt.Sprintf("%s:%s", ip, puerto)
	tcpAddress, _ := net.ResolveTCPAddr("tcp4", direccion)

	lectorReport := bufio.NewReader(os.Stdin)
	fmt.Println("Cliente# Intervalo de tiempo en el que desea recibir reportes: ")
	timeReport, _ := lectorReport.ReadString('\n')
	reportesServer := strings.TrimSpace(timeReport)
	reportes, _ = strconv.Atoi(reportesServer)

	fmt.Println("Cliente# Conectando con...", tcpAddress.IP)
	socket, _ = net.DialTCP("tcp", nil, tcpAddress)
	fmt.Println("Cliente# Autenticando con...", socket.RemoteAddr())

	lector := bufio.NewReader(os.Stdin)
	fmt.Println("Cliente# Usuario: ")
	usr, _ := lector.ReadString('\n')
	usr = strings.TrimSpace(usr)

	fmt.Println("Cliente# Contraseña: ")
	bytePassword, _ := terminal.ReadPassword(int(os.Stdin.Fd()))
	psw := string(bytePassword)
	psw = strings.TrimSpace(psw)
	pswLog := sha256.Sum256([]byte(psw))
	psw = hex.EncodeToString(pswLog[:])

	fmt.Fprintf(socket, "%s\n%s\n%s\n", usr, psw, reportesServer)

	login, _ := bufio.NewReader(socket).ReadString('\n')
	login = strings.TrimSpace(login)

	salir = true

	go RecibeReporteClient(&socket, reportes)

	for {
		
		fmt.Println("\n\nCliente# Digite el comando a enviar (o 'bye' para desconectarse): ")
		lector := bufio.NewReader(os.Stdin)
		comando, _ := lector.ReadString('\n')

		env := bufio.NewWriter(socket)
		env.WriteString(comando + "\n")
		env.Flush()
		fmt.Println("Cliente# Comando enviado:", comando)

		if strings.TrimSpace(comando) == "bye" {
			fmt.Println("Cliente# Desconectando... ")
			salir = false
			socket.Close()
			break
		}

		rec := bufio.NewReader(socket)
		for {
			sResComando, err := rec.ReadString('\n')
			if err != nil {
				fmt.Println("Cliente# Error al recibir la respuesta del comando: ", err)
				salir = false
				socket.Close()
				return
			}
			fmt.Print("Cliente#", sResComando)
			if sResComando == "\n" {
				break
			}
		}
		fmt.Println("-------------------------------------------")
	}
	socket.Close()
}

func RecibeReporteClient(socket *net.Conn, tiempo int) {
	for {
		if salir == false {
			break
		}
		time.Sleep(time.Duration(tiempo) * time.Second)
		rr, _ := bufio.NewReader(*socket).ReadString('\n')
		rr = strings.TrimRight(rr, "\r\n")
		fmt.Println("\nClient# From: ", (*socket).RemoteAddr(), "[", rr, "]")
	}
}
