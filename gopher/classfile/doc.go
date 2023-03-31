// Package classfile
package classfile

/**
tree -L 1 .
.
├── attr_code.go                    # attributes of code
├── attr_constant_value.go          # attributes of constant value
├── attr_exceptions.go              # attributes of expections
├── attr_line_number_table.go       # attributes of line number for debuggin, tracing
├── attr_local_variable_table.go    # attributes of local variable table
├── attr_markers.go                 # ...
├── attr_source_file.go             # ...
├── attr_unparsed.go                # unsupported
├── attribute_info.go               # ...
├── class_file.go                   # parse class file
├── class_reader.go                 # helper function to read byte stream of class file
├── constant_info.go                # unit of constant pool
├── constant_pool.go                # constant pool table
├── cp_class.go                     # constant class info
├── cp_invoke_dynamic.go            # constant invoke dynamic info
├── cp_member_ref.go                # constant member ref info
├── cp_name_and_type.go             # constant name and type info
├── cp_numeric.go                   # constant (int, lone, float, double) info
├── cp_string.go                    # constant (bytes) info
├── cp_utf8.go                      # constant (name) info
└── member_info.go                  # represent fields and methods
*/
