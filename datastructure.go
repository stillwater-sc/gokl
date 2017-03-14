/*
 File		:	$File: //depot/stillwater-sc/gokl/datastructure.go $

 Authors	:	E. Theodore L. Omtzigt
 Date		:	6 April 2014

 Source Control Information:
 Version	:	$Revision: #1 $
 Latest		:	$Date: 2014/04/06 $
 Location	:	$Id: //depot/stillwater-sc/gokl/datastructure.go#1 $

 Organization:
		Stillwater Supercomputing, Inc.
		P.O Box 720
		South Freeport, ME 04078-0720

Copyright (c) 2014-2017 E. Theodore L. Omtzigt.  All rights reserved.

Licence      : MIT license as defined in this directory

 */
package gokl

type DataStructureIdentifier int

// Data Structure Identifiers
const (
	KL_DENSE_VECTOR DataStructureIdentifier = iota
	KL_DENSE_MATRIX
	KL_DENSE_MATRIX_3D
	KL_SPARSE_VECTOR
	KL_SPARSE_MATRIX
	KL_SPARSE_MATRIX_3D
	KL_TREE
	KL_LIST
	KL_MAP
	KL_SET
	KL_BLOOMFILTER
)

type BitStrucureIdentifier int

// Bit structure identifiers
const (
	KL_NAN BitStrucureIdentifier = iota
	KL_STRING
	KL_INTEGER
	KL_FLOAT32
	KL_FLOAT64
	KL_UNUM
	KL_POSIT
)

type DataStructureDescriptor struct {
	Form     DataStructureIdentifier
	Elements BitStrucureIdentifier
	Bits     int
	N,M,K	 uint64
}

// DataStructure represents an abstract handle to a data structure managed by the runtime
type DataStructure struct {
	Descriptor DataStructureDescriptor
}