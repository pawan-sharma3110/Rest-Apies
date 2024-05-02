package handlers

import (
	"bank/models"
	"encoding/json"
	"net/http"
	"strconv"
)

var allCustomer = []models.Customer{

	{ID: 1,
		AccountNumber: 5678,
		Name:          "Jane Smith",
		DOB:           "12-25-1985",
		Address:       "456 Elm St",
	},

	{ID: 2,
		AccountNumber: 9876,
		Name:          "Alex Johnson",
		DOB:           "07-10-1978",
		Address:       "789 Oak Ave",
	},

	{ID: 3,
		AccountNumber: 2468,
		Name:          "Emily Brown",
		DOB:           "03-30-1993",
		Address:       "101 Pine Ln",
	},

	{ID: 4,
		AccountNumber: 1357,
		Name:          "Michael Davis",
		DOB:           "09-05-1980",
		Address:       "222 Cedar Rd",
	},

	{ID: 5,
		AccountNumber: 8642,
		Name:          "Sarah Wilson",
		DOB:           "11-20-1975",
		Address:       "333 Maple Blvd",
	},

	{ID: 6,
		AccountNumber: 7321,
		Name:          "Christopher Lee",
		DOB:           "06-12-1998",
		Address:       "444 Walnut Dr",
	},

	{ID: 7,
		AccountNumber: 9513,
		Name:          "Jessica Garcia",
		DOB:           "02-18-1987",
		Address:       "555 Birch St",
	},

	{ID: 8,
		AccountNumber: 3698,
		Name:          "David Martinez",
		DOB:           "08-08-1972",
		Address:       "666 Oakwood Ave",
	},

	{ID: 9,
		AccountNumber: 1470,
		Name:          "Amanda Thomas",
		DOB:           "04-01-1990",
		Address:       "777 Elm St",
	},

	{ID: 10,
		AccountNumber: 2589,
		Name:          "Ryan Clark",
		DOB:           "10-15-1983",
		Address:       "888 Cedar Rd",
	},
}

func CreatCustomer(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var new models.Customer
		json.NewDecoder(r.Body).Decode(&new)
		allCustomer = append(allCustomer, new)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(new)

	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"message": "Customer not created"})
	}
}
func AllCustomer(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		w.Header().Set("Content-Type", "aplication/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(allCustomer)
	} else {
		w.Header().Set("Content-Type", "aplication/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode([]string{"Only get reqest allowed"})
	}
}

func SearchById(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		paramsId := r.PathValue("id")
		idExists := false
		id, _ := strconv.Atoi(paramsId)
		for _, v := range allCustomer {
			if v.ID == id {
				idExists = true
				w.Header().Set("Content-Type", "aplication/json")
				w.WriteHeader(http.StatusOK)
				json.NewEncoder(w).Encode(v)
				break
			}

		}
		if !idExists {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(map[string]string{"message": "This id does not exist."})
		}

	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"message": "Customer not found,Please make an get request."})
	}

}
func DeleteById(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodDelete {
		paramsId := r.PathValue("id")
		id, _ := strconv.Atoi(paramsId)
		idExists := false
		for i, v := range allCustomer {
			if v.ID == id {
				idExists = true
				allCustomer = append(allCustomer[:i], allCustomer[i+1:]...)
				json.NewEncoder(w).Encode(map[string]string{"message": "Deleted"})
				w.WriteHeader(http.StatusOK)
				w.Header().Set("Request/Methord", "Delete recquired")
			}

		}
		if !idExists {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(map[string]string{"message": "This id does not exist."})
		}

	} else {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Request/Methord", "Delete recquired")
		json.NewEncoder(w).Encode("Delete request required only")
	}
}

func UpdateById(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPatch {
		paramsId := r.PathValue("id")
		id, _ := strconv.Atoi(paramsId)
		idExists := false
		for i, v := range allCustomer {
			if v.ID == id {
				idExists = true
				var updatedData models.Customer
				w.WriteHeader(http.StatusFound)
				w.Header().Set("Content-Type", "application/json")
				json.NewDecoder(r.Body).Decode(&updatedData)
				allCustomer[i].Name = updatedData.Name
				allCustomer[i].DOB = updatedData.DOB
				allCustomer[i].Address = updatedData.Address
				json.NewEncoder(w).Encode(allCustomer[i])

				return
			}

		}
		if !idExists {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(map[string]string{"message": "This id does not exist."})
		}

	} else {
		w.Header().Set("Content-Type", "aplication/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Patchrequest required only")
	}
}
