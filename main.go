package main

import (
	"fmt"
	"log"
	"os"

	_ "github.com/bytecodealliance/wasmtime-go/v14" // https://github.com/wasmerio/wasmer-go/issues/402

	"github.com/wasmerio/wasmer-go/wasmer"
)

const (
	havokWASMfilePath = "./HavokPhysics.wasm"
)

var (
	t0 = wasmer.NewFunctionType(
		wasmer.NewValueTypes(wasmer.I32, wasmer.I32, wasmer.I32),
		wasmer.NewValueTypes(),
	)
	t1 = wasmer.NewFunctionType(
		wasmer.NewValueTypes(wasmer.I32, wasmer.I32),
		wasmer.NewValueTypes(),
	)
	t3 = wasmer.NewFunctionType(
		wasmer.NewValueTypes(wasmer.I32),
		wasmer.NewValueTypes(wasmer.I32),
	)
	t4 = wasmer.NewFunctionType(
		wasmer.NewValueTypes(wasmer.I32, wasmer.I32),
		wasmer.NewValueTypes(wasmer.I32),
	)
	t5 = wasmer.NewFunctionType(
		wasmer.NewValueTypes(wasmer.I32),
		wasmer.NewValueTypes(),
	)
	t6 = wasmer.NewFunctionType(
		wasmer.NewValueTypes(wasmer.I32, wasmer.I32, wasmer.I32, wasmer.I32),
		wasmer.NewValueTypes(),
	)
	t7 = wasmer.NewFunctionType(
		wasmer.NewValueTypes(wasmer.I32, wasmer.I32, wasmer.I32, wasmer.I32, wasmer.I32, wasmer.I32),
		wasmer.NewValueTypes(),
	)
	t8 = wasmer.NewFunctionType(
		wasmer.NewValueTypes(wasmer.I32, wasmer.I32, wasmer.I32, wasmer.I32, wasmer.I32),
		wasmer.NewValueTypes(),
	)
	t9 = wasmer.NewFunctionType(
		wasmer.NewValueTypes(wasmer.I32, wasmer.I32, wasmer.I32, wasmer.I32, wasmer.I32, wasmer.I32, wasmer.I32, wasmer.I32, wasmer.I32),
		wasmer.NewValueTypes(),
	)
	t11 = wasmer.NewFunctionType(
		wasmer.NewValueTypes(wasmer.I32, wasmer.I32, wasmer.I32, wasmer.I32),
		wasmer.NewValueTypes(wasmer.I32),
	)
	t19 = wasmer.NewFunctionType(
		wasmer.NewValueTypes(),
		wasmer.NewValueTypes(wasmer.I32),
	)
	t35 = wasmer.NewFunctionType(
		wasmer.NewValueTypes(),
		wasmer.NewValueTypes(),
	)
	t78 = wasmer.NewFunctionType(
		wasmer.NewValueTypes(wasmer.I32, wasmer.I32, wasmer.I32, wasmer.I64, wasmer.I64),
		wasmer.NewValueTypes(),
	)
	t79 = wasmer.NewFunctionType(
		wasmer.NewValueTypes(),
		wasmer.NewValueTypes(wasmer.F64),
	)
)

func wasmerHavok() error {
	wasmBytes, err := os.ReadFile(havokWASMfilePath)
	if err != nil {
		log.Fatal("ReadFile(): err:", err)
		return err
	}

	engine := wasmer.NewEngine()
	store := wasmer.NewStore(engine)
	module, err := wasmer.NewModule(store, wasmBytes)
	if err != nil {
		log.Fatal("NewModule(): err:", err)
		return err
	}

	importObject := wasmer.NewImportObject()
	importObject.Register("env", map[string]wasmer.IntoExtern{
		"_emval_get_method_caller": wasmer.NewFunction(store, t4,
			func(args []wasmer.Value) ([]wasmer.Value, error) {
				return []wasmer.Value{}, nil
			},
		),
		"_emval_call_void_method": wasmer.NewFunction(store, t6,
			func(args []wasmer.Value) ([]wasmer.Value, error) {
				return []wasmer.Value{}, nil
			},
		),
		"_embind_register_value_array": wasmer.NewFunction(store, t7,
			func(args []wasmer.Value) ([]wasmer.Value, error) {
				return []wasmer.Value{}, nil
			},
		),
		"_embind_register_value_array_element": wasmer.NewFunction(store, t9,
			func(args []wasmer.Value) ([]wasmer.Value, error) {
				return []wasmer.Value{}, nil
			},
		),
		"_embind_finalize_value_array": wasmer.NewFunction(store, t5,
			func(args []wasmer.Value) ([]wasmer.Value, error) {
				return []wasmer.Value{}, nil
			},
		),
		"_embind_register_enum": wasmer.NewFunction(store, t6,
			func(args []wasmer.Value) ([]wasmer.Value, error) {
				return []wasmer.Value{}, nil
			},
		),
		"_embind_register_enum_value": wasmer.NewFunction(store, t0,
			func(args []wasmer.Value) ([]wasmer.Value, error) {
				return []wasmer.Value{}, nil
			},
		),
		"_embind_register_function": wasmer.NewFunction(store, t7,
			func(args []wasmer.Value) ([]wasmer.Value, error) {
				return []wasmer.Value{}, nil
			},
		),
		"_emval_decref": wasmer.NewFunction(store, t5,
			func(args []wasmer.Value) ([]wasmer.Value, error) {
				return []wasmer.Value{}, nil
			},
		),
		"abort": wasmer.NewFunction(store, t35,
			func(args []wasmer.Value) ([]wasmer.Value, error) {
				return []wasmer.Value{}, nil
			},
		),
		"_embind_register_void": wasmer.NewFunction(store, t1,
			func(args []wasmer.Value) ([]wasmer.Value, error) {
				return []wasmer.Value{}, nil
			},
		),
		"_embind_register_bool": wasmer.NewFunction(store, t8,
			func(args []wasmer.Value) ([]wasmer.Value, error) {
				return []wasmer.Value{}, nil
			},
		),
		"_embind_register_integer": wasmer.NewFunction(store, t8,
			func(args []wasmer.Value) ([]wasmer.Value, error) {
				return []wasmer.Value{}, nil
			},
		),
		"_embind_register_bigint": wasmer.NewFunction(store, t78,
			func(args []wasmer.Value) ([]wasmer.Value, error) {
				return []wasmer.Value{}, nil
			},
		),
		"_embind_register_float": wasmer.NewFunction(store, t0,
			func(args []wasmer.Value) ([]wasmer.Value, error) {
				return []wasmer.Value{}, nil
			},
		),
		"_embind_register_std_string": wasmer.NewFunction(store, t1,
			func(args []wasmer.Value) ([]wasmer.Value, error) {
				return []wasmer.Value{}, nil
			},
		),
		"_embind_register_std_wstring": wasmer.NewFunction(store, t0,
			func(args []wasmer.Value) ([]wasmer.Value, error) {
				return []wasmer.Value{}, nil
			},
		),
		"_embind_register_emval": wasmer.NewFunction(store, t1,
			func(args []wasmer.Value) ([]wasmer.Value, error) {
				return []wasmer.Value{}, nil
			},
		),
		"_embind_register_memory_view": wasmer.NewFunction(store, t0,
			func(args []wasmer.Value) ([]wasmer.Value, error) {
				return []wasmer.Value{}, nil
			},
		),
		"emscripten_memcpy_big": wasmer.NewFunction(store, t0,
			func(args []wasmer.Value) ([]wasmer.Value, error) {
				return []wasmer.Value{}, nil
			},
		),
		"_emscripten_get_now_is_monotonic": wasmer.NewFunction(store, t19,
			func(args []wasmer.Value) ([]wasmer.Value, error) {
				return []wasmer.Value{}, nil
			},
		),
		"emscripten_get_now": wasmer.NewFunction(store, t79,
			func(args []wasmer.Value) ([]wasmer.Value, error) {
				return []wasmer.Value{}, nil
			},
		),
		"emscripten_get_heap_max": wasmer.NewFunction(store, t19,
			func(args []wasmer.Value) ([]wasmer.Value, error) {
				return []wasmer.Value{}, nil
			},
		),
		"emscripten_resize_heap": wasmer.NewFunction(store, t3,
			func(args []wasmer.Value) ([]wasmer.Value, error) {
				return []wasmer.Value{}, nil
			},
		),
	})
	importObject.Register("wasi_snapshot_preview1", map[string]wasmer.IntoExtern{
		"fd_write": wasmer.NewFunction(store, t11,
			func(args []wasmer.Value) ([]wasmer.Value, error) {
				return []wasmer.Value{}, nil
			},
		),
	})

	instance, err := wasmer.NewInstance(module, importObject)
	if err != nil {
		log.Fatal("NewInstance(): err:", err)
		return err
	}

	// fmt.Println(fmt.Sprintf("instance: %v", instance))

	f, err := instance.Exports.GetFunction("HP_GetStatistics")
	if err != nil {
		log.Fatal("GetFunction(): err:", err)
		return err
	}

	result, err := f(0)
	if err != nil {
		log.Fatal("f(): err:", err)
		return err
	}

	fmt.Println(result)

	return nil
}

/*
go run ./main.go
*/
func main() {
	if err := wasmerHavok(); err != nil {
		log.Fatal("wasmerHavok(): err:", err)
	}
}
