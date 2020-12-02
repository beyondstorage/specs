package specs

import (
	"sort"
)

// Interface is the spec for interface.
type Interface struct {
	Name        string       `hcl:",label"`
	Description string       `hcl:"description,optional"`
	Internal    bool         `hcl:"internal,optional"`
	Embed       []string     `hcl:"embed,optional"`
	Ops         []*Operation `hcl:"op,block"`
}

// Sort will sort the Interface
func (i *Interface) Sort() {
	sort.Strings(i.Embed)
}

// Operations is the spec for operations.
type Operations struct {
	Interfaces []*Interface `hcl:"interface,block"`
	Fields     []*Field     `hcl:"field,block"`
}

// Sort will sort the Operations
func (o *Operations) Sort() {
	sort.Slice(o.Fields, func(i, j int) bool {
		x := o.Fields
		return x[i].Name < x[j].Name
	})

	sort.Slice(o.Interfaces, func(i, j int) bool {
		x := o.Interfaces
		return x[i].Name < x[j].Name
	})

	for _, v := range o.Interfaces {
		v.Sort()
	}
}

// Operation is the spec for operation.
type Operation struct {
	Name        string   `hcl:",label"`
	Description string   `hcl:"description,optional"`
	Params      []string `hcl:"params,optional"`
	Pairs       []string `hcl:"pairs,optional"`
	Results     []string `hcl:"results,optional"`
}

// Field is the spec for field.
type Field struct {
	Name string `hcl:",label"`
	Type string `hcl:"type"`
}

// Service is the data parsed from HCL.
type Service struct {
	Name       string       `hcl:"name"`
	Namespaces []*Namespace `hcl:"namespace,block"`
	Pairs      *Pairs       `hcl:"pairs,block"`
	Infos      *Infos       `hcl:"infos,block"`
}

// Sort will sort ther service spec.
func (s *Service) Sort() {
	s.Pairs.Sort()
	s.Infos.Sort()

	for _, v := range s.Namespaces {
		v.Sort()
	}
}

// Infos is the spec for infos.
type Infos struct {
	Infos []*Info `hcl:"info,block"`
}

// Sort will sort the pair spec.
func (p *Infos) Sort() {
	if p == nil || len(p.Infos) == 0 {
		return
	}

	sort.Slice(p.Infos, func(i, j int) bool {
		x := p.Infos
		return compareInfoSpec(x[i], x[j])
	})
}

// Info is the spec for info.
type Info struct {
	Scope       string `hcl:",label"`
	Category    string `hcl:",label"`
	Name        string `hcl:",label"`
	Type        string `hcl:"type"`
	DisplayName string `hcl:"display_name,optional"`
	ZeroValue   string `hcl:"zero_value,optional"`
	Export      bool   `hcl:"export,optional"`
	Comment     string `hcl:"comment,optional"`
}

// Pairs is the data parsed from HCL.
type Pairs struct {
	Pairs []*Pair `hcl:"pair,block"`
}

// Sort will sort the pair spec.
func (p *Pairs) Sort() {
	if p == nil || len(p.Pairs) == 0 {
		return
	}

	sort.Slice(p.Pairs, func(i, j int) bool {
		x := p.Pairs
		return x[i].Name < x[j].Name
	})
}

// Pair is the data parsed from HCL.
type Pair struct {
	Name        string `hcl:",label"`
	Type        string `hcl:"type"`
	Description string `hcl:"description,optional"`
	Parser      string `hcl:"parser,optional"`
	Default     string `hcl:"default,optional"`
}

// Namespace is the data parsed from HCL.
type Namespace struct {
	Name      string   `hcl:",label"`
	Implement []string `hcl:"implement,optional"`
	New       *New     `hcl:"new,block"`
	Op        []*Op    `hcl:"op,block"`
}

// Sort will sort the Namespace
func (n *Namespace) Sort() {
	n.New.Sort()

	sort.Strings(n.Implement)

	sort.Slice(n.Op, func(i, j int) bool {
		x := n.Op
		return x[i].Name < x[j].Name
	})

	for _, v := range n.Op {
		v.Sort()
	}
}

// Op means an operation definition.
type Op struct {
	Name     string   `hcl:",label"`
	Required []string `hcl:"required,optional"`
	Optional []string `hcl:"optional,optional"`
}

// Sort will sort the Op
func (o *Op) Sort() {
	sort.Strings(o.Required)
	sort.Strings(o.Optional)
}

// New is the spec for new function.
type New struct {
	Required []string `hcl:"required,optional"`
	Optional []string `hcl:"optional,optional"`
}

// Sort will sort the New
func (o *New) Sort() {
	sort.Strings(o.Required)
	sort.Strings(o.Optional)
}

func compareInfoSpec(x, y *Info) bool {
	if x.Scope != y.Scope {
		return x.Scope < y.Scope
	}
	if x.Category != y.Category {
		return x.Category < y.Category
	}
	return x.Name < y.Name
}
