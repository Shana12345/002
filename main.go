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
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) //To get parameters

	//Looping through every item to see if the item is equal to the params url
	for _, item := range pets {
		if item.Name ==params["name"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Pet{})
}
// Add a pet
func createPet(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	var Apet Pet
	_ = json.NewDecoder(r.Body).Decode(&Apet)
	//Apet.Name = rand.Animal //Generate a random ID -  not safe - would'nt use in production-it could generate the same name.
	pets = append(pets, Apet)
	json.NewEncoder(w).Encode(Apet)
}

//Update Pet
func updatePet(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range pets {
		if item.Name == params["name"] {
		pets = append(pets[:index], pets[index+1:]...)
		var Apet Pet
		_ = json.NewDecoder(r.Body).Decode(&Apet)
		//if you was to have a random ID you want add here "pet.ID = params["id"]"
		pets = append(pets, Apet)
		json.NewEncoder(w).Encode(Apet)
		return
	}
}
		json.NewEncoder(w).Encode(pets)
}

//Remove a pet
func deletePet(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range pets {
		if item.Name == params["name"] {
		pets = append(pets[:index], pets[index+1:]...) // This is like a slice in JS - so if you were to loop through an array and slice it up
		break
	}
}
	json.NewEncoder(w).Encode(pets) //repond with pets - delete book and then give you a response of all the books
}

func main(){
  // init router
  r := mux.NewRouter()

//Mock info - for DB
pets = append(pets, Pet{Name: "Burno", Animal: "puppy", Weight: 42, Age: 1, Owner: &Owner{Firstname: "Charlie", Surname: "Bucket", OwnerPayment: "Family"}})
pets = append(pets, Pet{Name: "Sammy", Animal: "Dog", Weight: 50, Age: 4, Owner: &Owner{Firstname: "Jiminy", Surname: "Cricket", OwnerPayment: "Let your conscience be your guide"}})

// Route Handler and EP
  r.HandleFunc("/api/pets", getPets).Methods("GET")
  r.HandleFunc("/api/pets/{name}", getPet).Methods("GET")
  r.HandleFunc("/api/pets", createPet).Methods("POST")
  r.HandleFunc("/api/pets/{name}", updatePet).Methods("PUT")
  r.HandleFunc("/api/pets/{name}", deletePet).Methods("DELETE")

  log.Fatal(http.ListenAndServe(":8000", r))
}
