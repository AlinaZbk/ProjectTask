package main
import (
    "bufio"
    "fmt"
    "os"
	"strings"
)

const (
    Create = "create"
    Read   = "read"
    Update = "update"
    Delete = "delete"
    Exit = "exit"
    Help = "help"
    FileName = "tasks.txt"
)

func main() {    
    tasks := loadTasks()
	reader := bufio.NewReader(os.Stdin)
    fmt.Println("Welcome to the TO DO List CLI app!")
    fmt.Println("Enter 'help' to view the list of available commands")
    for {
        fmt.Println("\nEnter your command (create, read, update, delete, help, exit): ")
        command, _ := reader.ReadString('\r')
		command = strings.TrimSpace(command) 
        switch command {
        case Help:
            showHelp()
        case Create:
            createTask(tasks, reader)
        case Read:
            readTasks()
        case Update:
            updateTask(tasks, reader)
        case Delete:
            deleteTask(tasks, reader)
        case Exit:
            fmt.Println("Exit the program. Goodbye!")
            clearFile()
            return
        default:
            fmt.Println("Invalid command! Please, try again!")
        }
    }
}

func createTask(tasks []string, reader *bufio.Reader){
    fmt.Println("Enter task name: ")
    newTask, _ := reader.ReadString('\r')
    newTask = strings.TrimSpace(newTask)  // remove '\n' at the end
    tasks = append(tasks, newTask)
    fmt.Println("Task added")
    saveTasks(tasks)
}

func readTasks() {
    tasks := loadTasks()
    if len(tasks) == 0 {
        fmt.Println("The task list is empty")
    } else {
        fmt.Println("Your tasks: ")
        for i, task := range tasks {
            fmt.Printf("%d. %s\n", i+1, task)
        }
    }
}

func updateTask(tasks []string, reader *bufio.Reader) {
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
    }
    fmt.Println("Enter new task name: ")
    newTaskName, _ := reader.ReadString('\r')
    newTaskName = strings.TrimSpace(newTaskName)  // remove '\n' at the end
    if len(newTaskName) < 3 {
        fmt.Println("The new task name is too short! Please, try again.")
    }
    tasks[indexToUpdate] = newTaskName
    saveTasks(tasks)
    fmt.Printf("Updated task #%d with name \"%s\" successfully!\n", indexToUpdate+1, newTaskName)
}

func deleteTask(tasks []string, reader *bufio.Reader) {
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
    }
    oldTaskName := tasks[indexToRemove]
    tasks = append(tasks[:indexToRemove], tasks[indexToRemove+1:]...)
    saveTasks(tasks)
    fmt.Printf("Removed task #%d with name \"%s\" successfully!\n", indexToRemove+1, oldTaskName)
}

func showHelp() {
    fmt.Println("Tasks list:")
    fmt.Println("create - add a new task")
    fmt.Println("read - show a list of tasks")
    fmt.Println("update - edit an existing task")
    fmt.Println("delete - delete a task from the list")
    fmt.Println("exit - exit the program")
}

func loadTasks() []string {
    file, err := os.Open(FileName)
    if err != nil {
        if os.IsNotExist(err) {
            os.Create(FileName)
            return []string{}
        }
        fmt.Println("File opening error: ", err)
        return []string{}
    }
    defer file.Close()
    var tasks []string
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        tasks = append(tasks, scanner.Text())
    }
    return tasks
}

func saveTasks(tasks []string) {
    file, err := os.OpenFile(FileName, os.O_APPEND|os.O_WRONLY, 0644)
    if err != nil {
        fmt.Println("Error saving to a file: ", err)
        return
    }
    defer file.Close()
    for _, task := range tasks {
        _, err := file.WriteString(task + "\n")
        if err != nil {
            fmt.Println("Error writing to a file: ", err)
            return
        }
    }
}

func clearFile() {
    file, err := os.Create(FileName)
    if err != nil {
        fmt.Println("File cleaning error: ", err)
        return
    }
    defer file.Close()
}