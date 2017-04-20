package main

import "fmt"

var currentId int

var estates Estates

// Give us some seed data
func init() {
    RepoCreateEstate(Estate { Name: "Departamento en Palermo", Lat: -34.594142, Long: -58.422036 })
    RepoCreateEstate(Estate { Name: "Casa en San Isidro", Lat: -34.467610, Long: -58.510191 })
}

func RepoFindEstate(id int) Estate {
    for _, t := range estates {
        if t.Id == id {
            return t
        }
    }
    // return empty Estate if not found
    return Estate{}
}

func RepoCreateEstate(e Estate) Estate {
    currentId += 1
    e.Id = currentId
    estates = append(estates, e)
    return e
}

func RepoDestroyEstate(id int) error {
    for i, e := range Estates {
        if e.Id == id {
            estates = append(estates[:i], estates[i+1:]...)
            return nil
        }
    }
    return fmt.Errorf("Could not find Estate with id of %d to delete", id)
}
