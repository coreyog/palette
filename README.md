# palette

Works with my godot shader that uses red values as an x coordindate and blue values as a y coordinate to index a palette image to emulate paletted/indexed images used in retro consoles with paletts up to 256 colors.

## godot shader

```
shader_type canvas_item;

uniform sampler2D color_palette : filter_nearest;

void fragment() {
	vec4 data = texture(TEXTURE, UV);
	COLOR = texture(color_palette, vec2(data.r+0.03125, data.b+0.03125));
}
```