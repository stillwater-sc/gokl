/*
 File		:	$File: //depot/stillwater-sc/gokl/kpu.go $

 Authors	:	E. Theodore L. Omtzigt
 Date		:	13 March 2017

 Source Control Information:
 Version	:	$Revision: #1 $
 Latest		:	$Date: 2017/03/13 $
 Location	:	$Id: //depot/stillwater-sc/gokl/kpu.go#1 $

 Organization:
		Stillwater Supercomputing, Inc.
		P.O Box 720
		South Freeport, ME 04078-0720

Copyright (c) 2014-2017 E. Theodore L. Omtzigt.  All rights reserved.

Licence      : MIT license as defined in this directory

 */
package gokl

import "log"

type KnowledgeProcessingUnit interface {
	// Device Management
	Intialize()
	Reset()
	Release()

	// Program Management
	LoadKernel(name string, program DomainFlowProgram)
	ReleaseKernel(name string)

	// Memory and Data Structure Management
	AllocateMemory(md MemoryDescriptor)
	LoadData(name string, ds *DataStructure)
	AcquireData(name string, ds *DataStructure)
	ReleaseMemory(md MemoryDescriptor)

	// Notifier Management
	AllocateNotifier(name string)
	ReleaseNotifier()

}

type KPU struct {
	name string
}

func NewKPU (name string) *KPU {
	var kpu *KPU = &KPU{name:name}
	return kpu
}

func (kpu *KPU) Initialize() {
	log.Printf("Initialize KPU %s\n", kpu.name)
}

func (kpu *KPU) Reset() {
	log.Printf("Reset KPU %s\n", kpu.name)
}

func (kpu *KPU) Release() {
	log.Printf("Release KPU %s\n", kpu.name)
}

func (kpu *KPU) LoadKernel() {
	log.Printf("LoadKernel KPU %s\n", kpu.name)
}

func (kpu *KPU) ReleaseKernel() {
	log.Printf("ReleaseKernel KPU %s\n", kpu.name)
}

func (kpu *KPU) AllocateMemory() {
	log.Printf("AllocateMemory on KPU %s\n", kpu.name)
}

func (kpu *KPU) LoadData() {
	log.Printf("LoadData on KPU %s\n", kpu.name)
}

func (kpu *KPU) AcquireData() {
	log.Printf("AcquireData from KPU %s\n", kpu.name)
}

func (kpu *KPU) ReleaseMemory() {
	log.Printf("ReleaseMemory KPU %s\n", kpu.name)
}

func (kpu *KPU) AllocateNotifier(name string) {
	log.Printf("AllocateNotifier on KPU %s\n", kpu.name)
}

func (kpu * KPU) ReleaseNotifier() {
	log.Printf("ReleaseNotifier on KPU %s\n", kpu.name)
}