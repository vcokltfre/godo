package api

import (
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/vcoltfre/godo/src/data"
	"github.com/vcoltfre/godo/src/ui"
)

type createTaskData struct {
	Content string `json:"content"`
}

type taskData struct {
	Id      int    `json:"id"`
	Content string `json:"content"`
	Done    bool   `json:"done"`
}

func convertTask(task data.Task) taskData {
	return taskData{
		Id:      task.Id,
		Content: task.Content,
		Done:    task.Done == 1,
	}
}

func createTask(c echo.Context) error {
	d := &createTaskData{}
	if err := c.Bind(d); err != nil {
		logrus.WithError(err).Error("Failed to bind request body")
		return echo.NewHTTPError(400, "Invalid request body")
	}

	task, err := data.CreateTask(d.Content)
	if err != nil {
		logrus.WithError(err).Error("Failed to create task")
		return echo.NewHTTPError(500, "Failed to create task")
	}

	return c.JSON(200, convertTask(task))
}

func getTasks(c echo.Context) error {
	tasks, err := data.GetTasks()
	if err != nil {
		logrus.WithError(err).Error("Failed to get tasks")
		return echo.NewHTTPError(500, "Failed to get tasks")
	}

	var data []taskData
	for _, task := range tasks {
		data = append(data, convertTask(task))
	}

	return c.JSON(200, data)
}

func getTask(c echo.Context) error {
	id := c.Param("id")

	parsedId, err := strconv.Atoi(id)
	if err != nil {
		logrus.WithError(err).Error("Failed to parse ID")
		return echo.NewHTTPError(400, "Invalid ID")
	}

	task, err := data.GetTask(parsedId)
	if err != nil {
		if err.Error() == "record not found" {
			return echo.NewHTTPError(404, "Task not found")
		}

		logrus.WithError(err).Error("Failed to get task")
		return echo.NewHTTPError(500, "Failed to get task")
	}

	return c.JSON(200, convertTask(task))
}

func markTaskDone(c echo.Context) error {
	id := c.Param("id")

	parsedId, err := strconv.Atoi(id)
	if err != nil {
		logrus.WithError(err).Error("Failed to parse ID")
		return echo.NewHTTPError(400, "Invalid ID")
	}

	if err := data.MarkTaskDone(parsedId); err != nil {
		logrus.WithError(err).Error("Failed to mark task as done")
		return echo.NewHTTPError(500, "Failed to mark task as done")
	}

	return c.NoContent(200)
}

func markTaskTodo(c echo.Context) error {
	id := c.Param("id")

	parsedId, err := strconv.Atoi(id)
	if err != nil {
		logrus.WithError(err).Error("Failed to parse ID")
		return echo.NewHTTPError(400, "Invalid ID")
	}

	if err := data.MarkTaskTodo(parsedId); err != nil {
		logrus.WithError(err).Error("Failed to mark task as todo")
		return echo.NewHTTPError(500, "Failed to mark task as todo")
	}

	return c.NoContent(200)
}

func deleteTask(c echo.Context) error {
	id := c.Param("id")

	parsedId, err := strconv.Atoi(id)
	if err != nil {
		logrus.WithError(err).Error("Failed to parse ID")
		return echo.NewHTTPError(400, "Invalid ID")
	}

	if err := data.DeleteTask(parsedId); err != nil {
		logrus.WithError(err).Error("Failed to delete task")
		return echo.NewHTTPError(500, "Failed to delete task")
	}

	return c.NoContent(200)
}

func Start(bind string) error {
	if err := data.Connect(); err != nil {
		return err
	}

	logrus.Info("Database connected")

	e := echo.New()

	e.HideBanner = true
	e.HidePort = true

	e.GET("/", ui.GetIndex)

	e.POST("/tasks", createTask)
	e.GET("/tasks", getTasks)
	e.GET("/tasks/:id", getTask)

	e.POST("/tasks/:id/done", markTaskDone)
	e.DELETE("/tasks/:id/done", markTaskTodo)

	e.DELETE("/tasks/:id", deleteTask)

	logrus.Info("Starting API on " + bind)

	return e.Start(bind)
}
