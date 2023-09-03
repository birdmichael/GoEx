package tupleext

// Tuple 是一个包含两个值的元组结构。
//
// 参数：
//   - T1: 第一个元素的类型。
//   - T2: 第二个元素的类型。
type Tuple[T1, T2 any] struct {
	S1 T1 // 第一个元素
	S2 T2 // 第二个元素
}

// Tuple3 是一个包含三个值的元组结构。
//
// 参数：
//   - T1: 第一个元素的类型。
//   - T2: 第二个元素的类型。
//   - T3: 第三个元素的类型。
type Tuple3[T1, T2, T3 any] struct {
	S1 T1 // 第一个元素
	S2 T2 // 第二个元素
	S3 T3 // 第三个元素
}

// Tuple4 是一个包含四个值的元组结构。
//
// 参数：
//   - T1: 第一个元素的类型。
//   - T2: 第二个元素的类型。
//   - T3: 第三个元素的类型。
//   - T4: 第四个元素的类型。
type Tuple4[T1, T2, T3, T4 any] struct {
	S1 T1
	S2 T2
	S3 T3
	S4 T4
}

// Tuple5 是一个包含五个值的元组结构。
//
// 参数：
//   - T1: 第一个元素的类型。
//   - T2: 第二个元素的类型。
//   - T3: 第三个元素的类型。
//   - T4: 第四个元素的类型。
//   - T5: 第五个元素的类型。
type Tuple5[T1, T2, T3, T4, T5 any] struct {
	S1 T1
	S2 T2
	S3 T3
	S4 T4
	S5 T5
}

// Tuple6 是一个包含六个值的元组结构。
//
// 参数：
//   - T1: 第一个元素的类型。
//   - T2: 第二个元素的类型。
//   - T3: 第三个元素的类型。
//   - T4: 第四个元素的类型。
//   - T5: 第五个元素的类型。
//   - T6: 第六个元素的类型。
type Tuple6[T1, T2, T3, T4, T5, T6 any] struct {
	S1 T1
	S2 T2
	S3 T3
	S4 T4
	S5 T5
	S6 T6
}
