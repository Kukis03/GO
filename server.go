package funciones

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
)

var socket net.Conn
var login = false

func IniciarServidor() {
	fmt.Println("=====================================================")
	fmt.Println("||              Servidor Comandos 	           ||")
	fmt.Println("=====================================================")

	tcpAddress, _ := net.ResolveTCPAddr("tcp4", ":1306")
	socketServer, _ := net.ListenTCP("tcp", tcpAddress)

	fmt.Println("Servidor# Esperando Conexion...")
	socket, _ = socketServer.Accept()
	fmt.Println("Servidor# Autenticando cliente...", socket.RemoteAddr())

	reader := bufio.NewReader(socket)
	usr, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}

	psw, errr := reader.ReadString('\n')
	if errr != nil {
		fmt.Println(err)
	}

	timeReport, errr := reader.ReadString('\n')
	if errr != nil {
		fmt.Println(err)
	}

	usr = strings.TrimSpace(usr)
	psw = strings.TrimSpace(psw)
	timeReport = strings.TrimSpace(timeReport)
	reporte, _ := strconv.Atoi(timeReport)

	archivo, _ := ioutil.ReadFile("users.pw")
	sArchivo := string(archivo)
	credenciales := strings.Split(sArchivo, ":")

	for i := 0; i < len(credenciales)-1; i++ {
		if credenciales[i] == usr && credenciales[i+1] == psw {
			login = true
			break
		}
	}

	if login {
		fmt.Println("Servidor# Cliente autenticado correctamente....")
		go EnviarReporteServer(&socket, reporte)
	} else {
		fmt.Println("Servidor# El cliente no se ha autenticado de manera correcta....")
		socket.Close()
	}

	for {

		rec := bufio.NewReader(socket)
		comando, err := rec.ReadString('\n')
		if strings.TrimSpace(comando) == "bye" {
			fmt.Println("Servidor# Cliente desconectado...")
			login = false
			socket.Close()
			return
		}

		if err != nil {
			fmt.Println("Servidor# Error al recibir el comando: ", err)
			login = false
			socket.Close()
			return
		}

		fmt.Println("Servidor# Comando recibido:", comando)

		array_datoIn := strings.Fields(comando)
		shell := exec.Command(array_datoIn[0], array_datoIn[1:]...)
		resComando, err := shell.Output()
		if err != nil {
			fmt.Println("Servidor# Error al ejecutar el comando:", err)
			login = false
			socket.Close()
			return
		}

		sResComando := string(resComando)
		fmt.Println("Servidor# Ejecutando comando recibido")

		env := bufio.NewWriter(socket)
		env.WriteString(sResComando + "\n")
		err = env.Flush()
		if err != nil {
			fmt.Println("Servidor# Error al enviar la respuesta del comando:", err)
			login = false
			socket.Close()
			return
		}
		fmt.Println("Servidor# Respuesta de ejecución enviada!")
	}
}

func EnviarReporteServer(socketS *net.Conn, tiempo int) {
	x := 0
	for {
		if !login {
			socket.Close()
			break
		}
		time.Sleep(time.Duration(tiempo) * time.Second)
		x++

		memInfo, _ := mem.VirtualMemory()
		repMem := int(memInfo.UsedPercent)

		diskInfo, _ := disk.Usage("/")
		repDisk := int(diskInfo.UsedPercent)

		cpuPercent, _ := cpu.Percent(0, false)
		repCpu := int(cpuPercent[0])

		msgReporte := fmt.Sprintf("*** Report #%d - [Procesador: %d%%][Memory: %d%%] [Disk: %d%%]", x, repCpu, repMem, repDisk)
		envRep := bufio.NewWriter(*socketS)
		envRep.WriteString(msgReporte + "\n")
		envRep.Flush()
		fmt.Println("Servidor# Reporte #", x, "Enviado a ", (*socketS).RemoteAddr())
	}

}
