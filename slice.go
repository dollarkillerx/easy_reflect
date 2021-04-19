package easy_reflect

import "github.com/pkg/errors"

func (r *ReflectItem) Len() (int, error) {
	if !r.IsSlice() {
		return 0, errors.New("Structure must be array")
	}

	return r.vType.Len(), nil
}
