in vec2 spriteCoord;
flat in int instance; 

out vec4 frag_colour;

in vec4 poss;

uniform sampler2D tex;
uniform mat3 uvtransforms[NUM_INSTANCES];

void main() {
		vec2 texCoord =
			 vec2(uvtransforms[instance] * vec3(spriteCoord,1) );

		// frag_colour = texture(tex, texCoord);
		// frag_colour = vec4(0,1,0,1);
		if (poss.x < -20) {
			frag_colour = vec4(0,1,0,1);
		}else{
			frag_colour = vec4(1,0,0,1);
		}
		if (frag_colour.a < 0.1) { discard; }
}
