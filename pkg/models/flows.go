package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Flows struct {
	bun.BaseModel `bun:"table:flows"`

	ID                  uuid.UUID         `bun:",pk,type:uuid,default:gen_random_uuid()" json:"id"`
	Name                string            `bun:"name,type:text,notnull" json:"name"`
	Description         string            `bun:"description,type:text,default:''" json:"description"`
	ProjectID           string            `bun:"project_id,type:text,notnull" json:"project_id"`
	RunnerID            string            `bun:"runner_id,type:text,default:''" json:"runner_id"`
	ExecParallel        bool              `bun:"exec_parallel,type:bool,default:false" json:"exec_parallel"`
	Actions             []Action          `bun:"type:jsonb,default:jsonb('[]')" json:"actions"`
	Maintenance         bool              `bun:"maintenance,type:bool,default:false" json:"maintenance"`
	MaintenanceMessage  string            `bun:"maintenance_message,type:text,default:''" json:"maintenance_message"`
	Disabled            bool              `bun:"disabled,type:bool,default:false" json:"disabled"`
	DisabledReason      string            `bun:"disabled_reason,type:text,default:''" json:"disabled_reason"`
	CreatedAt           time.Time         `bun:"created_at,type:timestamptz,default:now()" json:"created_at"`
	UpdatedAt           time.Time         `bun:"updated_at,type:timestamptz" json:"updated_at"`
	EncryptActionParams bool              `bun:"encrypt_action_params,type:bool,default:true" json:"encrypt_action_params"`
	EncryptExecutions   bool              `bun:"encrypt_executions,type:bool,default:true" json:"encrypt_executions"`
	FailurePipelines    []FailurePipeline `bun:"type:jsonb,default:jsonb('[]')" json:"failure_pipelines"`
	FailurePipelineID   string            `bun:"failure_pipeline_id,type:text,default:''" json:"failure_pipeline_id"`
}

type Action struct {
	ID                uuid.UUID `json:"id"`
	Name              string    `json:"name"`
	Description       string    `json:"description"`
	Plugin            string    `json:"plugin"`
	Version           string    `json:"version"`
	Icon              string    `json:"icon"`
	Category          string    `json:"category"`
	Active            bool      `json:"active"`
	Params            []Params  `json:"params"`
	CustomName        string    `json:"custom_name"`
	CustomDescription string    `json:"custom_description"`
	FailurePipelineID string    `json:"failure_pipeline_id"`
}

type Params struct {
	Key         string    `json:"key"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Category    string    `json:"category"`
	Required    bool      `json:"required"`
	Type        string    `json:"type"`
	Value       string    `json:"value"`
	Default     string    `json:"default"`
	Options     []Options `json:"options,omitempty"`
}

type Options struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type FailurePipeline struct {
	ID           uuid.UUID `json:"id"`
	Name         string    `json:"name"`
	Actions      []Action  `json:"actions"`
	ExecParallel bool      `json:"exec_parallel"`
}
