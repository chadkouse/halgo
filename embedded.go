package halgo

import (
)

// Links represents a collection of HAL links. You can embed this struct
// in your own structs for sweet, sweet HAL serialisation goodness.
//
//     type MyStruct struct {
//       halgo.Links
//     }
//
//     my := MyStruct{
//       Links: halgo.Links{}.
//         Self("http://example.com/").
//         Next("http://example.com/1"),
//     }
type Embedded struct {
	Items map[string]embeddedSet `json:"_embedded,omitempty"`
	// Curies CurieSet
}

// Link creates a link with a named rel. Optionally can act as a format
// string with parameters.
//
//     Link("abc", "http://example.com/a/1")
//     Link("abc", "http://example.com/a/%d", id)
func (l Embedded) Embed(rel string, item interface{}) Embedded {

	return l.Add(rel, item)
}

// Add creates multiple links with the same relation.
//
//     Add("abc", halgo.Link{Href: "/a/1"}, halgo.Link{Href: "/a/2"})
func (l Embedded) Add(rel string, items ...interface{}) Embedded {
	if l.Items == nil {
		l.Items = make(map[string]embeddedSet)
	}

	set, exists := l.Items[rel]

	if exists {
		set = append(set, items...)
	} else {
		set = make([]interface{}, len(items))
		copy(set, items)
	}

	l.Items[rel] = set

	return l
}

