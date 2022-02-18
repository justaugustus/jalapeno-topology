// Copyright (c) 2022 Cisco Systems, Inc. and its affiliates
// All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are
// met:
//
//     * Redistributions of source code must retain the above copyright
// notice, this list of conditions and the following disclaimer.
//
// The contents of this file are licensed under the Apache License, Version 2.0
// (the "License"); you may not use this file except in compliance with the
// License. You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations under
// the License.

package arangodb

// type StackItem interface {
// 	StackableItem()
// }

type FIFO interface {
	Push(DBRecord)
	Pop() DBRecord
	Len() int
}

type entry struct {
	next     *entry
	previous *entry
	data     DBRecord
}
type fifo struct {
	head *entry
	tail *entry
	len  int
}

func (f *fifo) Push(o DBRecord) {
	// Empty stack case
	e := &entry{
		next: f.tail,
		data: o,
	}
	if f.head == nil && f.tail == nil {
		f.head = e
	}
	f.tail = e
	if f.tail.next != nil {
		f.tail.next.previous = f.tail
	}
	f.len++
}

func (f *fifo) Pop() DBRecord {
	if f.head == nil {
		// Stack is empty
		return nil
	}
	data := f.head.data
	f.head = f.head.previous
	f.len--
	return data
}

func (f *fifo) Len() int {
	return f.len
}

func newFIFO() FIFO {
	return &fifo{
		head: nil,
		tail: nil,
	}
}
