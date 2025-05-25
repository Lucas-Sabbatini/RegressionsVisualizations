import { redrawPlot, redrawGradientPlot, addSurfaceToPlot, addPointToGradientPlot} from './plots.js';
import { updatePedictionFunction } from './parseFormat.js';
const spanElemento = document.getElementById('costValue');

export function start(input) {
    let featuresMatrix = featuresMatrixToJs(input)
    let output = generateRandomDots(featuresMatrix);
    redrawPlot(output.x,output.y,output.z)

    let costSurface = costSurfaceToJs(featuresMatrix,output.z)

    redrawGradientPlot(costSurface)

    setTimeout(gradientDescent(output.z,featuresMatrix, costSurface), 2000);
}


function generateRandom() {
    let max = 5;
    let min = -5;
    return Math.random() * (max - min) + min;
}

function gradientDescent(yAxis, featuresMatrix, costSurface){
    var b = generateRandom()
    var w = Array(featuresMatrix[0][0].length).fill(generateRandom())
    let gradientPointsTraceExists = false;

    var responseObject = gradientDescentToJs(featuresMatrix, yAxis, w, b)
    let [newW, newB, newJ] = adjustGradientPoint(responseObject.w[0],responseObject.b,costSurface)
    spanElemento.textContent = responseObject.j.toFixed(2);
    updatePedictionFunction(responseObject.w,responseObject.b)

    addSurfaceToPlot(responseObject.predictionPlot)
    addPointToGradientPlot(newW, newB, newJ,gradientPointsTraceExists)
}


function roundToNearestPointCostSurface(num) {
  return  Math.round(num + 10);
}

function adjustGradientPoint(wAxis,bAxis,costSurface){
    let roundedW = roundToNearestPointCostSurface(wAxis)
    let roundedB = roundToNearestPointCostSurface(bAxis)
    let roundedJ = costSurface[roundedW][roundedB]

    return [roundedW, roundedB, roundedJ];
}