package access

import (
	"context"
	"github.com/adverax/types"
	"time"
)

type reader struct {
	Getter
}

func (that *reader) GetBoolean(
	ctx context.Context,
	name string,
	defVal bool,
) (res bool, err error) {
	return types.GetBooleanProperty(ctx, that, name, defVal)
}

func (that *reader) GetString(
	ctx context.Context,
	name string,
	defVal string,
) (res string, err error) {
	return types.GetStringProperty(ctx, that, name, defVal)
}

func (that *reader) GetInteger(
	ctx context.Context,
	name string,
	defVal int64,
) (res int64, err error) {
	return types.GetIntegerProperty(ctx, that, name, defVal)
}

func (that *reader) GetFloat(
	ctx context.Context,
	name string,
	defVal float64,
) (res float64, err error) {
	return types.GetFloatProperty(ctx, that, name, defVal)
}

func (that *reader) GetDuration(
	ctx context.Context,
	name string,
	defVal time.Duration,
) (res time.Duration, err error) {
	return types.GetDurationProperty(ctx, that, name, defVal)
}

// NewReader is constructor for build Reader, based on the props
func NewReader(getter Getter) ReaderGetter {
	return &reader{
		Getter: getter,
	}
}
