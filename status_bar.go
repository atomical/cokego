// +build darwin

package cokego

// #cgo CFLAGS: -x objective-c -l/System/Library/Frameworks
// #cgo LDFLAGS: -framework Cocoa -framework Foundation
// #include <Cocoa/Cocoa.h>
// void * CreateNSStatusBar(){
//   return [NSStatusBar systemStatusBar];
// }
//   void CreateMenu(){
//     [NSAutoreleasePool new];
//     [NSApplication sharedApplication];
//     [NSApp setActivationPolicy:NSApplicationActivationPolicyRegular];
//     id menubar = [[NSMenu new] autorelease];
//     id appMenuItem = [[NSMenuItem new] autorelease];
//     [menubar addItem:appMenuItem];
//     [NSApp setMainMenu:menubar];
//     NSMenu *stackMenu = [[NSMenu alloc] initWithTitle:@"Status Menu"];
//     NSMenuItem *soMenuItem = 
//         [[NSMenuItem alloc] initWithTitle:@"Status Menu Item" action:nil keyEquivalent:@"S"];
//     [soMenuItem setEnabled:YES];
//     [stackMenu addItem:soMenuItem];
//     NSStatusItem  * statusItem = [[[NSStatusBar systemStatusBar]
//                    statusItemWithLength:NSVariableStatusItemLength]
//                   retain];
//     [statusItem setImage:[NSImage imageNamed:@"icon.png"]];
//     [statusItem setTitle:[stackMenu title]];
//     [statusItem setMenu:stackMenu];
// }
//   void * create_menu(){
//     [NSAutoreleasePool new];
//     [NSApplication sharedApplication];
//     [NSApp setActivationPolicy:NSApplicationActivationPolicyRegular];
//     id menubar = [[NSMenu new] autorelease];
//     id appMenuItem = [[NSMenuItem new] autorelease];
//     [menubar addItem:appMenuItem];
//     [NSApp setMainMenu:menubar];
//
//     NSMenu *stackMenu = [[NSMenu alloc] initWithTitle:@"Status Menu"];
//     NSMenuItem *soMenuItem = 
//         [[NSMenuItem alloc] initWithTitle:@"Status Menu Item" action:nil keyEquivalent:@"S"];
//     [soMenuItem setEnabled:YES];
//     [stackMenu addItem:soMenuItem];
//     NSStatusItem  * statusItem = [[[NSStatusBar systemStatusBar]
//                    statusItemWithLength:NSVariableStatusItemLength]
//                   retain];
//     [statusItem setImage:[NSImage imageNamed:@"icon.png"]];
//     //[statusItem setTitle:[stackMenu title]];
//     [statusItem setMenu:stackMenu];
//     return statusItem;
// }
//
// void set_image( NSStatusItem  * statusItem, NSString * path ) {
//   [statusItem setImage:[NSImage imageNamed:path]];
// }
import "C"
import "unsafe"


func NSStatusBar() ( unsafe.Pointer ){
  return C.CreateNSStatusBar()
}

func StatusBarCreateMenu() unsafe.Pointer {
  return C.create_menu()
}

func NSSStatusBarCreateItem( bar unsafe.Pointer ) {
}
