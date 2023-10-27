package api

import (
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/vcoltfre/godo/src/data"
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

func Start(bind string) error {
	if err := data.Connect(); err != nil {
		return err
	}

	logrus.Info("Database connected")

	e := echo.New()

	e.HideBanner = true
	e.HidePort = true

	e.POST("/tasks", createTask)

	logrus.Info("Starting API on " + bind)

	return e.Start(bind)
}
