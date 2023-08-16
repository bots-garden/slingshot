; ModuleID = 'probe6.d6b3674448a3f7cd-cgu.0'
source_filename = "probe6.d6b3674448a3f7cd-cgu.0"
target datalayout = "e-m:e-i8:8:32-i16:16:32-i64:64-i128:128-n32:64-S128"
target triple = "aarch64-unknown-linux-gnu"

; core::f64::<impl f64>::is_subnormal
; Function Attrs: inlinehint nonlazybind uwtable
define internal zeroext i1 @"_ZN4core3f6421_$LT$impl$u20$f64$GT$12is_subnormal17hccd4b92f3a0f70c1E"(double %self) unnamed_addr #0 {
start:
  %_2 = alloca i8, align 1
  %0 = alloca i8, align 1
; call core::f64::<impl f64>::classify
  %1 = call i8 @"_ZN4core3f6421_$LT$impl$u20$f64$GT$8classify17he751604a427fc5bbE"(double %self), !range !2
  store i8 %1, ptr %_2, align 1
  %2 = load i8, ptr %_2, align 1, !range !2, !noundef !3
  %_3 = zext i8 %2 to i64
  %3 = icmp eq i64 %_3, 3
  br i1 %3, label %bb3, label %bb2

bb3:                                              ; preds = %start
  store i8 1, ptr %0, align 1
  br label %bb4

bb2:                                              ; preds = %start
  store i8 0, ptr %0, align 1
  br label %bb4

bb4:                                              ; preds = %bb3, %bb2
  %4 = load i8, ptr %0, align 1, !range !4, !noundef !3
  %5 = trunc i8 %4 to i1
  ret i1 %5
}

; probe6::probe
; Function Attrs: nonlazybind uwtable
define void @_ZN6probe65probe17h75c442db33cc3c3aE() unnamed_addr #1 {
start:
; call core::f64::<impl f64>::is_subnormal
  %_1 = call zeroext i1 @"_ZN4core3f6421_$LT$impl$u20$f64$GT$12is_subnormal17hccd4b92f3a0f70c1E"(double 1.000000e+00)
  ret void
}

; core::f64::<impl f64>::classify
; Function Attrs: nonlazybind uwtable
declare i8 @"_ZN4core3f6421_$LT$impl$u20$f64$GT$8classify17he751604a427fc5bbE"(double) unnamed_addr #1

attributes #0 = { inlinehint nonlazybind uwtable "target-cpu"="generic" "target-features"="+v8a,+outline-atomics" }
attributes #1 = { nonlazybind uwtable "target-cpu"="generic" "target-features"="+v8a,+outline-atomics" }

!llvm.module.flags = !{!0, !1}

!0 = !{i32 8, !"PIC Level", i32 2}
!1 = !{i32 2, !"RtLibUseGOT", i32 1}
!2 = !{i8 0, i8 5}
!3 = !{}
!4 = !{i8 0, i8 2}
