package main

/*
void **getVTable(void *pthis, int size)
{
	return *((void***)(((char*)pthis)+size));
}

void hookFunction(void *target, void *replacement) {
    uint8_t *target_ptr = (uint8_t *)target;

    size_t jmp_size = 5;

    size_t page_size = sysconf(_SC_PAGESIZE);
    uintptr_t page_start = (uintptr_t)target_ptr & -page_size;
    if (mprotect((void *)page_start, page_size, PROT_READ | PROT_WRITE | PROT_EXEC) != 0) {
        perror("mprotect");
        return;
    }

    intptr_t offset = (intptr_t)replacement - (intptr_t)target - jmp_size;

    target_ptr[0] = 0xE9;
    *((int32_t *)(target_ptr + 1)) = offset;

    mprotect((void *)page_start, page_size, PROT_READ | PROT_EXEC);
}

extern const char* ReadString(globalvars_t *gpGlobals, int offset);
extern int goTakeDamage(entvars_t* inflictor, entvars_t* attacker, float damage, int bitsDamageType);
extern globalvars_t* goGetGlobalVars();

//int FNullEnt(entvars_t *pev) { return pev == NULL || FNullEnt(OFFSET(pev)); }

void *origFunc;
//(int(*)(entvars_t*, entvars_t*, float, int)) f;

int TakeDamage(entvars_t *pevInflictor, entvars_t *pevAttacker, float flDamage, int bitsDamageType) {
	printf("Took damage\n");
	printf("flDamage: %f\n", flDamage);

	printf("Inflictor classname: %d\n", pevInflictor->classname);
	printf("Inflictor globalname: %d\n", pevInflictor->globalname);
	printf("Inflictor model: %d\n", pevInflictor->model);
	printf("Inflictor netname: %d\n", pevInflictor->netname);
	printf("Inflictor message: %d\n", pevInflictor->message);
	printf("Inflictor weaponmodel: %d\n", pevInflictor->weaponmodel);
	printf(pevInflictor->pContainingEntity);
	printf("\n");

	if (pevInflictor == NULL) {
		printf("Inflictor is null\n");
	}

	if (pevInflictor->pContainingEntity == NULL) {
		printf("Containing entity is null\n");
	}

	if (pevAttacker->pContainingEntity == NULL) {
		printf("Attacker containing entity is null\n");
	}

	//char *globalname = ReadString(goGetGlobalVars(), (int)pevInflictor->globalname);
	//printf("Inflictor classname: %s\n", globalname);

	if (pevInflictor == NULL) {
		printf("Inflictor is null\n");
	}


	goTakeDamage(pevInflictor, pevAttacker, flDamage, bitsDamageType);

	//*f = malloc(sizeof(int(*)(entvars_t*, entvars_t*, float, int)));

	//f = (int(*)(entvars_t*, entvars_t*, float, int))origFunc;

	((int(*)(entvars_t*, entvars_t*, float, int))origFunc)(pevInflictor, pevAttacker, flDamage, bitsDamageType);
}

void hookEntity(edict_t *ed) {
	printf("called hookEntity\n");

	void** vtable = getVTable(ed->pvPrivateData, 0);

	int **ivtable=(int **)vtable;
	void *func=ivtable[11];

	getFuncName((uint32_t)func);

	origFunc = func;

	hookFunction(func, (void*)TakeDamage);
}
*/
import "C"

// takedamage 13
//        takehealth 14
//        killed 15

// traceattack 10
//	takedamage 11
//	takehealth 12

func hookEntity(ed *Edict) {
	C.hookEntity(ed.p)
}
