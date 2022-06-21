/*
 * Grafana HTTP API.
 *
 * The Grafana backend exposes an HTTP API, the same API is used by the frontend to do everything from saving dashboards, creating users and updating data sources.
 *
 * API version: 0.0.1
 * Contact: hello@grafana.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package goclient

import (
	"time"
)

type AlertListItemDtoModel struct {
	DashboardId    int64                `json:"dashboardId,omitempty"`
	DashboardSlug  string               `json:"dashboardSlug,omitempty"`
	DashboardUid   string               `json:"dashboardUid,omitempty"`
	EvalData       *JsonModel           `json:"evalData,omitempty"`
	EvalDate       time.Time            `json:"evalDate,omitempty"`
	ExecutionError string               `json:"executionError,omitempty"`
	Id             int64                `json:"id,omitempty"`
	Name           string               `json:"name,omitempty"`
	NewStateDate   time.Time            `json:"newStateDate,omitempty"`
	PanelId        int64                `json:"panelId,omitempty"`
	State          *AlertStateTypeModel `json:"state,omitempty"`
	Url            string               `json:"url,omitempty"`
}