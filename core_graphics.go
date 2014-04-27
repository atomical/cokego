// +build darwin

package cokego

// #cgo CFLAGS: -x objective-c -l/System/Library/Frameworks
// #cgo LDFLAGS: -framework Cocoa
// #include <QuartzCore/QuartzCore.h>
// #include <Cocoa/Cocoa.h>
// void ** window_ids_to_int_array_bridge( CGWindowID *window_ids, int length){
//   CGWindowID *windows = malloc(length * sizeof(CGWindowID));
//   memcpy(windows, window_ids, length * sizeof(CGWindowID));
//   return (void **)(windows);
// }
import "C"
import "unsafe"

// Quartz Window Services Reference
// https://developer.apple.com/library/mac/documentation/Carbon/reference/CGWindow_Reference/Reference/Functions.html

const (
  KCGNullWindowID           = C.kCGNullWindowID
  KCGWindowSharingNone      = C.kCGWindowSharingNone
  KCGWindowSharingReadOnly  = C.kCGWindowSharingReadOnly
  KCGWindowSharingReadWrite = C.kCGWindowSharingReadWrite
)
//https://developer.apple.com/library/mac/documentation/graphicsimaging/reference/CGImage/Reference/reference.html
const (
  KCGBitmapAlphaInfoMask   = C.kCGBitmapAlphaInfoMask
  KCGImageAlphaFirst       = C.kCGImageAlphaFirst
  KCGImageAlphaLast      = C.kCGImageAlphaLast

)
// Window List Option Constants
// https://developer.apple.com/library/mac/documentation/Carbon/reference/CGWindow_Reference/Constants/Constants.html#//apple_ref/doc/constant_group/Window_List_Option_Constants

const (
  KCGWindowListOptionAll                 = C.kCGWindowListOptionAll
  KCGWindowListOptionOnScreenOnly        = C.kCGWindowListOptionOnScreenOnly
  KCGWindowListOptionOnScreenAboveWindow = C.kCGWindowListOptionOnScreenAboveWindow
  KCGWindowListOptionOnScreenBelowWindow = C.kCGWindowListOptionOnScreenBelowWindow
  KCGWindowListOptionIncludingWindow     = C.kCGWindowListOptionIncludingWindow
  KCGWindowListExcludeDesktopElements    = C.kCGWindowListExcludeDesktopElements
)

// Window Image Types
// https://developer.apple.com/library/mac/documentation/Carbon/reference/CGWindow_Reference/Constants/Constants.html#//apple_ref/doc/constant_group/Window_Image_Types
const (
  KCGWindowImageDefault             = C.kCGWindowImageDefault
  KCGWindowImageBoundsIgnoreFraming = C.kCGWindowImageBoundsIgnoreFraming
  KCGWindowImageShouldBeOpaque      = C.kCGWindowImageShouldBeOpaque
  KCGWindowImageOnlyShadows         = C.kCGWindowImageOnlyShadows
  KCGWindowImageBestResolution      = C.kCGWindowImageBestResolution
  KCGWindowImageNominalResolution   = C.kCGWindowImageNominalResolution
)


// const (
//   // KCGWindowOwnerName  = C.kCGWindowOwnerName
//   // KCGWindowWorkspace  = C.kCGWindowWorkspace
//   // CGWindowOwnerPID    = C.kCGWindowOwnerPID

//   // Required Window Keys
//   // const CFStringRef kCGWindowNumber;
//   // const CFStringRef kCGWindowStoreType;
//   // const CFStringRef kCGWindowLayer;
//   // const CFStringRef kCGWindowBounds;
//   // const CFStringRef kCGWindowSharingState;
//   // const CFStringRef kCGWindowAlpha;
//   // const CFStringRef kCGWindowOwnerPID;
//   // const CFStringRef kCGWindowMemoryUsage;
// )
type CGWindowID C.CGWindowID

type CGImageRef struct {
 Ref C.CGImageRef
 Data []byte
}

type Image struct {
  Width int
  Height int
  BytesPerRow int
  Data []byte
  DataLength int
}

type CaptureArea struct {
  Rect Rect
  WindowIds []CGWindowID
}

func MakeRect(x float64, y float64, w float64, h float64 ) C.CGRect{
  return C.CGRectMake(C.CGFloat(x), C.CGFloat(y), C.CGFloat(w), C.CGFloat(h))
}

type Rect struct {
  X float64
  Y float64
  Width float64
  Height float64
}


type Window struct {
  OwnerName string
  WindowId  CGWindowID
  Rect Rect
}

func CGWindowListCopyWindowInfo( option C.CGWindowListOption, relativeToWindow C.CGWindowID) ( []Window ) {

  listArray  := C.CGWindowListCopyWindowInfo( option, relativeToWindow )
  count      := CFArrayGetCount( listArray ) 
  windows    := make([]Window, count)

  for iter := 0; iter < count; iter++ {
    entry := C.CFArrayGetValueAtIndex(listArray, C.CFIndex(iter))
    name := CFStringGet(CFDictionaryGetValue(entry, unsafe.Pointer(C.kCGWindowOwnerName)))
    windowId := CFNumberGetValue(
                  C.CFNumberRef(CFDictionaryGetValue(entry, unsafe.Pointer(C.kCGWindowNumber))),
                  C.kCGWindowIDCFNumberType,
                )
    bounds_value := CFDictionaryGetValue(entry, unsafe.Pointer(C.kCGWindowBounds))
    rect := CGRectMakeWithDictionaryRepresentation(bounds_value)

    windows[iter] = Window{
                      OwnerName: name, 
                      WindowId: CGWindowID(windowId),
                      Rect: rect,
                    }

    // const CFStringRef kCGWindowNumber;
    // const CFStringRef kCGWindowStoreType;
    // const CFStringRef kCGWindowLayer;
    // const CFStringRef kCGWindowBounds;
    // const CFStringRef kCGWindowSharingState;
    // const CFStringRef kCGWindowAlpha;
    // const CFStringRef kCGWindowOwnerPID;
    // const CFStringRef kCGWindowMemoryUsage;

  }

  return windows
}

func CGImageGetWidth( image C.CGImageRef ) int {
  return int(C.CGImageGetWidth( image ))
}

func CGImageGetHeight( image C.CGImageRef ) int {
  return int(C.CGImageGetHeight( image ))
}

func CGRectMakeWithDictionaryRepresentation( dict unsafe.Pointer ) Rect {
  var rect Rect
  var CGRect C.CGRect
  
  C.CGRectMakeWithDictionaryRepresentation((*[0]byte)(dict), &CGRect)

  rect.X      = (float64)(C.CGRectGetMinX(   CGRect ))
  rect.Y      = (float64)(C.CGRectGetMinY(   CGRect ))
  rect.Width  = (float64)(C.CGRectGetWidth(  CGRect ))
  rect.Height = (float64)(C.CGRectGetHeight( CGRect ))

  return rect
}

func Capture( area CaptureArea ) Image {
  rect := MakeRect( area.Rect.X, area.Rect.Y, area.Rect.Width, area.Rect.Height )
  numberOfWindows := len(area.WindowIds)
  var image C.CGImageRef

  if numberOfWindows == 0 {

    image = C.CGWindowListCreateImage(
               rect,
               C.kCGWindowListOptionOnScreenOnly,
               C.kCGNullWindowID,
               C.kCGWindowImageBoundsIgnoreFraming,
              )

  } else {

    arrayRef := C.CFArrayCreate( 
                  C.kCFAllocatorDefault,
                  C.window_ids_to_int_array_bridge((*C.CGWindowID)(WindowIdSliceToCArray(area.WindowIds)), C.int(numberOfWindows)), //&ptr, //cb,
                  C.CFIndex(numberOfWindows), 
                  nil,
                  )

    // fmt.Println(area)
    // fmt.Println(rect)
    // fmt.Println("Array count:", C.CFArrayGetCount(arrayRef))
    // fmt.Println(C.CFArrayGetValueAtIndex(arrayRef, 0))
 
    image = C.CGWindowListCreateImageFromArray(
              C.CGRectNull, //rect,
              arrayRef,
              C.kCGWindowImageBoundsIgnoreFraming,
            )
  }


  var width = C.CGImageGetWidth(image)
  var height= C.CGImageGetHeight(image)
  var bitmapBytesPerRow  =  (width * 4)
  var bitmapByteCount    =  (bitmapBytesPerRow * height)
  var colorSpace = C.CGColorSpaceCreateDeviceRGB()

  context := C.CGBitmapContextCreate (nil,
              width,
              height,
              C.CGImageGetBitsPerComponent(image),
              bitmapBytesPerRow, //stride
              colorSpace,
              C.kCGImageAlphaPremultipliedLast)
  
  drawingRect := C.CGRectMake(0.0, 0.0,C.CGFloat(width), C.CGFloat(height))

  C.CGContextDrawImage(context, drawingRect, image)
  
  var cbytes *C.char = (*C.char)(C.CGBitmapContextGetData(context))
  var bytes = C.GoBytes(unsafe.Pointer(cbytes),C.int(bitmapByteCount))

  defer C.CGImageRelease(image)
  defer C.CGContextRelease(context)
  defer C.CGColorSpaceRelease(colorSpace)

  return Image{
    Width: int(width),
    Height: int(height),
    BytesPerRow: int(bitmapBytesPerRow),
    Data: bytes,
    DataLength: len(bytes),
  }
 
}

func CGWindowListCreateImageFromArray( windowIds []CGWindowID )  C.CGImageRef {
  numberOfWindows := len(windowIds)
  arrayRef := C.CFArrayCreate( 
                C.kCFAllocatorDefault,
                C.window_ids_to_int_array_bridge(
                  (*C.CGWindowID)(WindowIdSliceToCArray(windowIds)), 
                  C.int(numberOfWindows)),
                C.CFIndex(numberOfWindows), 
                nil)

  image := C.CGWindowListCreateImageFromArray(
            C.CGRectNull,
            arrayRef,
            C.kCGWindowImageBoundsIgnoreFraming,
          )

  return image
}

func CGImageGetBitmapInfo( image C.CGImageRef ) C.CGBitmapInfo {
  return C.CGImageGetBitmapInfo( image )     // CGImage may return pixels in RGBA, BGRA, or ARGB order
}
func CGImageGetBitsPerComponent( image C.CGImageRef ) C.size_t {
  return C.CGImageGetBitsPerComponent( image )
}
//CGColorSpaceGetModel(CGImageGetColorSpace(CGImage));
//size_t CGImageGetBitsPerPixel(CGImage)

func CGImageRefToGoBytes( image C.CGImageRef )  []byte {
  data := C.CGDataProviderCopyData(C.CGImageGetDataProvider(image))
  ptr := C.CFDataGetBytePtr(data)

  defer C.CGImageRelease(image)
  defer C.CFRelease((C.CFTypeRef)(data))
  
  return C.GoBytes( unsafe.Pointer(ptr), C.int(C.CFDataGetLength(data)))
}


//   data := C.CGDataProviderCopyData(C.CGImageGetDataProvider(image))
//   ptr := C.CFDataGetBytePtr(data)
//   len := C.CFDataGetLength(data)
//   bytes := C.GoBytes(unsafe.Pointer(ptr), C.int(len) )

//   defer C.CGImageRelease(image)
//   defer C.CFRelease((C.CFTypeRef)(data))

//   return CGImageRef{ Ref: image, Data: bytes }
// }

func WindowIdSliceToCArray (byteSlice []CGWindowID ) unsafe.Pointer {
       var array = unsafe.Pointer(C.calloc(C.size_t(len(byteSlice)), 1))
       var arrayptr = uintptr(array)

       for i := 0; i < len(byteSlice); i ++ {
              *(*CGWindowID )(unsafe.Pointer(arrayptr)) = CGWindowID(byteSlice[i])
              arrayptr ++
       }

       return array
}

func CGImageRelease( image C.CGImageRef ) {
  C.CGImageRelease(image)
}

