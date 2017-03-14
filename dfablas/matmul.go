/*
 File		:	$File: //depot/stillwater-sc/gokl/dfablas/matmul.go $

 Authors	:	E. Theodore L. Omtzigt
 Date		:	6 April 2014

 Source Control Information:
 Version	:	$Revision: #1 $
 Latest		:	$Date: 2014/04/06 $
 Location	:	$Id: //depot/stillwater-sc/gokl/dfablas/matmul.go#1 $

 Organization:
		Stillwater Supercomputing, Inc.
		P.O Box 720
		South Freeport, ME 04078-0720

Copyright (c) 2014-2017 E. Theodore L. Omtzigt.  All rights reserved.

Licence      : MIT license as defined in this directory

 */
package dfablas

import (
	"github.com/stillwater-sc/gokl"
	"log"
)

func GEMM(A, B, C, D gokl.DataStructure) (err error) {
	kpu := rm.GetRemote("KPU-0")
	// marshal the A and B matrices into the KPU fabric
	var targets []*gokl.KPU = make([]*gokl.KPU, 1)
	targets[0] = kpu
	Akpu, err := rm.Marshal(A, targets)		// marshal a CPU data structure to a set of target KPUs
	if err != nil {
		log.Fatal(err)
	}
	Bkpu, err := rm.Marshal(B, targets)
	if err != nil {
		log.Fatal(err)
	}
	// Create the zero type input C matrix
	var Cdescriptor gokl.DataStructureDescriptor = gokl.DataStructureDescriptor{Form:gokl.KL_DENSE_MATRIX, Elements:gokl.KL_FLOAT64, Bits:64, N:5, M:5, K:0}
	Ckpu, err := rm.New(Cdescriptor)
	if err != nil {
		log.Fatal(err)
	}
	// Create the output D matrix to receive the results from the domain flow computation
	var Ddescriptor gokl.DataStructureDescriptor = gokl.DataStructureDescriptor{Form:gokl.KL_DENSE_MATRIX, Elements:gokl.KL_FLOAT64, Bits:64, N:5, M:5, K:0}
	Dkpu, err := rm.New(Ddescriptor)
	if err != nil {
		log.Fatal(err)
	}
	// now all the input and outputs have been defined, allocated, and initialized, call the operator chain
	dfablas.GEMM_fused(Akpu, Bkpu, Ckpu, Dkpu)
	// marshal the result data set back to the CPU address space
	return nil
}
