package ordered_map

import "iter"

type OrderedMap[K comparable, V any] struct {
 keys  []K
 items map[K]V
}

func New[K comparable, V any](capacity int) *OrderedMap[K, V] {
 return &OrderedMap[K, V]{
  keys:  make([]K, 0, capacity),
  items: make(map[K]V, capacity),
 }
}

// Add new items by key, value
func (om *OrderedMap[K, V]) Add(k K, v V) {
 om.keys = append(om.keys, k)
 om.items[k] = v
}

// Get present item by key
func (om *OrderedMap[K, V]) Get(k K) V {
 return om.items[k]
}

// Retrieve all keys for iteration.
func (om *OrderedMap[K, V]) Keys() []K {
 return om.keys
}

// only iterate values when use this function
func(om *OrderedMap[K, V]) Values() iter.Seq[V]{
    return func(yield func(V) bool){
        for _, k := range om.keys{
            if !yield(om.Get(k)){
                return 
            }
        }
    }
}



// iterate key and value
func (om *OrderedMap[K, V]) KeyValues() iter.Seq2[K, V] {
    return func(yield func(K, V) bool) {
       for _, k := range om.keys {
          if !yield(k, om.Get(k)) {
             return
          }
       }
    }
}

