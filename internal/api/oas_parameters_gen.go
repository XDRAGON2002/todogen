// Code generated by ogen, DO NOT EDIT.

package api

import (
	"net/http"
	"net/url"

	"github.com/go-faster/errors"

	"github.com/ogen-go/ogen/conv"
	"github.com/ogen-go/ogen/middleware"
	"github.com/ogen-go/ogen/ogenerrors"
	"github.com/ogen-go/ogen/uri"
	"github.com/ogen-go/ogen/validate"
)

// DeleteTodoByIdParams is parameters of deleteTodoById operation.
type DeleteTodoByIdParams struct {
	// ID of todo to delete.
	TodoId int64
}

func unpackDeleteTodoByIdParams(packed middleware.Parameters) (params DeleteTodoByIdParams) {
	{
		key := middleware.ParameterKey{
			Name: "todoId",
			In:   "path",
		}
		params.TodoId = packed[key].(int64)
	}
	return params
}

func decodeDeleteTodoByIdParams(args [1]string, argsEscaped bool, r *http.Request) (params DeleteTodoByIdParams, _ error) {
	// Decode path: todoId.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "todoId",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt64(val)
				if err != nil {
					return err
				}

				params.TodoId = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "todoId",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// GetTodoByIdParams is parameters of getTodoById operation.
type GetTodoByIdParams struct {
	// ID of todo to return.
	TodoId int64
}

func unpackGetTodoByIdParams(packed middleware.Parameters) (params GetTodoByIdParams) {
	{
		key := middleware.ParameterKey{
			Name: "todoId",
			In:   "path",
		}
		params.TodoId = packed[key].(int64)
	}
	return params
}

func decodeGetTodoByIdParams(args [1]string, argsEscaped bool, r *http.Request) (params GetTodoByIdParams, _ error) {
	// Decode path: todoId.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "todoId",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt64(val)
				if err != nil {
					return err
				}

				params.TodoId = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "todoId",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// UpdateTodoByIdParams is parameters of updateTodoById operation.
type UpdateTodoByIdParams struct {
	// ID of todo to update.
	TodoId int64
}

func unpackUpdateTodoByIdParams(packed middleware.Parameters) (params UpdateTodoByIdParams) {
	{
		key := middleware.ParameterKey{
			Name: "todoId",
			In:   "path",
		}
		params.TodoId = packed[key].(int64)
	}
	return params
}

func decodeUpdateTodoByIdParams(args [1]string, argsEscaped bool, r *http.Request) (params UpdateTodoByIdParams, _ error) {
	// Decode path: todoId.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "todoId",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt64(val)
				if err != nil {
					return err
				}

				params.TodoId = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "todoId",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}
