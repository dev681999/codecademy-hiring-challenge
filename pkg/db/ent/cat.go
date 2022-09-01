// Code generated by ent, DO NOT EDIT.

package ent

import (
	"catinator-backend/pkg/db/ent/cat"
	"catinator-backend/pkg/db/ent/user"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// Cat is the model entity for the Cat schema.
type Cat struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime time.Time `json:"update_time,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// ImageID holds the value of the "image_id" field.
	ImageID string `json:"image_id,omitempty"`
	// OwnerID holds the value of the "owner_id" field.
	OwnerID string `json:"owner_id,omitempty"`
	// Tags holds the value of the "tags" field.
	Tags []string `json:"tags,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the CatQuery when eager-loading is set.
	Edges CatEdges `json:"edges"`
}

// CatEdges holds the relations/edges for other nodes in the graph.
type CatEdges struct {
	// Owner holds the value of the owner edge.
	Owner *User `json:"owner,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// OwnerOrErr returns the Owner value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e CatEdges) OwnerOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.Owner == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.Owner, nil
	}
	return nil, &NotLoadedError{edge: "owner"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Cat) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case cat.FieldTags:
			values[i] = new([]byte)
		case cat.FieldID, cat.FieldName, cat.FieldDescription, cat.FieldImageID, cat.FieldOwnerID:
			values[i] = new(sql.NullString)
		case cat.FieldCreateTime, cat.FieldUpdateTime:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Cat", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Cat fields.
func (c *Cat) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case cat.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				c.ID = value.String
			}
		case cat.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				c.CreateTime = value.Time
			}
		case cat.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				c.UpdateTime = value.Time
			}
		case cat.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				c.Name = value.String
			}
		case cat.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				c.Description = value.String
			}
		case cat.FieldImageID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field image_id", values[i])
			} else if value.Valid {
				c.ImageID = value.String
			}
		case cat.FieldOwnerID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field owner_id", values[i])
			} else if value.Valid {
				c.OwnerID = value.String
			}
		case cat.FieldTags:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field tags", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &c.Tags); err != nil {
					return fmt.Errorf("unmarshal field tags: %w", err)
				}
			}
		}
	}
	return nil
}

// QueryOwner queries the "owner" edge of the Cat entity.
func (c *Cat) QueryOwner() *UserQuery {
	return (&CatClient{config: c.config}).QueryOwner(c)
}

// Update returns a builder for updating this Cat.
// Note that you need to call Cat.Unwrap() before calling this method if this Cat
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Cat) Update() *CatUpdateOne {
	return (&CatClient{config: c.config}).UpdateOne(c)
}

// Unwrap unwraps the Cat entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Cat) Unwrap() *Cat {
	_tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Cat is not a transactional entity")
	}
	c.config.driver = _tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Cat) String() string {
	var builder strings.Builder
	builder.WriteString("Cat(")
	builder.WriteString(fmt.Sprintf("id=%v, ", c.ID))
	builder.WriteString("create_time=")
	builder.WriteString(c.CreateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("update_time=")
	builder.WriteString(c.UpdateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(c.Name)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(c.Description)
	builder.WriteString(", ")
	builder.WriteString("image_id=")
	builder.WriteString(c.ImageID)
	builder.WriteString(", ")
	builder.WriteString("owner_id=")
	builder.WriteString(c.OwnerID)
	builder.WriteString(", ")
	builder.WriteString("tags=")
	builder.WriteString(fmt.Sprintf("%v", c.Tags))
	builder.WriteByte(')')
	return builder.String()
}

// Cats is a parsable slice of Cat.
type Cats []*Cat

func (c Cats) config(cfg config) {
	for _i := range c {
		c[_i].config = cfg
	}
}