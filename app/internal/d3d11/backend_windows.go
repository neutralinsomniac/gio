// SPDX-License-Identifier: Unlicense OR MIT

package d3d11

import (
	"fmt"
	"image"
	"unsafe"

	"gioui.org/gpu"
	"golang.org/x/sys/windows"
)

type Context struct {
	dev          *_ID3D11Device
	ctx          *_ID3D11DeviceContext
	swchain      *_IDXGISwapChain
	renderTarget *_ID3D11RenderTargetView
	featLvl      uint32
}

type Backend struct {
	clearColor [4]float32
	viewport   _D3D11_VIEWPORT

	ctx  *Context
	caps gpu.Caps
}

type Texture struct {
}

type Program struct {
	backend *Backend
	vertShader *_ID3D11VertexShader
	pixShader *_ID3D11PixelShader
}

type Framebuffer struct {
}

type Buffer struct {
	bind uint32
	buf  *_ID3D11Buffer
}

func NewContext(hwnd windows.Handle) (*Context, error) {
	dev, d3dctx, swchain, featLvl, err := _D3D11CreateDeviceAndSwapChain(_D3D_DRIVER_TYPE_HARDWARE, &_DXGI_SWAP_CHAIN_DESC{
		BufferDesc: _DXGI_MODE_DESC{
			Format: _DXGI_FORMAT_R8G8B8A8_UNORM_SRGB,
		},
		SampleDesc: _DXGI_SAMPLE_DESC{
			Count: 1,
		},
		BufferUsage:  _DXGI_USAGE_RENDER_TARGET_OUTPUT,
		BufferCount:  1,
		OutputWindow: hwnd,
		Windowed:     1,
		SwapEffect:   _DXGI_SWAP_EFFECT_DISCARD,
	})
	if err != nil {
		return nil, err
	}
	if featLvl < _D3D_FEATURE_LEVEL_9_1 {
		_IUnknownRelease(unsafe.Pointer(dev), dev.vtbl.Release)
		return nil, fmt.Errorf("d3d11: feature level too low: %d", featLvl)
	}
	backBuffer, err := swchain.GetBuffer(0, &_IID_ID3D11Texture2D)
	if err != nil {
		_IUnknownRelease(unsafe.Pointer(dev), dev.vtbl.Release)
		_IUnknownRelease(unsafe.Pointer(swchain), swchain.vtbl.Release)
		_IUnknownRelease(unsafe.Pointer(d3dctx), d3dctx.vtbl.Release)
		return nil, err
	}
	renderTarget, err := dev.CreateRenderTargetView((*_ID3D11Resource)(unsafe.Pointer(backBuffer)))
	_IUnknownRelease(unsafe.Pointer(backBuffer), backBuffer.vtbl.Release)
	if err != nil {
		_IUnknownRelease(unsafe.Pointer(dev), dev.vtbl.Release)
		_IUnknownRelease(unsafe.Pointer(swchain), swchain.vtbl.Release)
		_IUnknownRelease(unsafe.Pointer(d3dctx), d3dctx.vtbl.Release)
		return nil, err
	}
	d3dctx.OMSetRenderTargets(renderTarget)
	ctx := &Context{dev: dev, ctx: d3dctx, swchain: swchain, renderTarget: renderTarget, featLvl: featLvl}
	return ctx, nil
}

func (c *Context) Backend() (gpu.Backend, error) {
	return newBackend(c)
}

func (c *Context) Release() {
	_IUnknownRelease(unsafe.Pointer(c.renderTarget), c.renderTarget.vtbl.Release)
	_IUnknownRelease(unsafe.Pointer(c.swchain), c.swchain.vtbl.Release)
	_IUnknownRelease(unsafe.Pointer(c.ctx), c.ctx.vtbl.Release)
	_IUnknownRelease(unsafe.Pointer(c.dev), c.dev.vtbl.Release)
}

func (c *Context) Present() error {
	return c.swchain.Present(0, 0)
}

func newBackend(c *Context) (*Backend, error) {
	caps := gpu.Caps{
		MaxTextureSize: 2048, // 9.1 maximum
	}
	switch {
	case c.featLvl >= _D3D_FEATURE_LEVEL_11_0:
		caps.MaxTextureSize = 16384
	case c.featLvl >= _D3D_FEATURE_LEVEL_9_3:
		caps.MaxTextureSize = 4096
	}
	return &Backend{
		ctx:  c,
		caps: caps,
	}, nil
}

func (b *Backend) BeginFrame() {
}

func (b *Backend) EndFrame() {
}

func (b *Backend) Caps() gpu.Caps {
	return b.caps
}

func (b *Backend) NewTimer() gpu.Timer {
	panic("timers not supported")
}

func (b *Backend) IsTimeContinuous() bool {
	panic("timers not supported")
}

func (b *Backend) NewTexture(minFilter, magFilter gpu.TextureFilter) gpu.Texture {
	return &Texture{}
}

func (b *Backend) DefaultFramebuffer() gpu.Framebuffer {
	return &Framebuffer{}
}

func (b *Backend) NilTexture() gpu.Texture {
	return &Texture{}
}

func (b *Backend) NewFramebuffer() gpu.Framebuffer {
	return &Framebuffer{}
}

func (b *Backend) NewBuffer(typ gpu.BufferType, data []byte) gpu.Buffer {
	var bind uint32
	switch typ {
	case gpu.BufferTypeData:
		bind = _D3D11_BIND_VERTEX_BUFFER
	case gpu.BufferTypeIndices:
		bind = _D3D11_BIND_INDEX_BUFFER
	default:
		panic("unsupported buffer type")
	}
	buf, err := b.ctx.dev.CreateBuffer(&_D3D11_BUFFER_DESC{
		ByteWidth: uint32(len(data)),
		Usage:     _D3D11_USAGE_IMMUTABLE,
		BindFlags: bind,
	}, data)
	if err != nil {
		panic(err)
	}
	return &Buffer{buf: buf, bind: bind}
}

func (b *Backend) NewProgram(vertexShader, fragmentShader gpu.ShaderSources, attribMap []string) (gpu.Program, error) {
	vs, err := b.ctx.dev.CreateVertexShader(vertexShader.HLSL)
	if err != nil {
		return nil, err
	}
	ps, err := b.ctx.dev.CreatePixelShader(fragmentShader.HLSL)
	if err != nil {
		return nil, err
	}
	return &Program{backend: b, vertShader:vs, pixShader: ps}, nil
}

func (b *Backend) SetupVertexArray(slot int, size int, dataType gpu.DataType, stride, offset int) {
}

func (b *Backend) DepthFunc(f gpu.DepthFunc) {
}

func (b *Backend) ClearColor(colr, colg, colb, cola float32) {
	b.clearColor = [...]float32{colr, colg, colb, cola}
}

func (b *Backend) ClearDepth(d float32) {
}

func (b *Backend) Clear(buffers gpu.BufferAttachments) {
	if buffers&gpu.BufferAttachmentColor != 0 {
		b.ctx.ctx.ClearRenderTargetView(b.ctx.renderTarget, &b.clearColor)
	}
}

func (b *Backend) Viewport(x, y, width, height int) {
	b.viewport = _D3D11_VIEWPORT{
		TopLeftX: float32(x),
		TopLeftY: float32(y),
		Width:    float32(width),
		Height:   float32(height),
		MinDepth: 0.0,
		MaxDepth: 1.0,
	}
	b.ctx.ctx.RSSetViewports(&b.viewport)
}

func (b *Backend) DrawArrays(mode gpu.DrawMode, off, count int) {
}

func (b *Backend) DrawElements(mode gpu.DrawMode, off, count int) {
}

func (b *Backend) SetBlend(enable bool) {
}

func (b *Backend) SetDepthTest(enable bool) {
}

func (b *Backend) DepthMask(mask bool) {
}

func (b *Backend) BlendFunc(sfactor, dfactor gpu.BlendFactor) {
}

func (t *Texture) Upload(img *image.RGBA) {
}

func (t *Texture) Release() {
}

func (t *Texture) Bind(unit int) {
}

func (t *Texture) Resize(format gpu.TextureFormat, width, height int) {
}

func (p *Program) Bind() {
	p.backend.ctx.ctx.VSSetShader(p.vertShader)
	p.backend.ctx.ctx.PSSetShader(p.pixShader)
}

func (p *Program) Release() {
	_IUnknownRelease(unsafe.Pointer(p.vertShader), p.vertShader.vtbl.Release)
	_IUnknownRelease(unsafe.Pointer(p.pixShader), p.pixShader.vtbl.Release)
}

func (p *Program) UniformFor(uniform string) gpu.Uniform {
	return nil
}

func (p *Program) Uniform1i(u gpu.Uniform, v int) {
}

func (p *Program) Uniform1f(u gpu.Uniform, v float32) {
}

func (p *Program) Uniform2f(u gpu.Uniform, v0, v1 float32) {
}

func (p *Program) Uniform4f(u gpu.Uniform, v0, v1, v2, v3 float32) {
}

func (b *Buffer) Bind() {
}

func (b *Buffer) Release() {
	_IUnknownRelease(unsafe.Pointer(b.buf), b.buf.vtbl.Release)
}

func (f *Framebuffer) Bind() {
}

func (f *Framebuffer) BindTexture(t gpu.Texture) {
}

func (f *Framebuffer) Invalidate() {
}

func (f *Framebuffer) Release() {
}

func (f *Framebuffer) IsComplete() error {
	return nil
}
