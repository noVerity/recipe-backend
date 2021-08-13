// Code generated by entc, DO NOT EDIT.

package ingredient

import (
	"adomeit.xyz/recipe/ent/predicate"
	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Ingredient {
	return predicate.Ingredient(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Ingredient {
	return predicate.Ingredient(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Ingredient {
	return predicate.Ingredient(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Ingredient {
	return predicate.Ingredient(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.Ingredient {
	return predicate.Ingredient(func(s *sql.Selector) {
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
func IDGT(id int) predicate.Ingredient {
	return predicate.Ingredient(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Ingredient {
	return predicate.Ingredient(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Ingredient {
	return predicate.Ingredient(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Ingredient {
	return predicate.Ingredient(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Ingredient {
	return predicate.Ingredient(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// Calories applies equality check predicate on the "calories" field. It's identical to CaloriesEQ.
func Calories(v float32) predicate.Ingredient {
	return predicate.Ingredient(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCalories), v))
	})
}

// Fat applies equality check predicate on the "fat" field. It's identical to FatEQ.
func Fat(v float32) predicate.Ingredient {
	return predicate.Ingredient(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFat), v))
	})
}

// Carbohydrates applies equality check predicate on the "carbohydrates" field. It's identical to CarbohydratesEQ.
func Carbohydrates(v float32) predicate.Ingredient {
	return predicate.Ingredient(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCarbohydrates), v))
	})
}

// Protein applies equality check predicate on the "protein" field. It's identical to ProteinEQ.
func Protein(v float32) predicate.Ingredient {
	return predicate.Ingredient(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldProtein), v))
	})
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Ingredient {
	return predicate.Ingredient(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Ingredient {
	return predicate.Ingredient(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldName), v))
	})
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Ingredient {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Ingredient(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldName), v...))
	})
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Ingredient {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Ingredient(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldName), v...))
	})
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Ingredient {
	return predicate.Ingredient(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldName), v))
	})
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Ingredient {
	return predicate.Ingredient(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldName), v))
	})
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Ingredient {
	return predicate.Ingredient(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldName), v))
	})
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Ingredient {
	return predicate.Ingredient(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldName), v))
	})
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Ingredient {
	return predicate.Ingredient(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldName), v))
	})
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Ingredient {
	return predicate.Ingredient(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldName), v))
	})
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Ingredient {
	return predicate.Ingredient(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldName), v))
	})
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Ingredient {
	return predicate.Ingredient(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldName), v))
	})
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Ingredient {
	return predicate.Ingredient(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldName), v))
	})
}

// CaloriesEQ applies the EQ predicate on the "calories" field.
func CaloriesEQ(v float32) predicate.Ingredient {
	return predicate.Ingredient(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCalories), v))
	})
}

// CaloriesNEQ applies the NEQ predicate on the "calories" field.
func CaloriesNEQ(v float32) predicate.Ingredient {
	return predicate.Ingredient(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCalories), v))
	})
}

// CaloriesIn applies the In predicate on the "calories" field.
func CaloriesIn(vs ...float32) predicate.Ingredient {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Ingredient(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCalories), v...))
	})
}

// CaloriesNotIn applies the NotIn predicate on the "calories" field.
func CaloriesNotIn(vs ...float32) predicate.Ingredient {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Ingredient(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCalories), v...))
	})
}

// CaloriesGT applies the GT predicate on the "calories" field.
func CaloriesGT(v float32) predicate.Ingredient {
	return predicate.Ingredient(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCalories), v))
	})
}

// CaloriesGTE applies the GTE predicate on the "calories" field.
func CaloriesGTE(v float32) predicate.Ingredient {
	return predicate.Ingredient(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCalories), v))
	})
}

// CaloriesLT applies the LT predicate on the "calories" field.
func CaloriesLT(v float32) predicate.Ingredient {
	return predicate.Ingredient(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCalories), v))
	})
}

// CaloriesLTE applies the LTE predicate on the "calories" field.
func CaloriesLTE(v float32) predicate.Ingredient {
	return predicate.Ingredient(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCalories), v))
	})
}

// FatEQ applies the EQ predicate on the "fat" field.
func FatEQ(v float32) predicate.Ingredient {
	return predicate.Ingredient(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFat), v))
	})
}

// FatNEQ applies the NEQ predicate on the "fat" field.
func FatNEQ(v float32) predicate.Ingredient {
	return predicate.Ingredient(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldFat), v))
	})
}

// FatIn applies the In predicate on the "fat" field.
func FatIn(vs ...float32) predicate.Ingredient {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Ingredient(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldFat), v...))
	})
}

// FatNotIn applies the NotIn predicate on the "fat" field.
func FatNotIn(vs ...float32) predicate.Ingredient {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Ingredient(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldFat), v...))
	})
}

// FatGT applies the GT predicate on the "fat" field.
func FatGT(v float32) predicate.Ingredient {
	return predicate.Ingredient(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldFat), v))
	})
}

// FatGTE applies the GTE predicate on the "fat" field.
func FatGTE(v float32) predicate.Ingredient {
	return predicate.Ingredient(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldFat), v))
	})
}

// FatLT applies the LT predicate on the "fat" field.
func FatLT(v float32) predicate.Ingredient {
	return predicate.Ingredient(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldFat), v))
	})
}

// FatLTE applies the LTE predicate on the "fat" field.
func FatLTE(v float32) predicate.Ingredient {
	return predicate.Ingredient(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldFat), v))
	})
}

// CarbohydratesEQ applies the EQ predicate on the "carbohydrates" field.
func CarbohydratesEQ(v float32) predicate.Ingredient {
	return predicate.Ingredient(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCarbohydrates), v))
	})
}

// CarbohydratesNEQ applies the NEQ predicate on the "carbohydrates" field.
func CarbohydratesNEQ(v float32) predicate.Ingredient {
	return predicate.Ingredient(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCarbohydrates), v))
	})
}

// CarbohydratesIn applies the In predicate on the "carbohydrates" field.
func CarbohydratesIn(vs ...float32) predicate.Ingredient {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Ingredient(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCarbohydrates), v...))
	})
}

// CarbohydratesNotIn applies the NotIn predicate on the "carbohydrates" field.
func CarbohydratesNotIn(vs ...float32) predicate.Ingredient {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Ingredient(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCarbohydrates), v...))
	})
}

// CarbohydratesGT applies the GT predicate on the "carbohydrates" field.
func CarbohydratesGT(v float32) predicate.Ingredient {
	return predicate.Ingredient(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCarbohydrates), v))
	})
}

// CarbohydratesGTE applies the GTE predicate on the "carbohydrates" field.
func CarbohydratesGTE(v float32) predicate.Ingredient {
	return predicate.Ingredient(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCarbohydrates), v))
	})
}

// CarbohydratesLT applies the LT predicate on the "carbohydrates" field.
func CarbohydratesLT(v float32) predicate.Ingredient {
	return predicate.Ingredient(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCarbohydrates), v))
	})
}

// CarbohydratesLTE applies the LTE predicate on the "carbohydrates" field.
func CarbohydratesLTE(v float32) predicate.Ingredient {
	return predicate.Ingredient(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCarbohydrates), v))
	})
}

// ProteinEQ applies the EQ predicate on the "protein" field.
func ProteinEQ(v float32) predicate.Ingredient {
	return predicate.Ingredient(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldProtein), v))
	})
}

// ProteinNEQ applies the NEQ predicate on the "protein" field.
func ProteinNEQ(v float32) predicate.Ingredient {
	return predicate.Ingredient(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldProtein), v))
	})
}

// ProteinIn applies the In predicate on the "protein" field.
func ProteinIn(vs ...float32) predicate.Ingredient {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Ingredient(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldProtein), v...))
	})
}

// ProteinNotIn applies the NotIn predicate on the "protein" field.
func ProteinNotIn(vs ...float32) predicate.Ingredient {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Ingredient(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldProtein), v...))
	})
}

// ProteinGT applies the GT predicate on the "protein" field.
func ProteinGT(v float32) predicate.Ingredient {
	return predicate.Ingredient(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldProtein), v))
	})
}

// ProteinGTE applies the GTE predicate on the "protein" field.
func ProteinGTE(v float32) predicate.Ingredient {
	return predicate.Ingredient(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldProtein), v))
	})
}

// ProteinLT applies the LT predicate on the "protein" field.
func ProteinLT(v float32) predicate.Ingredient {
	return predicate.Ingredient(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldProtein), v))
	})
}

// ProteinLTE applies the LTE predicate on the "protein" field.
func ProteinLTE(v float32) predicate.Ingredient {
	return predicate.Ingredient(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldProtein), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Ingredient) predicate.Ingredient {
	return predicate.Ingredient(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Ingredient) predicate.Ingredient {
	return predicate.Ingredient(func(s *sql.Selector) {
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
func Not(p predicate.Ingredient) predicate.Ingredient {
	return predicate.Ingredient(func(s *sql.Selector) {
		p(s.Not())
	})
}