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
	"github.com/gonum/matrix/mat64"
)

// init configures the standard logger to record filename and linenumbers
func init() {
	// log with filename:lineNr info
	log.SetFlags(log.LstdFlags | log.Lmicroseconds | log.Lshortfile)
}

/*
ResourceManager aggregates all the KPU resources into a single resource abstraction.

The Go OpenKL runtime is organized around a resource manager that manages one or more KPU devices.
The application may have information how it wants to use the different KPU devices to organize a computation.
In particular, multiple-kernel-multiple-data pipelines are good solutions to create processing pipelines
that scale in terms of performance and capacity.

To enable this use case, the KPU device abstraction allows low level operations on its resources. The OpenKL
runtime providing resource management implementations to make the first use of the KPU device easy and productive.

The application will initialize the Go OpenKL runtime, which will query the underlying hardware and build a
map of available resources for the application to leverage. The OpenKL runtime will return a fully initialized
resource manager that the application can use to query for KPU devices and their resources.

TODO: Can there be multiple resource managers active at the same time?
If yes, then we can use it to support multiple concurrent applications.
If no, then the RM becomes the arbiter.

If yes, then the underlying system needs to have the ability to report on resource availability AND context.
 */
type ResourceManager struct {
	id             string
	managedMemory  int64
	kpus           []KnowledgeProcessingUnit
	// if we have KPU resources with their own memory allocation API then what is the value of explicit memory abstraction
	memorySegments []*MemorySegments
}

// Create a new resource manager to manage a collection of data structures on behalf of the application
func (rm *ResourceManager) Initialize(id string) {
	log.Printf("Initializing Resource Manager %s\n", id)
	rm.id = id
	rm.kpus = make([]KnowledgeProcessingUnit, 10)   // hardcoding right now
	rm.memorySegments = make([]*MemorySegments, 10)
}

// Release all the data structure assets held by this resource manager
func (rm *ResourceManager) Release() {
	log.Printf("Releasing Resource Manager %s\n", rm.id)
}

func (rm *ResourceManager) GetRemote(id string) *KPU {
	return rm.kpus[0]
}

// New takes a data structure descriptor and allocates a new data structure, returning a handle
func (rm *ResourceManager) New(descriptor DataStructureDescriptor) (handle DataStructure, err error) {
	log.Printf("Allocating new data structure on KPU %s\n", rm.id)
	handle.Descriptor = descriptor
	return handle, nil
}

// Marshal takes an interface of a data structure and marshals the data structure to the kpu, returning a handle on success
func (rm *ResourceManager) Marshal(ds interface{}, targets []*KPU) (handle DataStructure, err error) {
	switch ds.(type){
	case *mat64.Dense:
		var matrix *mat64.Dense = ds.(*mat64.Dense)
		rows, cols := matrix.Dims()
		log.Printf("Marshalling a mat64.Dense data struture of size %dx%d\n", rows, cols)
		var descriptor DataStructureDescriptor = DataStructureDescriptor{Form:KL_DENSE_MATRIX, Elements:KL_FLOAT64, Bits:64}
		handle, err = rm.New(descriptor)

	}
	return
}

// NewDense creates a new dense matrix
func (rm *ResourceManager) NewDense(row, cols int, values []float64) {

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
	if ds.Descriptor.Form != KL_DENSE_MATRIX && ds.Descriptor.Elements != KL_FLOAT32 && ds.Descriptor.Bits != 32 {
		return nil, errors.New("Type mismatch: Unable to acquire Dense Int32 Matrix")
	}
	var mat [][]float32 = make([][]float32,ds.Descriptor.N)
	for i := uint64(0); i < ds.Descriptor.N; i++ {
		mat[i] = make([]float32, ds.Descriptor.M)
	}
	return mat, nil
}

func (rm *ResourceManager) AcquireDenseFloat64Matrix(ds DataStructure) ([][]float64, error) {
	if ds.Descriptor.Form != KL_DENSE_MATRIX && ds.Descriptor.Elements != KL_FLOAT64 && ds.Descriptor.Bits != 64 {
		return nil, errors.New("Type mismatch: Unable to acquire Dense Int32 Matrix")
	}
	var mat [][]float64 = make([][]float64,ds.Descriptor.N)
	for i := uint64(0); i < ds.Descriptor.N; i++ {
		mat[i] = make([]float64, ds.Descriptor.M)
	}
	return mat, nil
}