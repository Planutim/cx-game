out vec4 frag_colour;

in vec2 texCoords;

uniform sampler2DArray u_texture;
uniform float imageLayer;


void main(){
    frag_colour = texture(u_texture, vec3(texCoords, imageLayer));
}


/*

0.09 0.18

0.2 0.4
*/