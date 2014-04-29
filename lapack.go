// Do not manually edit this file. It was created by the genLapack.pl script from clapack.h.

// Copyright ©2014 The gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package lapack

/*
#cgo CFLAGS: -g -O2 -fPIC -m64 -pthread
#cgo LDFLAGS: -L/usr/lib/ -llapack -lblas
#include "cblas.h"
#include "clapack.h"
*/
import "C"

import (
	"github.com/gonum/blas"
	"unsafe"
)

func Sgesv(o blas.Order, n int, nRHS int, a []float32, lda int, ipiv []int32, b []float32, ldb int) int {
	return int(C.clapack_sgesv(C.enum_CBLAS_ORDER(o), C.int(n), C.int(nRHS), (*C.float)(&a[0]), C.int(lda), (*C.int)(&ipiv[0]), (*C.float)(&b[0]), C.int(ldb)))
}
func Sgetrf(o blas.Order, m int, n int, a []float32, lda int, ipiv []int32) int {
	return int(C.clapack_sgetrf(C.enum_CBLAS_ORDER(o), C.int(m), C.int(n), (*C.float)(&a[0]), C.int(lda), (*C.int)(&ipiv[0])))
}
func Sgetrs(o blas.Order, t blas.Transpose, n int, nRHS int, a []float32, lda int, ipiv []int32, b []float32, ldb int) int {
	return int(C.clapack_sgetrs(C.enum_CBLAS_ORDER(o), C.enum_CBLAS_TRANSPOSE(t), C.int(n), C.int(nRHS), (*C.float)(&a[0]), C.int(lda), (*C.int)(&ipiv[0]), (*C.float)(&b[0]), C.int(ldb)))
}
func Sgetri(o blas.Order, n int, a []float32, lda int, ipiv []int32) int {
	return int(C.clapack_sgetri(C.enum_CBLAS_ORDER(o), C.int(n), (*C.float)(&a[0]), C.int(lda), (*C.int)(&ipiv[0])))
}
func Sposv(o blas.Order, ul blas.Uplo, n int, nRHS int, a []float32, lda int, b []float32, ldb int) int {
	return int(C.clapack_sposv(C.enum_ATLAS_ORDER(o), C.enum_ATLAS_UPLO(ul), C.int(n), C.int(nRHS), (*C.float)(&a[0]), C.int(lda), (*C.float)(&b[0]), C.int(ldb)))
}
func Spotrf(o blas.Order, ul blas.Uplo, n int, a []float32, lda int) int {
	return int(C.clapack_spotrf(C.enum_ATLAS_ORDER(o), C.enum_ATLAS_UPLO(ul), C.int(n), (*C.float)(&a[0]), C.int(lda)))
}
func Spotrs(o blas.Order, ul blas.Uplo, n int, nRHS int, a []float32, lda int, b []float32, ldb int) int {
	return int(C.clapack_spotrs(C.enum_CBLAS_ORDER(o), C.enum_CBLAS_UPLO(ul), C.int(n), C.int(nRHS), (*C.float)(&a[0]), C.int(lda), (*C.float)(&b[0]), C.int(ldb)))
}
func Spotri(o blas.Order, ul blas.Uplo, n int, a []float32, lda int) int {
	return int(C.clapack_spotri(C.enum_ATLAS_ORDER(o), C.enum_ATLAS_UPLO(ul), C.int(n), (*C.float)(&a[0]), C.int(lda)))
}
func Slauum(o blas.Order, ul blas.Uplo, n int, a []float32, lda int) int {
	return int(C.clapack_slauum(C.enum_ATLAS_ORDER(o), C.enum_ATLAS_UPLO(ul), C.int(n), (*C.float)(&a[0]), C.int(lda)))
}
func Strtri(o blas.Order, ul blas.Uplo, d blas.Diag, n int, a []float32, lda int) int {
	return int(C.clapack_strtri(C.enum_ATLAS_ORDER(o), C.enum_ATLAS_UPLO(ul), C.enum_ATLAS_DIAG(d), C.int(n), (*C.float)(&a[0]), C.int(lda)))
}
func Dgesv(o blas.Order, n int, nRHS int, a []float64, lda int, ipiv []int32, b []float64, ldb int) int {
	return int(C.clapack_dgesv(C.enum_CBLAS_ORDER(o), C.int(n), C.int(nRHS), (*C.double)(&a[0]), C.int(lda), (*C.int)(&ipiv[0]), (*C.double)(&b[0]), C.int(ldb)))
}
func Dgetrf(o blas.Order, m int, n int, a []float64, lda int, ipiv []int32) int {
	return int(C.clapack_dgetrf(C.enum_CBLAS_ORDER(o), C.int(m), C.int(n), (*C.double)(&a[0]), C.int(lda), (*C.int)(&ipiv[0])))
}
func Dgetrs(o blas.Order, t blas.Transpose, n int, nRHS int, a []float64, lda int, ipiv []int32, b []float64, ldb int) int {
	return int(C.clapack_dgetrs(C.enum_CBLAS_ORDER(o), C.enum_CBLAS_TRANSPOSE(t), C.int(n), C.int(nRHS), (*C.double)(&a[0]), C.int(lda), (*C.int)(&ipiv[0]), (*C.double)(&b[0]), C.int(ldb)))
}
func Dgetri(o blas.Order, n int, a []float64, lda int, ipiv []int32) int {
	return int(C.clapack_dgetri(C.enum_CBLAS_ORDER(o), C.int(n), (*C.double)(&a[0]), C.int(lda), (*C.int)(&ipiv[0])))
}
func Dposv(o blas.Order, ul blas.Uplo, n int, nRHS int, a []float64, lda int, b []float64, ldb int) int {
	return int(C.clapack_dposv(C.enum_ATLAS_ORDER(o), C.enum_ATLAS_UPLO(ul), C.int(n), C.int(nRHS), (*C.double)(&a[0]), C.int(lda), (*C.double)(&b[0]), C.int(ldb)))
}
func Dpotrf(o blas.Order, ul blas.Uplo, n int, a []float64, lda int) int {
	return int(C.clapack_dpotrf(C.enum_ATLAS_ORDER(o), C.enum_ATLAS_UPLO(ul), C.int(n), (*C.double)(&a[0]), C.int(lda)))
}
func Dpotrs(o blas.Order, ul blas.Uplo, n int, nRHS int, a []float64, lda int, b []float64, ldb int) int {
	return int(C.clapack_dpotrs(C.enum_CBLAS_ORDER(o), C.enum_CBLAS_UPLO(ul), C.int(n), C.int(nRHS), (*C.double)(&a[0]), C.int(lda), (*C.double)(&b[0]), C.int(ldb)))
}
func Dpotri(o blas.Order, ul blas.Uplo, n int, a []float64, lda int) int {
	return int(C.clapack_dpotri(C.enum_ATLAS_ORDER(o), C.enum_ATLAS_UPLO(ul), C.int(n), (*C.double)(&a[0]), C.int(lda)))
}
func Dlauum(o blas.Order, ul blas.Uplo, n int, a []float64, lda int) int {
	return int(C.clapack_dlauum(C.enum_ATLAS_ORDER(o), C.enum_ATLAS_UPLO(ul), C.int(n), (*C.double)(&a[0]), C.int(lda)))
}
func Dtrtri(o blas.Order, ul blas.Uplo, d blas.Diag, n int, a []float64, lda int) int {
	return int(C.clapack_dtrtri(C.enum_ATLAS_ORDER(o), C.enum_ATLAS_UPLO(ul), C.enum_ATLAS_DIAG(d), C.int(n), (*C.double)(&a[0]), C.int(lda)))
}
func Cgesv(o blas.Order, n int, nRHS int, a []complex64, lda int, ipiv []int32, b []complex64, ldb int) int {
	return int(C.clapack_cgesv(C.enum_CBLAS_ORDER(o), C.int(n), C.int(nRHS), unsafe.Pointer(&a[0]), C.int(lda), (*C.int)(&ipiv[0]), unsafe.Pointer(&b[0]), C.int(ldb)))
}
func Cgetrf(o blas.Order, m int, n int, a []complex64, lda int, ipiv []int32) int {
	return int(C.clapack_cgetrf(C.enum_CBLAS_ORDER(o), C.int(m), C.int(n), unsafe.Pointer(&a[0]), C.int(lda), (*C.int)(&ipiv[0])))
}
func Cgetrs(o blas.Order, t blas.Transpose, n int, nRHS int, a []complex64, lda int, ipiv []int32, b []complex64, ldb int) int {
	return int(C.clapack_cgetrs(C.enum_CBLAS_ORDER(o), C.enum_CBLAS_TRANSPOSE(t), C.int(n), C.int(nRHS), unsafe.Pointer(&a[0]), C.int(lda), (*C.int)(&ipiv[0]), unsafe.Pointer(&b[0]), C.int(ldb)))
}
func Cgetri(o blas.Order, n int, a []complex64, lda int, ipiv []int32) int {
	return int(C.clapack_cgetri(C.enum_CBLAS_ORDER(o), C.int(n), unsafe.Pointer(&a[0]), C.int(lda), (*C.int)(&ipiv[0])))
}
func Cposv(o blas.Order, ul blas.Uplo, n int, nRHS int, a []complex64, lda int, b []complex64, ldb int) int {
	return int(C.clapack_cposv(C.enum_ATLAS_ORDER(o), C.enum_ATLAS_UPLO(ul), C.int(n), C.int(nRHS), unsafe.Pointer(&a[0]), C.int(lda), unsafe.Pointer(&b[0]), C.int(ldb)))
}
func Cpotrf(o blas.Order, ul blas.Uplo, n int, a []complex64, lda int) int {
	return int(C.clapack_cpotrf(C.enum_ATLAS_ORDER(o), C.enum_ATLAS_UPLO(ul), C.int(n), unsafe.Pointer(&a[0]), C.int(lda)))
}
func Cpotrs(o blas.Order, ul blas.Uplo, n int, nRHS int, a []complex64, lda int, b []complex64, ldb int) int {
	return int(C.clapack_cpotrs(C.enum_CBLAS_ORDER(o), C.enum_CBLAS_UPLO(ul), C.int(n), C.int(nRHS), unsafe.Pointer(&a[0]), C.int(lda), unsafe.Pointer(&b[0]), C.int(ldb)))
}
func Cpotri(o blas.Order, ul blas.Uplo, n int, a []complex64, lda int) int {
	return int(C.clapack_cpotri(C.enum_ATLAS_ORDER(o), C.enum_ATLAS_UPLO(ul), C.int(n), unsafe.Pointer(&a[0]), C.int(lda)))
}
func Clauum(o blas.Order, ul blas.Uplo, n int, a []complex64, lda int) int {
	return int(C.clapack_clauum(C.enum_ATLAS_ORDER(o), C.enum_ATLAS_UPLO(ul), C.int(n), unsafe.Pointer(&a[0]), C.int(lda)))
}
func Ctrtri(o blas.Order, ul blas.Uplo, d blas.Diag, n int, a []complex64, lda int) int {
	return int(C.clapack_ctrtri(C.enum_ATLAS_ORDER(o), C.enum_ATLAS_UPLO(ul), C.enum_ATLAS_DIAG(d), C.int(n), unsafe.Pointer(&a[0]), C.int(lda)))
}
func Zgesv(o blas.Order, n int, nRHS int, a []complex128, lda int, ipiv []int32, b []complex128, ldb int) int {
	return int(C.clapack_zgesv(C.enum_CBLAS_ORDER(o), C.int(n), C.int(nRHS), unsafe.Pointer(&a[0]), C.int(lda), (*C.int)(&ipiv[0]), unsafe.Pointer(&b[0]), C.int(ldb)))
}
func Zgetrf(o blas.Order, m int, n int, a []complex128, lda int, ipiv []int32) int {
	return int(C.clapack_zgetrf(C.enum_CBLAS_ORDER(o), C.int(m), C.int(n), unsafe.Pointer(&a[0]), C.int(lda), (*C.int)(&ipiv[0])))
}
func Zgetrs(o blas.Order, t blas.Transpose, n int, nRHS int, a []complex128, lda int, ipiv []int32, b []complex128, ldb int) int {
	return int(C.clapack_zgetrs(C.enum_CBLAS_ORDER(o), C.enum_CBLAS_TRANSPOSE(t), C.int(n), C.int(nRHS), unsafe.Pointer(&a[0]), C.int(lda), (*C.int)(&ipiv[0]), unsafe.Pointer(&b[0]), C.int(ldb)))
}
func Zgetri(o blas.Order, n int, a []complex128, lda int, ipiv []int32) int {
	return int(C.clapack_zgetri(C.enum_CBLAS_ORDER(o), C.int(n), unsafe.Pointer(&a[0]), C.int(lda), (*C.int)(&ipiv[0])))
}
func Zposv(o blas.Order, ul blas.Uplo, n int, nRHS int, a []complex128, lda int, b []complex128, ldb int) int {
	return int(C.clapack_zposv(C.enum_ATLAS_ORDER(o), C.enum_ATLAS_UPLO(ul), C.int(n), C.int(nRHS), unsafe.Pointer(&a[0]), C.int(lda), unsafe.Pointer(&b[0]), C.int(ldb)))
}
func Zpotrf(o blas.Order, ul blas.Uplo, n int, a []complex128, lda int) int {
	return int(C.clapack_zpotrf(C.enum_ATLAS_ORDER(o), C.enum_ATLAS_UPLO(ul), C.int(n), unsafe.Pointer(&a[0]), C.int(lda)))
}
func Zpotrs(o blas.Order, ul blas.Uplo, n int, nRHS int, a []complex128, lda int, b []complex128, ldb int) int {
	return int(C.clapack_zpotrs(C.enum_CBLAS_ORDER(o), C.enum_CBLAS_UPLO(ul), C.int(n), C.int(nRHS), unsafe.Pointer(&a[0]), C.int(lda), unsafe.Pointer(&b[0]), C.int(ldb)))
}
func Zpotri(o blas.Order, ul blas.Uplo, n int, a []complex128, lda int) int {
	return int(C.clapack_zpotri(C.enum_ATLAS_ORDER(o), C.enum_ATLAS_UPLO(ul), C.int(n), unsafe.Pointer(&a[0]), C.int(lda)))
}
func Zlauum(o blas.Order, ul blas.Uplo, n int, a []complex128, lda int) int {
	return int(C.clapack_zlauum(C.enum_ATLAS_ORDER(o), C.enum_ATLAS_UPLO(ul), C.int(n), unsafe.Pointer(&a[0]), C.int(lda)))
}
func Ztrtri(o blas.Order, ul blas.Uplo, d blas.Diag, n int, a []complex128, lda int) int {
	return int(C.clapack_ztrtri(C.enum_ATLAS_ORDER(o), C.enum_ATLAS_UPLO(ul), C.enum_ATLAS_DIAG(d), C.int(n), unsafe.Pointer(&a[0]), C.int(lda)))
}
