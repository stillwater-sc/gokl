/*
 File		:	$File: //depot/stillwater-sc/gokl/resourcemanager.go $

 Authors	:	E. Theodore L. Omtzigt
 Date		:	6 April 2014

 Source Control Information:
 Version	:	$Revision: #1 $
 Latest		:	$Date: 2014/04/06 $
 Location	:	$Id: //depot/stillwater-sc/gokl/resourcemanager.go#1 $

 Organization:
		Stillwater Supercomputing, Inc.
		P.O Box 720
		South Freeport, ME 04078-0720

Copyright (c) 2014-2017 E. Theodore L. Omtzigt.  All rights reserved.

Licence      : MIT license as defined in this directory

 */
package gokl

import (
	"log"
	"errors"
)

/*
 Can there be multiple resource managers active at the same time?
 If yes, then we can use it to support multiple concurrent applications.
 If no, then the RM becomes the arbiter.

 If yes, then the underlying system needs to have the ability to report
 on resource availability AND context.
 */
type ResourceManager struct {
	id             string
	managedMemory  int64
	memorySegments []MemorySegments
}

// Create a new resource manager to manage a collection of data structures on behalf of the application
func (rm *ResourceManager) Initialize(id string) {
	log.Printf("Initializing Resource Manager %s\n", id)
	rm.id = id
}

// Release all the data structure assets held by this resource manager
func (rm *ResourceManager) Release() {
	log.Printf("Releasing Resource Manager %s\n", rm.id)
}

// New takes a data structure descriptor and allocates a new data structure, returning a handle
func (rm *ResourceManager) New(descriptor DataStructureDescriptor) (handle DataStructure, err error) {
	handle.Descriptor = descriptor
	return handle, nil
}

// Free takes a dsta structure and deallocates it from the resource list, returning error if unable
func (rm *ResourceManager) Free(descriptor DataStructure) error {

	return nil
}

func (rm *ResourceManager) AcquireDenseInt32Matrix(ds DataStructure) ([][]int32, error) {
	if ds.Descriptor.Form != KL_DENSE_MATRIX && ds.Descriptor.Elements != KL_INTEGER && ds.Descriptor.Bits != 32 {
		return nil, errors.New("Type mismatch: Unable to acquire Dense Int32 Matrix")
	}
	var mat [][]int32 = make([][]int32,ds.Descriptor.N)
	for i := uint64(0); i < ds.Descriptor.N; i++ {
		mat[i] = make([]int32, ds.Descriptor.M)
	}
	return mat, nil
}

func (rm *ResourceManager) AcquireDenseInt64Matrix(ds DataStructure) ([][]int64, error) {
	if ds.Descriptor.Form != KL_DENSE_MATRIX && ds.Descriptor.Elements != KL_INTEGER && ds.Descriptor.Bits != 64 {
		return nil, errors.New("Type mismatch: Unable to acquire Dense Int32 Matrix")
	}
	var mat [][]int64 = make([][]int64,ds.Descriptor.N)
	for i := uint64(0); i < ds.Descriptor.N; i++ {
		mat[i] = make([]int64, ds.Descriptor.M)
	}
	return mat, nil
}

func (rm *ResourceManager) AcquireDenseFloat32Matrix(ds DataStructure) ([][]float32, error) {
	if ds.Descriptor.Form != KL_DENSE_MATRIX && ds.Descriptor.Elements != KL_FLOAT && ds.Descriptor.Bits != 32 {
		return nil, errors.New("Type mismatch: Unable to acquire Dense Int32 Matrix")
	}
	var mat [][]float32 = make([][]float32,ds.Descriptor.N)
	for i := uint64(0); i < ds.Descriptor.N; i++ {
		mat[i] = make([]float32, ds.Descriptor.M)
	}
	return mat, nil
}

func (rm *ResourceManager) AcquireDenseFloat64Matrix(ds DataStructure) ([][]float64, error) {
	if ds.Descriptor.Form != KL_DENSE_MATRIX && ds.Descriptor.Elements != KL_FLOAT && ds.Descriptor.Bits != 64 {
		return nil, errors.New("Type mismatch: Unable to acquire Dense Int32 Matrix")
	}
	var mat [][]float64 = make([][]float64,ds.Descriptor.N)
	for i := uint64(0); i < ds.Descriptor.N; i++ {
		mat[i] = make([]float64, ds.Descriptor.M)
	}
	return mat, nil
}