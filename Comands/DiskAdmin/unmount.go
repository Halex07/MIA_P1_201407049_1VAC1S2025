package Comandos

import (
	Herramientas "MIA_P1_201407049/Analisis"
	"MIA_P1_201407049/Structs"
	"fmt"
	"strings"
)

func Unmount(parametros []string) {
	fmt.Println(">>Ejecutando Comando UNMOUNT")
	var id string
	paramOk := true
	temp2 := strings.TrimRight(parametros[1], " ")
	temp := strings.Split(temp2, "=")

	if len(temp) != 2 {
		fmt.Println("Error valor de parametro ", temp[0], " no reconocido")
		paramOk = false
		return
	}

	if strings.ToLower(temp[0]) == "id" {
		id = strings.ToUpper(temp[1])
	} else {
		fmt.Println("Error parametro ", temp[0], "no reconocido")
		paramOk = false
	}
	if paramOk {
		disk := id[0:1]
		folder := "./MIA/P1/"
		ext := ".dsk"
		dirDisk := folder + disk + ext
		disco, err := Herramientas.OpenFile(dirDisk)
		if err != nil {
			return
		}

		var mbr Structs.MBR
		if err := Herramientas.ReadObj(disco, &mbr, 0); err != nil {
			return
		}

		defer disco.Close()

		unmount := true
		for i := 0; i < 4; i++ {
			identificador := Structs.GETID(string(mbr.Partitions[i].Id[:]))
			if identificador == id {
				unmount = false
				name := Structs.GETNOM(string(mbr.Partitions[i].Name[:]))
				var unmount Structs.Partition
				mbr.Partitions[i].Id = unmount.Id
				copy(mbr.Partitions[i].Status[:], "I")
				if err := Herramientas.WrObj(disco, mbr, 0); err != nil {
					return
				}
				fmt.Println("La partición", name, " ha sido desmontada de manera exitosa")
				break
			}
		}

		if unmount {
			fmt.Println("Error al montar la partición ", id, "vuelva a intentar")
			fmt.Println("UNMOUNT Error. No existe el id")
		} else {
			fmt.Println("\nLista de particiones montadas en el sistema\n ")
			for i := 0; i < 4; i++ {
				estado := string(mbr.Partitions[i].Status[:])
				if estado == "A" {
					fmt.Printf("Partition %d: name: %s, status: %s, id: %s, tipo: %s, start: %d, size: %d, fit: %s, correlativo: %d\n", i, string(mbr.Partitions[i].Name[:]), string(mbr.Partitions[i].Status[:]), string(mbr.Partitions[i].Id[:]), string(mbr.Partitions[i].Type[:]), mbr.Partitions[i].Start, mbr.Partitions[i].Size, string(mbr.Partitions[i].Fit[:]), mbr.Partitions[i].Correlativo)
				}
			}
		}
	}
}
