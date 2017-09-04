package tensor

import "github.com/pkg/errors"

/*
GENERATED FILE. DO NOT EDIT
*/

// Gt performs t > other elementwise. Both t and other must have the same shape.
// Acceptable FuncOpts are: UseUnsafe(), AsSameType(), WithReuse().
//UseUnsafe() will ensure that the same type is returned.
// Tensors used in WithReuse has to have the same Dtype as the return value's Dtype.
func (t *Dense) Gt(other *Dense, opts ...FuncOpt) (retVal *Dense, err error) {

	var ret Tensor
	if t.oe != nil {
		if ret, err = t.oe.Gt(t, other, opts...); err != nil {
			return nil, errors.Wrapf(err, "Unable to do Gt()")
		}
		if retVal, err = assertDense(ret); err != nil {
			return nil, errors.Wrapf(err, opFail, "Gt")
		}
		return
	}

	if gter, ok := t.e.(Gter); ok {
		if ret, err = gter.Gt(t, other, opts...); err != nil {
			return nil, errors.Wrapf(err, "Unable to do Gt()")
		}
		if retVal, err = assertDense(ret); err != nil {
			return nil, errors.Wrapf(err, opFail, "Gt")
		}
		return
	}
	return nil, errors.Errorf("Engine does not support Gt()")
}

// Gte performs t ≥ other elementwise. Both t and other must have the same shape.
// Acceptable FuncOpts are: UseUnsafe(), AsSameType(), WithReuse().
//UseUnsafe() will ensure that the same type is returned.
// Tensors used in WithReuse has to have the same Dtype as the return value's Dtype.
func (t *Dense) Gte(other *Dense, opts ...FuncOpt) (retVal *Dense, err error) {

	var ret Tensor
	if t.oe != nil {
		if ret, err = t.oe.Gte(t, other, opts...); err != nil {
			return nil, errors.Wrapf(err, "Unable to do Gte()")
		}
		if retVal, err = assertDense(ret); err != nil {
			return nil, errors.Wrapf(err, opFail, "Gte")
		}
		return
	}

	if gteer, ok := t.e.(Gteer); ok {
		if ret, err = gteer.Gte(t, other, opts...); err != nil {
			return nil, errors.Wrapf(err, "Unable to do Gte()")
		}
		if retVal, err = assertDense(ret); err != nil {
			return nil, errors.Wrapf(err, opFail, "Gte")
		}
		return
	}
	return nil, errors.Errorf("Engine does not support Gte()")
}

// Lt performs t < other elementwise. Both t and other must have the same shape.
// Acceptable FuncOpts are: UseUnsafe(), AsSameType(), WithReuse().
//UseUnsafe() will ensure that the same type is returned.
// Tensors used in WithReuse has to have the same Dtype as the return value's Dtype.
func (t *Dense) Lt(other *Dense, opts ...FuncOpt) (retVal *Dense, err error) {

	var ret Tensor
	if t.oe != nil {
		if ret, err = t.oe.Lt(t, other, opts...); err != nil {
			return nil, errors.Wrapf(err, "Unable to do Lt()")
		}
		if retVal, err = assertDense(ret); err != nil {
			return nil, errors.Wrapf(err, opFail, "Lt")
		}
		return
	}

	if lter, ok := t.e.(Lter); ok {
		if ret, err = lter.Lt(t, other, opts...); err != nil {
			return nil, errors.Wrapf(err, "Unable to do Lt()")
		}
		if retVal, err = assertDense(ret); err != nil {
			return nil, errors.Wrapf(err, opFail, "Lt")
		}
		return
	}
	return nil, errors.Errorf("Engine does not support Lt()")
}

// Lte performs t ≤ other elementwise. Both t and other must have the same shape.
// Acceptable FuncOpts are: UseUnsafe(), AsSameType(), WithReuse().
//UseUnsafe() will ensure that the same type is returned.
// Tensors used in WithReuse has to have the same Dtype as the return value's Dtype.
func (t *Dense) Lte(other *Dense, opts ...FuncOpt) (retVal *Dense, err error) {

	var ret Tensor
	if t.oe != nil {
		if ret, err = t.oe.Lte(t, other, opts...); err != nil {
			return nil, errors.Wrapf(err, "Unable to do Lte()")
		}
		if retVal, err = assertDense(ret); err != nil {
			return nil, errors.Wrapf(err, opFail, "Lte")
		}
		return
	}

	if lteer, ok := t.e.(Lteer); ok {
		if ret, err = lteer.Lte(t, other, opts...); err != nil {
			return nil, errors.Wrapf(err, "Unable to do Lte()")
		}
		if retVal, err = assertDense(ret); err != nil {
			return nil, errors.Wrapf(err, opFail, "Lte")
		}
		return
	}
	return nil, errors.Errorf("Engine does not support Lte()")
}

// ElEq performs t == other elementwise. Both t and other must have the same shape.
// Acceptable FuncOpts are: UseUnsafe(), AsSameType(), WithReuse().
//UseUnsafe() will ensure that the same type is returned.
// Tensors used in WithReuse has to have the same Dtype as the return value's Dtype.
func (t *Dense) ElEq(other *Dense, opts ...FuncOpt) (retVal *Dense, err error) {

	var ret Tensor
	if t.oe != nil {
		if ret, err = t.oe.ElEq(t, other, opts...); err != nil {
			return nil, errors.Wrapf(err, "Unable to do Eq()")
		}
		if retVal, err = assertDense(ret); err != nil {
			return nil, errors.Wrapf(err, opFail, "Eq")
		}
		return
	}

	if eleqer, ok := t.e.(ElEqer); ok {
		if ret, err = eleqer.ElEq(t, other, opts...); err != nil {
			return nil, errors.Wrapf(err, "Unable to do Eq()")
		}
		if retVal, err = assertDense(ret); err != nil {
			return nil, errors.Wrapf(err, opFail, "Eq")
		}
		return
	}
	return nil, errors.Errorf("Engine does not support Eq()")
}

// ElNe performs t ≠ other elementwise. Both t and other must have the same shape.
// Acceptable FuncOpts are: UseUnsafe(), AsSameType(), WithReuse().
//UseUnsafe() will ensure that the same type is returned.
// Tensors used in WithReuse has to have the same Dtype as the return value's Dtype.
func (t *Dense) ElNe(other *Dense, opts ...FuncOpt) (retVal *Dense, err error) {

	var ret Tensor
	if t.oe != nil {
		if ret, err = t.oe.ElNe(t, other, opts...); err != nil {
			return nil, errors.Wrapf(err, "Unable to do Ne()")
		}
		if retVal, err = assertDense(ret); err != nil {
			return nil, errors.Wrapf(err, opFail, "Ne")
		}
		return
	}

	if eleqer, ok := t.e.(ElEqer); ok {
		if ret, err = eleqer.ElNe(t, other, opts...); err != nil {
			return nil, errors.Wrapf(err, "Unable to do Ne()")
		}
		if retVal, err = assertDense(ret); err != nil {
			return nil, errors.Wrapf(err, opFail, "Ne")
		}
		return
	}
	return nil, errors.Errorf("Engine does not support Ne()")
}

// GtScalar performs t > other elementwise. The leftTensor parameter indicates if the tensor is the left operand. Only scalar types are accepted in other
// Acceptable FuncOpts are: UseUnsafe(), AsSameType(), WithReuse().
// UseUnsafe() will ensure that the same type is returned.
// Tensors used in WithReuse has to have the same Dtype as the return value's Dtype.
func (t *Dense) GtScalar(other interface{}, leftTensor bool, opts ...FuncOpt) (retVal *Dense, err error) {
	var ret Tensor
	if t.oe != nil {
		if ret, err = t.oe.GtScalar(t, other, leftTensor, opts...); err != nil {
			return nil, errors.Wrapf(err, "Unable to do GtScalar()")
		}
		if retVal, err = assertDense(ret); err != nil {
			return nil, errors.Wrapf(err, opFail, "GtScalar")
		}
		return
	}

	if gter, ok := t.e.(Gter); ok {
		if ret, err = gter.GtScalar(t, other, leftTensor, opts...); err != nil {
			return nil, errors.Wrapf(err, "Unable to do GtScalar()")
		}
		if retVal, err = assertDense(ret); err != nil {
			return nil, errors.Wrapf(err, opFail, "GtScalar")
		}
		return
	}
	return nil, errors.Errorf("Engine does not support GtScalar()")
}

// GteScalar performs t ≥ other elementwise. The leftTensor parameter indicates if the tensor is the left operand. Only scalar types are accepted in other
// Acceptable FuncOpts are: UseUnsafe(), AsSameType(), WithReuse().
// UseUnsafe() will ensure that the same type is returned.
// Tensors used in WithReuse has to have the same Dtype as the return value's Dtype.
func (t *Dense) GteScalar(other interface{}, leftTensor bool, opts ...FuncOpt) (retVal *Dense, err error) {
	var ret Tensor
	if t.oe != nil {
		if ret, err = t.oe.GteScalar(t, other, leftTensor, opts...); err != nil {
			return nil, errors.Wrapf(err, "Unable to do GteScalar()")
		}
		if retVal, err = assertDense(ret); err != nil {
			return nil, errors.Wrapf(err, opFail, "GteScalar")
		}
		return
	}

	if gteer, ok := t.e.(Gteer); ok {
		if ret, err = gteer.GteScalar(t, other, leftTensor, opts...); err != nil {
			return nil, errors.Wrapf(err, "Unable to do GteScalar()")
		}
		if retVal, err = assertDense(ret); err != nil {
			return nil, errors.Wrapf(err, opFail, "GteScalar")
		}
		return
	}
	return nil, errors.Errorf("Engine does not support GteScalar()")
}

// LtScalar performs t < other elementwise. The leftTensor parameter indicates if the tensor is the left operand. Only scalar types are accepted in other
// Acceptable FuncOpts are: UseUnsafe(), AsSameType(), WithReuse().
// UseUnsafe() will ensure that the same type is returned.
// Tensors used in WithReuse has to have the same Dtype as the return value's Dtype.
func (t *Dense) LtScalar(other interface{}, leftTensor bool, opts ...FuncOpt) (retVal *Dense, err error) {
	var ret Tensor
	if t.oe != nil {
		if ret, err = t.oe.LtScalar(t, other, leftTensor, opts...); err != nil {
			return nil, errors.Wrapf(err, "Unable to do LtScalar()")
		}
		if retVal, err = assertDense(ret); err != nil {
			return nil, errors.Wrapf(err, opFail, "LtScalar")
		}
		return
	}

	if lter, ok := t.e.(Lter); ok {
		if ret, err = lter.LtScalar(t, other, leftTensor, opts...); err != nil {
			return nil, errors.Wrapf(err, "Unable to do LtScalar()")
		}
		if retVal, err = assertDense(ret); err != nil {
			return nil, errors.Wrapf(err, opFail, "LtScalar")
		}
		return
	}
	return nil, errors.Errorf("Engine does not support LtScalar()")
}

// LteScalar performs t ≤ other elementwise. The leftTensor parameter indicates if the tensor is the left operand. Only scalar types are accepted in other
// Acceptable FuncOpts are: UseUnsafe(), AsSameType(), WithReuse().
// UseUnsafe() will ensure that the same type is returned.
// Tensors used in WithReuse has to have the same Dtype as the return value's Dtype.
func (t *Dense) LteScalar(other interface{}, leftTensor bool, opts ...FuncOpt) (retVal *Dense, err error) {
	var ret Tensor
	if t.oe != nil {
		if ret, err = t.oe.LteScalar(t, other, leftTensor, opts...); err != nil {
			return nil, errors.Wrapf(err, "Unable to do LteScalar()")
		}
		if retVal, err = assertDense(ret); err != nil {
			return nil, errors.Wrapf(err, opFail, "LteScalar")
		}
		return
	}

	if lteer, ok := t.e.(Lteer); ok {
		if ret, err = lteer.LteScalar(t, other, leftTensor, opts...); err != nil {
			return nil, errors.Wrapf(err, "Unable to do LteScalar()")
		}
		if retVal, err = assertDense(ret); err != nil {
			return nil, errors.Wrapf(err, opFail, "LteScalar")
		}
		return
	}
	return nil, errors.Errorf("Engine does not support LteScalar()")
}

// EqScalar performs t == other elementwise. The leftTensor parameter indicates if the tensor is the left operand. Only scalar types are accepted in other
// Acceptable FuncOpts are: UseUnsafe(), AsSameType(), WithReuse().
// UseUnsafe() will ensure that the same type is returned.
// Tensors used in WithReuse has to have the same Dtype as the return value's Dtype.
func (t *Dense) ElEqScalar(other interface{}, leftTensor bool, opts ...FuncOpt) (retVal *Dense, err error) {
	var ret Tensor
	if t.oe != nil {
		if ret, err = t.oe.EqScalar(t, other, leftTensor, opts...); err != nil {
			return nil, errors.Wrapf(err, "Unable to do EqScalar()")
		}
		if retVal, err = assertDense(ret); err != nil {
			return nil, errors.Wrapf(err, opFail, "EqScalar")
		}
		return
	}

	if eleqer, ok := t.e.(ElEqer); ok {
		if ret, err = eleqer.EqScalar(t, other, leftTensor, opts...); err != nil {
			return nil, errors.Wrapf(err, "Unable to do EqScalar()")
		}
		if retVal, err = assertDense(ret); err != nil {
			return nil, errors.Wrapf(err, opFail, "EqScalar")
		}
		return
	}
	return nil, errors.Errorf("Engine does not support EqScalar()")
}

// NeScalar performs t ≠ other elementwise. The leftTensor parameter indicates if the tensor is the left operand. Only scalar types are accepted in other
// Acceptable FuncOpts are: UseUnsafe(), AsSameType(), WithReuse().
// UseUnsafe() will ensure that the same type is returned.
// Tensors used in WithReuse has to have the same Dtype as the return value's Dtype.
func (t *Dense) ElNeScalar(other interface{}, leftTensor bool, opts ...FuncOpt) (retVal *Dense, err error) {
	var ret Tensor
	if t.oe != nil {
		if ret, err = t.oe.NeScalar(t, other, leftTensor, opts...); err != nil {
			return nil, errors.Wrapf(err, "Unable to do NeScalar()")
		}
		if retVal, err = assertDense(ret); err != nil {
			return nil, errors.Wrapf(err, opFail, "NeScalar")
		}
		return
	}

	if eleqer, ok := t.e.(ElEqer); ok {
		if ret, err = eleqer.NeScalar(t, other, leftTensor, opts...); err != nil {
			return nil, errors.Wrapf(err, "Unable to do NeScalar()")
		}
		if retVal, err = assertDense(ret); err != nil {
			return nil, errors.Wrapf(err, opFail, "NeScalar")
		}
		return
	}
	return nil, errors.Errorf("Engine does not support NeScalar()")
}
