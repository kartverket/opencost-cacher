package rest

import (
	"github.com/gin-gonic/gin"
	"kartverket/skip/opencost/pkg/database"
	"math"
	"net/http"
	"strings"
	"time"
)

type RestHandler struct {
	dbClient   *database.Client
	reportType string
}

func NewRestHandler(dbClient *database.Client, reportType string) *RestHandler {
	return &RestHandler{
		dbClient:   dbClient,
		reportType: reportType,
	}
}

func (r *RestHandler) HandleGET(c *gin.Context) {
	window := c.Query("window")
	cluster := c.Query("cluster")

	if !isWindowValid(window) {
		c.String(http.StatusBadRequest, "Invalid window")
	}

	namespaceReports, err := r.getReports(window, cluster)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	c.JSON(http.StatusOK, namespaceReports)
}

func (r *RestHandler) getReports(window string, cluster string) ([]NamespaceCost, error) {
	startDate, err := time.Parse(time.RFC3339, strings.Split(window, ",")[0])
	if err != nil {
		return nil, err
	}
	endDate, err := time.Parse(time.RFC3339, strings.Split(window, ",")[1])
	if err != nil {
		return nil, err
	}

	adjustedStartTime := startDate.Add(-time.Hour)
	duration := endDate.Sub(adjustedStartTime)
	daysBetween := int(math.Max(duration.Hours()/24, 1))

	var reports []database.Report

	if cluster == "all" {
		reports, err = r.dbClient.GetReportsForByTypeAndDate(r.reportType, adjustedStartTime, endDate)
	} else {
		reports, err = r.dbClient.GetReportsByClusterAndTypeAndDate(cluster, r.reportType, adjustedStartTime, endDate)
	}

	if err != nil {
		return nil, err
	}

	namespaceCosts := mapDatabaseReportsToNamespaceCosts(reports, daysBetween)

	return namespaceCosts, nil
}

func isWindowValid(window string) bool {
	if window == "" {
		return false
	}
	dateSplits := strings.Split(window, ",")
	_, errStart := time.Parse(time.RFC3339, dateSplits[0])
	_, errEnd := time.Parse(time.RFC3339, dateSplits[1])

	if errStart != nil || errEnd != nil {
		return false
	}
	return true
}
