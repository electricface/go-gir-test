package main

import (
	"github.com/electricface/go-gir3/gi"

	"github.com/electricface/go-gir/gtk-3.0"
	gsv "github.com/electricface/go-gir/gtksource-4"
)

func main() {
	gtk.Init(0, 0)
	window := gtk.NewWindow(gtk.WindowTypeToplevel)
	window.SetTitle("SourceView")
	window.Connect(gtk.SigDestroy, gtk.MainQuit)

	swin := gtk.NewScrolledWindow(nil, nil)
	swin.SetPolicy(gtk.PolicyTypeAutomatic, gtk.PolicyTypeAutomatic)
	swin.SetShadowType(gtk.ShadowTypeIn)
	sourcebuffer := gsv.NewBufferWithLanguage(gsv.NewLanguageManager().GetLanguage("cpp"))
	sourceview := gsv.NewViewWithBuffer(sourcebuffer)

	start := gtk.TextIter{P: gi.Malloc(gtk.SizeOfStructTextIter)}
	defer start.Free()
	sourcebuffer.GetStartIter(start)
	sourcebuffer.BeginNotUndoableAction()
	const code = `#include <iostream>
template<class T>
struct foo_base {
  T operator+(T const &rhs) const {
    T tmp(static_cast<T const &>(*this));
    tmp += rhs;
    return tmp;
  }
};

class foo : public foo_base<foo> {
private:
  int v;
public:
  foo(int v) : v(v) {}
  foo &operator+=(foo const &rhs){
    this->v += rhs.v;
    return *this;
  }
  operator int() { return v; }
};

int main(void) {
  foo a(1), b(2);
  a += b;
  std::cout << (int)a << std::endl;
}
`
	sourcebuffer.Insert(start, code, int32(len(code)))
	sourcebuffer.EndNotUndoableAction()

	swin.Add(sourceview)

	window.Add(swin)
	window.SetSizeRequest(400, 300)
	window.ShowAll()

	gtk.Main()
}
