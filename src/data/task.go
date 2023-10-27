package data

type Task struct {
	Id      int    `gorm:"int;primaryKey"`
	Content string `gorm:"text;notnull;index"`
	Done    int    `gorm:"int;notnull;default:0;index"`
}

func CreateTask(content string) (Task, error) {
	task := Task{Content: content}
	result := db.Create(&task)
	return task, result.Error
}

func GetTasks() ([]Task, error) {
	var tasks []Task
	result := db.Find(&tasks)
	return tasks, result.Error
}

func GetTask(id int) (Task, error) {
	var task Task
	result := db.First(&task, id)
	return task, result.Error
}

func MarkTaskDone(id int) error {
	result := db.Model(&Task{}).Where("id = ?", id).Update("done", 1)
	return result.Error
}

func MarkTaskTodo(id int) error {
	result := db.Model(&Task{}).Where("id = ?", id).Update("done", 0)
	return result.Error
}

func DeleteTask(id int) error {
	result := db.Delete(&Task{}, id)
	return result.Error
}
