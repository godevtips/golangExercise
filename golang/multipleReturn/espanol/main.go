package main

import ("fmt"
        "os"
        "strconv")

// En este ejercicio vamos a demostrar cómo implementar una función que devuelve multiples Valores de Retorno
func main(){

  // Este es un simple regresa 2 valores de tipo string
  fmt.Println("\n|------------- Multi retorno ------------------|")
  valor1, valor2 := multiRetorno()
  fmt.Println("Valor 1: "+valor1,"\nValor 2: "+ valor2)
  fmt.Println("|---------------------------------------------|")

   // Esta sección maneja los argumentos que se pasan
   // Para ejecutar esta sección, ingrese la siguiente tarea:
   // 1. Navegue a la carpeta 'multipleReturn' -> español a través del terminal e ingrese el siguiente comando: 'go run main.go 10'
   argument := os.Args
   if len(argument) != 2 {
     fmt.Println("El programa necesita al menos 1 argumento para realizar el calculación")
     return
   }

   argumento, error := strconv.Atoi(argument[1]) //Recupere el valor del argumento en el índice 1 (en índice 0 es la ruta al programa)
   if error != nil {  // En caso de que no haya podido recuperar el argumento el error será captura
     fmt.Println(error) // Mensaje de error
     return
   }

  // Este simple ejemplo demuestra una función de dos valores int en los cuales uno es el valor que se está pasando y el otro se multiplicará por 10
   fmt.Println("\n|------------- Multi retorno de tipo int ------------------|")
   argumento, valorMultiplicar := multiplicarValores(argumento)
   argumentoConvertidoEnString := strconv.Itoa(argumento)
   ValorMultiplicadoConvertidoEnString := strconv.Itoa(valorMultiplicar)

   fmt.Println("Valor del argumento (int): "+argumentoConvertidoEnString)
   fmt.Println("Argumento de multiplicación (int): "+ValorMultiplicadoConvertidoEnString)
   fmt.Println("|-------------------------------------------------|")

   // Este ejemplo demuestra retorno de múltiples valores de diferentes tipos
   fmt.Println("\n|------------ Retorno múltiple de different tipos --------------------|")
   valorDeArgumento, valorMultiplicadoPor10, resultadoDeLaTarea := multiplicarValorDeDifferenteTipos(argumento)
   valorDeArgumentoConvertidoEnString := strconv.Itoa(valorDeArgumento)valorMultiplicadoPor10ConvertidoEnString
   valorMultiplicadoPor10ConvertidoEnString := strconv.Itoa(valorMultiplicadoPor10)

   fmt.Println("Valor del argumento de tipo int : "+valorDeArgumentoConvertidoEnString)
   fmt.Println("Valor del argumento multiplicado y convertido en un tipo int: "+valorMultiplicadoPor10ConvertidoEnString)
   fmt.Println("Resultado de la tarea en tipo string: "+resultadoDeLaTarea)
   fmt.Println("|------------------------------------------------------------------|")
}

// Return 2 valores de tipo string
func multiRetorno() (valor1 string, valor2 string) {
  valor1 = "valores de retorno 1 de tipo string"
  valor2 = "valores de retorno 2 de tipo string"
  return
}

// Multiplica el valor de tipo int por 10
func multiplicarValores(valor int) (valorDeParametro int, valorMultiplicado int) {
  valorDeParametro = valor
  valorMultiplicado = valor * 10
  return
}

// Multiplique el valor de tipo int y devuelva múltiples valores de diferente tipo
func multiplicarValorDeDifferenteTipos(valor int) (int, int, string) {
  return valor, valor * 10, "calculado"
}
