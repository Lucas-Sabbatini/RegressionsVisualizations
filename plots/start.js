import { redrawPlot, redrawGradientPlot } from './plots.js';


export function start(input) {
    let featuresMatrix = featuresMatrixToJs(input)
    let output = generateRandomDots(featuresMatrix);

    redrawPlot(output.x,output.y,output.z)

    let costSurface = costSurfaceToJs(featuresMatrix,output.z)

    redrawGradientPlot(costSurface)
}