// +build darwin

package cokego

// #cgo CFLAGS: -x objective-c -l/System/Library/Frameworks
// #cgo LDFLAGS: -framework Cocoa
// #include <QuartzCore/QuartzCore.h>
// #include <Cocoa/Cocoa.h>
// //http://stackoverflow.com/questions/8348627/what-is-the-correct-way-to-identify-the-currently-active-application-in-osx-10
// NSString * get_active_application(){
//   for (NSRunningApplication *currApp in [[NSWorkspace sharedWorkspace] runningApplications]) {
//       if ([currApp isActive]) {
//         return [currApp localizedName];
//         //NSLog(@"* %@", [currApp localizedName]);
//       } else {
//         //NSLog(@"  %@", [currApp localizedName]);
//       }
//   }
//   return NULL;
// }
import "C"


func GetActiveApplication() string {
  activeApplication := C.get_active_application()
  return NSStringToGoString(activeApplication)
}