// Code generated by ent, DO NOT EDIT.

package entgen

import (
	"fmt"
	"inventory/ent/entgen/tblcart"
	"inventory/ent/entgen/tblinventory"
	"inventory/ent/entgen/tbluser"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// TblCart is the model entity for the TblCart schema.
type TblCart struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// ProductId holds the value of the "ProductId" field.
	ProductId string `json:"ProductId,omitempty"`
	// UserId holds the value of the "UserId" field.
	UserId string `json:"UserId,omitempty"`
	// Status holds the value of the "Status" field.
	Status tblcart.Status `json:"Status,omitempty"`
	// CreatedAt holds the value of the "Created_at" field.
	CreatedAt time.Time `json:"Created_at,omitempty"`
	// UpdatedAt holds the value of the "Updated_at" field.
	UpdatedAt time.Time `json:"Updated_at,omitempty"`
	// DeletedAt holds the value of the "Deleted_at" field.
	DeletedAt *time.Time `json:"Deleted_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the TblCartQuery when eager-loading is set.
	Edges        TblCartEdges `json:"edges"`
	selectValues sql.SelectValues
}

// TblCartEdges holds the relations/edges for other nodes in the graph.
type TblCartEdges struct {
	// Inventory holds the value of the Inventory edge.
	Inventory *TblInventory `json:"Inventory,omitempty"`
	// User holds the value of the User edge.
	User *TblUser `json:"User,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// InventoryOrErr returns the Inventory value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e TblCartEdges) InventoryOrErr() (*TblInventory, error) {
	if e.Inventory != nil {
		return e.Inventory, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: tblinventory.Label}
	}
	return nil, &NotLoadedError{edge: "Inventory"}
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e TblCartEdges) UserOrErr() (*TblUser, error) {
	if e.User != nil {
		return e.User, nil
	} else if e.loadedTypes[1] {
		return nil, &NotFoundError{label: tbluser.Label}
	}
	return nil, &NotLoadedError{edge: "User"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*TblCart) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case tblcart.FieldID, tblcart.FieldProductId, tblcart.FieldUserId, tblcart.FieldStatus:
			values[i] = new(sql.NullString)
		case tblcart.FieldCreatedAt, tblcart.FieldUpdatedAt, tblcart.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the TblCart fields.
func (tc *TblCart) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case tblcart.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				tc.ID = value.String
			}
		case tblcart.FieldProductId:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field ProductId", values[i])
			} else if value.Valid {
				tc.ProductId = value.String
			}
		case tblcart.FieldUserId:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field UserId", values[i])
			} else if value.Valid {
				tc.UserId = value.String
			}
		case tblcart.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field Status", values[i])
			} else if value.Valid {
				tc.Status = tblcart.Status(value.String)
			}
		case tblcart.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field Created_at", values[i])
			} else if value.Valid {
				tc.CreatedAt = value.Time
			}
		case tblcart.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field Updated_at", values[i])
			} else if value.Valid {
				tc.UpdatedAt = value.Time
			}
		case tblcart.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field Deleted_at", values[i])
			} else if value.Valid {
				tc.DeletedAt = new(time.Time)
				*tc.DeletedAt = value.Time
			}
		default:
			tc.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the TblCart.
// This includes values selected through modifiers, order, etc.
func (tc *TblCart) Value(name string) (ent.Value, error) {
	return tc.selectValues.Get(name)
}

// QueryInventory queries the "Inventory" edge of the TblCart entity.
func (tc *TblCart) QueryInventory() *TblInventoryQuery {
	return NewTblCartClient(tc.config).QueryInventory(tc)
}

// QueryUser queries the "User" edge of the TblCart entity.
func (tc *TblCart) QueryUser() *TblUserQuery {
	return NewTblCartClient(tc.config).QueryUser(tc)
}

// Update returns a builder for updating this TblCart.
// Note that you need to call TblCart.Unwrap() before calling this method if this TblCart
// was returned from a transaction, and the transaction was committed or rolled back.
func (tc *TblCart) Update() *TblCartUpdateOne {
	return NewTblCartClient(tc.config).UpdateOne(tc)
}

// Unwrap unwraps the TblCart entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (tc *TblCart) Unwrap() *TblCart {
	_tx, ok := tc.config.driver.(*txDriver)
	if !ok {
		panic("entgen: TblCart is not a transactional entity")
	}
	tc.config.driver = _tx.drv
	return tc
}

// String implements the fmt.Stringer.
func (tc *TblCart) String() string {
	var builder strings.Builder
	builder.WriteString("TblCart(")
	builder.WriteString(fmt.Sprintf("id=%v, ", tc.ID))
	builder.WriteString("ProductId=")
	builder.WriteString(tc.ProductId)
	builder.WriteString(", ")
	builder.WriteString("UserId=")
	builder.WriteString(tc.UserId)
	builder.WriteString(", ")
	builder.WriteString("Status=")
	builder.WriteString(fmt.Sprintf("%v", tc.Status))
	builder.WriteString(", ")
	builder.WriteString("Created_at=")
	builder.WriteString(tc.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("Updated_at=")
	builder.WriteString(tc.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	if v := tc.DeletedAt; v != nil {
		builder.WriteString("Deleted_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteByte(')')
	return builder.String()
}

// TblCarts is a parsable slice of TblCart.
type TblCarts []*TblCart
