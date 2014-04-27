// +build darwin

package cokego

// #cgo CFLAGS: -x objective-c -l/System/Library/Frameworks
// #cgo LDFLAGS: -framework Cocoa -framework Accelerate
// #include <QuartzCore/QuartzCore.h>
// #include <Cocoa/Cocoa.h>
// #include <Accelerate/Accelerate.h>
import "C"
import "unsafe"
import "log"
// type vImage_Buffer C.vImage_Buffer
// type vImage_Buffer struct {
//   data unsafe.Pointer
//   width C.vImagePixelCount
//   height C.vImagePixelCount
//   rowBytes C.size_t
// }

// type Pixel_8888  *C.Pixel_8888
// type uint8_t C.uint8_t


func VImageFlatten_BGRA8888ToRGB888( image *[]byte, width, height, rowBytes int ){
  // n := C.CString(string(image))

  buffer := &C.vImage_Buffer{
      data: unsafe.Pointer(&image),
      width:  C.vImagePixelCount(width),
      height: C.vImagePixelCount(height),
      rowBytes: C.size_t(rowBytes),
    }
  // b := make([]C.uint8_t, 4)
  var b C.uint8_t = 0
  //(*C.vImage_Buffer)(unsafe.Pointer(&buffer))
  ret := C.vImageFlatten_BGRA8888ToRGB888(
    (*C.vImage_Buffer)(unsafe.Pointer(&buffer)),
    (*C.vImage_Buffer)(unsafe.Pointer(&buffer)),
    (*C.uint8_t)(&b),
    false,
    C.kvImageDoNotTile)

  log.Println("ret:", ret )


}
 // kvImageNoError                  =   0,
 // kvImageRoiLargerThanInputBuffer =   -21766,
 // kvImageInvalidKernelSize        =   -21767,
 // kvImageInvalidEdgeStyle         =   -21768,
 // kvImageInvalidOffset_X          =   -21769,
 // kvImageInvalidOffset_Y          =   -21770,
 // kvImageMemoryAllocationError    =   -21771,
 // kvImageNullPointerArgument      =   -21772,
 // kvImageInvalidParameter         =   -21773,
 // kvImageBufferSizeMismatch       =   -21774,
 // kvImageUnknownFlagsBit           =    -21775


func VImageFlatten_RGBA8888ToRGB888( image *[]byte, width, height, rowBytes int ){
  // n := C.CString(string(image))

  buffer := &C.vImage_Buffer{
      data: unsafe.Pointer(image),
      width:  C.vImagePixelCount(width),
      height: C.vImagePixelCount(height),
      rowBytes: C.size_t(rowBytes),
    }
  // b := make([]C.uint8_t, 4)
  var b C.uint8_t = 0
  //(*C.vImage_Buffer)(unsafe.Pointer(&buffer))
  C.vImageFlatten_RGBA8888ToRGB888(
    (*C.vImage_Buffer)(unsafe.Pointer(&buffer)),
    (*C.vImage_Buffer)(unsafe.Pointer(&buffer)),
    (*C.uint8_t)(&b),
    false,
    C.kvImageDoNotTile)


}