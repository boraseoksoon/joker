// This file is generated by generate-std.joke script. Do not edit manually!

package filepath

import (
	. "github.com/candid82/joker/core"
	"path/filepath"
)

var list_separator_ String
var separator_ String
var __abs__P ProcFn = __abs_
var abs_ Proc = Proc{Fn: __abs__P, Name: "abs_", Package: "std/filepath"}

func __abs_(_args []Object) Object {
	_c := len(_args)
	switch {
	case _c == 1:
		path := ExtractString(_args, 0)
		 _res, err := filepath.Abs(path)
		PanicOnErr(err)
		return MakeString(_res)

	default:
		PanicArity(_c)
	}
	return NIL
}

var __isabs__P ProcFn = __isabs_
var isabs_ Proc = Proc{Fn: __isabs__P, Name: "isabs_", Package: "std/filepath"}

func __isabs_(_args []Object) Object {
	_c := len(_args)
	switch {
	case _c == 1:
		path := ExtractString(_args, 0)
		_res := filepath.IsAbs(path)
		return MakeBoolean(_res)

	default:
		PanicArity(_c)
	}
	return NIL
}

var __base__P ProcFn = __base_
var base_ Proc = Proc{Fn: __base__P, Name: "base_", Package: "std/filepath"}

func __base_(_args []Object) Object {
	_c := len(_args)
	switch {
	case _c == 1:
		path := ExtractString(_args, 0)
		_res := filepath.Base(path)
		return MakeString(_res)

	default:
		PanicArity(_c)
	}
	return NIL
}

var __clean__P ProcFn = __clean_
var clean_ Proc = Proc{Fn: __clean__P, Name: "clean_", Package: "std/filepath"}

func __clean_(_args []Object) Object {
	_c := len(_args)
	switch {
	case _c == 1:
		path := ExtractString(_args, 0)
		_res := filepath.Clean(path)
		return MakeString(_res)

	default:
		PanicArity(_c)
	}
	return NIL
}

var __dir__P ProcFn = __dir_
var dir_ Proc = Proc{Fn: __dir__P, Name: "dir_", Package: "std/filepath"}

func __dir_(_args []Object) Object {
	_c := len(_args)
	switch {
	case _c == 1:
		path := ExtractString(_args, 0)
		_res := filepath.Dir(path)
		return MakeString(_res)

	default:
		PanicArity(_c)
	}
	return NIL
}

var __eval_symlinks__P ProcFn = __eval_symlinks_
var eval_symlinks_ Proc = Proc{Fn: __eval_symlinks__P, Name: "eval_symlinks_", Package: "std/filepath"}

func __eval_symlinks_(_args []Object) Object {
	_c := len(_args)
	switch {
	case _c == 1:
		path := ExtractString(_args, 0)
		 _res, err := filepath.EvalSymlinks(path)
		PanicOnErr(err)
		return MakeString(_res)

	default:
		PanicArity(_c)
	}
	return NIL
}

var __ext__P ProcFn = __ext_
var ext_ Proc = Proc{Fn: __ext__P, Name: "ext_", Package: "std/filepath"}

func __ext_(_args []Object) Object {
	_c := len(_args)
	switch {
	case _c == 1:
		path := ExtractString(_args, 0)
		_res := filepath.Ext(path)
		return MakeString(_res)

	default:
		PanicArity(_c)
	}
	return NIL
}

var __file_seq__P ProcFn = __file_seq_
var file_seq_ Proc = Proc{Fn: __file_seq__P, Name: "file_seq_", Package: "std/filepath"}

func __file_seq_(_args []Object) Object {
	_c := len(_args)
	switch {
	case _c == 1:
		root := ExtractString(_args, 0)
		_res := fileSeq(root)
		return _res

	default:
		PanicArity(_c)
	}
	return NIL
}

var __from_slash__P ProcFn = __from_slash_
var from_slash_ Proc = Proc{Fn: __from_slash__P, Name: "from_slash_", Package: "std/filepath"}

func __from_slash_(_args []Object) Object {
	_c := len(_args)
	switch {
	case _c == 1:
		path := ExtractString(_args, 0)
		_res := filepath.FromSlash(path)
		return MakeString(_res)

	default:
		PanicArity(_c)
	}
	return NIL
}

var __glob__P ProcFn = __glob_
var glob_ Proc = Proc{Fn: __glob__P, Name: "glob_", Package: "std/filepath"}

func __glob_(_args []Object) Object {
	_c := len(_args)
	switch {
	case _c == 1:
		pattern := ExtractString(_args, 0)
		 _res, err := filepath.Glob(pattern)
		PanicOnErr(err)
		return MakeStringVector(_res)

	default:
		PanicArity(_c)
	}
	return NIL
}

var __join__P ProcFn = __join_
var join_ Proc = Proc{Fn: __join__P, Name: "join_", Package: "std/filepath"}

func __join_(_args []Object) Object {
	_c := len(_args)
	switch {
	case true:
		CheckArity(_args, 0, 999)
		elems := ExtractStrings(_args, 0)
		_res := filepath.Join(elems...)
		return MakeString(_res)

	default:
		PanicArity(_c)
	}
	return NIL
}

var __ismatches__P ProcFn = __ismatches_
var ismatches_ Proc = Proc{Fn: __ismatches__P, Name: "ismatches_", Package: "std/filepath"}

func __ismatches_(_args []Object) Object {
	_c := len(_args)
	switch {
	case _c == 2:
		pattern := ExtractString(_args, 0)
		name := ExtractString(_args, 1)
		 _res, err := filepath.Match(pattern, name)
		PanicOnErr(err)
		return MakeBoolean(_res)

	default:
		PanicArity(_c)
	}
	return NIL
}

var __rel__P ProcFn = __rel_
var rel_ Proc = Proc{Fn: __rel__P, Name: "rel_", Package: "std/filepath"}

func __rel_(_args []Object) Object {
	_c := len(_args)
	switch {
	case _c == 2:
		basepath := ExtractString(_args, 0)
		targpath := ExtractString(_args, 1)
		 _res, err := filepath.Rel(basepath, targpath)
		PanicOnErr(err)
		return MakeString(_res)

	default:
		PanicArity(_c)
	}
	return NIL
}

var __split__P ProcFn = __split_
var split_ Proc = Proc{Fn: __split__P, Name: "split_", Package: "std/filepath"}

func __split_(_args []Object) Object {
	_c := len(_args)
	switch {
	case _c == 1:
		path := ExtractString(_args, 0)
		 _dir, _file := filepath.Split(path)
		_res := NewVectorFrom(MakeString(_dir), MakeString(_file))
		return _res

	default:
		PanicArity(_c)
	}
	return NIL
}

var __split_list__P ProcFn = __split_list_
var split_list_ Proc = Proc{Fn: __split_list__P, Name: "split_list_", Package: "std/filepath"}

func __split_list_(_args []Object) Object {
	_c := len(_args)
	switch {
	case _c == 1:
		path := ExtractString(_args, 0)
		_res := filepath.SplitList(path)
		return MakeStringVector(_res)

	default:
		PanicArity(_c)
	}
	return NIL
}

var __to_slash__P ProcFn = __to_slash_
var to_slash_ Proc = Proc{Fn: __to_slash__P, Name: "to_slash_", Package: "std/filepath"}

func __to_slash_(_args []Object) Object {
	_c := len(_args)
	switch {
	case _c == 1:
		path := ExtractString(_args, 0)
		_res := filepath.ToSlash(path)
		return MakeString(_res)

	default:
		PanicArity(_c)
	}
	return NIL
}

var __volume_name__P ProcFn = __volume_name_
var volume_name_ Proc = Proc{Fn: __volume_name__P, Name: "volume_name_", Package: "std/filepath"}

func __volume_name_(_args []Object) Object {
	_c := len(_args)
	switch {
	case _c == 1:
		path := ExtractString(_args, 0)
		_res := filepath.VolumeName(path)
		return MakeString(_res)

	default:
		PanicArity(_c)
	}
	return NIL
}

func Init() {
	list_separator_ = MakeString(string(filepath.ListSeparator))
	separator_ = MakeString(string(filepath.Separator))
	InternsOrThunks()
}

var filepathNamespace = GLOBAL_ENV.EnsureNamespace(MakeSymbol("joker.filepath"))

func init() {
	filepathNamespace.Lazy = Init
}
