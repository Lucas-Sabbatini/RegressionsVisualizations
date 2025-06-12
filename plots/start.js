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
    costSurface,
    maxIterations,
    delay
  );
}

async function runGradientDescent(
  yAxis,
  featuresMatrix,
  costSurface,
  maxIterations,
  delay
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

    const [plotW, plotB, plotJ] = adjustGradientPoint(newW[0], newB, costSurface);
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
  let numRounded = Math.round(num/2 + 9)
  return Math.min(19, Math.max(0, numRounded));
}

function adjustGradientPoint(wAxis, bAxis, costSurface) {
  const i = roundToNearestPointCostSurface(wAxis);
  const j = roundToNearestPointCostSurface(bAxis);
  return [i, j, costSurface[j][i]+10];
}
