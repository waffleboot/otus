package usecase

import "context"

type TestUseCase interface {
	Test(ctx context.Context) error
}
