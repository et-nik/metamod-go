package main

/*
const char* ReadString(globalvars_t *gpGlobals, int offset);
*/
import "C"
import "fmt"

//export goTakeDamage
func goTakeDamage(inflictor *C.entvars_t, attacker *C.entvars_t, damage C.float, bitsDamageType C.int) C.int {
	fmt.Println("")
	fmt.Println("")
	fmt.Println("Go Take damage")

	fmt.Println("inflictor.p.classname: ", inflictor.classname)
	fmt.Println("attacker.p.classname: ", attacker.classname)

	inf := entVarsFromC(P.GlobalVars.p, inflictor)
	att := entVarsFromC(P.GlobalVars.p, attacker)

	fmt.Println("Inflictor class name: ", inf.ClassName())
	fmt.Println("Attacker net name: ", att.NetName())
	fmt.Println("Damage: ", damage)
	fmt.Println("Damage: ", float32(damage))

	fmt.Println("")

	//ent := P.EngineFuncs.CreateEntity()
	//if ent.EntVars().IsValid() {
	//	fmt.Println("New Entity is valid")
	//} else {
	//	fmt.Println("New Entity is not valid")
	//}
	//
	//P.EngineFuncs.RemoveEntity(ent)
	//
	//if ent.EntVars().IsValid() {
	//	fmt.Println("Removed Entity is valid")
	//} else {
	//	fmt.Println("Removed Entity is not valid")
	//}

	return 0
}

//export getFuncName
func getFuncName(addr C.uint32) {
	fmt.Println("funcName: ", P.EngineFuncs.NameForFunction(uint32(addr)))
}

//export goGetGlobalVars
func goGetGlobalVars() *C.globalvars_t {
	return P.GlobalVars.p
}
