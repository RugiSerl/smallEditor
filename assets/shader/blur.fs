#version 330



// Input vertex attributes (from vertex shader)
in vec2 fragTexCoord;
in vec4 fragColor;

// Input uniform values
uniform sampler2D texture0;
uniform vec4 colDiffuse;


// Output fragment color
out vec4 finalColor;

// NOTE: Add here your custom variables
uniform float size;

// NOTE: Render size values must be passed from code
const float renderWidth = 1920;
const float renderHeight = 1080;



void main()
{
    // Texel color fetching from texture sampler
    vec4 texelColor = vec4(0.0, 0.0, 0.0, 0.0);
    int radius = int(size);
    float weightSum = 0.0;
    float maxDistance = size*sqrt(2);
    float weight;
    vec4 color;

    for (int i = -radius; i <= radius; i++) 
    {
        for (int j = -radius; j <= radius; j++) 
        {
            color = texture(texture0, fragTexCoord + vec2(i/renderWidth, j/renderHeight));
            weight = (maxDistance-sqrt(int(i*i+j*j)));
            weightSum += weight;
            
            texelColor += color*weight;
        }
    }

    

    finalColor = texelColor/weightSum*(1-size/30.0);

    if (radius == 0) {
        finalColor = texture(texture0, fragTexCoord);
    } 
    
}