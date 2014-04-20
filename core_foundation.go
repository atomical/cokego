// +build darwin

package cokego

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework Cocoa
// #include <QuartzCore/QuartzCore.h>
//
// void * DictionaryObjectForKey(NSDictionary *dict, id key ){
//   return [dict objectForKey: key];
// }
//
// static NSString* GetDictValue( NSDictionary* dict, CFStringRef key )
// {
//     if ([dict objectForKey:(id)kCGWindowOwnerPID])
//     {
//        NSString *string = 
//          [NSString stringWithFormat:@"(%@)", [dict objectForKey:(id)key]];
//        return string;
//      }
//     else 
//     {
//         return nil;
//     }
// }
//
// const char * NSStringToCString( NSString * str ){
//   const char * c_str = [str UTF8String];
//   return c_str;
// }
import "C"
import "unsafe"

func CFDictionaryContainsKey( entry unsafe.Pointer, key unsafe.Pointer ) int {
  return (int(C.CFDictionaryContainsKey((*[0]byte)(entry),
    unsafe.Pointer(C.kCGWindowOwnerName),
  )))
}

func CFDictionaryGetValue( entry unsafe.Pointer, key unsafe.Pointer ) unsafe.Pointer {
  return C.CFDictionaryGetValue( (*[0]byte)(entry), key )
}

func CFNumberGetValue( number C.CFNumberRef, theType C.CFNumberType ) int {
  var val int
  C.CFNumberGetValue ( number, theType, unsafe.Pointer(&val) )
  return val
}

func CFStringGet( ptr unsafe.Pointer ) string {
  length := C.CFStringGetLength((*[0]byte)(ptr))
  buffer := C.malloc(C.size_t(length + 1))

  C.CFStringGetCString( 
    (*[0]byte)(ptr), 
    (*C.char)(unsafe.Pointer(buffer)),
    length + 1,
    C.kCFStringEncodingUTF8 )

  str := C.GoStringN((*C.char)(buffer), C.int(length))
  defer C.free(buffer)

  return str
}

// CFArray Reference
// https://developer.apple.com/library/mac/documentation/corefoundation/Reference/CFArrayRef/Reference/reference.html

func CFArrayGetCount( theArray C.CFArrayRef ) ( int ) {
  return int(C.CFArrayGetCount( theArray ))
}

func CFArrayGetValueAtIndex( theArray C.CFArrayRef, idx C.CFIndex ) ( unsafe.Pointer ) {
  return C.CFArrayGetValueAtIndex(theArray, idx)
}

func NSStringToGoString( str *C.NSString ) string {
  return C.GoString(C.NSStringToCString( str ))
}

func CFRelease( p unsafe.Pointer ) {
  C.CFRelease((C.CFTypeRef)(p))
}
