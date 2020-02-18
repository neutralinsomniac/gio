// SPDX-License-Identifier: Unlicense OR MIT

package window

import (
	"gioui.org/app/internal/d3d11"
)

type d3d11Context struct {
	win *window
	*d3d11.Context
}

func init() {
	backends = append(backends, backend{
		priority: 1,
		initializer: func(w *window) (Context, error) {
			hwnd, _, _ := w.HWND()
			ctx, err := d3d11.NewContext(hwnd)
			if err != nil {
				return nil, err
			}
			return &d3d11Context{win: w, Context: ctx}, nil
		},
	})
}

func (c *d3d11Context) MakeCurrent() error {
	return nil
}

func (c *d3d11Context) Lock() {}

func (c *d3d11Context) Unlock() {}
