// Copyright 2018 The gVisor Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package hostarch

import (
	"fmt"
)

// Addr represents an address in an unspecified address space.
//
// +stateify savable
type Addr uintptr

// AddLength adds the given length to start and returns the result. ok is true
// iff adding the length did not overflow the range of Addr.
//
// Note: This function is usually used to get the end of an address range
// defined by its start address and length. Since the resulting end is
// exclusive, end == 0 is technically valid, and corresponds to a range that
// extends to the end of the address space, but ok will be false. This isn't
// expected to ever come up in practice.
func (v Addr) AddLength(length uint64) (end Addr, ok bool) {
	end = v + Addr(length)
	// The second half of the following check is needed in case uintptr is
	// smaller than 64 bits.
	ok = end >= v && length <= uint64(^Addr(0))
	return
}

// RoundDown is equivalent to function PageRoundDown.
func (v Addr) RoundDown() Addr {
	return PageRoundDown(v)
}

// RoundUp is equivalent to function PageRoundUp.
func (v Addr) RoundUp() (Addr, bool) {
	return PageRoundUp(v)
}

// MustRoundUp is equivalent to function MustPageRoundUp.
func (v Addr) MustRoundUp() Addr {
	return MustPageRoundUp(v)
}

// HugeRoundDown is equivalent to function HugePageRoundDown.
func (v Addr) HugeRoundDown() Addr {
	return HugePageRoundDown(v)
}

// HugeRoundUp is equivalent to function HugePageRoundUp.
func (v Addr) HugeRoundUp() (Addr, bool) {
	return HugePageRoundUp(v)
}

// PageOffset is equivalent to function PageOffset, except that it casts the
// result to uint64.
func (v Addr) PageOffset() uint64 {
	return uint64(PageOffset(v))
}

// IsPageAligned is equivalent to function IsPageAligned.
func (v Addr) IsPageAligned() bool {
	return IsPageAligned(v)
}

// AddrRange is a range of Addrs.
//
// type AddrRange <generated by go_generics>

// ToRange returns [v, v+length).
func (v Addr) ToRange(length uint64) (AddrRange, bool) {
	end, ok := v.AddLength(length)
	return AddrRange{v, end}, ok
}

// IsPageAligned returns true if ar.Start.IsPageAligned() and
// ar.End.IsPageAligned().
func (ar AddrRange) IsPageAligned() bool {
	return ar.Start.IsPageAligned() && ar.End.IsPageAligned()
}

// String implements fmt.Stringer.String.
func (ar AddrRange) String() string {
	return fmt.Sprintf("[%#x, %#x)", ar.Start, ar.End)
}
