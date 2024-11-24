package three

type WebGLCapabilities struct {
	gl            GLContext
	extensions    *WebGLExtensions
	parameters    *WebGLRendererParams
	utils         *WebGLUtils
	maxAnisotropy int
}

func NewWebGLCapabilities(gl GLContext, extensions *WebGLExtensions, parameters *WebGLRendererParams, utils *WebGLUtils) *WebGLCapabilities {
	r := &WebGLCapabilities{
		gl:            gl,
		extensions:    extensions,
		parameters:    parameters,
		utils:         utils,
		maxAnisotropy: -1,
	}

	return r
}

func (c *WebGLCapabilities) GetMaxAnisotropy() int {
	if c.maxAnisotropy < 0 {
		if c.extensions.Has("EXT_texture_filter_anisotropic") {
			extension := c.extensions.Get("EXT_texture_filter_anisotropic")
			c.maxAnisotropy = c.gl.getParameterInt(extension.Const("MAX_TEXTURE_MAX_ANISOTROPY_EXT"))
		} else {
			c.maxAnisotropy = 0
		}
	}
	return c.maxAnisotropy
}

//todo
//	function textureFormatReadable( textureFormat ) {
//
//		if ( textureFormat !== RGBAFormat && utils.convert( textureFormat ) !== gl.getParameter( gl.IMPLEMENTATION_COLOR_READ_FORMAT ) ) {
//
//			return false;
//
//		}
//
//		return true;
//
//	}
//
//	function textureTypeReadable( textureType ) {
//
//		const halfFloatSupportedByExt = ( textureType === HalfFloatType ) && ( extensions.has( 'EXT_color_buffer_half_float' ) || extensions.has( 'EXT_color_buffer_float' ) );
//
//		if ( textureType !== UnsignedByteType && utils.convert( textureType ) !== gl.getParameter( gl.IMPLEMENTATION_COLOR_READ_TYPE ) && // Edge and Chrome Mac < 52 (#9513)
//			textureType !== FloatType && ! halfFloatSupportedByExt ) {
//
//			return false;
//
//		}
//
//		return true;
//
//	}
//
//	function getMaxPrecision( precision ) {
//
//		if ( precision === 'highp' ) {
//
//			if ( gl.getShaderPrecisionFormat( gl.VERTEX_SHADER, gl.HIGH_FLOAT ).precision > 0 &&
//				gl.getShaderPrecisionFormat( gl.FRAGMENT_SHADER, gl.HIGH_FLOAT ).precision > 0 ) {
//
//				return 'highp';
//
//			}
//
//			precision = 'mediump';
//
//		}
//
//		if ( precision === 'mediump' ) {
//
//			if ( gl.getShaderPrecisionFormat( gl.VERTEX_SHADER, gl.MEDIUM_FLOAT ).precision > 0 &&
//				gl.getShaderPrecisionFormat( gl.FRAGMENT_SHADER, gl.MEDIUM_FLOAT ).precision > 0 ) {
//
//				return 'mediump';
//
//			}
//
//		}
//
//		return 'lowp';
//
//	}
//
//	let precision = parameters.precision !== undefined ? parameters.precision : 'highp';
//	const maxPrecision = getMaxPrecision( precision );
//
//	if ( maxPrecision !== precision ) {
//
//		console.warn( 'THREE.WebGLRenderer:', precision, 'not supported, using', maxPrecision, 'instead.' );
//		precision = maxPrecision;
//
//	}
//
//	const logarithmicDepthBuffer = parameters.logarithmicDepthBuffer === true;
//	const reverseDepthBuffer = parameters.reverseDepthBuffer === true && extensions.has( 'EXT_clip_control' );
//
//	const maxTextures = gl.getParameter( gl.MAX_TEXTURE_IMAGE_UNITS );
//	const maxVertexTextures = gl.getParameter( gl.MAX_VERTEX_TEXTURE_IMAGE_UNITS );
//	const maxTextureSize = gl.getParameter( gl.MAX_TEXTURE_SIZE );
//	const maxCubemapSize = gl.getParameter( gl.MAX_CUBE_MAP_TEXTURE_SIZE );
//
//	const maxAttributes = gl.getParameter( gl.MAX_VERTEX_ATTRIBS );
//	const maxVertexUniforms = gl.getParameter( gl.MAX_VERTEX_UNIFORM_VECTORS );
//	const maxVaryings = gl.getParameter( gl.MAX_VARYING_VECTORS );
//	const maxFragmentUniforms = gl.getParameter( gl.MAX_FRAGMENT_UNIFORM_VECTORS );
//
//	const vertexTextures = maxVertexTextures > 0;
//
//	const maxSamples = gl.getParameter( gl.MAX_SAMPLES );
//
//	return {
//
//		isWebGL2: true, // keeping this for backwards compatibility
//
//		getMaxAnisotropy: getMaxAnisotropy,
//		getMaxPrecision: getMaxPrecision,
//
//		textureFormatReadable: textureFormatReadable,
//		textureTypeReadable: textureTypeReadable,
//
//		precision: precision,
//		logarithmicDepthBuffer: logarithmicDepthBuffer,
//		reverseDepthBuffer: reverseDepthBuffer,
//
//		maxTextures: maxTextures,
//		maxVertexTextures: maxVertexTextures,
//		maxTextureSize: maxTextureSize,
//		maxCubemapSize: maxCubemapSize,
//
//		maxAttributes: maxAttributes,
//		maxVertexUniforms: maxVertexUniforms,
//		maxVaryings: maxVaryings,
//		maxFragmentUniforms: maxFragmentUniforms,
//
//		vertexTextures: vertexTextures,
//
//		maxSamples: maxSamples
//
//	};
