package job_handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rattapon001/porter-management-demo/internal/job/app"
	"github.com/rattapon001/porter-management-demo/internal/job/domain"
)

type JobHandler struct {
	JobUsecase app.JobServicePort
}

func NewJobHandler(jobUsecase app.JobServicePort) *JobHandler {
	return &JobHandler{
		JobUsecase: jobUsecase,
	}
}

func (h *JobHandler) CreateJob(c *gin.Context) {
	var job domain.Job
	if err := c.ShouldBindJSON(&job); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	location := domain.Location{
		From: job.Location.From,
		To:   job.Location.To,
	}
	patient := domain.Patient{
		Name: job.Patient.Name,
		HN:   job.Patient.HN,
	}
	newJob, err := h.JobUsecase.CreateJob(location, patient)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, newJob)

}
