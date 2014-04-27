// +build darwin

package cokego

// #cgo CFLAGS: -x objective-c -l/System/Library/Frameworks
// #cgo LDFLAGS: -framework Cocoa 
// #include <Cocoa/Cocoa.h>
// #include <mach/mach_vm.h>
// #include <sys/types.h>
// #include <mach/mach_init.h>
// #include <mach/mach_traps.h>
// #include <mach/mach_types.h>
// #include <mach/mach_vm.h>
// #include <mach/mach_error.h>
// #define _ALIGN_UP(addr, size)   (((addr)+((size)-1))&(~((size)-1)))
// mach_vm_address_t align_up( int address, int size){
//   return _ALIGN_UP(address, size);
// }
// mach_port_name_t mach_task_self_alias(){
//   return mach_task_self();
// }
// uintptr_t page_down( uintptr_t addr, int page_size){
//   return ((addr)&(~((page_size)-1)));
// }
// uintptr_t page_align_up( uintptr_t addr, int size ){
//   return (((addr)+((size)-1))&(~((size)-1)));
// }
// vm_offset_t * convert_to_ptr( vm_offset_t *data ){
//   return *(vm_offset_t *)data;
// }
import "C"

import (
  "unsafe"
  "fmt"
)

const (
  KERN_SUCCESS = C.KERN_SUCCESS
  PAGE_SIZE    = 4096
)

func VMCopy( pid int, address uintptr, length int )  []byte  {

  var task C.mach_port_name_t
  ret := C.task_for_pid(C.mach_task_self_alias(), C.int(pid), &task)

  if ret != KERN_SUCCESS {
    panic(formatMachError(ret))
  }
  
  var dataCount C.mach_msg_type_number_t //mach_msg_type_number_t
  bottom := C.page_down(C.uintptr_t(address), C.int(PAGE_SIZE))
  top    := C.page_align_up(C.uintptr_t(address) + C.uintptr_t(length), C.int(PAGE_SIZE))
  length = int(top - bottom)

  offset := int(address) - int(bottom)


  if int(top) % PAGE_SIZE != 0 || int(bottom) % PAGE_SIZE != 0 {
    panic(fmt.Sprintf("Address is not a multiple of %v: (%v, %v)", PAGE_SIZE, top, bottom))
  } 

  data := (*C.vm_offset_t)(C.malloc(C.size_t(length)))
  example := "ABC"
  C.memcpy(unsafe.Pointer(data), unsafe.Pointer(&example), 3)
  // defer C.free(unsafe.Pointer(data))
  // var data C.mach_vm_address_t = 0

  // ret = C.mach_vm_allocate(
  //         (C.vm_map_t)(C.mach_task_self_alias()), 
  //         &data, 
  //         C.mach_vm_size_t(length), 
  //         C.int(1))

  // if ret != KERN_SUCCESS {
  //   panic(formatMachError(ret))
  // }


  // fmt.Println("Length:", length)

  ret = C.mach_vm_read( (C.vm_map_t)(task), 
                         (C.mach_vm_address_t)(bottom),
                         (C.mach_vm_size_t)(length), 
                         (*C.vm_offset_t)(unsafe.Pointer(&data)),  // (*C.vm_offset_t)(unsafe.Pointer(&data)),
                         &dataCount)

  // var dataTemp [35000]C.char
  // ret = C.mach_vm_read_overwrite( (C.vm_map_t)(task), 
  //                        (C.mach_vm_address_t)(bottom),
  //                        (C.mach_vm_size_t)(length), 
  //                       (C.mach_vm_address_t)(data),  // (*C.vm_offset_t)(unsafe.Pointer(&data)),
  //                        &dataCount)

  if ret != KERN_SUCCESS {
    panic(formatMachError(ret))
  }
  
  // fmt.Println("Data:", *(data))

  // data2 := (*C.vm_offset_t)(C.malloc(C.size_t(length)))
  // C.memcpy(unsafe.Pointer(data2), unsafe.Pointer(&data), C.size_t(length))
  offset = offset
  return C.GoBytes(unsafe.Pointer(&*data), C.int(length))[offset:]

}

func formatMachError( ret C.kern_return_t ) string {
  return C.GoString(C.mach_error_string(C.mach_error_t(ret)))
}
