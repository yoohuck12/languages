package main

import "fmt"

func main() {
    //    var x map[string]int
    //    x = make(map[string]int)
    x := make(map[string]int)
    x["key"] = 10
    x["key2"] = 20
    fmt.Println(x)  // map[key:10 key2:20]
    fmt.Println(x["key"])

    name, ok := x["OK"]
    fmt.Println(name, ok) // 0, false

    if name, ok = x["OK"]; ok {
        fmt.Println(name, ok)  //출력안됨
    }

    elements := map[string]string{
        "H": "Hydrogen",
        "He": "Helium",
        "Li": "Lithium",
        "Be": "Beryllium",
        "B": "Boron",
        "C": "Carbon",
        "N": "Nitrogen",
        "O": "Oxygen",
        "F": "Fluorine",
        "Ne": "Neon",
    }
    fmt.Println(elements)


    elements2 := map[string]map[string]string{
        "H": map[string]string{
            "name":"Hydrogen",
            "state":"gas",
        },
        "He": map[string]string{
            "name":"Helium",
            "state":"gas",
        },
        "Li": map[string]string{
            "name":"Lithium",
            "state":"solid",
        },
        "Be": map[string]string{
            "name":"Beryllium",
            "state":"solid",
        },
        "B":  map[string]string{
            "name":"Boron",
            "state":"solid",
        },
        "C":  map[string]string{
            "name":"Carbon",
            "state":"solid",
        },
        "N":  map[string]string{
            "name":"Nitrogen",
            "state":"gas",
        },
        "O":  map[string]string{
            "name":"Oxygen",
            "state":"gas",
        },
        "F":  map[string]string{
            "name":"Fluorine",
            "state":"gas",
        },
        "Ne":  map[string]string{
            "name":"Neon",
            "state":"gas",
        },
    }

    if el, ok := elements2["Li"]; ok {
        fmt.Println(el["name"], el["state"])
    }
}
