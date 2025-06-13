import {
  redrawPlot,
  redrawGradientPlot,
  addSurfaceToPlot,
  addPointToGradientPlot
} from './plots.js';
import { updatePedictionFunction } from './parseFormat.js';

const spanElemento = document.getElementById('costValue');

export async function start(input, {
  maxIterations = 50,   
  delay = 200           
} = {}) {
  const featuresMatrix = featuresMatrixToJs(input);
  const outputWB = generateRandomWeightsBias(featuresMatrix[0][0].length)
  const datasetDots = generateRandomDots(featuresMatrix,outputWB.w,outputWB.b);
  redrawPlot(datasetDots.x, datasetDots.y, datasetDots.z);
  let costSurface = costSurfaceToJs(featuresMatrix, datasetDots.z,outputWB.w);
  redrawGradientPlot(costSurface);

  await runGradientDescent(
    datasetDots.z,
    featuresMatrix,
    maxIterations,
    delay,
    outputWB.w
  );
}

async function runGradientDescent(
  yAxis,
  featuresMatrix,
  maxIterations,
  delay,
  targetW,
) {
  let b = generateRandom();
  let w = Array(featuresMatrix[0][0].length)
            .fill(0)
            .map(() => generateRandom());
  let prevCost = Infinity;
  
  for (let iter = 0; iter < maxIterations; iter++) {
    var numIteractions =  Math.trunc(Math.pow(1.07,iter))
    const response = gradientDescentToJs(featuresMatrix, yAxis, w, b,numIteractions);
    const { w: newW, b: newB, j: costJ, predictionPlot } = response;

    spanElemento.textContent = costJ.toFixed(2);
    updatePedictionFunction(newW, newB);

    addSurfaceToPlot(predictionPlot);

    const [plotW, plotB, plotJ] = adjustGradientPoint(newW, targetW, newB, costJ);
    addPointToGradientPlot(plotW, plotB, plotJ, iter > 0);

    w = newW;
    b = newB;

    if (Math.abs(prevCost - costJ) < 1e-6) break;
    prevCost = costJ;

    await sleep(delay);
  }
}

function sleep(ms) {
  return new Promise(resolve => setTimeout(resolve, ms));
}

function generateRandom() {
  const max = 5, min = -5;
  return Math.random() * (max - min) + min;
}

function roundToNearestPointCostSurface(num) {
  let numRounded = num + 10
  return Math.min(20, Math.max(0, numRounded));
}

function adjustGradientPoint(w ,w0, b, costJ) {
  const wAxis = scalerSignedDistanceToJs(w,w0) + 10.5;
  const bAxis = roundToNearestPointCostSurface(b);
  const j = costJ + 30
  console.log(b,wAxis,j)
  return [bAxis, wAxis, j];
}
