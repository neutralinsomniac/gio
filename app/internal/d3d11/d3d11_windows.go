// SPDX-License-Identifier: Unlicense OR MIT

package d3d11

import (
	"fmt"
	"unsafe"

	"syscall"

	gunsafe "gioui.org/internal/unsafe"

	"golang.org/x/sys/windows"
)

type _DXGI_SWAP_CHAIN_DESC struct {
	BufferDesc   _DXGI_MODE_DESC
	SampleDesc   _DXGI_SAMPLE_DESC
	BufferUsage  uint32
	BufferCount  uint32
	OutputWindow windows.Handle
	Windowed     uint32
	SwapEffect   uint32
	Flags        uint32
}

type _DXGI_SAMPLE_DESC struct {
	Count   uint32
	Quality uint32
}

type _DXGI_MODE_DESC struct {
	Width            uint32
	Height           uint32
	RefreshRate      _DXGI_RATIONAL
	Format           uint32
	ScanlineOrdering uint32
	Scaling          uint32
}

type _DXGI_RATIONAL struct {
	Numerator   uint32
	Denominator uint32
}

type _IDXGISwapChain struct {
	vtbl *struct {
		_IUnknownVTbl
		SetPrivateData          uintptr
		SetPrivateDataInterface uintptr
		GetPrivateData          uintptr
		GetParent               uintptr
		GetDevice               uintptr
		Present                 uintptr
		GetBuffer               uintptr
		SetFullscreenState      uintptr
		GetFullscreenState      uintptr
		GetDesc                 uintptr
		ResizeBuffers           uintptr
		ResizeTarget            uintptr
		GetContainingOutput     uintptr
		GetFrameStatistics      uintptr
		GetLastPresentCount     uintptr
	}
}

type _ID3D11Device struct {
	vtbl *struct {
		_IUnknownVTbl
		CreateBuffer                         uintptr
		CreateTexture1D                      uintptr
		CreateTexture2D                      uintptr
		CreateTexture3D                      uintptr
		CreateShaderResourceView             uintptr
		CreateUnorderedAccessView            uintptr
		CreateRenderTargetView               uintptr
		CreateDepthStencilView               uintptr
		CreateInputLayout                    uintptr
		CreateVertexShader                   uintptr
		CreateGeometryShader                 uintptr
		CreateGeometryShaderWithStreamOutput uintptr
		CreatePixelShader                    uintptr
		CreateHullShader                     uintptr
		CreateDomainShader                   uintptr
		CreateComputeShader                  uintptr
		CreateClassLinkage                   uintptr
		CreateBlendState                     uintptr
		CreateDepthStencilState              uintptr
		CreateRasterizerState                uintptr
		CreateSamplerState                   uintptr
		CreateQuery                          uintptr
		CreatePredicate                      uintptr
		CreateCounter                        uintptr
		CreateDeferredContext                uintptr
		OpenSharedResource                   uintptr
		CheckFormatSupport                   uintptr
		CheckMultisampleQualityLevels        uintptr
		CheckCounterInfo                     uintptr
		CheckCounter                         uintptr
		CheckFeatureSupport                  uintptr
		GetPrivateData                       uintptr
		SetPrivateData                       uintptr
		SetPrivateDataInterface              uintptr
		GetFeatureLevel                      uintptr
		GetCreationFlags                     uintptr
		GetDeviceRemovedReason               uintptr
		GetImmediateContext                  uintptr
		SetExceptionMode                     uintptr
		GetExceptionMode                     uintptr
	}
}

type _ID3D11DeviceContext struct {
	vtbl *struct {
		_IUnknownVTbl
		GetDevice                                 uintptr
		GetPrivateData                            uintptr
		SetPrivateData                            uintptr
		SetPrivateDataInterface                   uintptr
		VSSetConstantBuffers                      uintptr
		PSSetShaderResources                      uintptr
		PSSetShader                               uintptr
		PSSetSamplers                             uintptr
		VSSetShader                               uintptr
		DrawIndexed                               uintptr
		Draw                                      uintptr
		Map                                       uintptr
		Unmap                                     uintptr
		PSSetConstantBuffers                      uintptr
		IASetInputLayout                          uintptr
		IASetVertexBuffers                        uintptr
		IASetIndexBuffer                          uintptr
		DrawIndexedInstanced                      uintptr
		DrawInstanced                             uintptr
		GSSetConstantBuffers                      uintptr
		GSSetShader                               uintptr
		IASetPrimitiveTopology                    uintptr
		VSSetShaderResources                      uintptr
		VSSetSamplers                             uintptr
		Begin                                     uintptr
		End                                       uintptr
		GetData                                   uintptr
		SetPredication                            uintptr
		GSSetShaderResources                      uintptr
		GSSetSamplers                             uintptr
		OMSetRenderTargets                        uintptr
		OMSetRenderTargetsAndUnorderedAccessViews uintptr
		OMSetBlendState                           uintptr
		OMSetDepthStencilState                    uintptr
		SOSetTargets                              uintptr
		DrawAuto                                  uintptr
		DrawIndexedInstancedIndirect              uintptr
		DrawInstancedIndirect                     uintptr
		Dispatch                                  uintptr
		DispatchIndirect                          uintptr
		RSSetState                                uintptr
		RSSetViewports                            uintptr
		RSSetScissorRects                         uintptr
		CopySubresourceRegion                     uintptr
		CopyResource                              uintptr
		UpdateSubresource                         uintptr
		CopyStructureCount                        uintptr
		ClearRenderTargetView                     uintptr
		ClearUnorderedAccessViewUint              uintptr
		ClearUnorderedAccessViewFloat             uintptr
		ClearDepthStencilView                     uintptr
		GenerateMips                              uintptr
		SetResourceMinLOD                         uintptr
		GetResourceMinLOD                         uintptr
		ResolveSubresource                        uintptr
		ExecuteCommandList                        uintptr
		HSSetShaderResources                      uintptr
		HSSetShader                               uintptr
		HSSetSamplers                             uintptr
		HSSetConstantBuffers                      uintptr
		DSSetShaderResources                      uintptr
		DSSetShader                               uintptr
		DSSetSamplers                             uintptr
		DSSetConstantBuffers                      uintptr
		CSSetShaderResources                      uintptr
		CSSetUnorderedAccessViews                 uintptr
		CSSetShader                               uintptr
		CSSetSamplers                             uintptr
		CSSetConstantBuffers                      uintptr
		VSGetConstantBuffers                      uintptr
		PSGetShaderResources                      uintptr
		PSGetShader                               uintptr
		PSGetSamplers                             uintptr
		VSGetShader                               uintptr
		PSGetConstantBuffers                      uintptr
		IAGetInputLayout                          uintptr
		IAGetVertexBuffers                        uintptr
		IAGetIndexBuffer                          uintptr
		GSGetConstantBuffers                      uintptr
		GSGetShader                               uintptr
		IAGetPrimitiveTopology                    uintptr
		VSGetShaderResources                      uintptr
		VSGetSamplers                             uintptr
		GetPredication                            uintptr
		GSGetShaderResources                      uintptr
		GSGetSamplers                             uintptr
		OMGetRenderTargets                        uintptr
		OMGetRenderTargetsAndUnorderedAccessViews uintptr
		OMGetBlendState                           uintptr
		OMGetDepthStencilState                    uintptr
		SOGetTargets                              uintptr
		RSGetState                                uintptr
		RSGetViewports                            uintptr
		RSGetScissorRects                         uintptr
		HSGetShaderResources                      uintptr
		HSGetShader                               uintptr
		HSGetSamplers                             uintptr
		HSGetConstantBuffers                      uintptr
		DSGetShaderResources                      uintptr
		DSGetShader                               uintptr
		DSGetSamplers                             uintptr
		DSGetConstantBuffers                      uintptr
		CSGetShaderResources                      uintptr
		CSGetUnorderedAccessViews                 uintptr
		CSGetShader                               uintptr
		CSGetSamplers                             uintptr
		CSGetConstantBuffers                      uintptr
		ClearState                                uintptr
		Flush                                     uintptr
		GetType                                   uintptr
		GetContextFlags                           uintptr
		FinishCommandList                         uintptr
	}
}

type _ID3D11RenderTargetView struct {
	vtbl *struct {
		_IUnknownVTbl
	}
}

type _ID3D11Resource struct {
	vtbl *struct {
		_IUnknownVTbl
	}
}

type _ID3D11Texture2D struct {
	vtbl *struct {
		_IUnknownVTbl
	}
}

type _ID3D11Buffer struct {
	vtbl *struct {
		_IUnknownVTbl
	}
}

type _ID3D11PixelShader struct {
	vtbl *struct {
		_IUnknownVTbl
	}
}

type _ID3D11VertexShader struct {
	vtbl *struct {
		_IUnknownVTbl
	}
}

type _ID3DBlob struct {
	vtbl *struct {
		_IUnknownVTbl
		GetBufferPointer uintptr
		GetBufferSize    uintptr
	}
}

type _IUnknown struct {
	vtbl *struct {
		_IUnknownVTbl
	}
}

type _IUnknownVTbl struct {
	QueryInterface uintptr
	AddRef         uintptr
	Release        uintptr
}

type _D3D11_BUFFER_DESC struct {
	ByteWidth           uint32
	Usage               uint32
	BindFlags           uint32
	CPUAccessFlags      uint32
	MiscFlags           uint32
	StructureByteStride uint32
}

type _GUID struct {
	Data1   uint32
	Data2   uint16
	Data3   uint16
	Data4_0 uint8
	Data4_1 uint8
	Data4_2 uint8
	Data4_3 uint8
	Data4_4 uint8
	Data4_5 uint8
	Data4_6 uint8
	Data4_7 uint8
}

type _D3D11_VIEWPORT struct {
	TopLeftX float32
	TopLeftY float32
	Width    float32
	Height   float32
	MinDepth float32
	MaxDepth float32
}

type _D3D11_SUBRESOURCE_DATA struct {
	pSysMem *byte
}

var (
	_IID_ID3D11Texture2D = _GUID{0x6f15aaf2, 0xd208, 0x4e89, 0x9a, 0xb4, 0x48, 0x95, 0x35, 0xd3, 0x4f, 0x9c}
)

var (
	d3d11 = windows.NewLazySystemDLL("d3d11.dll")

	__D3D11CreateDevice             = d3d11.NewProc("D3D11CreateDevice")
	__D3D11CreateDeviceAndSwapChain = d3d11.NewProc("D3D11CreateDeviceAndSwapChain")

	d3dcompiler_47 = windows.NewLazySystemDLL("d3dcompiler_47.dll")

	__D3DCompile = d3dcompiler_47.NewProc("D3DCompile")
)

const (
	_D3D11_SDK_VERSION        = 7
	_D3D_DRIVER_TYPE_HARDWARE = 1

	_DXGI_FORMAT_R8G8B8A8_UNORM_SRGB = 29

	_DXGI_USAGE_RENDER_TARGET_OUTPUT = 1 << (1 + 4)

	_DXGI_SWAP_EFFECT_DISCARD = 0

	_D3D_FEATURE_LEVEL_9_1  = 0x9100
	_D3D_FEATURE_LEVEL_9_3  = 0x9300
	_D3D_FEATURE_LEVEL_11_0 = 0xb000

	_D3D11_USAGE_IMMUTABLE = 1

	_D3D11_BIND_VERTEX_BUFFER = 1
	_D3D11_BIND_INDEX_BUFFER  = 2
)

func _D3D11CreateDevice(driverType uint32) (*_ID3D11Device, *_ID3D11DeviceContext, error) {
	var (
		dev *_ID3D11Device
		ctx *_ID3D11DeviceContext
	)
	r, _, _ := __D3D11CreateDevice.Call(
		0,                             // pAdapter
		uintptr(driverType),           // driverType
		0,                             // Software
		0,                             // Flags
		0,                             // pFeatureLevels
		0,                             // FeatureLevels
		_D3D11_SDK_VERSION,            // SDKVersion
		uintptr(unsafe.Pointer(&dev)), // ppDevice
		0,                             // pFeatureLevel
		uintptr(unsafe.Pointer(&ctx)), // ppImmediateContext
	)
	if r != 0 {
		return nil, nil, fmt.Errorf("D3D11CreateDevice: %#x", r)
	}
	return dev, ctx, nil
}

func _D3D11CreateDeviceAndSwapChain(driverType uint32, swapDesc *_DXGI_SWAP_CHAIN_DESC) (*_ID3D11Device, *_ID3D11DeviceContext, *_IDXGISwapChain, uint32, error) {
	var (
		dev     *_ID3D11Device
		ctx     *_ID3D11DeviceContext
		swchain *_IDXGISwapChain
		featLvl uint32
	)
	r, _, _ := __D3D11CreateDeviceAndSwapChain.Call(
		0,                                 // pAdapter
		uintptr(driverType),               // driverType
		0,                                 // Software
		0,                                 // Flags
		0,                                 // pFeatureLevels
		0,                                 // FeatureLevels
		_D3D11_SDK_VERSION,                // SDKVersion
		uintptr(unsafe.Pointer(swapDesc)), // pSwapChainDesc
		uintptr(unsafe.Pointer(&swchain)), // ppSwapChain
		uintptr(unsafe.Pointer(&dev)),     // ppDevice
		uintptr(unsafe.Pointer(&featLvl)), // pFeatureLevel
		uintptr(unsafe.Pointer(&ctx)),     // ppImmediateContext
	)
	if r != 0 {
		return nil, nil, nil, 0, fmt.Errorf("D3D11CreateDeviceAndSwapChain: %#x", r)
	}
	return dev, ctx, swchain, featLvl, nil
}

func _D3DCompile(src []byte, entryPoint, target string) ([]byte, error) {
	var (
		code   *_ID3DBlob
		errors *_ID3DBlob
	)
	entryPoint0 := []byte(entryPoint + "\x00")
	target0 := []byte(target + "\x00")
	r, _, _ := __D3DCompile.Call(
		uintptr(unsafe.Pointer(&src[0])),
		uintptr(len(src)),
		0, // pSourceName
		0, // pDefines
		0, // pInclude
		uintptr(unsafe.Pointer(&entryPoint0[0])),
		uintptr(unsafe.Pointer(&target0[0])),
		0, // Flags1
		0, // Flags2
		uintptr(unsafe.Pointer(&code)),
		uintptr(unsafe.Pointer(&errors)),
	)
	if r != 0 {
		compileErr := errors.data()
		return nil, fmt.Errorf("D3D11Compile: %#x: %s", r, compileErr)
	}
	return code.data(), nil
}

func (d *_ID3D11Device) CreateBuffer(desc *_D3D11_BUFFER_DESC, data []byte) (*_ID3D11Buffer, error) {
	var buf *_ID3D11Buffer
	r, _, _ := syscall.Syscall6(
		d.vtbl.CreateBuffer,
		4,
		uintptr(unsafe.Pointer(d)),
		uintptr(unsafe.Pointer(desc)),
		uintptr(unsafe.Pointer(&_D3D11_SUBRESOURCE_DATA{
			pSysMem: &data[0],
		})),
		uintptr(unsafe.Pointer(&buf)),
		0, 0,
	)
	if r != 0 {
		return nil, fmt.Errorf("ID3D11DeviceCreateBuffer: %#x", r)
	}
	return buf, nil
}

func (d *_ID3D11Device) CreatePixelShader(bytecode []byte) (*_ID3D11PixelShader, error) {
	var shader *_ID3D11PixelShader
	r, _, _ := syscall.Syscall6(
		d.vtbl.CreatePixelShader,
		5,
		uintptr(unsafe.Pointer(d)),
		uintptr(unsafe.Pointer(&bytecode[0])),
		uintptr(len(bytecode)),
		0, // pClassLinkage
		uintptr(unsafe.Pointer(&shader)),
		0,
	)
	if r != 0 {
		return nil, fmt.Errorf("ID3D11DeviceCreatePixelShader; %#x", r)
	}
	return shader, nil
}

func (d *_ID3D11Device) CreateVertexShader(bytecode []byte) (*_ID3D11VertexShader, error) {
	var shader *_ID3D11VertexShader
	r, _, _ := syscall.Syscall6(
		d.vtbl.CreateVertexShader,
		5,
		uintptr(unsafe.Pointer(d)),
		uintptr(unsafe.Pointer(&bytecode[0])),
		uintptr(len(bytecode)),
		0, // pClassLinkage
		uintptr(unsafe.Pointer(&shader)),
		0,
	)
	if r != 0 {
		return nil, fmt.Errorf("ID3D11DeviceCreateVertexShader; %#x", r)
	}
	return shader, nil
}

func (d *_ID3D11Device) CreateRenderTargetView(res *_ID3D11Resource) (*_ID3D11RenderTargetView, error) {
	var target *_ID3D11RenderTargetView
	r, _, _ := syscall.Syscall6(
		d.vtbl.CreateRenderTargetView,
		4,
		uintptr(unsafe.Pointer(d)),
		uintptr(unsafe.Pointer(res)),
		0, // pDesc
		uintptr(unsafe.Pointer(&target)),
		0, 0,
	)
	if r != 0 {
		return nil, fmt.Errorf("ID3D11DeviceCreateRenderTargetView: %#x", r)
	}
	return target, nil
}

func (s *_IDXGISwapChain) Present(SyncInterval int, Flags uint32) error {
	r, _, _ := syscall.Syscall(
		s.vtbl.Present,
		3,
		uintptr(unsafe.Pointer(s)),
		uintptr(SyncInterval),
		uintptr(Flags),
	)
	if r != 0 {
		return fmt.Errorf("IDXGISwapChainPresent: %#x", r)
	}
	return nil
}

func (s *_IDXGISwapChain) GetBuffer(index int, riid *_GUID) (*_IUnknown, error) {
	var buf *_IUnknown
	r, _, _ := syscall.Syscall6(
		s.vtbl.GetBuffer,
		4,
		uintptr(unsafe.Pointer(s)),
		uintptr(index),
		uintptr(unsafe.Pointer(riid)),
		uintptr(unsafe.Pointer(&buf)),
		0,
		0,
	)
	if r != 0 {
		return nil, fmt.Errorf("IDXGISwapChainGetBuffer: %#x", r)
	}
	return buf, nil
}

func (c *_ID3D11DeviceContext) ClearRenderTargetView(target *_ID3D11RenderTargetView, color *[4]float32) {
	syscall.Syscall(
		c.vtbl.ClearRenderTargetView,
		3,
		uintptr(unsafe.Pointer(c)),
		uintptr(unsafe.Pointer(target)),
		uintptr(unsafe.Pointer(color)),
	)
}

func (c *_ID3D11DeviceContext) RSSetViewports(viewport *_D3D11_VIEWPORT) {
	syscall.Syscall(
		c.vtbl.RSSetViewports,
		3,
		uintptr(unsafe.Pointer(c)),
		1, // NumViewports
		uintptr(unsafe.Pointer(viewport)),
	)
}

func (c *_ID3D11DeviceContext) VSSetShader(s *_ID3D11VertexShader) {
	syscall.Syscall6(
		c.vtbl.VSSetShader,
		4,
		uintptr(unsafe.Pointer(c)),
		uintptr(unsafe.Pointer(s)),
		0, // ppClassInstances
		0, // NumClassInstances
		0, 0,
	)
}

func (c *_ID3D11DeviceContext) PSSetShader(s *_ID3D11PixelShader) {
	syscall.Syscall6(
		c.vtbl.PSSetShader,
		4,
		uintptr(unsafe.Pointer(c)),
		uintptr(unsafe.Pointer(s)),
		0, // ppClassInstances
		0, // NumClassInstances
		0, 0,
	)
}

func (c *_ID3D11DeviceContext) OMSetRenderTargets(target *_ID3D11RenderTargetView) {
	syscall.Syscall6(
		c.vtbl.OMSetRenderTargets,
		4,
		uintptr(unsafe.Pointer(c)),
		1, // NumViews
		uintptr(unsafe.Pointer(&target)),
		0, // pDepthStencilView
		0, 0,
	)
}

func (b *_ID3DBlob) GetBufferPointer() uintptr {
	ptr, _, _ := syscall.Syscall(
		b.vtbl.GetBufferPointer,
		1,
		uintptr(unsafe.Pointer(b)),
		0,
		0,
	)
	return ptr
}

func (b *_ID3DBlob) GetBufferSize() uintptr {
	sz, _, _ := syscall.Syscall(
		b.vtbl.GetBufferSize,
		1,
		uintptr(unsafe.Pointer(b)),
		0,
		0,
	)
	return sz
}

func (b *_ID3DBlob) data() []byte {
	data := gunsafe.SliceOf(b.GetBufferPointer())
	n := int(b.GetBufferSize())
	return data[:n:n]
}

func _IUnknownRelease(obj unsafe.Pointer, releaseMethod uintptr) {
	syscall.Syscall(
		releaseMethod,
		1,
		uintptr(obj),
		0,
		0,
	)
}
