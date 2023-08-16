; ModuleID = 'probe4.a5a8de2fde549693-cgu.0'
source_filename = "probe4.a5a8de2fde549693-cgu.0"
target datalayout = "e-m:e-i8:8:32-i16:16:32-i64:64-i128:128-n32:64-S128"
target triple = "aarch64-unknown-linux-gnu"

@alloc_15c1356ab95b0d2b98ea1d91ca33573b = private unnamed_addr constant <{ [75 x i8] }> <{ [75 x i8] c"/rustc/eb26296b556cef10fb713a38f3d16b9886080f26/library/core/src/num/mod.rs" }>, align 1
@alloc_42e4369f469ed050e0a451fe83604fcd = private unnamed_addr constant <{ ptr, [16 x i8] }> <{ ptr @alloc_15c1356ab95b0d2b98ea1d91ca33573b, [16 x i8] c"K\00\00\00\00\00\00\00~\04\00\00\05\00\00\00" }>, align 8
@str.0 = internal constant [25 x i8] c"attempt to divide by zero"

; probe4::probe
; Function Attrs: nonlazybind uwtable
define void @_ZN6probe45probe17h60e7ddd0b92b4af0E() unnamed_addr #0 {
start:
  %0 = call i1 @llvm.expect.i1(i1 false, i1 false)
  br i1 %0, label %panic.i, label %"_ZN4core3num21_$LT$impl$u20$u32$GT$10div_euclid17h1dccc8db2f3d1802E.exit"

panic.i:                                          ; preds = %start
; call core::panicking::panic
  call void @_ZN4core9panicking5panic17h398ac5fd442a9c3fE(ptr align 1 @str.0, i64 25, ptr align 8 @alloc_42e4369f469ed050e0a451fe83604fcd) #3
  unreachable

"_ZN4core3num21_$LT$impl$u20$u32$GT$10div_euclid17h1dccc8db2f3d1802E.exit": ; preds = %start
  ret void
}

; Function Attrs: nocallback nofree nosync nounwind willreturn memory(none)
declare i1 @llvm.expect.i1(i1, i1) #1

; core::panicking::panic
; Function Attrs: cold noinline noreturn nonlazybind uwtable
declare void @_ZN4core9panicking5panic17h398ac5fd442a9c3fE(ptr align 1, i64, ptr align 8) unnamed_addr #2

attributes #0 = { nonlazybind uwtable "target-cpu"="generic" "target-features"="+v8a,+outline-atomics" }
attributes #1 = { nocallback nofree nosync nounwind willreturn memory(none) }
attributes #2 = { cold noinline noreturn nonlazybind uwtable "target-cpu"="generic" "target-features"="+v8a,+outline-atomics" }
attributes #3 = { noreturn }

!llvm.module.flags = !{!0, !1}

!0 = !{i32 8, !"PIC Level", i32 2}
!1 = !{i32 2, !"RtLibUseGOT", i32 1}
