#if !defined EDICT_H
#define EDICT_H
#ifdef _WIN32
#ifndef __MINGW32__
#pragma once
#endif /* not __MINGW32__ */
#endif
#define	MAX_ENT_LEAFS	48

#include "progdefs.h"

typedef struct edict_s
{
	qboolean	free;
	int			serialnumber;
	link_t		area;				// linked to a division node or leaf
	
	int			headnode;			// -1 to use normal leaf check
	int			num_leafs;
	short		leafnums[MAX_ENT_LEAFS];

	float		freetime;			// sv.time when the object was freed

	void*		pvPrivateData;		// Alloced and freed by engine, used by DLLs

	entvars_t	v;					// C exported fields from progs

	// other fields from progs come immediately after
};

#endif
