// Code generated by entc, DO NOT EDIT.

package subjectdigest

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/testifysec/archivist/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.SubjectDigest {
	return predicate.SubjectDigest(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.SubjectDigest {
	return predicate.SubjectDigest(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.SubjectDigest {
	return predicate.SubjectDigest(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.SubjectDigest {
	return predicate.SubjectDigest(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.SubjectDigest {
	return predicate.SubjectDigest(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.SubjectDigest {
	return predicate.SubjectDigest(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.SubjectDigest {
	return predicate.SubjectDigest(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.SubjectDigest {
	return predicate.SubjectDigest(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.SubjectDigest {
	return predicate.SubjectDigest(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Algorithm applies equality check predicate on the "algorithm" field. It's identical to AlgorithmEQ.
func Algorithm(v string) predicate.SubjectDigest {
	return predicate.SubjectDigest(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAlgorithm), v))
	})
}

// Value applies equality check predicate on the "value" field. It's identical to ValueEQ.
func Value(v string) predicate.SubjectDigest {
	return predicate.SubjectDigest(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldValue), v))
	})
}

// AlgorithmEQ applies the EQ predicate on the "algorithm" field.
func AlgorithmEQ(v string) predicate.SubjectDigest {
	return predicate.SubjectDigest(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAlgorithm), v))
	})
}

// AlgorithmNEQ applies the NEQ predicate on the "algorithm" field.
func AlgorithmNEQ(v string) predicate.SubjectDigest {
	return predicate.SubjectDigest(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAlgorithm), v))
	})
}

// AlgorithmIn applies the In predicate on the "algorithm" field.
func AlgorithmIn(vs ...string) predicate.SubjectDigest {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SubjectDigest(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldAlgorithm), v...))
	})
}

// AlgorithmNotIn applies the NotIn predicate on the "algorithm" field.
func AlgorithmNotIn(vs ...string) predicate.SubjectDigest {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SubjectDigest(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldAlgorithm), v...))
	})
}

// AlgorithmGT applies the GT predicate on the "algorithm" field.
func AlgorithmGT(v string) predicate.SubjectDigest {
	return predicate.SubjectDigest(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAlgorithm), v))
	})
}

// AlgorithmGTE applies the GTE predicate on the "algorithm" field.
func AlgorithmGTE(v string) predicate.SubjectDigest {
	return predicate.SubjectDigest(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAlgorithm), v))
	})
}

// AlgorithmLT applies the LT predicate on the "algorithm" field.
func AlgorithmLT(v string) predicate.SubjectDigest {
	return predicate.SubjectDigest(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAlgorithm), v))
	})
}

// AlgorithmLTE applies the LTE predicate on the "algorithm" field.
func AlgorithmLTE(v string) predicate.SubjectDigest {
	return predicate.SubjectDigest(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAlgorithm), v))
	})
}

// AlgorithmContains applies the Contains predicate on the "algorithm" field.
func AlgorithmContains(v string) predicate.SubjectDigest {
	return predicate.SubjectDigest(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldAlgorithm), v))
	})
}

// AlgorithmHasPrefix applies the HasPrefix predicate on the "algorithm" field.
func AlgorithmHasPrefix(v string) predicate.SubjectDigest {
	return predicate.SubjectDigest(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldAlgorithm), v))
	})
}

// AlgorithmHasSuffix applies the HasSuffix predicate on the "algorithm" field.
func AlgorithmHasSuffix(v string) predicate.SubjectDigest {
	return predicate.SubjectDigest(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldAlgorithm), v))
	})
}

// AlgorithmEqualFold applies the EqualFold predicate on the "algorithm" field.
func AlgorithmEqualFold(v string) predicate.SubjectDigest {
	return predicate.SubjectDigest(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldAlgorithm), v))
	})
}

// AlgorithmContainsFold applies the ContainsFold predicate on the "algorithm" field.
func AlgorithmContainsFold(v string) predicate.SubjectDigest {
	return predicate.SubjectDigest(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldAlgorithm), v))
	})
}

// ValueEQ applies the EQ predicate on the "value" field.
func ValueEQ(v string) predicate.SubjectDigest {
	return predicate.SubjectDigest(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldValue), v))
	})
}

// ValueNEQ applies the NEQ predicate on the "value" field.
func ValueNEQ(v string) predicate.SubjectDigest {
	return predicate.SubjectDigest(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldValue), v))
	})
}

// ValueIn applies the In predicate on the "value" field.
func ValueIn(vs ...string) predicate.SubjectDigest {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SubjectDigest(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldValue), v...))
	})
}

// ValueNotIn applies the NotIn predicate on the "value" field.
func ValueNotIn(vs ...string) predicate.SubjectDigest {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SubjectDigest(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldValue), v...))
	})
}

// ValueGT applies the GT predicate on the "value" field.
func ValueGT(v string) predicate.SubjectDigest {
	return predicate.SubjectDigest(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldValue), v))
	})
}

// ValueGTE applies the GTE predicate on the "value" field.
func ValueGTE(v string) predicate.SubjectDigest {
	return predicate.SubjectDigest(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldValue), v))
	})
}

// ValueLT applies the LT predicate on the "value" field.
func ValueLT(v string) predicate.SubjectDigest {
	return predicate.SubjectDigest(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldValue), v))
	})
}

// ValueLTE applies the LTE predicate on the "value" field.
func ValueLTE(v string) predicate.SubjectDigest {
	return predicate.SubjectDigest(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldValue), v))
	})
}

// ValueContains applies the Contains predicate on the "value" field.
func ValueContains(v string) predicate.SubjectDigest {
	return predicate.SubjectDigest(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldValue), v))
	})
}

// ValueHasPrefix applies the HasPrefix predicate on the "value" field.
func ValueHasPrefix(v string) predicate.SubjectDigest {
	return predicate.SubjectDigest(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldValue), v))
	})
}

// ValueHasSuffix applies the HasSuffix predicate on the "value" field.
func ValueHasSuffix(v string) predicate.SubjectDigest {
	return predicate.SubjectDigest(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldValue), v))
	})
}

// ValueEqualFold applies the EqualFold predicate on the "value" field.
func ValueEqualFold(v string) predicate.SubjectDigest {
	return predicate.SubjectDigest(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldValue), v))
	})
}

// ValueContainsFold applies the ContainsFold predicate on the "value" field.
func ValueContainsFold(v string) predicate.SubjectDigest {
	return predicate.SubjectDigest(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldValue), v))
	})
}

// HasSubject applies the HasEdge predicate on the "subject" edge.
func HasSubject() predicate.SubjectDigest {
	return predicate.SubjectDigest(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(SubjectTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, SubjectTable, SubjectColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasSubjectWith applies the HasEdge predicate on the "subject" edge with a given conditions (other predicates).
func HasSubjectWith(preds ...predicate.Subject) predicate.SubjectDigest {
	return predicate.SubjectDigest(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(SubjectInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, SubjectTable, SubjectColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.SubjectDigest) predicate.SubjectDigest {
	return predicate.SubjectDigest(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.SubjectDigest) predicate.SubjectDigest {
	return predicate.SubjectDigest(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.SubjectDigest) predicate.SubjectDigest {
	return predicate.SubjectDigest(func(s *sql.Selector) {
		p(s.Not())
	})
}