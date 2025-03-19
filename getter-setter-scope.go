package access

import "context"

type scopeGetterSetter struct {
	prefix string
	parent GetterSetter
}

func (scope *scopeGetterSetter) GetProperty(
	ctx context.Context,
	name string,
) (interface{}, error) {
	return scope.parent.GetProperty(ctx, scope.prefix+name)
}

func (scope *scopeGetterSetter) SetProperty(
	ctx context.Context,
	name string,
	value interface{},
) error {
	return scope.parent.SetProperty(ctx, scope.prefix+name, value)
}

// NewScopeGetterSetter is constructor for build scope GetterSetter.
// It transparent adds prefix to the property name
func NewScopeGetterSetter(
	prefix string,
	parent GetterSetter,
) GetterSetter {
	return &scopeGetterSetter{
		prefix: prefix,
		parent: parent,
	}
}
