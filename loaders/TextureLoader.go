package loaders

// class TextureLoader extends Loader {
//
// 	constructor( manager ) {
//
// 		super( manager );
//
// 	}
//
// 	load( url, onLoad, onProgress, onError ) {
//
// 		const texture = new Texture();
//
// 		const loader = new ImageLoader( this.manager );
// 		loader.setCrossOrigin( this.crossOrigin );
// 		loader.setPath( this.path );
//
// 		loader.load( url, function ( image ) {
//
// 			texture.image = image;
// 			texture.needsUpdate = true;
//
// 			if ( onLoad !== undefined ) {
//
// 				onLoad( texture );
//
// 			}
//
// 		}, onProgress, onError );
//
// 		return texture;
//
// 	}
//
// }
