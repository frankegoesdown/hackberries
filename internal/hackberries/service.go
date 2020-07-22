package hackberries

import (
	"context"
	v1 "github.com/frankegoesdown/hackberries/internal/api/v1"
)

// Service ...
// @gtg http-server log metrics http-client
type Service interface {
	GetProjects(ctx context.Context, token *string) (projects v1.Projects, err error)
	CreateProject(ctx context.Context, token *string) (project v1.Project, err error)
	UpdateProject(ctx context.Context, token *string) (project v1.Project, err error)
	DeleteProject(ctx context.Context, token *string) (err error)
	JoinProject(ctx context.Context, token *string) (err error)
	ShowUsers(ctx context.Context, token *string) (users v1.Users, err error)
	LeaveProject(ctx context.Context, token *string) (err error)
	InviteUsers(ctx context.Context, token *string) (err error)
	PromoteUser(ctx context.Context, token *string) (err error)
}
