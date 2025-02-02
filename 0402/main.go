package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	//conectar al servidor SMTP de Mailhog en el puerto 1025
	conn, err := net.Dial("tcp", "localhost:1025")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	//Crear un escritor para la conexión
	writer := bufio.NewWriter(conn)

	//Leer la respuesta del servidor
	reader := bufio.NewReader(conn)

	//Leer la bienvenida del servidor
	responde, _ := reader.ReadString('\n')
	fmt.Println(responde) // 220 mailhog at your service

	//Enviar el comando HELO
	fmt.Fprintf(writer, "HELO localhost\r\n")
	writer.Flush()
	responde, _ = reader.ReadString('\n')
	fmt.Println(responde) // 250 mailhog at your service

	//Enviar MAIL FROM (Remitente)
	from := "alumno@fempa.local"
	fmt.Fprintf(writer, "MAIL FROM: <%s>\r\n", from)
	writer.Flush()
	responde, _ = reader.ReadString('\n')
	fmt.Println(responde) // 250 Sender OK

	//Enviar RCPT TO (Destinatario)
	to := "demo@fempa.local"
	fmt.Fprintf(writer, "RCPT TO: <%s>\r\n", to)
	writer.Flush()
	responde, _ = reader.ReadString('\n')
	fmt.Println(responde) // 250 Recipient OK

	//Iniciar la transmisión del cuerpo del mensaje con DATA
	fmt.Fprintf(writer, "DATA\r\n")
	writer.Flush()
	responde, _ = reader.ReadString('\n')
	fmt.Println(responde) // 354 Start mail input; end with <CRLF>.<CRLF>

	//Enviar el contenido del correo (asunto, cuerpo)
	subject := "Prueba de envio de correo: PSP"
	contentType := "Content-Type: text/html; charset=\"UTF-8\"" // Content-Type: text/plain; charset="UTF-8"
	body := `<html>
	<body>
		<h1>Hola, esto es un correo de prueba</h1>
		<p>Este es un mensaje enviado desde <b>MailHog</b>.</p>
	</body>
	</html>`
	fmt.Fprintf(writer, "Subject: %s\r\n%s\r\n\r\n%s\r\n.\r\n", subject, contentType, body)
	writer.Flush()
	responde, _ = reader.ReadString('\n')
	fmt.Println(responde) // 250 OK:  Queued as 12345

	//Cerrar la sesión SMTP con el comando QUIT
	fmt.Fprintf(writer, "QUIT\r\n")
	writer.Flush()
	responde, _ = reader.ReadString('\n')
	fmt.Println(responde) // 221 bye
}
