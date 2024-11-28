package engine

const (
	EdictFlagFly  = 1 << iota // Changes the SV_Movestep() behavior to not need to be on ground
	EdictFlagSwim             // Changes the SV_Movestep() behavior to not need to be on ground (but stay in water)
	EdictFlagConveyor
	EdictFlagClient
	EdictFlagInWater
	EdictFlagMonster
	EdictFlagGodMode
	EdictFlagNoTarget
	EdictFlagSkipLocalHost // Don't send entity to local host, it's predicting this entity itself
	EdictFlagOnGround      // At rest / on the ground
	EdictFlagPartialGround // not all corners are valid
	EdictFlagWaterJump     // player jumping out of water
	EdictFlagFrozen        // Player is frozen for 3rd person camera
	EdictFlagFakeClient    // JAC: fake client, simulated server side; don't send network messages to them
	EdictFlagDucking       // Player flag -- Player is fully crouched
	EdictFlagFloat         // Apply floating force to this entity when in water
	EdictFlagGraphed       // worldgraph has this ent listed as something that blocks a connection

	EdictFlagImmuneWater
	EdictFlagImmuneSlime
	EdictFlagImmuneLava

	EdictFlagProxy
	EdictFlagAlwaysThink
	EdictFlagBaseVelocity
	EdictFlagMonsterClip
	EdictFlagOnTrain
	EdictFlagWorldBrush
	EdictFlagSpectator
	EdictFlagCustomEntity // This is a custom entity
	EdictFlagKillMe       // This entity is marked for death -- This allows the engine to kill ents at the appropriate time
	EdictFlagDormant      // Entity is dormant, no updates to client
)
