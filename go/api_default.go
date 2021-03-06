/*
 * Care Plan Generator
 *
 * Care Plan Generator holds logic for generating CarePlan, Goal and Task from Templates together with auto assignment to correct providers from given Care Team
 *
 * API version: 0.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

func PlansCarePlanIdGoalsInstantiatePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	resource := "Goal"
	resource_id := uuid.New()
	response := SuccessResponse{
		Resource: resource,
		Url:      fmt.Sprintf("%s/%s", resource, resource_id),
	}
	json.NewEncoder(w).Encode(response)
}

func PlansCarePlanIdTasksInstantiatePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	resource := "Task"
	resource_id := uuid.New()
	response := SuccessResponse{
		Resource: resource,
		Url:      fmt.Sprintf("%s/%s", resource, resource_id),
	}
	json.NewEncoder(w).Encode(response)
}

func PlansInstantiatePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	resource := "CarePlan"
	resource_id := uuid.New()
	response := SuccessResponse{
		Resource: resource,
		Url:      fmt.Sprintf("%s/%s", resource, resource_id),
	}
	json.NewEncoder(w).Encode(response)
}
