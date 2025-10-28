package utilidades

import (
	"fmt"
	"io"
	"log"
	"time"

	pb "servidor-streaming/serviciosStreaming"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

// DecodificarReproducir decodifica y reproduce el audio MP3 en tiempo real
func DecodificarReproducir(reader io.Reader, canalSincronizacion chan struct{}) {
	streamer, format, err := mp3.Decode(io.NopCloser(reader))
	if err != nil {
		log.Printf("Error decodificando MP3: %v", err)
		close(canalSincronizacion)
		return
	}
	defer streamer.Close()

	// Inicializar el speaker con la tasa de muestreo del formato
	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/2))

	//fmt.Println("Reproduciendo audio...")

	// Reproducir el audio y notificar cuando termine
	speaker.Play(beep.Seq(streamer, beep.Callback(func() {
		fmt.Println("\n Reproducción de audio completada")
		close(canalSincronizacion)
	})))
}

// RecibirCancion recibe fragmentos del servidor y los escribe en el pipe
func RecibirCancion(stream pb.AudioService_EnviarCancionMedianteStreamClient, writer *io.PipeWriter, canalSincronizacion chan struct{}) {
	noFragmento := 0
	totalBytes := 0

	for {
		fragmento, err := stream.Recv()
		if err == io.EOF {
			fmt.Println("\n Todos los fragmentos recibidos")
			writer.Close()
			break
		}
		if err != nil {
			log.Printf(" Error recibiendo fragmento: %v", err)
			writer.Close()
			close(canalSincronizacion)
			return
		}

		noFragmento++
		totalBytes += len(fragmento.Data)

		// Mostrar progreso
		fmt.Printf("\r🎵 Fragmento #%d recibido (%d bytes) | Total: %d KB",
			noFragmento, len(fragmento.Data), totalBytes/1024)

		// Escribir datos en el pipe para que sean reproducidos
		if _, err := writer.Write(fragmento.Data); err != nil {
			log.Printf(" Error escribiendo en pipe: %v", err)
			break
		}
	}

	// Esperar hasta que termine la reproducción
	<-canalSincronizacion
	fmt.Println("\n Streaming finalizado")
}
