package main

/**
typedef struct plugin_info_t
{
    const char*     ifvers;     // meta_interface version
    const char*     name;       // full name of plugin
    const char*     version;    // version
    const char*     date;       // date
    const char*     author;     // author name/email
    const char*     url;        // URL
    const char*     logtag;     // log message prefix (unused right now)
    int   loadable;   // when loadable
    int   unloadable; // when unloadable
} plugin_info_t;
*/
import "C"
import "fmt"

func main() {}

//export call
func call() {
	fmt.Println("Hi from Go!")
}

//export call2
func call2(s *C.char) {
	v := C.GoString(s)
	fmt.Println("(call2) Hi from Go! You passed:", v)
}

//export call3
func call3(s C.int) C.int {
	fmt.Println("(call3) Hi from Go! You passed:", s)
	return s
}

//export Meta_Attach
func Meta_Attach(now C.int, pFunctionTable *C.void, pMGlobals *C.void, pGamedllFuncs *C.void) C.int {
	fmt.Println("=====================================")
	fmt.Println("(Meta_Attach) Hi from Go!")
	fmt.Println("=====================================")

	return 1
}

//export Meta_Query
func Meta_Query(interfaceVersion *C.char, plinfo **C.plugin_info_t, pMetaUtilFuncs *C.void) C.int {
	*plinfo = &C.plugin_info_t{
		ifvers:     C.CString(MetaInterfaceVersion),
		name:       C.CString("GoPlugin"),
		version:    C.CString("1.0"),
		date:       C.CString("2024-11-08"),
		author:     C.CString("Author"),
		url:        C.CString(""),
		logtag:     C.CString("EXAMPLE"),
		loadable:   3,
		unloadable: 3,
	}

	return 1
}

//export Meta_Detach
func Meta_Detach(now C.int, reason C.int) C.int {
	fmt.Println("=====================================")
	fmt.Println("(Meta_Detach) Hi from Go!")
	fmt.Println("=====================================")

	return 1
}

//export Meta_Init
func Meta_Init() {
	fmt.Println("=====================================")
	fmt.Println("(Meta_Init) Hi from Go!")
	fmt.Println("=====================================")
}

//export GiveFnptrsToDll
func GiveFnptrsToDll(pengfuncsFromEngine *C.void, pGlobals *C.void) {
	fmt.Println("=====================================")
	fmt.Println("(GiveFnptrsToDll) Hi from Go!")
	fmt.Println("=====================================")
}
