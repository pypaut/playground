_Source: https://refactoring.guru/design-patterns/memento_

# Context

New feature for a text editing application: undo.

We want to save the state of the current edition at each action, to be able to
retrieve them with "undo".

# Problem

Most objects will use private fields, thus can't be copied.

Even if we decide to make them public, refactoring will be a pain: changing any
fields will also require changing the function taking care of saving. There is
also a retro-compatibility issue.

# Solution

All problems that we’ve just experienced are caused by broken encapsulation.
Some objects try to do more than they are supposed to. To collect the data
required to perform some action, they invade the private space of other objects
instead of letting these objects perform the actual action.

-> Broken encapsulation: each object should only interact with their own private
   space.

The Memento pattern delegates creating the state snapshots to the actual owner
of that state, the originator object. Hence, instead of other objects trying to
copy the editor’s state from the “outside,” the editor class itself can make the
snapshot since it has full access to its own state.

-> Memento pattern: each object creates its own snapshot.

The pattern suggests storing the copy of the object’s state in a special
object called memento. The contents of the memento aren’t accessible to any
other object except the one that produced it. Other objects must communicate
with mementos using a limited interface which may allow fetching the
snapshot’s metadata (creation time, the name of the performed operation,
etc.), but not the original object’s state contained in the snapshot.

-> Store an object inside the object, to save the owner object's state. Only
   the owner has full access to the memento. Others see a limited interface.

# Usage

```go
// The user does any actions, saves current state
snapshot = editor.MakeSnapshot()
history.AddSnapshot(snapshot)

// The user uses "undo"
lastSnapshot = history.GetLastSnapshot()
editor.Restore(lastSnapshot)
```

# Structure

## Nested classes

## Intermediate interface

## Stricter encapsulation
