package metamod_go

import (
	"github.com/et-nik/metamod-go/engine"
	"iter"
)

// This file contains extended engine functions. More convenient to use in Go.
// For example, you can iterate over entities using a for loop.

func (ef *EngineFuncs) EntitiesByFieldValue(
	field engine.FindEntityField,
	value string,
) iter.Seq[*Edict] {
	return func(yield func(*Edict) bool) {
		var entity *Edict

		for {
			entity = ef.FindEntityByString(entity, field, value)

			if entity == nil {
				break
			}

			if IsNullEntity(entity) {
				break
			}

			if !yield(entity) {
				break
			}
		}
	}
}

func (ef *EngineFuncs) EntitiesByClassname(classname string) iter.Seq[*Edict] {
	return ef.EntitiesByFieldValue(engine.FindEntityFieldClassname, classname)
}

func (ef *EngineFuncs) EntitiesByModel(model string) iter.Seq[*Edict] {
	return ef.EntitiesByFieldValue(engine.FindEntityFieldModel, model)
}

func (ef *EngineFuncs) EntitiesByViewModel(model string) iter.Seq[*Edict] {
	return ef.EntitiesByFieldValue(engine.FindEntityFieldViewModel, model)
}
