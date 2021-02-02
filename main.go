package main
import "encoding/json"
import "log"
import "net/http"
import "github.com/gorilla/mux"

//init
var pets []Pet

//Struct for Pet
type Pet struct {
	Name string `json:"name"`
  Animal string `json:"animal"`
  Weight float64 `json:"weight"`
	Age int `json:"age"`
  Owner *Owner `json:"owner"`
}

//Struct for OwnerPayment
type Owner struct {
  Firstname string `json:"firstname"`
  Surname string `json:"surname"`
  OwnerPayment string `json:"-"`
}

// Get every pets
func getPets(w http.ResponseWriter, r *http.Request){
  w.Header().Set("Content-Type", "application/json")
  json.NewEncoder(w).Encode(pets)
}

//Get a pet
func getPet(w http.ResponseWriter, r *http.Request){

}
// Add a pet
func createPet(w http.ResponseWriter, r *http.Request){

}

//Update Pet
func updatePet(w http.ResponseWriter, r *http.Request){

}

//Remove a pet
func deletePet(w http.ResponseWriter, r *http.Request){

}

func main(){
  // init router
  r := mux.NewRouter()

//Mock info - for DB
pets = append(pets, Pet{Name: "Burno", Animal: "puppy", Weight: 42, Age: 1, Owner: &Owner{Firstname: "Charlie", Surname: "Bucket", OwnerPayment: "it's totally a secret"}})

// Route Handler and EP
  r.HandleFunc("/api/pets", getPets).Methods("GET")
  r.HandleFunc("/api/pets/{name}", getPet).Methods("GET")
  r.HandleFunc("/api/pets", createPet).Methods("POST")
  r.HandleFunc("/api/pets/{name}", updatePet).Methods("PUT")
  r.HandleFunc("/api/pets/{name}", deletePet).Methods("DELETE")

  log.Fatal(http.ListenAndServe(":8000", r))
}
