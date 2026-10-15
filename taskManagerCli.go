package main
import (
	"encoding/json"
	"os"
	"io"
	"time"
	"fmt"
)

type TastStatus int

const(
	Todo TastStatus=iota
	InProgress
	Done
)

func (s TastStatus) String() string{
	switch s{
		case Todo:
			return "todo"
		case InProgress:
			return "in-progress"
		case Done:
			return "done"
		default:
			return "processing"
	}
}

// 4. Custom JSON Marshaler so it exports as text ("todo") instead of an int (0)
func (s TastStatus) MarshalJson() ([]byte, error){
	return json.Marshal(s.String())
}

type Task struct{
	ID    int     `json:"id"`
	Description    string     `json:"id"`
	Status    TastStatus     `json:"status"`
	CreatedAt    string     `json:"id"`
	UpdatedAT    string     `json:"id"`
}

func addData(task Task) (bool, error){
	jsondata, err := json.MarshalIndent(task, "", "	")
	if err != nil{
		fmt.Println("Error: %v Occured", err)
		return false, err
	}
	// Write JSON bytes to a local file named "task.json"
	// 0644 sets standard file read/write permissions
	err = os.WriteFile("task.json", jsondata,0644)
	if err != nil{
		fmt.Println("Error: %v Occured", err)
		return false, err
	}

	fmt.Println("Task Added Sucefully!")
	return true, err
}

func updateTasK(taskId int) (Task, bool, err){
	file, err := os.Open("task.json")
	if err != nil{
		fmt.Println("Error: %v Occured", err)
		return Task, false, err
	}
	defer file.close()
	decoder := json.NewDecoder(file)
	t, err := decoder.Token()
	if err != nil{
		fmt.Println("Error: %v Occured", err)
		return Task, false, err
	}
	//search the data
	for decoder.more() {
		var task Task
		err := decoder.Decode(&task)
		if err != nil{
			fmt.Println("Error: %v Occured", err)
			return task, false, err
		}
		if task.id == taskId{
			return task, true, err
		}
	}
}

func main() {

}
