// package ini deals with parsing and adapting toml files
// to a structure.
//
// ABNF:
/*
	id -> value stmt
	stmt -> expr stmt'
	stmt' -> nop | op stmt
	value -> number | string | boolean

	table -> [ table' | [ array_table
	table' -> label array_close
	array_close -> ] epsilon

	array_table -> [ table_nested
	table_nested -> label nested_array_close
	nested_array_close -> ] array_close

	epsilon -> nop
*/
package ini
