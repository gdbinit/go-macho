// Code generated by "stringer -type=FieldDescriptorKind,FieldRecordFlags -output fields_string.go"; DO NOT EDIT.

package fields

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Struct-0]
	_ = x[Class-1]
	_ = x[Enum-2]
	_ = x[MultiPayloadEnum-3]
	_ = x[Protocol-4]
	_ = x[ClassProtocol-5]
	_ = x[ObjCProtocol-6]
	_ = x[ObjCClass-7]
}

const _FieldDescriptorKind_name = "StructClassEnumMultiPayloadEnumProtocolClassProtocolObjCProtocolObjCClass"

var _FieldDescriptorKind_index = [...]uint8{0, 6, 11, 15, 31, 39, 52, 64, 73}

func (i FieldDescriptorKind) String() string {
	if i >= FieldDescriptorKind(len(_FieldDescriptorKind_index)-1) {
		return "FieldDescriptorKind(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _FieldDescriptorKind_name[_FieldDescriptorKind_index[i]:_FieldDescriptorKind_index[i+1]]
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[IsIndirectCase-1]
	_ = x[IsVar-2]
	_ = x[IsArtificial-4]
}

const (
	_FieldRecordFlags_name_0 = "IsIndirectCaseIsVar"
	_FieldRecordFlags_name_1 = "IsArtificial"
)

var (
	_FieldRecordFlags_index_0 = [...]uint8{0, 14, 19}
)

func (i FieldRecordFlags) String() string {
	switch {
	case 1 <= i && i <= 2:
		i -= 1
		return _FieldRecordFlags_name_0[_FieldRecordFlags_index_0[i]:_FieldRecordFlags_index_0[i+1]]
	case i == 4:
		return _FieldRecordFlags_name_1
	default:
		return "FieldRecordFlags(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}
