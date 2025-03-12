  // Create a new school struct
type School struct {
    Name string
    Location string
    Students []string
}

func main() {
    // Define a new school
    var mySchool = School{
        Name: "My School",
        Location: "Anytown, USA",
        Students: []string{"Alice", "Bob", "Charlie"},
    }

    // Print the name of the school and its location
    fmt.Printf("The name of the school is %s and it's located in %s\n", mySchool.Name, mySchool.Location)

    // Print the names of the students enrolled in the school
    for _, student := range mySchool.Students {
        fmt.Printf("The student's name is %s\n", student)
    }
}