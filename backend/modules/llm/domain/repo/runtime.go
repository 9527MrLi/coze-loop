// Copyright (c) 2025 Bytedance Ltd. and/or its affiliates
// SPDX-License-Identifier: Apache-2.0

package repo

import (
	"context"

	"github.com/coze-dev/cozeloop/backend/modules/llm/domain/entity"
)

//go:generate mockgen -destination=mocks/runtime.go -package=mocks . IRuntimeRepo
type IRuntimeRepo interface {
	CreateModelRequestRecord(ctx context.Context, record *entity.ModelRequestRecord) (err error)
}
