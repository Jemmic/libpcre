package libpcre

func NotBOOL(x BOOL) BOOL {
	if x == 0 {
		return 1
	}

	return 0
}

//func InterfaceToByteSliceSlice(x interface{}) [][]byte {
//	return x.([][]byte)
//}

//func InterfaceToPcre_uint8SliceSlice(x interface{}) [][]pcre_uint8 {
//	return x.([][]pcre_uint8)
//}

//func InterfaceToPcre_ucharSliceSlice(x interface {}) [][]pcre_uchar {
//	return x.([][]pcre_uchar)
//}

//func Pcre_ucharSliceToUint8(x []pcre_uchar) []uint8 {
//	return *(*[]uint8)(unsafe.Pointer(UnsafeSliceToSlice(x, int32(1), int32(1))))
//}
