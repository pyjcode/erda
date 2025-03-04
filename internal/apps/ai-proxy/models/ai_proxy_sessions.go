// Copyright (c) 2021 Terminus, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package models

import (
	"database/sql"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/erda-project/erda-infra/providers/mysql/v2/plugins/fields"
	"github.com/erda-project/erda-proto-go/apps/aiproxy/pb"
)

// AIProxySessions is the table ai_proxy_sessions
type AIProxySessions struct {
	Id        fields.UUID      `json:"id" yaml:"id" gorm:"id"`
	CreatedAt time.Time        `json:"createdAt" yaml:"createdAt" gorm:"created_at"`
	UpdatedAt time.Time        `json:"updatedAt" yaml:"updatedAt" gorm:"updated_at"`
	DeletedAt fields.DeletedAt `json:"deletedAt" yaml:"deletedAt" gorm:"deleted_at"`

	UserId        string `json:"userId" yaml:"userId" gorm:"user_id"`
	Name          string
	Topic         string
	ContextLength uint32 `json:"contextLength" yaml:"contextLength" gorm:"context_length"`
	Source        string
	IsArchived    bool         `json:"isArchived" yaml:"isArchived" gorm:"is_archived"`
	ResetAt       sql.NullTime `json:"resetAt" yaml:"resetAt" gorm:"reset_at"`
	Model         string
	Temperature   float64
}

func (*AIProxySessions) TableName() string {
	return "ai_proxy_sessions"
}

func (session *AIProxySessions) ToProtobuf() *pb.Session {
	return &pb.Session{
		UserId:        session.UserId,
		Id:            session.Id.String,
		Name:          session.Name,
		Topic:         session.Topic,
		ContextLength: session.ContextLength,
		IsArchived:    session.IsArchived,
		Source:        session.Source,
		ResetAt:       timestamppb.New(session.ResetAt.Time),
		Model:         session.Model,
		Temperature:   session.Temperature,
	}
}
