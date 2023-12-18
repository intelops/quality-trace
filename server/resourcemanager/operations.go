package resourcemanager

import (
	"context"
	"fmt"

	"github.com/kubeshop/tracetest/server/pkg/id"
	"golang.org/x/exp/slices"
)

type Operation string

const (
	OperationNoop   Operation = ""
	OperationList   Operation = "list"
	OperationCreate Operation = "create"
	OperationUpdate Operation = "update"
	OperationGet    Operation = "get"
	OperationDelete Operation = "delete"

	OperationGetAugmented  Operation = "getAugmented"
	OperationListAugmented Operation = "listAugmented"
)

var availableOperations = []Operation{
	OperationList,
	OperationCreate,
	OperationUpdate,
	OperationGet,
	OperationDelete,
}

var augmentedOperations = []Operation{
	OperationGetAugmented,
	OperationListAugmented,
}

type SortableHandler interface {
	SortingFields() []string
}

type List[T ResourceSpec] interface {
	SortableHandler
	List(_ context.Context, take, skip int, query, sortBy, sortDirection string) ([]T, error)
	Count(_ context.Context, query string) (int, error)
}

type IDSetter[T ResourceSpec] interface {
	SetID(T, id.ID) T
}

type Create[T ResourceSpec] interface {
	Create(context.Context, T) (T, error)
	IDSetter[T]
}

type Update[T ResourceSpec] interface {
	Update(context.Context, T) (T, error)
}

type Get[T ResourceSpec] interface {
	Get(context.Context, id.ID) (T, error)
}

type Delete[T ResourceSpec] interface {
	Delete(context.Context, id.ID) error
}

type Provision[T ResourceSpec] interface {
	Provision(context.Context, T) error
	IDSetter[T]
}

type Current[T ResourceSpec] interface {
	Current(context.Context) (T, error)
}

type GetAugmented[T ResourceSpec] interface {
	GetAugmented(context.Context, id.ID) (T, error)
}

type ListAugmented[T ResourceSpec] interface {
	ListAugmented(_ context.Context, take, skip int, query, sortBy, sortDirection string) ([]T, error)
}

type resourceHandler[T ResourceSpec] struct {
	SetID         func(T, id.ID) T
	List          func(_ context.Context, take, skip int, query, sortBy, sortDirection string) ([]T, error)
	Count         func(_ context.Context, query string) (int, error)
	SortingFields func() []string
	Create        func(context.Context, T) (T, error)
	Update        func(context.Context, T) (T, error)
	Get           func(context.Context, id.ID) (T, error)
	Delete        func(context.Context, id.ID) error
	Provision     func(context.Context, T) error

	GetAugmented  func(context.Context, id.ID) (T, error)
	ListAugmented func(_ context.Context, take, skip int, query, sortBy, sortDirection string) ([]T, error)
}

func (rh *resourceHandler[T]) bindOperations(enabledOperations []Operation, handler any) error {
	if len(enabledOperations) < 1 {

		fmt.Println("Error: no operations enabled") //debug
		return fmt.Errorf("no operations enabled")
	}

	fmt.Println("Binding operations for handler...") //debug

	if slices.Contains(enabledOperations, OperationList) {
		fmt.Println("Binding List operation...") //debug

		err := rh.bindListOperation(handler)
		if err != nil {
			return err
		}
	}

	if slices.Contains(enabledOperations, OperationCreate) {
		fmt.Println("Binding Create operation...") //debug

		err := rh.bindCreateOperation(handler)
		if err != nil {
			return err
		}
	}

	if slices.Contains(enabledOperations, OperationUpdate) {

		fmt.Println("Binding Update operation...") //debug
		err := rh.bindUpdateOperation(handler)
		if err != nil {
			return err
		}
	}

	if slices.Contains(enabledOperations, OperationGet) {

		fmt.Println("Binding Get operation...") //debug
		err := rh.bindGetOperation(handler)
		if err != nil {
			return err
		}
	}

	if slices.Contains(enabledOperations, OperationDelete) {

		fmt.Println("Binding Delete operation...") //debug
		err := rh.bindDeleteOperation(handler)
		if err != nil {
			return err
		}
	}

	if slices.Contains(enabledOperations, OperationGetAugmented) {
		err := rh.bindGetAugmented(handler)
		if err != nil {
			return err
		}
	}

	if slices.Contains(enabledOperations, OperationListAugmented) {
		err := rh.bindListAugmented(handler)
		if err != nil {
			return err
		}
	}
	fmt.Println("Binding Provision operation...") //debug

	err := rh.bindProvisionOperation(handler)
	if err != nil {
		return err
	}

	fmt.Println("Operations binding completed.") //debug
	return nil
}

func (rh *resourceHandler[T]) bindListOperation(handler any) error {
	casted, ok := handler.(List[T])

	if !ok {
		fmt.Println("Error: handler does not implement interface `List[T]`") //debug
		return fmt.Errorf("handler does not implement interface `List[T]`")
	}

	fmt.Println("List operation binding completed.") //debug
	rh.List = casted.List
	rh.Count = casted.Count
	rh.SortingFields = casted.SortingFields

	return nil
}

func (rh *resourceHandler[T]) bindCreateOperation(handler any) error {
	casted, ok := handler.(Create[T])
	if !ok {
		return fmt.Errorf("handler does not implement interface `Create[T]`")
	}

	fmt.Println("Create operation binding completed.") //debug
	rh.Create = casted.Create
	rh.SetID = casted.SetID

	return nil
}

func (rh *resourceHandler[T]) bindUpdateOperation(handler any) error {
	casted, ok := handler.(Update[T])
	if !ok {
		return fmt.Errorf("handler does not implement interface `Update[T]`")
	}

	fmt.Println("Update operation binding completed.") //debug
	rh.Update = casted.Update

	return nil
}

func (rh *resourceHandler[T]) bindGetOperation(handler any) error {
	casted, ok := handler.(Get[T])
	if !ok {
		return fmt.Errorf("handler does not implement interface `Get[T]`")
	}

	fmt.Println("Get operation binding completed.") //debug
	rh.Get = casted.Get

	return nil
}

func (rh *resourceHandler[T]) bindDeleteOperation(handler any) error {
	casted, ok := handler.(Delete[T])
	if !ok {
		return fmt.Errorf("handler does not implement interface `Delete[T]`")
	}

	fmt.Println("Delete operation binding completed.") //debug
	rh.Delete = casted.Delete

	return nil
}

func (rh *resourceHandler[T]) bindProvisionOperation(handler any) error {
	casted, ok := handler.(Provision[T])
	if !ok {
		return fmt.Errorf("handler does not implement interface `Provision[T]`")
	}

	fmt.Println("Provision operation binding completed.") //debug
	rh.Provision = casted.Provision
	rh.SetID = casted.SetID

	return nil
}

func (rh *resourceHandler[T]) bindGetAugmented(handler any) error {
	casted, ok := handler.(GetAugmented[T])
	if !ok {
		return fmt.Errorf("handler does not implement interface `GetAugmented[T]`")
	}
	rh.GetAugmented = casted.GetAugmented

	return nil
}

func (rh *resourceHandler[T]) bindListAugmented(handler any) error {
	casted, ok := handler.(ListAugmented[T])
	if !ok {
		return fmt.Errorf("handler does not implement interface `ListAugmented[T]`")
	}
	rh.ListAugmented = casted.ListAugmented

	return nil
}
