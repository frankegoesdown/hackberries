package hackberries

import (
	"context"
)

// Service ...
// @gtg http-server log metrics http-client
type Service interface {
	// @gtg http-server-method GET
	// @gtg http-server-uri-path /api/v1/projects
	// @gtg http-server-query limit={limit}
	// @gtg http-server-header X-Auth-Token {token}
	// @gtg http-server-content-type application/json
	// @gtg http-server-response-status http.StatusOK
	// @gtg http-server-response-content-type application/json
	GetProjects(ctx context.Context, limit uint32, token *string) (det v1.Detail, ns v1.Namespace, id *string, err error)
	// @gtg http-server-method PUT
	// @gtg http-server-uri-path /api/v1/namespaces/{namespace}/details/{detail}
	// @gtg http-server-query test={testID}&bla={blaID}
	// @gtg http-server-header X-Auth-Token {token}
	// @gtg http-server-json-tag pretty ThePretty
	// @gtg http-server-content-type application/json
	// @gtg http-server-response-header X-Auth-ID {id}
	// @gtg http-server-response-status http.StatusOK
	// @gtg http-server-response-content-type application/json
	PutDetails(ctx context.Context, namespace string, detail string, testID string, blaID *string, token *string, pretty v1.Detail, yang v1.Namespace) (cool v1.Detail, nothing v1.Namespace, id *string, err error)
}
