package entity

// and entity is a struct with

// - Movementcomponent int
// - MovementParamater int or struct*
// - PhysicsUpdateFunction (the id of function that updates the movement)
// - PhysicsState struct (like position and so on )
// - PhysicsParameter int or struct* (size, bounding box, etc all parameters in this struct)
// - DrawUpdateFunction int
// - MetaData struct (misc like name, and other stuff go in here, creation date, etc)
// - Attributes Struct (health, oxygen, etc)

// and we have for each object

// Synth, [02.07.21 07:16]
// and just add one thing at a time

// Synth, [02.07.21 07:16]
// - start with draw state

// Synth, [02.07.21 07:16]
// and movement state

// Synth, [02.07.21 07:16]
// add attributes

// Synth, [02.07.21 07:16]
// and phytsics update and physics parameters

// - object type id (the type, particle, agent, furniture, etc)
// - object id (the id or index for that object in the list for that object type)
// - the draw component id (that will be passed the object type id and the object id); the parameters will be fetchable from the object id and object type id, and most draw components will only work for a specific object type
// - the physics update id (an integer) describing the type of physics or state update to do on the object

// **COMPONENT SYSTEM**
// Each object has:
// - Object type id (the type, particle, agent, furniture, etc)
// - Object id (the id or index for that object in the list for that object type)
// - Draw component id (that will be passed the object type id and the object id); the parameters will be fetchable from the object id and object type id, and most draw components will only work for a specific object type
// - Physics update id (an integer) describing the type of physics or state update to do on the object

// These components are integers. Indexes into a list.
// The integer can reference a struct (data) or it can reference a function (something to be called, like UpdatePhysics() for frame or DrawObject() for that object type).

// */

type Entity struct {
	// EntityType            int
	EntityType int
	EntityId   int
	// DrawComponentId int
	// PhysicsUpdateId int
	// MovementComponent     int
	// MovementParameter     int
	// PhysicsUpdateFunction int
	// PhysicsParameter      int
	// DrawUpdateFunction int
	// Metadata Metadata
	// Attributes Attribute
}

type EntityType int

//we should create array of entities for each agent, particle, furniture etc

type Metadata struct {
}

type Attribute struct {
}

// and we have for each object

// - object type id (the type, particle, agent, furniture, etc)
// - object id (the id or index for that object in the list for that object type)
// - the draw component id (that will be passed the object type id and the object id);
// the parameters will be fetchable from the object id and object type id, and most draw components will only work for a specific object type
// - the physics update id (an integer) describing the type of physics or state update to do on the object

// particles[]
// agents[]
// furniture[]

//particle
//agent
//furniture
//draw component id
