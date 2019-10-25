package faunadb

// Aliases for commonly used reducers, which may be used on arrays, pages or sets.
// Note: Min, Max already exist
func Count(coll interface{}) Expr { return fn1("count", coll) }
func Mean(coll interface{}) Expr  { return fn1("mean", coll) }
func Sum(coll interface{}) Expr   { return fn1("sum", coll) }
func Any(coll interface{}) Expr   { return fn1("any", coll) }
func All(coll interface{}) Expr   { return fn1("all", coll) }

// Reverse(set/array/page)
// We will add a Reverse() function which can take an array, page, or set, and return the reversed version.
func Reverse(coll interface{}) Expr {
	return fn1("reduce", coll)
}

// Documents(set/array/page)
// We will add a Documents() built-in which will allow iterating through all of the documents in a collection. Combined with Filter(), Reduce(), Count(), etc.
// This will allow for arbitrary querying of a collection without the need for indexes. Initially this functionality will be backed by a scan of the collection.
func Documents(coll interface{}) Expr {
	return fn1("documents", coll)
}


func StartsWith(value, search interface{}) Expr { return fn2("startswith", value, "search", search) }
func ContainsStr(value, search interface{}) Expr { return fn2("containsstr", value, "search", search) }
func EndsWith(value, search interface{}) Expr { return fn2("endswith", value, "search", search) }