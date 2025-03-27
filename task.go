package projectTask

import (
    "bufio"
    "fmt"
    "os"
	"strings"
)

func main() {
    const (
        Create = "create"
        Read   = "read"
        Update = "update"
        Delete = "delete"
    )
	reader := bufio.NewReader(os.Stdin)
    tasks := make([]string, 0)

    fmt.Println("Welcome to the TO DO List CLI app!")

    for {
        fmt.Println()
        fmt.Println("Enter your command (create, read, update, delete): ")
        command, _ := reader.ReadString('\r')
		command = strings.TrimSpace(command) 
        switch command {
        case Create:
            fmt.Println("Enter task name: ")
            newTask, _ := reader.ReadString('\r')
            newTask = strings.TrimSpace(newTask)  // remove '\n' at the end

            tasks = append(tasks, newTask)

        case Read:
            for i, task := range tasks {
                fmt.Printf("%d. %s\n", i+1, task)
            }

        case Update:
            fmt.Println("Enter task name to update: ")
            input, _ := reader.ReadString('\r')
            input = strings.TrimSpace(input)  // remove '\n' at the end

            indexToUpdate := -1
            for i, task := range tasks {
                if task == input {
                    indexToUpdate = i
                    break
                }
            }

            if indexToUpdate == -1 {
                fmt.Println("Invalid name. Please try again.")
                continue
            }

            fmt.Println("Enter new task name: ")
            newTaskName, _ := reader.ReadString('\r')
            newTaskName = strings.TrimSpace(newTaskName)  // remove '\n' at the end

            if len(newTaskName) < 3 {
                fmt.Println("The new task name is too short! Please, try again.")
                continue
            }

            tasks[indexToUpdate] = newTaskName
            fmt.Printf("Updated task #%d with name \"%s\" successfully!\n", indexToUpdate+1, newTaskName)

        case Delete:
            fmt.Println("Enter task name to remove: ")
            input, _ := reader.ReadString('\r')
            input = strings.TrimSpace(input) // remove '\n' at the end

            indexToRemove := -1
            for i, task := range tasks {
                if task == input {
                    indexToRemove = i
                    break
                }
            }

            if indexToRemove == -1 {
                fmt.Println("Invalid name. Please try again.")
                continue
            }

            oldTaskName := tasks[indexToRemove]
            tasks = append(tasks[:indexToRemove], tasks[indexToRemove+1:]...)
            fmt.Printf("Removed task #%d with name \"%s\" successfully!\n", indexToRemove+1, oldTaskName)
        default:
            fmt.Println("Invalid command! Please, try again!")
        }
    }
}
