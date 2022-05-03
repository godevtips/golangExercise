package main

import ("fmt"
        "os"
        "strconv")

// In dit voorbeeld gaan we een functie implementeren die meerdere waarden retourneert
func main(){

  // Dit is een voorbeeld waarbij twee strings waardes worden terug verstuurd
  fmt.Println("\n|------------- Twee string waardes terug sturen ------------------|")
  waarde1, waarde2 := tweeStringWaardesRetouneren()
  fmt.Println("Waarde 1: "+waarde1,"\nWaarde 2: "+ waarde2)
  fmt.Println("|-------------------------------------------------------------------|")

   // Deze onderdeel vang de argumenten die doorgegeven worden op
   // Deze onderdeel wordt uitgevoerd door de volgende taken uit te voeren:
   // 1. Navigeer naar de map 'multipleReturn' -> nederlands via de terminal en voer de volgende opdracht uit: 'go run main.go 10'
   argument := os.Args
   if len(argument) != 2 {
     fmt.Println("Het programma heeft minimaal 1 argument nodig om de berekening uit te voeren")
     return
   }

  // value
   argumentWaarde, error := strconv.Atoi(argument[1]) // Haal argument waarde in index 1 op (index 0 is het pad naar het programma)
   if error != nil {  // In het geval dat het niet lukt vang de error bericht op
     fmt.Println(error) // Error bericht
     return
   }

   // Dit voorbeeld is een voorbeeld van een functie waarmee twee int-waardes worden terug gestuurd. De ene waarde wordt doorgegeven en de andere wordt vermenigvuldigd met 10.
   fmt.Println("\n|------------- Twee int waardes retouneren ------------------|")
   doorgegevenWaarde, waardeVermenigVuldigenMet10 := waardeVermenigvuldigen(argumentWaarde)

   doorgegevenWaardeOmgezetNaarString := strconv.Itoa(doorgegevenWaarde)
   waardeVermenigVuldigenMet10OmgezetNaarString := strconv.Itoa(waardeVermenigVuldigenMet10)

   fmt.Println("Argument waarde (int): "+ doorgegevenWaardeOmgezetNaarString)
   fmt.Println("Vermenigvuldig waarde (int): "+waardeVermenigVuldigenMet10OmgezetNaarString)
   fmt.Println("|-------------------------------------------------|")

   // Dit voorbeeld demonstreert hoe een functie meerdere retourwaarden van verschillende typen kan retourneren
   fmt.Println("\n|------------ Waardes retouren van verschillende types --------------------|")

  //doorgegevenWaarde
   doorgegevenWaardes, doorgegevenWaardeVermenigvuldigMet10, taakResultaat := waardesRetounerenVanVerschillendeTypes(argumentWaarde)
   doorgegevenWaardesOmgezetNaarString := strconv.Itoa(doorgegevenWaardes)
   vermenigvuldigWaardeOmgezetNaarToString := strconv.Itoa(doorgegevenWaardeVermenigvuldigMet10)

   fmt.Println("Doorgegeven int waarde : "+doorgegevenWaardesOmgezetNaarString)
   fmt.Println("Vermenigvuldigen int waarde : "+vermenigvuldigWaardeOmgezetNaarToString)
   fmt.Println("Taak resultaat string waarde: "+taakResultaat)
   fmt.Println("|------------------------------------------------------------------|")
}

// Twee string retourneren
func tweeStringWaardesRetouneren() (waarde1 string, waarde2 string) {
  waarde1 = "waarde 1"
  waarde2 = "waarde 2"
  return
}

// Vermenigvuldig de doorgegeven int waardes met 10
func waardeVermenigvuldigen(waarde int) (parameter int, parameterVermenigVuldigenMet10 int) {
  parameter = waarde
  parameterVermenigVuldigenMet10 = waarde * 10
  return
}

// waardesRetounerenVanVerschillendeTypes
func waardesRetounerenVanVerschillendeTypes(waarde int) (int, int, string) {
  return waarde, waarde * 10, "Berekend"
}
